package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/beevik/etree"
	gxv "github.com/terminalstatic/go-xsd-validate"
	xsdvalidate "github.com/terminalstatic/go-xsd-validate"
)

// nolint: unused
func newInv(doc *etree.Document) *etree.Element {
	return doc.CreateElement("Invoice").CreateAttr("xmlns", urnInvoice2).Element().
		CreateAttr("xmlns:cac", urnCommonAggregateComponents2).Element().
		CreateAttr("xmlns:cbc", urnCommonBasicComponents2).Element()
}

// nolint: unused
func invTop(inv *etree.Element, id, iD, iT, lvid, iTC, dCC, tCC string) {
	inv.CreateElement(cbcID).SetText(id)
	inv.CreateElement(cbcIssueDate).SetText(iD)
	inv.CreateElement(cbcIssueTime).SetText(iT)
	inv.CreateElement(cbcInvoiceTypeCode).CreateAttr(lVID, lvid).Element().SetText(iTC)
	inv.CreateElement(cbcDocumentCurrencyCode).SetText(dCC)
	inv.CreateElement(cbcTaxCurrencyCode).SetText(tCC)
}

// nolint: unused
func invPer(inv *etree.Element, sD, eD, d string) {
	p := inv.CreateElement(cacInvoicePeriod)
	p.CreateElement(cbcStartDate).SetText(sD)
	p.CreateElement(cbcEndDate).SetText(eD)
	p.CreateElement(cbcDescription).SetText(d)
}

// nolint: unused
func invBilRefAddDocRef(inv *etree.Element, id string) {
	inv.CreateElement(cacBillingReference).CreateElement(cacAdditionalDocumentReference).
		CreateElement(cbcID).SetText(id)
}

// nolint: unused
func invAddDocRef(inv *etree.Element, id, dT, dD string) {
	l1 := inv.CreateElement(cacAdditionalDocumentReference)
	l1.CreateElement(cbcID).SetText(id)
	if dT != NUL {
		l1.CreateElement(cbcDocumentType).SetText(dT)
	}
	if dD != NUL {
		l1.CreateElement(cbcDocumentDescription).SetText(dD)
	}
}

// nolint: unused
func invAccSupPar(inv *etree.Element, aaid *valAttrs) *etree.Element {
	asp := inv.CreateElement(cacAccountingSupplierParty)
	if aaid != nil {
		_aaid := asp.CreateElement(cbcAdditionalAccountID)
		_aaid.SetText(aaid.Val)
		for _, duo := range aaid.Attrs {
			_aaid.CreateAttr(duo[0], duo[1])
		}
	}
	return asp
}

// nolint: unused
func invSupPar(asp *etree.Element, iCC string, pI []valAttr) *etree.Element {
	sp := asp.CreateElement(cacParty)
	sp.CreateElement(cbcIndustryClassificationCode).
		CreateAttr(nAME, cMSICSC[iCC].Description).Element().SetText(iCC)
	for _, pi := range pI {
		sp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
			CreateAttr(pi.Attr[0], pi.Attr[1]).Element().SetText(pi.Val)
	}
	return sp
}

// nolint: unused
func invXxxParPosAdd(xp *etree.Element, cN, pZ, cSC string, aL []string, cIC string) {
	pa := xp.CreateElement(cacPostalAddress)
	pa.CreateElement(cbcCityName).SetText(cN)
	pa.CreateElement(cbcPostalZone).SetText(pZ)
	pa.CreateElement(cbcCountrySubentityCode).SetText(cStateRev[cSC].Code)
	for _, line := range aL {
		pa.CreateElement(cacAddressLine).CreateElement(cbcLine).SetText(line)
	}
	pa.CreateElement(cacCountry).CreateElement(cbcIdentificationCode).
		CreateAttr(lID, "ISO3166-1").Element().CreateAttr(lAID, "6").
		Element().SetText(cCountryRev[cIC].Code)
}

// nolint: unused
func invParLegEntRegNam(p *etree.Element, n string) {
	p.CreateElement(cacPartyLegalEntity).CreateElement(cbcRegistrationName).SetText(n)
}

// nolint: unused
func invCon(p *etree.Element, t, eM string) {
	c := p.CreateElement(cacContact)
	if t != NUL {
		c.CreateElement(cbcTelephone).SetText(t)
	}
	if eM != NUL {
		c.CreateElement(cbcElectronicMail).SetText(eM)
	}
}

