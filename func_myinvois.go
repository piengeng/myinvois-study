package main

import (
	"context"
	"database/sql"
	"dev/store"
	"encoding/base64"
	"fmt"
	"log"
	"maps"
	"net/http"
	"strings"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/golang-jwt/jwt/v5"
	"github.com/imroc/req/v3"
	"github.com/valkey-io/valkey-go"
)

// nolint: unused
func renewTokens() {
	vkc, err := valkey.NewClient(vkCO)
	if err != nil {
		log.Fatalln("vkc failed on renewTokens")
	}
	defer vkc.Close()
	ctx := context.Background()

	for _, sc := range clients {
		tokenKey := "tok:" + getUUIDLastSegment(sc.ID)
		_post := func(sc Client) {
			if rpl, err := callPostLogin(sc); err == nil {
				fields := map[string]string{"jwt": rpl.AccessToken}
				cmd := vkc.B().Hset().Key(tokenKey).FieldValue().FieldValueIter(maps.All(fields))
				if _, err := vkc.Do(ctx, cmd.Build()).ToInt64(); err != nil {
					log.Println("Failed HSET on tok's fields", err)
				}
			}
		}
		_, err = vkc.Do(ctx, vkc.B().Get().Key(tokenKey).Build()).ToString()
		if err != nil && valkey.IsValkeyNil(err) { // if key missing, create with new jwt
			_post(sc)
		} else {
			// if key exist, _jwt missing callLogin, insert _jwt
			_jwt, err := vkc.Do(ctx, vkc.B().Hget().Key(tokenKey).Field("jwt").Build()).ToString()
			if err != nil {
				log.Println("Failed HGET on jwt field", err)
			}
			if isJwtExpired(_jwt) {
				_post(sc)
				log.Println(tokenKey + " renew'd")
			}
		}
	}
}

func callPostLogin(sc Client) (respPostLogin, error) {
	c := req.C().SetCommonHeaders(comHead)
	res, err := c.R().
		SetFormData(map[string]string{
			"client_id":     sc.ID,
			"client_secret": sc.Secret,
			"grant_type":    "client_credentials",
			"scope":         "InvoicingAPI",
		}).
		Post(posLogin)
	if err != nil {
		log.Println("failed to POST, do something?!")
	}
	var rpl respPostLogin
	if err = res.UnmarshalJson(&rpl); err != nil {
		return respPostLogin{}, err
	}
	return rpl, nil
}

func jwtByUUID(uuid string) string {
	tokenKey := "tok:" + getUUIDLastSegment(uuid)
	vkc, err := valkey.NewClient(vkCO)
	if err != nil {
		log.Println("vkc failed in jwtByUUID")
	}
	defer vkc.Close()
	ctx := context.Background()
	_jwt, _ := vkc.Do(ctx, vkc.B().Hget().Key(tokenKey).Field("jwt").Build()).ToString()
	return _jwt
}

func isJwtExpired(_jwt string) bool {
	isExpired := true
	parser := jwt.NewParser(jwt.WithoutClaimsValidation())
	claims := &jwt.RegisteredClaims{}
	if _, _, err := parser.ParseUnverified(_jwt, claims); err != nil {
		log.Printf("Error parsing JWT (unverified): %v\n", err)
	}
	// if claims.IssuedAt != nil { fmt.Printf("iat: %s\n", claims.IssuedAt.Time.Format(time.RFC3339)) }
	if claims.ExpiresAt != nil {
		expiryTime := claims.ExpiresAt.Time
		// fmt.Printf("exp: %s in: %v\n", expiryTime.Format(time.RFC3339), time.Until(expiryTime))
		if time.Until(expiryTime) > 5*time.Minute { // update jwt when close to expiry
			isExpired = false
		}
	} else {
		log.Println("No 'exp' (expiration) claim found in the token.")
	}
	return isExpired
}

func callGetDocTypes(uuid string) {
	_jwt := jwtByUUID(uuid) // assume jwt is not expired and valid here. :) will trust cron
	c := req.C().SetCommonHeaders(comHead)
	res, err := c.R().SetBearerAuthToken(_jwt).Get(getDocTypes)
	if err != nil {
		log.Println("Failed getDocTypes, erm...", err)
	}
	fmt.Println(res) // not much use as it's been hard-defined
}

func callGetDocType(uuid, did string) {
	_jwt := jwtByUUID(uuid)
	c := req.C().SetCommonHeaders(comHead)
	res, err := c.R().SetBearerAuthToken(_jwt).SetPathParam("id", did).Get(getDocType)
	if err != nil {
		log.Println("Failed getDocType, erm...", err)
	}
	fmt.Println(res)
}

func callGetDocTypeVer(uuid, did, ver string) {
	_jwt := jwtByUUID(uuid)
	c := req.C().SetCommonHeaders(comHead)
	res, err := c.R().SetBearerAuthToken(_jwt).
		SetPathParam("id", did).SetPathParam("vid", ver).Get(getDocTypeVer)
	if err != nil {
		log.Println("Failed getDocTypeVer, erm...", err)
	}
	fmt.Println(res)
}

func callGetNoticesDefault(uuid string) {
	_jwt := jwtByUUID(uuid)
	c := req.C().SetCommonHeaders(comHead)
	res, err := c.R().SetBearerAuthToken(_jwt).Get(getNoticesDefault)
	if err != nil {
		log.Println("Failed getNoticesDefault, erm...", err)
	}
	fmt.Println(res)
}

