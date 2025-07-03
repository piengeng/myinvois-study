package main

import (
	"github.com/beevik/etree"
)

var (
	egSPI_1 = []valAttr{
		{Attr: [2]string{sID, TIN}, Val: "Supplier's TIN"},
		{Attr: [2]string{sID, BRN}, Val: "Supplier's BRN"},
		{Attr: [2]string{sID, SST}, Val: "NA"},
		{Attr: [2]string{sID, TTX}, Val: "NA"},
	}
	egCPI_1 = []valAttr{
		{Attr: [2]string{sID, TIN}, Val: "Buyer's TIN"},
		{Attr: [2]string{sID, BRN}, Val: "Buyer's BRN"},
		{Attr: [2]string{sID, SST}, Val: "NA"},
		{Attr: [2]string{sID, TTX}, Val: "NA"},
	}
	egDPI_1 = []valAttr{
		{Attr: [2]string{sID, TIN}, Val: "Recipient's TIN"},
		{Attr: [2]string{sID, BRN}, Val: "Recipient's BRN"},
	}
	egCPI_2 = []valAttr{
		{Attr: [2]string{sID, TIN}, Val: gpTIN}, // consolidated
		{Attr: [2]string{sID, BRN}, Val: "NA"},
		{Attr: [2]string{sID, SST}, Val: "NA"},
		{Attr: [2]string{sID, TTX}, Val: "NA"},
	}
	egAddrLines_1 = []string{"Lot 66", "Bangunan Merdeka", "Persiaran Jaya"}
)

// nolint:unused
func samplingInvoice10(out string) {
	doc := etree.NewDocument()
	inv := newInv(doc)
	invTop(inv, "XML-INV12345", "2024-07-23", "00:40:00Z", _10, iTCInvoice, MYR, MYR)
	invPer(inv, "2024-07-01", "2024-07-31", "Monthly")
	invBilRefAddDocRef(inv, "151891-1981")
	invAddDocRef(inv, "L1", "CustomsImportForm", NUL)
	invAddDocRef(inv, "FTA", "FreeTradeAgreement", "Sample Description")
	invAddDocRef(inv, "L1", "K2", NUL)
	invAddDocRef(inv, "L1", NUL, NUL)

	asp := invAccSupPar(inv, &valAttrs{Val: "CPT-CCN-W-211111-KL-000002", Attrs: [][2]string{{sAN, "CertEX"}}})

	sp := invSupPar(asp, spMSICSC, egSPI_1)
	invXxxParPosAdd(sp, reSubWP.ReplaceAllString(cState[spCSC].State, ""), "50480", spCSC, egAddrLines_1, spCIC)
	invParLegEntRegNam(sp, "Supplier's Name")
	invCon(sp, "+60123456789", "supplier@email.com")

	acp := invAccCusPar(inv, nil)
	cp := invCusPar(acp, egCPI_1)
	invXxxParPosAdd(cp, reSubWP.ReplaceAllString(cState[spCSC].State, ""), "50480", spCSC, egAddrLines_1, "MALAYSIA")
	invParLegEntRegNam(cp, "Buyer's Name")
	invCon(cp, "+60123456780", "buyer@email.com")

	d := invDel(inv)
	dp := invDelPar(d, egDPI_1)
	invXxxParPosAdd(dp, reSubWP.ReplaceAllString(cState[spCSC].State, ""), "50480", spCSC, egAddrLines_1, "MALAYSIA")
	invParLegEntRegNam(dp, "Recipient's Name")
	ds := invDelShi(d, "1234")
	invDelShiFreAllCha(ds, "true", "Service charge", MYR, "100")

	invPayMea(inv, "01", "1234567890")
	invPayTerNot(inv, "Payment method is cash")
	invPrePay(inv, "E12345678912", MYR, "1.00", "2024-07-23", "00:30:00Z")

	invAllCha(inv, "false", "Sample Description", MYR, "100")
	invAllCha(inv, "true", "Service charge", MYR, "100")

	tt := invTaxTot(inv, MYR, "87.63")
	invTaxTotTaxSub(tt, MYR, "87.63", MYR, "87.63", "01", OTH)
	invLegMonTot(inv, MYR, "1436.50", "1436.50", "1436.50", "1436.50", "1436.50", "0.30", "1436.50")

	il := invInvLin(inv, "1234", MYR, "1436.50", &valAttr{Val: "1", Attr: [2]string{uC, "C62"}})
	invInvLinAllCha(il, "false", "Sample Description", "0.15", MYR, "100")
	invInvLinAllCha(il, "true", "Sample Description", "0.10", MYR, "100")
	iltt := invInvLinTaxTot(il, MYR, "0")
	invInvLinTaxTotTaxSubPer(iltt, MYR, "1460.50", MYR, "0", "6.00", "E", "Exempt New Means of Transport", OTH)
	ptc := &valAttr{Val: "9800.00.0010", Attr: [2]string{lID, "PTC"}}
	cls := &valAttr{Val: "003", Attr: [2]string{lID, "CLASS"}}
	invInvLinIte(il, "Laptop Peripherals", "MALAYSIA", ptc, cls)
	invInvLinPriPriAmo(il, MYR, "17")
	invInvLinItePriExtAmo(il, MYR, "100")

	invSav(doc, out)
}