// nolint: unused
func invAccCusPar(inv *etree.Element, aaid *valAttrs) *etree.Element {
	asp := inv.CreateElement(cacAccountingCustomerParty)
	if aaid != nil {
		_aaid := asp.CreateElement(cbcAdditionalAccountID)
		_aaid.SetText(aaid.Val)
		for _, duo := range aaid.Attrs {
			_aaid.CreateAttr(duo[0], duo[1])
		}
	}
	return asp
}

// nolint: unused
func invCusPar(acp *etree.Element, pI []valAttr) *etree.Element {
	cp := acp.CreateElement(cacParty)
	for _, pi := range pI {
		cp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
			CreateAttr(pi.Attr[0], pi.Attr[1]).Element().SetText(pi.Val)
	}
	return cp
}

// nolint: unused
func invDel(inv *etree.Element) *etree.Element {
	return inv.CreateElement(cacDelivery)
}

// nolint: unused
func invDelPar(d *etree.Element, pI []valAttr) *etree.Element {
	dp := d.CreateElement(cacDeliveryParty)
	for _, pi := range pI {
		dp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
			CreateAttr(pi.Attr[0], pi.Attr[1]).Element().SetText(pi.Val)
	}
	return dp
}

// nolint: unused
func invDelShi(d *etree.Element, id string) *etree.Element {
	ds := d.CreateElement(cacShipment)
	ds.CreateElement(cbcID).SetText(id)
	return ds
}

// nolint: unused
func invDelShiFreAllCha(ds *etree.Element, cI, aCR, c, a string) {
	fac := ds.CreateElement(cacFreightAllowanceCharge)
	fac.CreateElement(cbcChargeIndicator).SetText(cI)
	fac.CreateElement(cbcAllowanceChargeReason).SetText(aCR)
	fac.CreateElement(cbcAmount).CreateAttr(cID, c).Element().SetText(a)
}

// nolint: unused
func invPayMea(inv *etree.Element, pMC, pFA string) {
	pm := inv.CreateElement(cacPaymentMeans)
	pm.CreateElement(cbcPaymentMeansCode).SetText(pMC) // PaymentMethods.json
	if pFA != NUL {
		pm.CreateElement(cacPayeeFinancialAccount).CreateElement(cbcID).SetText(pFA)
	}
}

// nolint: unused
func invPayTerNot(inv *etree.Element, v string) {
	inv.CreateElement(cacPaymentTerms).CreateElement(cbcNote).SetText(v)
}

// nolint: unused
func invPrePay(inv *etree.Element, id, c, pA, pD, pT string) {
	pp := inv.CreateElement(cacPrepaidPayment)
	pp.CreateElement(cbcID).SetText(id)
	pp.CreateElement(cbcPaidAmount).CreateAttr(cID, c).Element().SetText(pA)
	pp.CreateElement(cbcPaidDate).SetText(pD)
	pp.CreateElement(cbcPaidTime).SetText(pT)
}

// nolint: unused
func invAllCha(inv *etree.Element, cI, aCR, c, a string) {
	ac := inv.CreateElement(cacAllowanceCharge)
	ac.CreateElement(cbcChargeIndicator).SetText(cI)
	ac.CreateElement(cbcAllowanceChargeReason).SetText(aCR)
	ac.CreateElement(cbcAmount).CreateAttr(cID, c).Element().SetText(a)
}

// nolint: unused
func invTaxTot(inv *etree.Element, c, a string) *etree.Element {
	tt := inv.CreateElement(cacTaxTotal)
	tt.CreateElement(cbcTaxAmount).CreateAttr(cID, c).Element().SetText(a)
	return tt
}

// nolint: unused
func invTaxTotTaxSub(tt *etree.Element, cTAA, tAA, cTA, tA, tCID, tS string) {
	ts := tt.CreateElement(cacTaxSubtotal)
	ts.CreateElement(cbcTaxableAmount).CreateAttr(cID, cTAA).Element().SetText(tAA)
	ts.CreateElement(cbcTaxAmount).CreateAttr(cID, cTA).Element().SetText(tA)
	tstc := ts.CreateElement(cacTaxCategory)
	tstc.CreateElement(cbcID).SetText(tCID)
	tstc.CreateElement(cacTaxScheme).CreateElement(cbcID).
		CreateAttr(sID, UE5).Element().CreateAttr(sAID, "6").Element().SetText(tS)
}

