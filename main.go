package main

//go:generate ./ublns.py

import (
	"log"
	"regexp"

	gxv "github.com/terminalstatic/go-xsd-validate"
)

// nolint: unused
var (
	hInvoice *gxv.XsdHandler
	cMSICSC  map[string]CodeMSICSubCategory = make(map[string]CodeMSICSubCategory)
	cCountry map[string]CodeCountry         = make(map[string]CodeCountry)
	cState   map[string]CodeState           = make(map[string]CodeState)
	// cCurrency map[string]CodeCurrency = make(map[string]CodeCurrency)
	// cClassification map[string]CodeClassification = make(map[string]CodeClassification)

	reSubWP = regexp.MustCompile("(?i)" + regexp.QuoteMeta("wilayah persekutuan "))
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
	cMSICSC, cState, cCountry = codes()
}

func codes() (map[string]CodeMSICSubCategory, map[string]CodeState, map[string]CodeCountry) {
	var jd []byte

	jd, _ = readToBytes(D_C + F_C_MSICSubCategory)
	var categories []CodeMSICSubCategory
	unmarshalTo(jd, &categories)
	for _, category := range categories {
		cMSICSC[category.Code] = category // ignore duplicated Code here
	}

	jd, _ = readToBytes(D_C + F_C_State)
	var states []CodeState
	unmarshalTo(jd, &states)
	for _, state := range states {
		cState[state.Code] = state
	}

	jd, _ = readToBytes(D_C + F_C_Country)
	var countries []CodeCountry
	unmarshalTo(jd, &countries)
	for _, country := range countries {
		cCountry[country.Country] = country
	}

	return cMSICSC, cState, cCountry
}

func freeCleanup() {
	hInvoice.Free()
	gxv.Cleanup()
}

func main() {
	// sdkSampleInvoice10(D_BASE + "refs/samples/invoice10.xml")
	// sdkSampleInvoiceMultiLine10(D_BASE + "refs/samples/invoicemultiline10.xml")
	// sdkSampleInvoiceConsolidated10(D_BASE + "refs/samples/invoiceconsolidated10.xml")
	// sdkSampleInvoiceForeignCurrency10(D_BASE + "refs/samples/invoiceforeigncurrency10.xml")
	// sdkSampleCreditNote10(D_BASE + "refs/samples/creditnote10.xml")
	// sdkSampleDebitNote10(D_BASE + "refs/samples/debitnote10.xml")
	// sdkSampleRefundNote10(D_BASE + "refs/samples/refundnote10.xml")
	// sdkSampleSelfBilledInvoice10(D_BASE + "refs/samples/selfbilledinvoice10.xml")
	// sdkSampleSelfBilledCreditNote10(D_BASE + "refs/samples/selfbilledcreditnote10.xml")
	// sdkSampleSelfBilledDebitNote10(D_BASE + "refs/samples/selfbilleddebitnote10.xml")
	// sdkSampleSelfBilledRefundNote10(D_BASE + "refs/samples/selfbilledrefundnote10.xml")

	// samplingInvoice10(D_BASE + "refs/sampling/invoice10.xml")
	// samplingInvoiceMultiLine10(D_BASE + "refs/sampling/invoicemultiline10.xml")
	// samplingInvoiceConsolidated10(D_BASE + "refs/sampling/invoiceconsolidated10.xml")
	// samplingInvoiceForeignCurrency10(D_BASE + "refs/sampling/invoiceforeigncurrency10.xml")
	// samplingCreditNote10(D_BASE + "refs/sampling/creditnote10.xml")
	// samplingDebitNote10(D_BASE + "refs/sampling/debitnote10.xml")
	// samplingRefundNote10(D_BASE + "refs/sampling/refundnote10.xml")
	// samplingSelfBilledInvoice10(D_BASE + "refs/sampling/selfbilledinvoice10.xml")
	// samplingSelfBilledCreditNote10(D_BASE + "refs/sampling/selfbilledcreditnote10.xml")
	// samplingSelfBilledDebitNote10(D_BASE + "refs/sampling/selfbilleddebitnote10.xml")
	// samplingSelfBilledRefundNote10(D_BASE + "refs/sampling/selfbilledrefundnote10.xml")

	// checks([]string{
	// 	D_EG + "1.0-Credit-Note-Sample.xml",
	// 	D_EG + "1.0-Debit-Note-Sample.xml",
	// 	D_EG + "1.0-Invoice-Consolidated-Sample.xml",
	// 	D_EG + "1.0-Invoice-ForeignCurrency-Sample.xml", // fail validation on TaxExchangeRate
	// 	D_EG + "1.0-Invoice-MultiLineItem-Sample.xml",
	// 	D_EG + "1.0-Invoice-Sample.xml",
	// 	D_EG + "1.0-Refund-Note-Sample.xml",
	// 	D_EG + "1.0-Self-Billed-Credit-Sample.xml",
	// 	D_EG + "1.0-Self-Billed-Debit-Sample.xml",
	// 	D_EG + "1.0-Self-Billed-Invoice-Sample.xml",
	// 	D_EG + "1.0-Self-Billed-Refund-Sample.xml",
	// })

	// checks([]string{
	// 	D_BASE + "refs/samples/invoice10.xml",
	// 	D_BASE + "refs/samples/invoicemultiline10.xml",
	// 	D_BASE + "refs/samples/invoiceconsolidated10.xml",
	// 	D_BASE + "refs/samples/invoiceforeigncurrency10.xml", // fail validation on TaxExchangeRate
	// 	D_BASE + "refs/samples/creditnote10.xml",
	// 	D_BASE + "refs/samples/debitnote10.xml",
	// 	D_BASE + "refs/samples/refundnote10.xml",
	// 	D_BASE + "refs/samples/selfbilledinvoice10.xml",
	// 	D_BASE + "refs/samples/selfbilledcreditnote10.xml",
	// 	D_BASE + "refs/samples/selfbilleddebitnote10.xml",
	// 	D_BASE + "refs/samples/selfbilledrefundnote10.xml",
	// })

	// checks([]string{
	// 	D_BASE + "refs/sampling/invoice10.xml",
	// 	D_BASE + "refs/sampling/invoicemultiline10.xml",
	// 	D_BASE + "refs/sampling/invoiceconsolidated10.xml",
	// 	D_BASE + "refs/sampling/invoiceforeigncurrency10.xml", // fixed on TaxExchangeRate
	// 	D_BASE + "refs/sampling/creditnote10.xml",
	// 	D_BASE + "refs/sampling/debitnote10.xml",
	// 	D_BASE + "refs/sampling/refundnote10.xml",
	// 	D_BASE + "refs/sampling/selfbilledinvoice10.xml",
	// 	D_BASE + "refs/sampling/selfbilledcreditnote10.xml",
	// 	D_BASE + "refs/sampling/selfbilleddebitnote10.xml",
	// 	D_BASE + "refs/sampling/selfbilledrefundnote10.xml",
	// })

	// checks([]string{D_BASE + "failed/1.0-Invoice-ForeignCurrency-Sample.xml"})
	freeCleanup()
}