// nolint:unused
func samplingInvoiceMultiLine10(out string) {
	doc := etree.NewDocument()
	inv := newInv(doc)
	invTop(inv, "XML-INV12345", "2025-02-06", "00:30:00Z", _10, iTCInvoice, MYR, MYR)
	invPer(inv, "2025-01-01", "2025-01-31", "Monthly")
	invBilRefAddDocRef(inv, "E12345678912")
	invAddDocRef(inv, "E23456789123,E98765432123", "CustomsImportForm", NUL)
	invAddDocRef(inv, "FTA", "FreeTradeAgreement", "ASEAN-Australia-New Zealand FTA (AANZFTA)")
	invAddDocRef(inv, "E12345678912,E23456789123", "K2", NUL)
	invAddDocRef(inv, "CIF", NUL, NUL)

	asp := invAccSupPar(inv, &valAttrs{Val: "CPT-CCN-W-211111-KL-000002", Attrs: [][2]string{{sAN, "CertEX"}}})
	sp := invSupPar(asp, spMSICSC, egSPI_1)
	invXxxParPosAdd(sp, reSubWP.ReplaceAllString(cState[spCSC].State, ""), "50480", spCSC, egAddrLines_1, spCIC)
	invParLegEntRegNam(sp, "Supplier's Name")
	invCon(sp, "+60123456789", "supplier@email.com")

	acp := invAccCusPar(inv, nil)
	cp := invCusPar(acp, egCPI_1)
	invXxxParPosAdd(cp, reSubWP.ReplaceAllString(cState[spCSC].State, ""), "50480", spCSC, egAddrLines_1, "MALAYSIA")
	invParLegEntRegNam(cp, "Buyer's Name")
	invCon(cp, "+60123456780", "buyer@email.com")

	invAllCha(inv, "false", "", MYR, "0.00")
	invAllCha(inv, "true", "", MYR, "0.00")

	tt := invTaxTot(inv, MYR, "85.00")
	invTaxTotTaxSub(tt, MYR, "1000.00", MYR, "10.00", "01", OTH)
	invTaxTotTaxSub(tt, MYR, "1500.00", MYR, "75.00", "02", OTH)
	invTaxTotTaxSub(tt, MYR, "2000.00", MYR, "100.00", "E", OTH)

	invLegMonTot(inv, MYR, "4500.00", "4500.00", "4585.00", "0.00", "0.00", "0.00", "4585.00")

	il1 := invInvLin(inv, "001", MYR, "1000.00", &valAttr{Val: "1", Attr: [2]string{uC, "C62"}})
	invInvLinAllCha(il1, "false", "", "0", MYR, "0.00")
	invInvLinAllCha(il1, "true", "", "0", MYR, "0.00")
	il1tt := invInvLinTaxTot(il1, MYR, "10.00")
	bum1 := &valAttr{Val: "1", Attr: [2]string{uC, "C62"}}
	pum1 := &valAttr{Val: "10.00", Attr: [2]string{cID, MYR}}
	invInvLinTaxTotTaxSubUni(il1tt, MYR, "1000.00", MYR, "10.00", bum1, pum1, "01", NUL, OTH)
	ptc1 := &valAttr{Val: "9800.00.0010", Attr: [2]string{lID, "PTC"}}
	cls1 := &valAttr{Val: "003", Attr: [2]string{lID, "CLASS"}}
	invInvLinIte(il1, "Laptop Peripherals", spCIC, ptc1, cls1)
	invInvLinPriPriAmo(il1, MYR, "1000.00")
	invInvLinItePriExtAmo(il1, MYR, "1000.00")

	il2 := invInvLin(inv, "002", MYR, "1500.00", &valAttr{Val: "1", Attr: [2]string{uC, "C62"}})
	invInvLinAllCha(il2, "false", "", "0", MYR, "0.00")
	invInvLinAllCha(il2, "true", "", "0", MYR, "0.00")
	il2tt := invInvLinTaxTot(il2, MYR, "75.00")
	invInvLinTaxTotTaxSubPer(il2tt, MYR, "1500.00", MYR, "75.00", "5.00", "02", NUL, OTH)
	ptc2 := &valAttr{Val: "9800.00.0011", Attr: [2]string{lID, "PTC"}}
	cls2 := &valAttr{Val: "003", Attr: [2]string{lID, "CLASS"}}
	invInvLinIte(il2, "Computer Monitor", spCIC, ptc2, cls2)
	invInvLinPriPriAmo(il2, MYR, "1500.00")
	invInvLinItePriExtAmo(il2, MYR, "1500.00")

	il3 := invInvLin(inv, "003", MYR, "2000.00", &valAttr{Val: "1", Attr: [2]string{uC, "C62"}})
	invInvLinAllCha(il3, "false", "", "0", MYR, "0.00")
	invInvLinAllCha(il3, "true", "", "0", MYR, "0.00")
	il3tt := invInvLinTaxTot(il3, MYR, "0.00")
	invInvLinTaxTotTaxSubPer(il3tt, MYR, "0.00", MYR, "0.00", "5.00", "01", NUL, OTH)
	invInvLinTaxTotTaxSubPer(il3tt, MYR, "2000.00", MYR, "100.00", NUL, "E", "Special Case", OTH)
	ptc3 := &valAttr{Val: "9800.00.0012", Attr: [2]string{lID, "PTC"}}
	cls3 := &valAttr{Val: "003", Attr: [2]string{lID, "CLASS"}}
	invInvLinIte(il3, "Wireless Mouse", spCIC, ptc3, cls3)
	invInvLinPriPriAmo(il3, MYR, "2000.00")
	invInvLinItePriExtAmo(il3, MYR, "2000.00")

	invSav(doc, out)
}