// nolint: unused
func invLegMonTot(inv *etree.Element, c, lEA, tEA, tIA, aTA, cTA, pRA, pA string) {
	lmt := inv.CreateElement(cacLegalMonetaryTotal)
	if lEA != NUL {
		lmt.CreateElement(cbcLineExtensionAmount).CreateAttr(cID, c).Element().SetText(lEA)
	}
	if tEA != NUL {
		lmt.CreateElement(cbcTaxExclusiveAmount).CreateAttr(cID, c).Element().SetText(tEA)
	}
	if tIA != NUL {
		lmt.CreateElement(cbcTaxInclusiveAmount).CreateAttr(cID, c).Element().SetText(tIA)
	}
	if aTA != NUL {
		lmt.CreateElement(cbcAllowanceTotalAmount).CreateAttr(cID, c).Element().SetText(aTA)
	}
	if cTA != NUL {
		lmt.CreateElement(cbcChargeTotalAmount).CreateAttr(cID, c).Element().SetText(cTA)
	}
	if pRA != NUL {
		lmt.CreateElement(cbcPayableRoundingAmount).CreateAttr(cID, c).Element().SetText(pRA)
	}
	if pA != NUL {
		lmt.CreateElement(cbcPayableAmount).CreateAttr(cID, c).Element().SetText(pA)
	}
}

// nolint: unused
func invInvLin(inv *etree.Element, id, c, a string, iQ *valAttr) *etree.Element {
	il := inv.CreateElement(cacInvoiceLine)
	il.CreateElement(cbcID).SetText(id)
	il.CreateElement(cbcInvoicedQuantity).SetText(iQ.Val)
	if iQ.Attr[0] != NUL {
		il.FindElement(cbcInvoicedQuantity).CreateAttr(iQ.Attr[0], iQ.Attr[1]) // what's C62?
	}
	il.CreateElement(cbcLineExtensionAmount).CreateAttr(cID, c).Element().SetText(a)
	return il
}

// nolint: unused
func invInvLinAllCha(il *etree.Element, cI, aCR, mFN, c, a string) {
	ac := il.CreateElement(cacAllowanceCharge)
	ac.CreateElement(cbcChargeIndicator).SetText(cI)
	ac.CreateElement(cbcAllowanceChargeReason).SetText(aCR)
	ac.CreateElement(cbcMultiplierFactorNumeric).SetText(mFN)
	ac.CreateElement(cbcAmount).CreateAttr(cID, c).Element().SetText(a)
}

// nolint: unused
func invInvLinTaxTot(il *etree.Element, c, tA string) *etree.Element {
	tt := il.CreateElement(cacTaxTotal)
	tt.CreateElement(cbcTaxAmount).CreateAttr(cID, c).Element().SetText(tA)
	return tt
}

// nolint: unused
func invInvLinTaxTotTaxSubPer(iltt *etree.Element, cTAA, tAA, cTA, tA, p, tCID, tER, tS string) {
	ts := iltt.CreateElement(cacTaxSubtotal)
	ts.CreateElement(cbcTaxableAmount).CreateAttr(cID, cTAA).Element().SetText(tAA)
	ts.CreateElement(cbcTaxAmount).CreateAttr(cID, cTA).Element().SetText(tA)
	if p != NUL {
		ts.CreateElement(cbcPercent).SetText(p)
	}
	tstc := ts.CreateElement(cacTaxCategory)
	tstc.CreateElement(cbcID).SetText(tCID)
	if tER != NUL {
		tstc.CreateElement(cbcTaxExemptionReason).SetText(tER)
	}
	tstc.CreateElement(cacTaxScheme).CreateElement(cbcID).
		CreateAttr(sID, UE5).Element().CreateAttr(sAID, "6").Element().SetText(tS)
}