func callGetNotices(uuid, df, dt, t, l, s, pn, ps string) {
	_jwt := jwtByUUID(uuid)
	c := req.C().SetCommonHeaders(comHead)
	res, err := c.R().SetBearerAuthToken(_jwt).
		SetPathParam("dateFrom", df).
		SetPathParam("dateTo", dt).
		SetPathParam("type", t).
		SetPathParam("language", l).
		SetPathParam("status", s).
		SetPathParam("pageNo", pn).
		SetPathParam("pageSize", ps).
		Get(getNotices)
	if err != nil {
		log.Println("Failed getNotices, erm...", err)
	}
	fmt.Println(res)
}

func callGetValTaxTin(uuid, tin, t, v string) {
	_jwt := jwtByUUID(uuid)
	c := req.C().SetCommonHeaders(comHead)
	res, err := c.R().SetBearerAuthToken(_jwt).
		SetPathParam("tin", tin).
		SetPathParam("idType", t).
		SetPathParam("idValue", v).
		Get(getTaxValTin)
	if err != nil {
		log.Println("Failed getValTaxTin, erm...", err)
	}
	fmt.Println(res) // 200 when found?! BRN always ok on preprod.
}

func callPosDocSubmit(uuid string, d Docs) (duid string) {
	inv, _ := q.InvoiceExistsByID(ctx, d.Documents[0].CodeNumber)
	if inv == 0 {
		var rps respPostSubmit
		c := req.C().SetCommonHeaders(comHead).
			SetCommonHeader("Content-Type", "application/json")
		_, err := c.R().SetBearerAuthToken(jwtByUUID(uuid)).SetBody(d).
			SetSuccessResult(&rps).
			EnableDumpToFile(D_BASE + "refs/req.txt").
			Post(posDocSubmit)
		if err != nil {
			log.Println("Failed callPosDocSubmit, erm...", err)
			return
		}
		// uno, single, singular, one only!
		p := store.InsertSubmissionParams{
			SubUid:   rps.SubmissionUid,
			Accepted: sql.NullInt64{Int64: 1, Valid: true},
			DocUid:   rps.AcceptedDocuments[0].UUID,
			InvID:    rps.AcceptedDocuments[0].InvoiceCodeNumber,
		}
		if duid, err = q.InsertSubmission(ctx, p); err != nil {
			return duid
		}
		// fmt.Println(res) // 200 when found?! BRN always ok on preprod.
	} else {
		log.Println("Drop repeated InvoiceID:", d.Documents[0].CodeNumber)
	}
	return
}

func callGetDocDetailThenUpdate(uuid, duid string) {
	c := req.C().SetCommonHeaders(comHead)
	var dd respGetDocumentDetails
	_, err := c.R().SetBearerAuthToken(jwtByUUID(uuid)).
		SetPathParam("uuid", duid).
		SetSuccessResult(&dd).
		Get(getDocDetail)
	if err != nil {
		log.Println("Failed getDocDetail, erm...", err)
		return
	}
	// log.Println(duid, dd.Status, dd.LongID)
	p := store.UpdateSubmissionStatusLongIDParams{
		DocUid: duid,
		Status: sql.NullString{Valid: true, String: dd.Status},
		LongID: sql.NullString{Valid: true, String: dd.LongID},
	}
	if e := q.UpdateSubmissionStatusLongID(ctx, p); e != nil {
		logPrintln(e)
	}
}

func callGetTaxQRInfo(uuid, qrct string) string {
	vkc, err := valkey.NewClient(vkCO)
	if err != nil {
		log.Fatalln("vkc failed on renewTokens")
	}
	defer vkc.Close()
	ctx := context.Background()
	key := "tpi:" + qrct

	_get := func() string {
		// var tpi respGetTaxpayersQRCodeInfo
		c := req.C().SetCommonHeaders(comHead)
		res, err := c.R().SetBearerAuthToken(jwtByUUID(uuid)).
			SetPathParam("qrCodeText", qrct).Get(getTaxQRInfo)
		if err != nil {
			log.Println("Failed getTaxQRInfo, erm...", err)
		}
		ba, _ := res.ToBytes()
		k, _ := loadKey("keys.tpi")
		b64, err := encAESGCM(ba, k)
		logPrintln(err)

		_, err = vkc.Do(ctx, vkc.B().Set().Key(key).Value(b64).Ex(time.Hour*3).Build()).ToString()
		logPrintln(err)
		return b64
	}

	_json, err := vkc.Do(ctx, vkc.B().Get().Key(key).Build()).ToString()
	if err != nil && valkey.IsValkeyNil(err) { // missing?
		return _get()
	} else {
		return _json
	}
}

func scanMyInvoisQRCode() string {
	encoded := k.String("private.qr") // to simulate after scan, base64-encoded
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		log.Println(encoded, err)
	}
	return string(decoded)
}

func buildValidationLink(iid string) (vl string) {
	s, err := q.GetSubmissionByInvoiceID(ctx, iid)
	logPrintln(err)
	vl = strings.Join([]string{urlPortal, s.DocUid, "share", s.LongID.String}, "/")
	return
}

func isAuth(sc int, h http.Header) {
	spew.Dump(sc, h.Get("x-rate-limit-remaining"))
}