// nolint:unused
func samplingInvoiceConsolidated10(out string) {
	doc := etree.NewDocument()
	inv := newInv(doc)
	invTop(inv, "XML-INV12345", "2024-07-23", "00:40:00Z", _10, iTCInvoice, MYR, MYR)

	asp := invAccSupPar(inv, nil)
	sp := invSupPar(asp, spMSICSC, egSPI_1)
	invXxxParPosAdd(sp, reSubWP.ReplaceAllString(cState[spCSC].State, ""), "50480", spCSC, egAddrLines_1, spCIC)
	invParLegEntRegNam(sp, "Supplier's Name")
	invCon(sp, "+60123456789", "supplier@email.com")

	acp := invAccCusPar(inv, nil)
	cp := invCusPar(acp, egCPI_2)
	invXxxParPosAdd(cp, "", "", "", []string{"NA", "", ""}, "")
	invParLegEntRegNam(cp, "Consolidated Buyers")
	invCon(cp, "NA", "NA")

	tt := invTaxTot(inv, MYR, "3000")
	invTaxTotTaxSub(tt, MYR, "30000", MYR, "3000", "01", OTH)

	invLegMonTot(inv, MYR, "30000", "30000", "33000", "0", "0", "0", "33000")

	il1 := invInvLin(inv, "1", MYR, "10000", &valAttr{Val: "1", Attr: [2]string{uC, "C62"}})
	il1tt := invInvLinTaxTot(il1, MYR, "1000")
	invInvLinTaxTotTaxSubPer(il1tt, MYR, "10000", MYR, "1000", "10.00", "01", NUL, OTH)
	ptc1 := &valAttr{Val: "", Attr: [2]string{lID, "PTC"}}
	cls1 := &valAttr{Val: "004", Attr: [2]string{lID, "CLASS"}}
	invInvLinIte(il1, "Receipt 001 - 100", spCIC, ptc1, cls1)
	invInvLinPriPriAmo(il1, MYR, "10000")
	invInvLinItePriExtAmo(il1, MYR, "10000")

	il2 := invInvLin(inv, "2", MYR, "20000", &valAttr{Val: "1", Attr: [2]string{uC, "C62"}})
	il2tt := invInvLinTaxTot(il2, MYR, "2000")
	invInvLinTaxTotTaxSubPer(il2tt, MYR, "20000", MYR, "2000", "10.00", "01", NUL, OTH)
	ptc2 := &valAttr{Val: "", Attr: [2]string{lID, "PTC"}}
	cls2 := &valAttr{Val: "004", Attr: [2]string{lID, "CLASS"}}
	invInvLinIte(il2, "Receipt 101 - 200", spCIC, ptc2, cls2)
	invInvLinPriPriAmo(il2, MYR, "20000")
	invInvLinItePriExtAmo(il2, MYR, "20000")

	invSav(doc, out)
}

// nolint:unused
func samplingInvoiceForeignCurrency10(out string) {
	doc := etree.NewDocument()
	inv := newInv(doc)
	invTop(inv, "XML-INV12345", "2024-07-23", "00:40:00Z", _10, iTCInvoice, USD, MYR)
	invPer(inv, "2024-07-01", "2024-07-31", "Monthly")
	invBilRefAddDocRef(inv, "151891-1981")
	invAddDocRef(inv, "L1", "CustomsImportForm", NUL)
	invAddDocRef(inv, "FTA", "FreeTradeAgreement", "Sample Description")
	invAddDocRef(inv, "L1", "K2", NUL)
	invAddDocRef(inv, "L1", NUL, NUL)

	asp := invAccSupPar(inv, &valAttrs{Val: "CPT-CCN-W-211111-KL-000002", Attrs: [][2]string{{sAN, "CertEX"}}})
	sp := invSupPar(asp, spMSICSC, egSPI_1)
	invXxxParPosAdd(sp, reSubWP.ReplaceAllString(cState[spCSC].State, ""), "50480", spCSC, egAddrLines_1, spCIC)
	invParLegEntRegNam(sp, "Supplier's Name")
	invCon(sp, "+60123456789", "supplier@email.com")

	acp := invAccCusPar(inv, nil)
	cp := invCusPar(acp, egCPI_1)
	invXxxParPosAdd(cp, reSubWP.ReplaceAllString(cState[spCSC].State, ""), "50480", spCSC, egAddrLines_1, "MALAYSIA")
	invParLegEntRegNam(cp, "Buyer's Name")
	invCon(cp, "+60123456780", "buyer@email.com")

	d := invDel(inv)
	dp := invDelPar(d, egDPI_1)
	invXxxParPosAdd(dp, reSubWP.ReplaceAllString(cState[spCSC].State, ""), "50480", spCSC, egAddrLines_1, "MALAYSIA")
	invParLegEntRegNam(dp, "Recipient's Name")
	ds := invDelShi(d, "1234")
	invDelShiFreAllCha(ds, "true", "Service charge", USD, "100")

	invPayMea(inv, "01", "1234567890")
	invPayTerNot(inv, "Payment method is cash")

	invPrePay(inv, "E12345678912", USD, "1.00", "2024-07-23", "00:30:00Z")

	invAllCha(inv, "false", "Sample Description", USD, "100")
	invAllCha(inv, "true", "Service charge", USD, "100")

	invTaxExtRat(inv, USD, MYR, "4.72")
	tt := invTaxTot(inv, MYR, "87.63")
	invTaxTotTaxSub(tt, USD, "87.63", MYR, "87.63", "01", OTH)

	invLegMonTot(inv, USD, "1436.50", "1436.50", "1436.50", "1436.50", "1436.50", "0.30", "1436.50")

	il1 := invInvLin(inv, "1234", USD, "1436.50", &valAttr{Val: "1", Attr: [2]string{uC, "C62"}})
	invInvLinAllCha(il1, "false", "Sample Description", "0.15", USD, "100")
	invInvLinAllCha(il1, "true", "Sample Description", "0.10", USD, "100")
	il1tt := invInvLinTaxTot(il1, MYR, "0")
	invInvLinTaxTotTaxSubPer(il1tt, USD, "1460.50", MYR, "0", "6.00", "E", "Exempt New Means of Transport", OTH)
	ptc1 := &valAttr{Val: "9800.00.0010", Attr: [2]string{lID, "PTC"}}
	cls1 := &valAttr{Val: "003", Attr: [2]string{lID, "CLASS"}}
	invInvLinIte(il1, "Laptop Peripherals", spCIC, ptc1, cls1)
	invInvLinPriPriAmo(il1, USD, "17")
	invInvLinItePriExtAmo(il1, USD, "100")

	invSav(doc, out)
}