// nolint: unused
func invInvLinTaxTotTaxSubUni(iltt *etree.Element, cTAA, tAA, cTA, tA string, bUM, pUM *valAttr, tCID, tER, tS string) {
	ts := iltt.CreateElement(cacTaxSubtotal)
	ts.CreateElement(cbcTaxableAmount).CreateAttr(cID, cTAA).Element().SetText(tAA)
	ts.CreateElement(cbcTaxAmount).CreateAttr(cID, cTA).Element().SetText(tA)
	ts.CreateElement(cbcBaseUnitMeasure).CreateAttr(bUM.Attr[0], bUM.Attr[1]).Element().SetText(bUM.Val)
	ts.CreateElement(cbcPerUnitAmount).CreateAttr(pUM.Attr[0], pUM.Attr[1]).Element().SetText(pUM.Val)
	tstc := ts.CreateElement(cacTaxCategory)
	tstc.CreateElement(cbcID).SetText(tCID)
	if tER != NUL {
		tstc.CreateElement(cbcTaxExemptionReason).SetText(tER)
	}
	tstc.CreateElement(cacTaxScheme).CreateElement(cbcID).
		CreateAttr(sID, UE5).Element().CreateAttr(sAID, "6").Element().SetText(tS)
}

// nolint: unused
func invInvLinIte(il *etree.Element, d, oCIC string, cCICCp, cCICCc *valAttr) {
	i := il.CreateElement(cacItem)
	i.CreateElement(cbcDescription).SetText(d)
	i.CreateElement(cacOriginCountry).CreateElement(cbcIdentificationCode).SetText(cCountryRev[oCIC].Code)
	if cCICCp != nil {
		i.CreateElement(cacCommodityClassification).CreateElement(cbcItemClassificationCode).
			CreateAttr(cCICCp.Attr[0], cCICCp.Attr[1]).Element().SetText(cCICCp.Val)
	}
	if cCICCc != nil {
		i.CreateElement(cacCommodityClassification).CreateElement(cbcItemClassificationCode).
			CreateAttr(cCICCc.Attr[0], cCICCc.Attr[1]).Element().SetText(cCICCc.Val)
	}
}

// nolint: unused
func invInvLinPriPriAmo(il *etree.Element, c, pPA string) {
	il.CreateElement(cacPrice).CreateElement(cbcPriceAmount).CreateAttr(cID, c).
		Element().SetText(pPA)
}

// nolint: unused
func invInvLinItePriExtAmo(il *etree.Element, c, iPEA string) {
	il.CreateElement(cacItemPriceExtension).CreateElement(cbcAmount).CreateAttr(cID, c).
		Element().SetText(iPEA)
}

func buildInvLines(inv *etree.Element, ils []InvLine) {
	for idx, il := range ils {
		idx_ := fmt.Sprintf("%d", idx+1)
		il_ := invInvLin(inv, idx_, il.Currency, il.TotalPrice, &valAttr{Val: il.Quantity, Attr: [2]string{uC, il.UnitCode}})

		iltt_ := invInvLinTaxTot(il_, il.Currency, il.TaxAmount)
		invInvLinTaxTotTaxSubPer(iltt_, il.Currency, il.TaxTaxable, il.Currency, il.TaxAmount, il.TaxPercentage, cTaxNotApplicable, "uSME, not eligible yet", OTH)

		invInvLinIte(il_, il.Description, il.OriginCountry, il.ProductTariff, il.Classification)
		invInvLinPriPriAmo(il_, il.Currency, il.UnitPrice)
		invInvLinItePriExtAmo(il_, il.Currency, il.TotalPrice)
	}
}

// nolint: unused
func invTaxExtRat(il *etree.Element, sCC, tCC, cR string) {
	ter := il.CreateElement(cacTaxExchangeRate)
	ter.CreateElement(cbcSourceCurrencyCode).SetText(sCC)
	ter.CreateElement(cbcTargetCurrencyCode).SetText(tCC)
	ter.CreateElement(cbcCalculationRate).SetText(cR)
}

// nolint: unused
func invBilRefInvDocRef(inv *etree.Element, id, uuid string) {
	idr := inv.CreateElement(cacBillingReference).CreateElement(cacInvoiceDocumentReference)
	idr.CreateElement(cbcID).SetText(id)
	idr.CreateElement(cbcUUID).SetText(uuid)
}

