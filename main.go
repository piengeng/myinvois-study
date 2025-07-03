package main

//go:generate ./ublns.py
//go:generate ./codes.py
//go:generate sqlc generate --file store/sqlc.yml

import (
	"context"
	"database/sql"
	"dev/store"
	_ "embed"
	"fmt"
	"log"
	"os"
	"os/signal"
	"regexp"
	"sync"
	"syscall"

	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
	"github.com/robfig/cron/v3"
	gxv "github.com/terminalstatic/go-xsd-validate"
	"github.com/valkey-io/valkey-go"
	_ "modernc.org/sqlite"
)

// nolint: unused
var (
	hInvoice    *gxv.XsdHandler
	cMSICSC     map[string]CodeMSICSubCategory = make(map[string]CodeMSICSubCategory)
	cCountryRev map[string]CodeCountry         = make(map[string]CodeCountry)
	cState      map[string]CodeState           = make(map[string]CodeState)
	cStateRev   map[string]CodeState           = make(map[string]CodeState)
	// cCurrency map[string]CodeCurrency = make(map[string]CodeCurrency)
	// cClassification map[string]CodeClassification = make(map[string]CodeClassification)

	reSubWP = regexp.MustCompile("(?i)" + regexp.QuoteMeta("wilayah persekutuan "))

	k = koanf.New(".")

	clients []Client // oauth clients

	urlApi, urlPortal string
	// Platform endpoints
	posLogin, getDocTypes, getDocType, getDocTypeVer, getNoticesDefault, getNotices string
	// EInvoicing endpoints
	getTaxValTin, posDocSubmit, putDocCancel, putDocReject, getDocDetail, getTaxQRInfo string

	vkCO valkey.ClientOption
	// http common headers // "onbehalfof": "tin", // for Intermediary System
	comHead = map[string]string{"Accept-Encoding": "gzip,deflate,br", "Connection": "keep-alive", "User-Agent": "req/v3"}

	//go:embed store/schema.sql
	ddl string

	db  *sql.DB
	dsn string
	ctx context.Context
	q   *store.Queries

	wg sync.WaitGroup
)

func init() {
	_ = gxv.Init()
	// defer xsdvalidate.Cleanup()
	var err error
	hInvoice, err = gxv.NewXsdHandlerUrl(D_XSDRT_MAINDOC+F_XSD_INVOICE, gxv.ParsErrDefault)
	if err != nil {
		log.Fatalln(err)
	}
	// defer xsdhandler.Free()
	cMSICSC, cState, cStateRev, cCountryRev = codes()

	os.MkdirAll(D_BASE+D_XML, os.FileMode(0750))

	ctx = context.Background()
	dsn = fmt.Sprintf("file:%s?_journal_mode=WAL", D_BASE+"invoices/myinvois.db")
	db, err = sql.Open("sqlite", dsn)
	q = store.New(db)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	// defer db.Close() // single context
	if _, err := db.ExecContext(ctx, ddl); err != nil {
		log.Fatalf("Error executing schema: %v", err)
	}
}

func konf() {
	if err := k.Load(file.Provider("config.yml"), yaml.Parser()); err != nil {
		log.Fatalf("error loading config: %v", err)
	}
	if err := k.Unmarshal("clients", &clients); err != nil {
		log.Fatalf("error unmarshalling config: %v", err)
	}
	urlApi = k.String(fmt.Sprintf("urls.%s_api", k.String("mode")))
	urlPortal = k.String(fmt.Sprintf("urls.%s_portal", k.String("mode")))

	// valkey
	vkCO = valkey.ClientOption{InitAddress: []string{k.String("valkey.url")}}

	genApiUrls()
}

func main() {
	konf()        // configs
	renewTokens() // once on start, every 2m after with cron
	studyCalls()
	os.Exit(0)

	c := cron.New(cron.WithSeconds())
	_, err := c.AddFunc("* */2 * * * *", renewTokens)
	if err != nil {
		log.Fatalln(err)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	go c.Start()
	fmt.Println("Cron jobs started. Press Ctrl+C to stop.")
	<-sigChan
	fmt.Println("Received shutdown signal. Stopping cron scheduler...")
	c.Stop()
	fmt.Println("Cron scheduler stopped. Exiting.")

	freeCleanupClose()
}