// nolint: unused
func samplingCreditNote10(out string) {
	doc := etree.NewDocument()
	inv := newInv(doc)
	invTop(inv, "XML-CN12345", "2024-07-23", "00:40:00Z", _10, iTCCreditNote, MYR, MYR)
	invPer(inv, "2024-07-01", "2024-07-31", "Monthly")
	invBilRefInvDocRef(inv, "XML-INV12345", "Reference Invoice UUID")
	invBilRefAddDocRef(inv, "151891-1981")

	invAddDocRef(inv, "L1", "CustomsImportForm", NUL)
	invAddDocRef(inv, "FTA", "FreeTradeAgreement", "Sample Description")
	invAddDocRef(inv, "L1", "K2", NUL)
	invAddDocRef(inv, "L1", NUL, NUL)

	asp := invAccSupPar(inv, &valAttrs{Val: "CPT-CCN-W-211111-KL-000002", Attrs: [][2]string{{sAN, "CertEX"}}})
	sp := invSupPar(asp, spMSICSC, egSPI_1)
	invXxxParPosAdd(sp, reSubWP.ReplaceAllString(cState[spCSC].State, ""), "50480", spCSC, egAddrLines_1, spCIC)
	invParLegEntRegNam(sp, "Supplier's Name")
	invCon(sp, "+60123456789", "supplier@email.com")

	acp := invAccCusPar(inv, nil)
	cp := invCusPar(acp, egCPI_1)
	invXxxParPosAdd(cp, reSubWP.ReplaceAllString(cState[spCSC].State, ""), "50480", spCSC, egAddrLines_1, "MALAYSIA")
	invParLegEntRegNam(cp, "Buyer's Name")
	invCon(cp, "+60123456780", "buyer@email.com")

	d := invDel(inv)
	dp := invDelPar(d, egDPI_1)
	invXxxParPosAdd(dp, reSubWP.ReplaceAllString(cState[spCSC].State, ""), "50480", spCSC, egAddrLines_1, "MALAYSIA")
	invParLegEntRegNam(dp, "Recipient's Name")
	ds := invDelShi(d, "1234")
	invDelShiFreAllCha(ds, "true", "Service charge", MYR, "100")

	invPayMea(inv, "01", "1234567890")
	invPayTerNot(inv, "Payment method is cash")
	invPrePay(inv, "E12345678912", MYR, "1.00", "2024-07-23", "00:30:00Z")

	invAllCha(inv, "false", "Sample Description", MYR, "100")
	invAllCha(inv, "true", "Service charge", MYR, "100")

	tt := invTaxTot(inv, MYR, "87.63")
	invTaxTotTaxSub(tt, MYR, "87.63", MYR, "87.63", "01", OTH)
	invLegMonTot(inv, MYR, "1436.50", "1436.50", "1436.50", "1436.50", "1436.50", "0.30", "1436.50")

	il := invInvLin(inv, "1234", MYR, "1436.50", &valAttr{Val: "1", Attr: [2]string{uC, "C62"}})
	invInvLinAllCha(il, "false", "Sample Description", "0.15", MYR, "100")
	invInvLinAllCha(il, "true", "Sample Description", "0.10", MYR, "100")
	iltt := invInvLinTaxTot(il, MYR, "0")
	invInvLinTaxTotTaxSubPer(iltt, MYR, "1460.50", MYR, "0", "6.00", "E", "Exempt New Means of Transport", OTH)
	ptc := &valAttr{Val: "9800.00.0010", Attr: [2]string{lID, "PTC"}}
	cls := &valAttr{Val: "003", Attr: [2]string{lID, "CLASS"}}
	invInvLinIte(il, "Laptop Peripherals", "MALAYSIA", ptc, cls)
	invInvLinPriPriAmo(il, MYR, "17")
	invInvLinItePriExtAmo(il, MYR, "100")

	invSav(doc, out)
}