// nolint: unused
func invSav(doc *etree.Document, out string) {
	tx, err := os.OpenFile(out, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer tx.Close()
	doc.WriteSettings.CanonicalEndTags = true // <T/> -> <T></T> prod=false
	doc.WriteSettings.CanonicalText = true    // &apos; -> '
	doc.Indent(2)
	_, _ = doc.WriteTo(tx)
	// _, _ = doc.WriteTo(os.Stdout)
}

// var b bytes.Buffer
// if _, err := doc.WriteTo(&b); err != nil {
// 	return nil, fmt.Errorf("failed to write XML to buffer: %w", err)
// }
// return b.Bytes(), nil

func logFatalln(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}

func logPrintln(e error) {
	if e != nil {
		log.Println(e)
	}
}

func readToBytes(fp string) []byte {
	jf, err := os.Open(fp)
	logFatalln(err)
	defer jf.Close()
	ba, _ := io.ReadAll(jf)
	return ba
}

func unmarshalTo(jd []byte, v any) {
	if err := json.Unmarshal(jd, &v); err != nil {
		log.Fatalln("Error unmarshaling JSON:", err)
	}
}

// nolint: unused
func isXsdInvoice(b []byte) bool {
	err := hInvoice.ValidateMem(b, xsdvalidate.ValidErrDefault)
	if err != nil {
		switch e := err.(type) {
		case xsdvalidate.ValidationError:
			log.Println(err)
			log.Printf("Error in line: %d\n", e.Errors[0].Line)
			log.Println(e.Errors[0].Message)
		default:
			log.Println(err)
		}
		log.Println("failed to validate")
		return false
	}
	return true
}

// nolint: unused
func checks(fps []string) {
	for _, fp := range fps {
		fmt.Println(isXsdInvoice(readToBytes(fp)), fp)
	}
}

// nolint: unused
func studyXml() {
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
}

func studyCalls() {
	// callGetDocTypes(clients[0].ID)
	// callGetDocType(clients[0].ID, "8")
	// callGetDocTypeVer(clients[0].ID, "8", "16")
	// callGetNoticesDefault(clients[0].ID)
	// callGetNotices(clients[0].ID, etc)
	// callGetValTaxTin(clients[0].ID, k.String("oh.my"), "BRN", "202512345")
	// log.Println(D_BASE + D_EG + "1.0-Invoice-ForeignCurrency-Sample.xml")
	// rds, err := genDocSub(D_BASE + D_EG + "1.0-Invoice-ForeignCurrency-Sample.xml")

	// generate customer's invoice
	aid, bid := int64(8888), int64(8888)
	ba := genCusInvMth(aid, bid)
	// log.Println("isValidXSD:", isXsdInvoice(ba)) // , calcSha256(ba)
	if isXsdInvoice(ba) {
		xml := fmt.Sprintf("%s/%d-%d.xml", D_BASE+D_XML, aid, bid)
		rds, err := genDocSub(xml)
		logPrintln(err)

		duid := callPosDocSubmit(clients[0].ID, Docs{[]Doc{rds}})
		// log.Println("duid insert'd:", duid)
		if duid != "" {
			wg.Add(1)
			go delaySetStatus(&wg, time.Second*3, clients[0].ID, duid)
			wg.Wait()
		}

		qrc := scanMyInvoisQRCode()
		// log.Println("qrc:", qrc)
		jsn := callGetTaxQRInfo(clients[0].ID, qrc) // ttl 3 hours, hardcoded for now
		// log.Println("jsn:", jsn)
		key, _ := loadKey("keys.tpi")
		j, _ := decAESGCM(jsn, key)
		log.Println(string(j))
		// this is basically the Accounting Supplier/Customer Party data
	}

	// store response for QR and other logics
}

func getUUIDLastSegment(uuidStr string) string {
	parts := strings.Split(uuidStr, "-")
	if len(parts) == 5 {
		return parts[4]
	}
	return "" // Return empty string for invalid UUID
}

func calcSha256(data []byte) string {
	hasher := sha256.New()
	hasher.Write(data)
	hashBytes := hasher.Sum(nil)
	return hex.EncodeToString(hashBytes)
}

func loadKey(keyPath string) ([]byte, error) {
	keyBytes := k.Bytes(keyPath)
	if len(keyBytes) == 0 {
		return nil, fmt.Errorf("key not found or is empty at path: %s", keyPath)
	} else if len(keyBytes) != 32 {
		return nil, fmt.Errorf("key length should be 32 bytes (256 bits): %s", keyPath)
	}
	return keyBytes, nil
}

func encAESGCM(plaintext []byte, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", fmt.Errorf("could not create new cipher: %w", err)
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", fmt.Errorf("could not create GCM: %w", err)
	}

	// Never reuse a nonce (or IV). It's randomly generated here.
	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", fmt.Errorf("could not generate nonce: %w", err)
	}

	// Encrypt and authenticate the plaintext.
	// The nonce is prepended to the ciphertext for easy retrieval during decryption.
	ciphertext := aesGCM.Seal(nonce, nonce, plaintext, nil) // additionalData is nil here
	return base64.URLEncoding.EncodeToString(ciphertext), nil
}

