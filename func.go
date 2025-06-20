package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/beevik/etree"
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
func invSupPar(asp *etree.Element, iCC, tin, brn, sst, ttx string) *etree.Element {
	sp := asp.CreateElement(cacParty)
	sp.CreateElement(cbcIndustryClassificationCode).
		CreateAttr(nAME, cMSICSC[iCC].Description).Element().SetText(iCC)
	sp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, TIN).Element().SetText(tin)
	sp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, BRN).Element().SetText(brn)
	if sst != NUL {
		sp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
			CreateAttr(sID, SST).Element().SetText(sst)
	}
	if ttx != NUL {
		sp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
			CreateAttr(sID, TTX).Element().SetText(ttx)
	}
	return sp
}

// nolint: unused
func invXxxParPosAdd(xp *etree.Element, pZ, cSC string, aL []string, cIC string) {
	pa := xp.CreateElement(cacPostalAddress)
	pa.CreateElement(cbcCityName).SetText(reSubWP.ReplaceAllString(cState[cSC].State, ""))
	pa.CreateElement(cbcPostalZone).SetText(pZ)
	pa.CreateElement(cbcCountrySubentityCode).SetText(cSC)
	for _, line := range aL {
		pa.CreateElement(cacAddressLine).CreateElement(cbcLine).SetText(line)
	}
	pa.CreateElement(cacCountry).CreateElement(cbcIdentificationCode).
		CreateAttr(lID, "ISO3166-1").Element().CreateAttr(lAID, "6").
		Element().SetText(cCountry[cIC].Code)
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
func invCusPar(acp *etree.Element, tin, brn, sst, ttx string) *etree.Element {
	cp := acp.CreateElement(cacParty)
	cp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, TIN).Element().SetText(tin)
	cp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, BRN).Element().SetText(brn)
	if sst != NUL {
		cp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
			CreateAttr(sID, SST).Element().SetText(sst)
	}
	if ttx != NUL {
		cp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
			CreateAttr(sID, TTX).Element().SetText(ttx)
	}
	return cp
}

// nolint: unused
func invDel(inv *etree.Element) *etree.Element {
	return inv.CreateElement(cacDelivery)
}

// nolint: unused
func invDelPar(d *etree.Element, tin, brn string) *etree.Element {
	dp := d.CreateElement(cacDeliveryParty)
	dp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, TIN).Element().SetText(tin)
	dp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, BRN).Element().SetText(brn)
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
	pm.CreateElement(cacPayeeFinancialAccount).CreateElement(cbcID).SetText(pFA)
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
	i.CreateElement(cacOriginCountry).CreateElement(cbcIdentificationCode).SetText(cCountry[oCIC].Code)
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

func readToBytes(fp string) ([]byte, error) {
	jf, err := os.Open(fp)
	if err != nil {
		log.Fatalln(err)
	}
	defer jf.Close()
	return io.ReadAll(jf)
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
		if b, e := readToBytes(fp); e == nil {
			fmt.Println(isXsdInvoice(b), fp)
		}
	}
}