// nolint: unused
func samplingDebitNote10(out string) {
	doc := etree.NewDocument()
	inv := newInv(doc)
	invTop(inv, "XML-DN12345", "2024-07-23", "00:40:00Z", _10, iTCDebitNote, MYR, MYR)
	invPer(inv, "2024-07-01", "2024-07-31", "Monthly")
	invBilRefInvDocRef(inv, "XML-INV12345", "Reference Invoice UUID")
	invBilRefAddDocRef(inv, "151891-1981")

	invAddDocRef(inv, "L1", "CustomsImportForm", NUL)
	invAddDocRef(inv, "FTA", "FreeTradeAgreement", "Sample Description")
	invAddDocRef(inv, "L1", "K2", NUL)
	invAddDocRef(inv, "L1", NUL, NUL)

	asp := invAccSupPar(inv, &valAttrs{Val: "CPT-CCN-W-211111-KL-000002", Attrs: [][2]string{{sAN, "CertEX"}}})
	sp := invSupPar(asp, spMSICSC, egSPI_1)
	invXxxParPosAdd(sp, reSubWP.ReplaceAllString(cState[spCSC].State, ""), "50480", spCSC, egAddrLines_1, spCIC)
	invParLegEntRegNam(sp, "Supplier's Name")
	invCon(sp, "+60123456789", "supplier@email.com")

	acp := invAccCusPar(inv, nil)
	cp := invCusPar(acp, egCPI_1)
	invXxxParPosAdd(cp, reSubWP.ReplaceAllString(cState[spCSC].State, ""), "50480", spCSC, egAddrLines_1, "MALAYSIA")
	invParLegEntRegNam(cp, "Buyer's Name")
	invCon(cp, "+60123456780", "buyer@email.com")

	d := invDel(inv)
	dp := invDelPar(d, egDPI_1)
	invXxxParPosAdd(dp, reSubWP.ReplaceAllString(cState[spCSC].State, ""), "50480", spCSC, egAddrLines_1, "MALAYSIA")
	invParLegEntRegNam(dp, "Recipient's Name")
	ds := invDelShi(d, "1234")
	invDelShiFreAllCha(ds, "true", "Service charge", MYR, "100")

	invPayMea(inv, "01", "1234567890")
	invPayTerNot(inv, "Payment method is cash")
	invPrePay(inv, "E12345678912", MYR, "1.00", "2024-07-23", "00:30:00Z")

	invAllCha(inv, "false", "Sample Description", MYR, "100")
	invAllCha(inv, "true", "Service charge", MYR, "100")

	tt := invTaxTot(inv, MYR, "87.63")
	invTaxTotTaxSub(tt, MYR, "87.63", MYR, "87.63", "01", OTH)
	invLegMonTot(inv, MYR, "1436.50", "1436.50", "1436.50", "1436.50", "1436.50", "0.30", "1436.50")

	il := invInvLin(inv, "1234", MYR, "1436.50", &valAttr{Val: "1", Attr: [2]string{uC, "C62"}})
	invInvLinAllCha(il, "false", "Sample Description", "0.15", MYR, "100")
	invInvLinAllCha(il, "true", "Sample Description", "0.10", MYR, "100")
	iltt := invInvLinTaxTot(il, MYR, "0")
	invInvLinTaxTotTaxSubPer(iltt, MYR, "1460.50", MYR, "0", "6.00", "E", "Exempt New Means of Transport", OTH)
	ptc := &valAttr{Val: "9800.00.0010", Attr: [2]string{lID, "PTC"}}
	cls := &valAttr{Val: "003", Attr: [2]string{lID, "CLASS"}}
	invInvLinIte(il, "Laptop Peripherals", "MALAYSIA", ptc, cls)
	invInvLinPriPriAmo(il, MYR, "17")
	invInvLinItePriExtAmo(il, MYR, "100")

	invSav(doc, out)
}

// nolint: unused
func samplingRefundNote10(out string) {
	doc := etree.NewDocument()
	inv := newInv(doc)
	invTop(inv, "XML-RN12345", "2024-07-23", "00:40:00Z", _10, iTCRefundNote, MYR, MYR)
	invPer(inv, "2024-07-01", "2024-07-31", "Monthly")
	invBilRefInvDocRef(inv, "XML-INV12345", "Reference Invoice UUID")
	invBilRefAddDocRef(inv, "151891-1981")

	invAddDocRef(inv, "L1", "CustomsImportForm", NUL)
	invAddDocRef(inv, "FTA", "FreeTradeAgreement", "Sample Description")
	invAddDocRef(inv, "L1", "K2", NUL)
	invAddDocRef(inv, "L1", NUL, NUL)

	asp := invAccSupPar(inv, &valAttrs{Val: "CPT-CCN-W-211111-KL-000002", Attrs: [][2]string{{sAN, "CertEX"}}})
	sp := invSupPar(asp, spMSICSC, egSPI_1)
	invXxxParPosAdd(sp, reSubWP.ReplaceAllString(cState[spCSC].State, ""), "50480", spCSC, egAddrLines_1, spCIC)
	invParLegEntRegNam(sp, "Supplier's Name")
	invCon(sp, "+60123456789", "supplier@email.com")

	acp := invAccCusPar(inv, nil)
	cp := invCusPar(acp, egCPI_1)
	invXxxParPosAdd(cp, reSubWP.ReplaceAllString(cState[spCSC].State, ""), "50480", spCSC, egAddrLines_1, "MALAYSIA")
	invParLegEntRegNam(cp, "Buyer's Name")
	invCon(cp, "+60123456780", "buyer@email.com")

	d := invDel(inv)
	dp := invDelPar(d, egDPI_1)
	invXxxParPosAdd(dp, reSubWP.ReplaceAllString(cState[spCSC].State, ""), "50480", spCSC, egAddrLines_1, "MALAYSIA")
	invParLegEntRegNam(dp, "Recipient's Name")
	ds := invDelShi(d, "1234")
	invDelShiFreAllCha(ds, "true", "Service charge", MYR, "100")

	invPayMea(inv, "01", "1234567890")
	invPayTerNot(inv, "Payment method is cash")
	invPrePay(inv, "E12345678912", MYR, "1.00", "2024-07-23", "00:30:00Z")

	invAllCha(inv, "false", "Sample Description", MYR, "100")
	invAllCha(inv, "true", "Service charge", MYR, "100")

	tt := invTaxTot(inv, MYR, "87.63")
	invTaxTotTaxSub(tt, MYR, "87.63", MYR, "87.63", "01", OTH)
	invLegMonTot(inv, MYR, "1436.50", "1436.50", "1436.50", "1436.50", "1436.50", "0.30", "1436.50")

	il := invInvLin(inv, "1234", MYR, "1436.50", &valAttr{Val: "1", Attr: [2]string{uC, "C62"}})
	invInvLinAllCha(il, "false", "Sample Description", "0.15", MYR, "100")
	invInvLinAllCha(il, "true", "Sample Description", "0.10", MYR, "100")
	iltt := invInvLinTaxTot(il, MYR, "0")
	invInvLinTaxTotTaxSubPer(iltt, MYR, "1460.50", MYR, "0", "6.00", "E", "Exempt New Means of Transport", OTH)
	ptc := &valAttr{Val: "9800.00.0010", Attr: [2]string{lID, "PTC"}}
	cls := &valAttr{Val: "003", Attr: [2]string{lID, "CLASS"}}
	invInvLinIte(il, "Laptop Peripherals", "MALAYSIA", ptc, cls)
	invInvLinPriPriAmo(il, MYR, "17")
	invInvLinItePriExtAmo(il, MYR, "100")

	invSav(doc, out)
}