func decAESGCM(encodedCiphertext string, key []byte) ([]byte, error) {
	ciphertext, err := base64.URLEncoding.DecodeString(encodedCiphertext)
	if err != nil {
		return nil, fmt.Errorf("could not decode base64 ciphertext: %w", err)
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("could not create new cipher: %w", err)
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("could not create GCM: %w", err)
	}

	nonceSize := aesGCM.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, fmt.Errorf("ciphertext too short")
	}

	nonce, encryptedMessage := ciphertext[:nonceSize], ciphertext[nonceSize:]

	// Decrypt and authenticate. If the data or tag is tampered with, this will return an error.
	plaintext, err := aesGCM.Open(nil, nonce, encryptedMessage, nil) // additionalData is nil here
	if err != nil {
		return nil, fmt.Errorf("could not decrypt/authenticate: %w", err)
	}

	return plaintext, nil
}

func genSPI() []valAttr {
	var PIs []PartyId
	if err := k.Unmarshal("ubl.pi", &PIs); err != nil {
		log.Fatalf("error unmarshalling 'pi': %v", err)
	}
	var spi []valAttr
	for _, pi := range PIs {
		spi = append(spi, valAttr{Val: pi.Val, Attr: [2]string{sID, pi.Sid}})
	}
	return spi
}

func genApiUrls() {
	// Platform
	apiV := fmt.Sprintf("%s/api/v%s", urlApi, _10)
	posLogin = fmt.Sprintf("%s%s", urlApi, "/connect/token")
	getDocTypes = apiV + "/documenttypes"
	getDocType = apiV + "/documenttypes/{id}"
	getDocTypeVer = apiV + "/documenttypes/{id}/versions/{vid}"
	getNoticesDefault = apiV + "/notifications/taxpayer"
	getNotices = apiV + "/notifications/taxpayer?dateFrom={dateFrom}&dateTo={dateTo}&type={type}&language={language}&status={status}&pageNo={pageNo}&pageSize={pageSize}"
	// EInvoicing
	getTaxValTin = apiV + "/taxpayer/validate/{tin}?idType={idType}&idValue={idValue}"
	posDocSubmit = apiV + "/documentsubmissions"
	putDocCancel = apiV + "/documents/state/{uuid}/state"
	putDocReject = apiV + "/documents/state/{uuid}/state"
	getDocDetail = apiV + "/documents/{uuid}/details"
	getTaxQRInfo = apiV + "/taxpayers/qrcodeinfo/{qrCodeText}"
}

func genDocSub(fp string) (Doc, error) {
	file, err := os.Open(fp)
	if err != nil {
		log.Println("genDocSub", fp)
		return Doc{}, err
	}
	defer file.Close()

	hasher := sha256.New() // not 'like', but 'must' for documentHash
	var contentBuf bytes.Buffer
	reader := io.TeeReader(file, hasher)
	if _, err = io.Copy(&contentBuf, reader); err != nil {
		return Doc{}, err
	}
	basename := filepath.Base(fp)
	basename = strings.TrimSuffix(basename, filepath.Ext(basename)) // rem. ext.
	return Doc{
		Format:       XML,
		DocumentHash: fmt.Sprintf("%x", hasher.Sum(nil)),
		Document:     base64.StdEncoding.EncodeToString(contentBuf.Bytes()),
		CodeNumber:   basename, // acc#-inv#
	}, nil
}

func delaySetStatus(wg *sync.WaitGroup, delay time.Duration, uuid, duid string) {
	defer wg.Done()
	// log.Printf("waiting %s for %s's status\n", delay, duid)
	time.Sleep(delay)
	callGetDocDetailThenUpdate(uuid, duid) // maybe should handle error, don't care for now
}