// nolint: unused
func samplingSelfBilledInvoice10(out string) {
	doc := etree.NewDocument()
	inv := newInv(doc)
	invTop(inv, "XML-SBINV12345", "2024-07-23", "00:40:00Z", _10, iTCSBInvoice, MYR, MYR)
	invPer(inv, "2024-07-01", "2024-07-31", "Monthly")
	invBilRefAddDocRef(inv, "151891-1981")
	invAddDocRef(inv, "L1", "CustomsImportForm", NUL)
	invAddDocRef(inv, "FTA", "FreeTradeAgreement", "Sample Description")
	invAddDocRef(inv, "L1", "K2", NUL)
	invAddDocRef(inv, "L1", NUL, NUL)

	asp := invAccSupPar(inv, &valAttrs{Val: "CPT-CCN-W-211111-KL-000002", Attrs: [][2]string{{sAN, "CertEX"}}})
	sp := invSupPar(asp, spMSICSC, egSPI_1)
	invXxxParPosAdd(sp, reSubWP.ReplaceAllString(cState[spCSC].State, ""), "50480", spCSC, egAddrLines_1, spCIC)
	invParLegEntRegNam(sp, "Supplier's Name")
	invCon(sp, "+60123456789", "supplier@email.com")

	acp := invAccCusPar(inv, nil)
	cp := invCusPar(acp, egCPI_1)
	invXxxParPosAdd(cp, reSubWP.ReplaceAllString(cState[spCSC].State, ""), "50480", spCSC, egAddrLines_1, "MALAYSIA")
	invParLegEntRegNam(cp, "Buyer's Name")
	invCon(cp, "+60123456780", "buyer@email.com")

	d := invDel(inv)
	dp := invDelPar(d, egDPI_1)
	invXxxParPosAdd(dp, reSubWP.ReplaceAllString(cState[spCSC].State, ""), "50480", spCSC, egAddrLines_1, "MALAYSIA")
	invParLegEntRegNam(dp, "Recipient's Name")
	ds := invDelShi(d, "1234")
	invDelShiFreAllCha(ds, "true", "Service charge", MYR, "100")

	invPayMea(inv, "01", "1234567890")
	invPayTerNot(inv, "Payment method is cash")
	invPrePay(inv, "E12345678912", MYR, "1.00", "2024-07-23", "00:30:00Z")

	invAllCha(inv, "false", "Sample Description", MYR, "100")
	invAllCha(inv, "true", "Service charge", MYR, "100")

	tt := invTaxTot(inv, MYR, "87.63")
	invTaxTotTaxSub(tt, MYR, "87.63", MYR, "87.63", "01", OTH)
	invLegMonTot(inv, MYR, "1436.50", "1436.50", "1436.50", "1436.50", "1436.50", "0.30", "1436.50")

	il := invInvLin(inv, "1234", MYR, "1436.50", &valAttr{Val: "1", Attr: [2]string{uC, "C62"}})
	invInvLinAllCha(il, "false", "Sample Description", "0.15", MYR, "100")
	invInvLinAllCha(il, "true", "Sample Description", "0.10", MYR, "100")
	iltt := invInvLinTaxTot(il, MYR, "0")
	invInvLinTaxTotTaxSubPer(iltt, MYR, "1460.50", MYR, "0", "6.00", "E", "Exempt New Means of Transport", OTH)
	ptc := &valAttr{Val: "9800.00.0010", Attr: [2]string{lID, "PTC"}}
	cls := &valAttr{Val: "003", Attr: [2]string{lID, "CLASS"}}
	invInvLinIte(il, "Laptop Peripherals", "MALAYSIA", ptc, cls)
	invInvLinPriPriAmo(il, MYR, "17")
	invInvLinItePriExtAmo(il, MYR, "100")

	invSav(doc, out)
}

// nolint: unused
func samplingSelfBilledCreditNote10(out string) {
	doc := etree.NewDocument()
	inv := newInv(doc)
	invTop(inv, "XML-SBCN12345", "2024-07-23", "00:40:00Z", _10, iTCSBCreditNote, MYR, MYR)
	invPer(inv, "2024-07-01", "2024-07-31", "Monthly")
	invBilRefInvDocRef(inv, "XML-SBINV12345", "Reference Self-billed Invoice UUID")
	invBilRefAddDocRef(inv, "151891-1981")

	invAddDocRef(inv, "L1", "CustomsImportForm", NUL)
	invAddDocRef(inv, "FTA", "FreeTradeAgreement", "Sample Description")
	invAddDocRef(inv, "L1", "K2", NUL)
	invAddDocRef(inv, "L1", NUL, NUL)

	asp := invAccSupPar(inv, &valAttrs{Val: "CPT-CCN-W-211111-KL-000002", Attrs: [][2]string{{sAN, "CertEX"}}})
	sp := invSupPar(asp, spMSICSC, egSPI_1)
	invXxxParPosAdd(sp, reSubWP.ReplaceAllString(cState[spCSC].State, ""), "50480", spCSC, egAddrLines_1, spCIC)
	invParLegEntRegNam(sp, "Supplier's Name")
	invCon(sp, "+60123456789", "supplier@email.com")

	acp := invAccCusPar(inv, nil)
	cp := invCusPar(acp, egCPI_1)
	invXxxParPosAdd(cp, reSubWP.ReplaceAllString(cState[spCSC].State, ""), "50480", spCSC, egAddrLines_1, "MALAYSIA")
	invParLegEntRegNam(cp, "Buyer's Name")
	invCon(cp, "+60123456780", "buyer@email.com")

	d := invDel(inv)
	dp := invDelPar(d, egDPI_1)
	invXxxParPosAdd(dp, reSubWP.ReplaceAllString(cState[spCSC].State, ""), "50480", spCSC, egAddrLines_1, "MALAYSIA")
	invParLegEntRegNam(dp, "Recipient's Name")
	ds := invDelShi(d, "1234")
	invDelShiFreAllCha(ds, "true", "Service charge", MYR, "100")

	invPayMea(inv, "01", "1234567890")
	invPayTerNot(inv, "Payment method is cash")
	invPrePay(inv, "E12345678912", MYR, "1.00", "2024-07-23", "00:30:00Z")

	invAllCha(inv, "false", "Sample Description", MYR, "100")
	invAllCha(inv, "true", "Service charge", MYR, "100")

	tt := invTaxTot(inv, MYR, "87.63")
	invTaxTotTaxSub(tt, MYR, "87.63", MYR, "87.63", "01", OTH)
	invLegMonTot(inv, MYR, "1436.50", "1436.50", "1436.50", "1436.50", "1436.50", "0.30", "1436.50")

	il := invInvLin(inv, "1234", MYR, "1436.50", &valAttr{Val: "1", Attr: [2]string{uC, "C62"}})
	invInvLinAllCha(il, "false", "Sample Description", "0.15", MYR, "100")
	invInvLinAllCha(il, "true", "Sample Description", "0.10", MYR, "100")
	iltt := invInvLinTaxTot(il, MYR, "0")
	invInvLinTaxTotTaxSubPer(iltt, MYR, "1460.50", MYR, "0", "6.00", "E", "Exempt New Means of Transport", OTH)
	ptc := &valAttr{Val: "9800.00.0010", Attr: [2]string{lID, "PTC"}}
	cls := &valAttr{Val: "003", Attr: [2]string{lID, "CLASS"}}
	invInvLinIte(il, "Laptop Peripherals", "MALAYSIA", ptc, cls)
	invInvLinPriPriAmo(il, MYR, "17")
	invInvLinItePriExtAmo(il, MYR, "100")

	invSav(doc, out)
}

// nolint: unused
func samplingSelfBilledDebitNote10(out string) {
	doc := etree.NewDocument()
	inv := newInv(doc)
	invTop(inv, "XML-SBDN12345", "2024-07-23", "00:40:00Z", _10, iTCSBDebitNote, MYR, MYR)
	invPer(inv, "2024-07-01", "2024-07-31", "Monthly")
	invBilRefInvDocRef(inv, "XML-SBINV12345", "Reference Self-billed Invoice UUID")
	invBilRefAddDocRef(inv, "151891-1981")

	invAddDocRef(inv, "L1", "CustomsImportForm", NUL)
	invAddDocRef(inv, "FTA", "FreeTradeAgreement", "Sample Description")
	invAddDocRef(inv, "L1", "K2", NUL)
	invAddDocRef(inv, "L1", NUL, NUL)

	asp := invAccSupPar(inv, &valAttrs{Val: "CPT-CCN-W-211111-KL-000002", Attrs: [][2]string{{sAN, "CertEX"}}})
	sp := invSupPar(asp, spMSICSC, egSPI_1)
	invXxxParPosAdd(sp, reSubWP.ReplaceAllString(cState[spCSC].State, ""), "50480", spCSC, egAddrLines_1, spCIC)
	invParLegEntRegNam(sp, "Supplier's Name")
	invCon(sp, "+60123456789", "supplier@email.com")

	acp := invAccCusPar(inv, nil)
	cp := invCusPar(acp, egCPI_1)
	invXxxParPosAdd(cp, reSubWP.ReplaceAllString(cState[spCSC].State, ""), "50480", spCSC, egAddrLines_1, "MALAYSIA")
	invParLegEntRegNam(cp, "Buyer's Name")
	invCon(cp, "+60123456780", "buyer@email.com")

	d := invDel(inv)
	dp := invDelPar(d, egDPI_1)
	invXxxParPosAdd(dp, reSubWP.ReplaceAllString(cState[spCSC].State, ""), "50480", spCSC, egAddrLines_1, "MALAYSIA")
	invParLegEntRegNam(dp, "Recipient's Name")
	ds := invDelShi(d, "1234")
	invDelShiFreAllCha(ds, "true", "Service charge", MYR, "100")

	invPayMea(inv, "01", "1234567890")
	invPayTerNot(inv, "Payment method is cash")
	invPrePay(inv, "E12345678912", MYR, "1.00", "2024-07-23", "00:30:00Z")

	invAllCha(inv, "false", "Sample Description", MYR, "100")
	invAllCha(inv, "true", "Service charge", MYR, "100")

	tt := invTaxTot(inv, MYR, "87.63")
	invTaxTotTaxSub(tt, MYR, "87.63", MYR, "87.63", "01", OTH)
	invLegMonTot(inv, MYR, "1436.50", "1436.50", "1436.50", "1436.50", "1436.50", "0.30", "1436.50")

	il := invInvLin(inv, "1234", MYR, "1436.50", &valAttr{Val: "1", Attr: [2]string{uC, "C62"}})
	invInvLinAllCha(il, "false", "Sample Description", "0.15", MYR, "100")
	invInvLinAllCha(il, "true", "Sample Description", "0.10", MYR, "100")
	iltt := invInvLinTaxTot(il, MYR, "0")
	invInvLinTaxTotTaxSubPer(iltt, MYR, "1460.50", MYR, "0", "6.00", "E", "Exempt New Means of Transport", OTH)
	ptc := &valAttr{Val: "9800.00.0010", Attr: [2]string{lID, "PTC"}}
	cls := &valAttr{Val: "003", Attr: [2]string{lID, "CLASS"}}
	invInvLinIte(il, "Laptop Peripherals", "MALAYSIA", ptc, cls)
	invInvLinPriPriAmo(il, MYR, "17")
	invInvLinItePriExtAmo(il, MYR, "100")

	invSav(doc, out)
}