func genCusInvMth(aid, bid int64) []byte {
	fn := fmt.Sprintf("%d%s%d", aid, invSep, bid)
	t := time.Now().UTC()

	doc := etree.NewDocument()
	inv := newInv(doc)
	// t = t.Truncate(time.Second * 5) // reduce resolution so that myinvois system can help capture duplicates?
	invTop(inv, fn, t.Format("2006-01-02"), t.Format("15:04:05Z"), _10, iTCInvoice, MYR, MYR)
	invPer(inv, "2025-05-26", "2025-06-25", "Monthly")

	asp := invAccSupPar(inv, nil)
	sp := invSupPar(asp, k.String("ubl.icc"), genSPI())
	invXxxParPosAdd(sp, k.String("ubl.pa.cn"), k.String("ubl.pa.pz"), k.String("ubl.pa.csc"), k.Strings("ubl.pa.al"), spCIC)
	invParLegEntRegNam(sp, k.String("ubl.rn"))
	invCon(sp, k.String("ubl.c.t"), k.String("ubl.c.em"))

	//🟠 acp from qr
	csp := invAccCusPar(inv, nil)
	cp := invCusPar(csp, genSPI())
	invXxxParPosAdd(cp, k.String("ubl.pa.cn"), k.String("ubl.pa.pz"), k.String("ubl.pa.csc"), k.Strings("ubl.pa.al"), spCIC)
	invParLegEntRegNam(cp, k.String("ubl.rn"))
	invCon(cp, k.String("ubl.c.t"), k.String("ubl.c.em"))

	invPayMea(inv, cPMOthers, NUL)
	invPayTerNot(inv, "Please refer the email'd invoice for JomPAY details.")
	tt := invTaxTot(inv, MYR, "0.00")                               // required nevertheless
	invTaxTotTaxSub(tt, MYR, "0", MYR, "0", cTaxNotApplicable, OTH) // uMSE for now

	total := "55.90" // 50 + (2.95x2)
	invLegMonTot(inv, MYR, total, total, total, NUL, NUL, NUL, total)

	clsSvc := &valAttr{Val: "022", Attr: [2]string{lID, CLASS}} // Others , as it's consider service fee
	// local only m'sian -> m'sian // imports(self-billed) //ignore export(n/a)

	il := []InvLine{
		{Quantity: "1", UnitCode: cUOne, UnitPrice: "50.00", TotalPrice: "50.00", OriginCountry: spCIC, ProductTariff: nil, Classification: clsSvc, Currency: MYR, Description: "VPS Rental", TaxType: cTaxNotApplicable, TaxAmount: "0", TaxTaxable: "0", TaxPercentage: "0.00"},
		{Quantity: "2", UnitCode: cUOne, UnitPrice: "2.95", TotalPrice: "5.90", OriginCountry: spCIC, ProductTariff: nil, Classification: clsSvc, Currency: MYR, Description: "Users/Month", TaxType: cTaxNotApplicable, TaxAmount: "0", TaxTaxable: "0", TaxPercentage: "0.00"},
	}
	buildInvLines(inv, il)

	doc.WriteSettings.CanonicalText = true // &apos; -> '
	if !k.Bool("ubl.pack") {
		doc.Indent(2)
	}
	xml := fmt.Sprintf("%s/%d-%d.xml", D_BASE+D_XML, aid, bid)
	if err := doc.WriteToFile(xml); err != nil {
		log.Println(xml, err)
	}
	ba, _ := doc.WriteToBytes()
	return ba
}

func codes() (map[string]CodeMSICSubCategory, map[string]CodeState, map[string]CodeState, map[string]CodeCountry) {

	var categories []CodeMSICSubCategory
	unmarshalTo(readToBytes(D_C+F_C_MSICSubCategory), &categories)
	for _, category := range categories {
		cMSICSC[category.Code] = category // ignore duplicated Code here
	}

	var states []CodeState
	unmarshalTo(readToBytes(D_C+F_C_State), &states)
	for _, state := range states {
		cState[state.Code] = state
	}

	var states_rev []CodeState
	unmarshalTo(readToBytes(D_C+F_C_State), &states_rev)
	for _, state := range states_rev {
		cStateRev[state.State] = state
	}

	var countries []CodeCountry
	unmarshalTo(readToBytes(D_C+F_C_Country), &countries)
	for _, country := range countries {
		cCountryRev[country.Country] = country
	}

	return cMSICSC, cState, cStateRev, cCountryRev
}

func freeCleanupClose() {
	hInvoice.Free()
	gxv.Cleanup()
	db.Close()
}