// nolint: unused
func samplingSelfBilledRefundNote10(out string) {
	doc := etree.NewDocument()
	inv := newInv(doc)
	invTop(inv, "XML-SBRN12345", "2024-07-23", "00:40:00Z", _10, iTCSBRefundNote, MYR, MYR)
	invPer(inv, "2024-07-01", "2024-07-31", "Monthly")
	invBilRefInvDocRef(inv, "XML-SBINV12345", "Reference Self-billed Invoice UUID")
	invBilRefAddDocRef(inv, "151891-1981")

	invAddDocRef(inv, "L1", "CustomsImportForm", NUL)
	invAddDocRef(inv, "FTA", "FreeTradeAgreement", "Sample Description")
	invAddDocRef(inv, "L1", "K2", NUL)
	invAddDocRef(inv, "L1", NUL, NUL)

	asp := invAccSupPar(inv, &valAttrs{Val: "CPT-CCN-W-211111-KL-000002", Attrs: [][2]string{{sAN, "CertEX"}}})
	sp := invSupPar(asp, spMSICSC, egSPI_1)
	invXxxParPosAdd(sp, reSubWP.ReplaceAllString(cState[spCSC].State, ""), "50480", spCSC, egAddrLines_1, spCIC)
	invParLegEntRegNam(sp, "Supplier's Name")
	invCon(sp, "+60123456789", "supplier@email.com")

	acp := invAccCusPar(inv, nil)
	cp := invCusPar(acp, egCPI_1)
	invXxxParPosAdd(cp, reSubWP.ReplaceAllString(cState[spCSC].State, ""), "50480", spCSC, egAddrLines_1, "MALAYSIA")
	invParLegEntRegNam(cp, "Buyer's Name")
	invCon(cp, "+60123456780", "buyer@email.com")

	d := invDel(inv)
	dp := invDelPar(d, egDPI_1)
	invXxxParPosAdd(dp, reSubWP.ReplaceAllString(cState[spCSC].State, ""), "50480", spCSC, egAddrLines_1, "MALAYSIA")
	invParLegEntRegNam(dp, "Recipient's Name")
	ds := invDelShi(d, "1234")
	invDelShiFreAllCha(ds, "true", "Service charge", MYR, "100")

	invPayMea(inv, "01", "1234567890")
	invPayTerNot(inv, "Payment method is cash")
	invPrePay(inv, "E12345678912", MYR, "1.00", "2024-07-23", "00:30:00Z")

	invAllCha(inv, "false", "Sample Description", MYR, "100")
	invAllCha(inv, "true", "Service charge", MYR, "100")

	tt := invTaxTot(inv, MYR, "87.63")
	invTaxTotTaxSub(tt, MYR, "87.63", MYR, "87.63", "01", OTH)
	invLegMonTot(inv, MYR, "1436.50", "1436.50", "1436.50", "1436.50", "1436.50", "0.30", "1436.50")

	il := invInvLin(inv, "1234", MYR, "1436.50", &valAttr{Val: "1", Attr: [2]string{uC, "C62"}})
	invInvLinAllCha(il, "false", "Sample Description", "0.15", MYR, "100")
	invInvLinAllCha(il, "true", "Sample Description", "0.10", MYR, "100")
	iltt := invInvLinTaxTot(il, MYR, "0")
	invInvLinTaxTotTaxSubPer(iltt, MYR, "1460.50", MYR, "0", "6.00", "E", "Exempt New Means of Transport", OTH)
	ptc := &valAttr{Val: "9800.00.0010", Attr: [2]string{lID, "PTC"}}
	cls := &valAttr{Val: "003", Attr: [2]string{lID, "CLASS"}}
	invInvLinIte(il, "Laptop Peripherals", "MALAYSIA", ptc, cls)
	invInvLinPriPriAmo(il, MYR, "17")
	invInvLinItePriExtAmo(il, MYR, "100")

	invSav(doc, out)
}
