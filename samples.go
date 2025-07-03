package main

import (
	"github.com/beevik/etree"
)

// nolint:unused
func sdkSampleInvoice10(out string) {
	// https://sdk.myinvois.hasil.gov.my/documents/invoice-v1-1/
	// https://sdk.myinvois.hasil.gov.my/documents/invoice-v1-0/ target mandatory
	doc := etree.NewDocument()
	// doc.CreateProcInst("xml", verEnc) // perhaps implied thus excluded
	inv := doc.CreateElement("Invoice")
	inv.CreateAttr("xmlns", urnInvoice2)
	inv.CreateAttr("xmlns:cac", urnCommonAggregateComponents2)
	inv.CreateAttr("xmlns:cbc", urnCommonBasicComponents2)

	inv.CreateElement(cbcID).SetText("XML-INV12345")
	inv.CreateElement(cbcIssueDate).SetText("2024-07-23")
	inv.CreateElement(cbcIssueTime).SetText("00:40:00Z")
	// 01:Invoice 02:Credit 03:Debit 04:Refund 11:SelfBilledInvoice 12:SBCredit 13:SBDebit 14:SBRefund
	inv.CreateElement(cbcInvoiceTypeCode).CreateAttr("listVersionID", "1.0").Element().SetText("01") // 1.1 requires signature!
	inv.CreateElement(cbcDocumentCurrencyCode).SetText(MYR)
	inv.CreateElement(cbcTaxCurrencyCode).SetText(MYR)

	period := inv.CreateElement(cacInvoicePeriod)
	period.CreateElement(cbcStartDate).SetText("2024-07-01")
	period.CreateElement(cbcEndDate).SetText("2024-07-31")
	period.CreateElement(cbcDescription).SetText("Monthly")

	inv.CreateElement(cacBillingReference).CreateElement(cacAdditionalDocumentReference).
		CreateElement(cbcID).SetText("151891-1981")

	l1 := inv.CreateElement(cacAdditionalDocumentReference)
	l1.CreateElement(cbcID).SetText("L1")
	l1.CreateElement(cbcDocumentType).SetText("CustomsImportForm")

	fta := inv.CreateElement(cacAdditionalDocumentReference)
	fta.CreateElement(cbcID).SetText("FTA")
	fta.CreateElement(cbcDocumentType).SetText("FreeTradeAgreement")
	fta.CreateElement(cbcDocumentDescription).SetText("Sample Description")

	l1_a := inv.CreateElement(cacAdditionalDocumentReference)
	l1_a.CreateElement(cbcID).SetText("L1")
	l1_a.CreateElement(cbcDocumentType).SetText("K2")

	inv.CreateElement(cacAdditionalDocumentReference).CreateElement(cbcID).SetText("L1")

	asp := inv.CreateElement(cacAccountingSupplierParty)
	asp.CreateElement(cbcAdditionalAccountID).CreateAttr("schemeAgencyName", "CertEX").
		Element().SetText("CPT-CCN-W-211111-KL-000002")

	sp := asp.CreateElement(cacParty)
	sp.CreateElement(cbcIndustryClassificationCode).
		CreateAttr("name", cMSICSC[spMSICSC].Description).Element().SetText(spMSICSC)

	sp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, TIN).Element().SetText("Supplier's TIN")
	sp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, BRN).Element().SetText("Supplier's BRN")
	sp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, SST).Element().SetText("NA")
	sp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, TTX).Element().SetText("NA")

	spa := sp.CreateElement(cacPostalAddress)
	spa.CreateElement(cbcCityName).SetText(reSubWP.ReplaceAllString(cState[spCSC].State, ""))
	spa.CreateElement(cbcPostalZone).SetText("50480")
	spa.CreateElement(cbcCountrySubentityCode).SetText(spCSC)
	spa.CreateElement(cacAddressLine).CreateElement(cbcLine).SetText("Lot 66")
	spa.CreateElement(cacAddressLine).CreateElement(cbcLine).SetText("Bangunan Merdeka")
	spa.CreateElement(cacAddressLine).CreateElement(cbcLine).SetText("Persiaran Jaya")
	spa.CreateElement(cacCountry).CreateElement(cbcIdentificationCode).
		CreateAttr(lID, "ISO3166-1").Element().CreateAttr(lAID, "6").
		Element().SetText(cCountryRev[spCIC].Code)
	sp.CreateElement(cacPartyLegalEntity).CreateElement(cbcRegistrationName).SetText("Supplier's Name")
	sc := sp.CreateElement(cacContact)
	sc.CreateElement(cbcTelephone).SetText("+60123456789")
	sc.CreateElement(cbcElectronicMail).SetText("supplier@email.com")

	acp := inv.CreateElement(cacAccountingCustomerParty)
	cp := acp.CreateElement(cacParty)

	cp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, TIN).Element().SetText("Buyer's TIN")
	cp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, BRN).Element().SetText("Buyer's BRN")
	cp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, SST).Element().SetText("NA")
	cp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, TTX).Element().SetText("NA")

	cpa := cp.CreateElement(cacPostalAddress)
	cpa.CreateElement(cbcCityName).SetText(reSubWP.ReplaceAllString(cState[spCSC].State, ""))
	cpa.CreateElement(cbcPostalZone).SetText("50480")
	cpa.CreateElement(cbcCountrySubentityCode).SetText(spCSC)
	cpa.CreateElement(cacAddressLine).CreateElement(cbcLine).SetText("Lot 66")
	cpa.CreateElement(cacAddressLine).CreateElement(cbcLine).SetText("Bangunan Merdeka")
	cpa.CreateElement(cacAddressLine).CreateElement(cbcLine).SetText("Persiaran Jaya")
	cpa.CreateElement(cacCountry).CreateElement(cbcIdentificationCode).
		CreateAttr(lID, "ISO3166-1").Element().CreateAttr(lAID, "6").
		Element().SetText(cCountryRev["MALAYSIA"].Code)
	cp.CreateElement(cacPartyLegalEntity).CreateElement(cbcRegistrationName).SetText("Buyer's Name")
	cc := cp.CreateElement(cacContact)
	cc.CreateElement(cbcTelephone).SetText("+60123456780")
	cc.CreateElement(cbcElectronicMail).SetText("buyer@email.com")

	d := inv.CreateElement(cacDelivery)
	dp := d.CreateElement(cacDeliveryParty)

	dp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, TIN).Element().SetText("Recipient's TIN")
	dp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, BRN).Element().SetText("Recipient's BRN")

	dpa := dp.CreateElement(cacPostalAddress)
	dpa.CreateElement(cbcCityName).SetText(reSubWP.ReplaceAllString(cState[spCSC].State, ""))
	dpa.CreateElement(cbcPostalZone).SetText("50480")
	dpa.CreateElement(cbcCountrySubentityCode).SetText(spCSC)
	dpa.CreateElement(cacAddressLine).CreateElement(cbcLine).SetText("Lot 66")
	dpa.CreateElement(cacAddressLine).CreateElement(cbcLine).SetText("Bangunan Merdeka")
	dpa.CreateElement(cacAddressLine).CreateElement(cbcLine).SetText("Persiaran Jaya")
	dpa.CreateElement(cacCountry).CreateElement(cbcIdentificationCode).
		CreateAttr(lID, "ISO3166-1").Element().CreateAttr(lAID, "6").
		Element().SetText(cCountryRev["MALAYSIA"].Code)
	dp.CreateElement(cacPartyLegalEntity).CreateElement(cbcRegistrationName).SetText("Recipient's Name")

	ds := d.CreateElement(cacShipment)
	ds.CreateElement(cbcID).SetText("1234")
	dsfac := ds.CreateElement(cacFreightAllowanceCharge)
	dsfac.CreateElement(cbcChargeIndicator).SetText("true")
	dsfac.CreateElement(cbcAllowanceChargeReason).SetText("Service charge")
	dsfac.CreateElement(cbcAmount).CreateAttr(cID, MYR).Element().SetText("100")

	pm := inv.CreateElement(cacPaymentMeans)
	pm.CreateElement(cbcPaymentMeansCode).SetText("01") // PaymentMethods.json
	pm.CreateElement(cacPayeeFinancialAccount).CreateElement(cbcID).SetText("1234567890")

	inv.CreateElement(cacPaymentTerms).CreateElement(cbcNote).SetText("Payment method is cash")

	pp := inv.CreateElement(cacPrepaidPayment)
	pp.CreateElement(cbcID).SetText("E12345678912")
	pp.CreateElement(cbcPaidAmount).CreateAttr(cID, MYR).Element().SetText("1.00")
	pp.CreateElement(cbcPaidDate).SetText("2024-07-23")
	pp.CreateElement(cbcPaidTime).SetText("00:30:00Z")

	ac1 := inv.CreateElement(cacAllowanceCharge)
	ac1.CreateElement(cbcChargeIndicator).SetText("false")
	ac1.CreateElement(cbcAllowanceChargeReason).SetText("Sample Description")
	ac1.CreateElement(cbcAmount).CreateAttr(cID, MYR).Element().SetText("100")

	ac2 := inv.CreateElement(cacAllowanceCharge)
	ac2.CreateElement(cbcChargeIndicator).SetText("true")
	ac2.CreateElement(cbcAllowanceChargeReason).SetText("Service charge")
	ac2.CreateElement(cbcAmount).CreateAttr(cID, MYR).Element().SetText("100")

	tt := inv.CreateElement(cacTaxTotal)
	tt.CreateElement(cbcTaxAmount).CreateAttr(cID, MYR).Element().SetText("87.63")

	ts := tt.CreateElement(cacTaxSubtotal)
	ts.CreateElement(cbcTaxableAmount).CreateAttr(cID, MYR).Element().SetText("87.63")
	ts.CreateElement(cbcTaxAmount).CreateAttr(cID, MYR).Element().SetText("87.63")
	tstc := ts.CreateElement(cacTaxCategory)
	tstc.CreateElement(cbcID).SetText("01")
	tstc.CreateElement(cacTaxScheme).CreateElement(cbcID).
		CreateAttr(sID, UE5).Element().CreateAttr(sAID, "6").Element().SetText(OTH)

	lmt := inv.CreateElement(cacLegalMonetaryTotal)
	lmt.CreateElement(cbcLineExtensionAmount).CreateAttr(cID, MYR).Element().SetText("1436.50")
	lmt.CreateElement(cbcTaxExclusiveAmount).CreateAttr(cID, MYR).Element().SetText("1436.50")
	lmt.CreateElement(cbcTaxInclusiveAmount).CreateAttr(cID, MYR).Element().SetText("1436.50")
	lmt.CreateElement(cbcAllowanceTotalAmount).CreateAttr(cID, MYR).Element().SetText("1436.50")
	lmt.CreateElement(cbcChargeTotalAmount).CreateAttr(cID, MYR).Element().SetText("1436.50")
	lmt.CreateElement(cbcPayableRoundingAmount).CreateAttr(cID, MYR).Element().SetText("0.30")
	lmt.CreateElement(cbcPayableAmount).CreateAttr(cID, MYR).Element().SetText("1436.50")

	il := inv.CreateElement(cacInvoiceLine)
	il.CreateElement(cbcID).SetText("1234")
	il.CreateElement(cbcInvoicedQuantity).CreateAttr("unitCode", "C62").Element().SetText("1")
	il.CreateElement(cbcLineExtensionAmount).CreateAttr(cID, MYR).Element().SetText("1436.50")

	ilac1 := il.CreateElement(cacAllowanceCharge)
	ilac1.CreateElement(cbcChargeIndicator).SetText("false")
	ilac1.CreateElement(cbcAllowanceChargeReason).SetText("Sample Description")
	ilac1.CreateElement(cbcMultiplierFactorNumeric).SetText("0.15")
	ilac1.CreateElement(cbcAmount).CreateAttr(cID, MYR).Element().SetText("100")

	ilac2 := il.CreateElement(cacAllowanceCharge)
	ilac2.CreateElement(cbcChargeIndicator).SetText("true")
	ilac2.CreateElement(cbcAllowanceChargeReason).SetText("Sample Description")
	ilac2.CreateElement(cbcMultiplierFactorNumeric).SetText("0.10")
	ilac2.CreateElement(cbcAmount).CreateAttr(cID, MYR).Element().SetText("100")

	iltt := il.CreateElement(cacTaxTotal)
	iltt.CreateElement(cbcTaxAmount).CreateAttr(cID, MYR).Element().SetText("0")
	ilttts := iltt.CreateElement(cacTaxSubtotal)
	ilttts.CreateElement(cbcTaxableAmount).CreateAttr(cID, MYR).Element().SetText("1460.50")
	ilttts.CreateElement(cbcTaxAmount).CreateAttr(cID, MYR).Element().SetText("0")
	ilttts.CreateElement(cbcPercent).SetText("6.00")
	ittttstc := ilttts.CreateElement(cacTaxCategory)
	ittttstc.CreateElement(cbcID).SetText("E")
	ittttstc.CreateElement(cbcTaxExemptionReason).SetText("Exempt New Means of Transport")
	ittttstc.CreateElement(cacTaxScheme).CreateElement(cbcID).
		CreateAttr(sID, UE5).Element().CreateAttr(sAID, "6").Element().SetText(OTH)

	ili := il.CreateElement(cacItem)
	ili.CreateElement(cbcDescription).SetText("Laptop Peripherals")
	ili.CreateElement(cacOriginCountry).CreateElement(cbcIdentificationCode).SetText(cCountryRev["MALAYSIA"].Code)
	ili.CreateElement(cacCommodityClassification).CreateElement(cbcItemClassificationCode).
		CreateAttr(lID, "PTC").Element().SetText("9800.00.0010")
	ili.CreateElement(cacCommodityClassification).CreateElement(cbcItemClassificationCode).
		CreateAttr(lID, "CLASS").Element().SetText("003")

	il.CreateElement(cacPrice).CreateElement(cbcPriceAmount).CreateAttr(cID, MYR).
		Element().SetText("17")
	il.CreateElement(cacItemPriceExtension).CreateElement(cbcAmount).CreateAttr(cID, MYR).
		Element().SetText("100")

	invSav(doc, out)
}

// nolint:unused
func sdkSampleInvoiceMultiLine10(out string) {
	doc := etree.NewDocument()
	inv := doc.CreateElement("Invoice")
	inv.CreateAttr("xmlns", urnInvoice2)
	inv.CreateAttr("xmlns:cac", urnCommonAggregateComponents2)
	inv.CreateAttr("xmlns:cbc", urnCommonBasicComponents2)

	inv.CreateElement(cbcID).SetText("XML-INV12345")
	inv.CreateElement(cbcIssueDate).SetText("2025-02-06")
	inv.CreateElement(cbcIssueTime).SetText("00:30:00Z")
	inv.CreateElement(cbcInvoiceTypeCode).CreateAttr("listVersionID", "1.0").Element().SetText("01") // 1.1 requires signature!
	inv.CreateElement(cbcDocumentCurrencyCode).SetText(MYR)
	inv.CreateElement(cbcTaxCurrencyCode).SetText(MYR)

	period := inv.CreateElement(cacInvoicePeriod)
	period.CreateElement(cbcStartDate).SetText("2025-01-01")
	period.CreateElement(cbcEndDate).SetText("2025-01-31")
	period.CreateElement(cbcDescription).SetText("Monthly")

	inv.CreateElement(cacBillingReference).CreateElement(cacAdditionalDocumentReference).
		CreateElement(cbcID).SetText("E12345678912")

	adf1 := inv.CreateElement(cacAdditionalDocumentReference)
	adf1.CreateElement(cbcID).SetText("E23456789123,E98765432123")
	adf1.CreateElement(cbcDocumentType).SetText("CustomsImportForm")

	adf2 := inv.CreateElement(cacAdditionalDocumentReference)
	adf2.CreateElement(cbcID).SetText("FTA")
	adf2.CreateElement(cbcDocumentType).SetText("FreeTradeAgreement")
	adf2.CreateElement(cbcDocumentDescription).SetText("ASEAN-Australia-New Zealand FTA (AANZFTA)")

	adf3 := inv.CreateElement(cacAdditionalDocumentReference)
	adf3.CreateElement(cbcID).SetText("E12345678912,E23456789123")
	adf3.CreateElement(cbcDocumentType).SetText("K2")

	inv.CreateElement(cacAdditionalDocumentReference).CreateElement(cbcID).SetText("CIF")

	asp := inv.CreateElement(cacAccountingSupplierParty)
	asp.CreateElement(cbcAdditionalAccountID).CreateAttr("schemeAgencyName", "CertEX").
		Element().SetText("CPT-CCN-W-211111-KL-000002")

	sp := asp.CreateElement(cacParty)

	sp.CreateElement(cbcIndustryClassificationCode).
		CreateAttr("name", cMSICSC[spMSICSC].Description).Element().SetText(spMSICSC)

	sp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, TIN).Element().SetText("Supplier's TIN")
	sp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, BRN).Element().SetText("Supplier's BRN")
	sp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, SST).Element().SetText("NA")
	sp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, TTX).Element().SetText("NA")

	spa := sp.CreateElement(cacPostalAddress)
	spa.CreateElement(cbcCityName).SetText(reSubWP.ReplaceAllString(cState[spCSC].State, ""))
	spa.CreateElement(cbcPostalZone).SetText("50480")
	spa.CreateElement(cbcCountrySubentityCode).SetText(spCSC)
	spa.CreateElement(cacAddressLine).CreateElement(cbcLine).SetText("Lot 66")
	spa.CreateElement(cacAddressLine).CreateElement(cbcLine).SetText("Bangunan Merdeka")
	spa.CreateElement(cacAddressLine).CreateElement(cbcLine).SetText("Persiaran Jaya")
	spa.CreateElement(cacCountry).CreateElement(cbcIdentificationCode).
		CreateAttr(lID, "ISO3166-1").Element().CreateAttr(lAID, "6").
		Element().SetText(cCountryRev[spCIC].Code)
	sp.CreateElement(cacPartyLegalEntity).CreateElement(cbcRegistrationName).SetText("Supplier's Name")
	sc := sp.CreateElement(cacContact)
	sc.CreateElement(cbcTelephone).SetText("+60123456789")
	sc.CreateElement(cbcElectronicMail).SetText("supplier@email.com")

	acp := inv.CreateElement(cacAccountingCustomerParty)
	cp := acp.CreateElement(cacParty)

	cp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, TIN).Element().SetText("Buyer's TIN")
	cp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, BRN).Element().SetText("Buyer's BRN")
	cp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, SST).Element().SetText("NA")
	cp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, TTX).Element().SetText("NA")

	cpa := cp.CreateElement(cacPostalAddress)
	cpa.CreateElement(cbcCityName).SetText(reSubWP.ReplaceAllString(cState[spCSC].State, ""))
	cpa.CreateElement(cbcPostalZone).SetText("50480")
	cpa.CreateElement(cbcCountrySubentityCode).SetText(spCSC)
	cpa.CreateElement(cacAddressLine).CreateElement(cbcLine).SetText("Lot 66")
	cpa.CreateElement(cacAddressLine).CreateElement(cbcLine).SetText("Bangunan Merdeka")
	cpa.CreateElement(cacAddressLine).CreateElement(cbcLine).SetText("Persiaran Jaya")
	cpa.CreateElement(cacCountry).CreateElement(cbcIdentificationCode).
		CreateAttr(lID, "ISO3166-1").Element().CreateAttr(lAID, "6").
		Element().SetText(cCountryRev["MALAYSIA"].Code)
	cp.CreateElement(cacPartyLegalEntity).CreateElement(cbcRegistrationName).SetText("Buyer's Name")
	cc := cp.CreateElement(cacContact)
	cc.CreateElement(cbcTelephone).SetText("+60123456780")
	cc.CreateElement(cbcElectronicMail).SetText("buyer@email.com")

	ac1 := inv.CreateElement(cacAllowanceCharge)
	ac1.CreateElement(cbcChargeIndicator).SetText("false")
	ac1.CreateElement(cbcAllowanceChargeReason).SetText("")
	ac1.CreateElement(cbcAmount).CreateAttr(cID, MYR).Element().SetText("0.00")

	ac2 := inv.CreateElement(cacAllowanceCharge)
	ac2.CreateElement(cbcChargeIndicator).SetText("true")
	ac2.CreateElement(cbcAllowanceChargeReason).SetText("")
	ac2.CreateElement(cbcAmount).CreateAttr(cID, MYR).Element().SetText("0.00")

	tt := inv.CreateElement(cacTaxTotal)
	tt.CreateElement(cbcTaxAmount).CreateAttr(cID, MYR).Element().SetText("85.00")

	ts1 := tt.CreateElement(cacTaxSubtotal)
	ts1.CreateElement(cbcTaxableAmount).CreateAttr(cID, MYR).Element().SetText("1000.00")
	ts1.CreateElement(cbcTaxAmount).CreateAttr(cID, MYR).Element().SetText("10.00")
	ts1tc := ts1.CreateElement(cacTaxCategory)
	ts1tc.CreateElement(cbcID).SetText("01")
	ts1tc.CreateElement(cacTaxScheme).CreateElement(cbcID).
		CreateAttr(sID, UE5).Element().CreateAttr(sAID, "6").Element().SetText(OTH)

	ts2 := tt.CreateElement(cacTaxSubtotal)
	ts2.CreateElement(cbcTaxableAmount).CreateAttr(cID, MYR).Element().SetText("1500.00")
	ts2.CreateElement(cbcTaxAmount).CreateAttr(cID, MYR).Element().SetText("75.00")
	ts2tc := ts2.CreateElement(cacTaxCategory)
	ts2tc.CreateElement(cbcID).SetText("02")
	ts2tc.CreateElement(cacTaxScheme).CreateElement(cbcID).
		CreateAttr(sID, UE5).Element().CreateAttr(sAID, "6").Element().SetText(OTH)

	ts3 := tt.CreateElement(cacTaxSubtotal)
	ts3.CreateElement(cbcTaxableAmount).CreateAttr(cID, MYR).Element().SetText("2000.00")
	ts3.CreateElement(cbcTaxAmount).CreateAttr(cID, MYR).Element().SetText("100.00")
	ts3tc := ts3.CreateElement(cacTaxCategory)
	ts3tc.CreateElement(cbcID).SetText("E")
	ts3tc.CreateElement(cacTaxScheme).CreateElement(cbcID).
		CreateAttr(sID, UE5).Element().CreateAttr(sAID, "6").Element().SetText(OTH)

	lmt := inv.CreateElement(cacLegalMonetaryTotal)
	lmt.CreateElement(cbcLineExtensionAmount).CreateAttr(cID, MYR).Element().SetText("4500.00")
	lmt.CreateElement(cbcTaxExclusiveAmount).CreateAttr(cID, MYR).Element().SetText("4500.00")
	lmt.CreateElement(cbcTaxInclusiveAmount).CreateAttr(cID, MYR).Element().SetText("4585.00")
	lmt.CreateElement(cbcAllowanceTotalAmount).CreateAttr(cID, MYR).Element().SetText("0.00")
	lmt.CreateElement(cbcChargeTotalAmount).CreateAttr(cID, MYR).Element().SetText("0.00")
	lmt.CreateElement(cbcPayableRoundingAmount).CreateAttr(cID, MYR).Element().SetText("0.00")
	lmt.CreateElement(cbcPayableAmount).CreateAttr(cID, MYR).Element().SetText("4585.00")

	il1 := inv.CreateElement(cacInvoiceLine)
	il1.CreateElement(cbcID).SetText("001")
	il1.CreateElement(cbcInvoicedQuantity).CreateAttr("unitCode", "C62").Element().SetText("1")
	il1.CreateElement(cbcLineExtensionAmount).CreateAttr(cID, MYR).Element().SetText("1000.00")

	il1ac1 := il1.CreateElement(cacAllowanceCharge)
	il1ac1.CreateElement(cbcChargeIndicator).SetText("false")
	il1ac1.CreateElement(cbcAllowanceChargeReason).SetText("")
	il1ac1.CreateElement(cbcMultiplierFactorNumeric).SetText("0")
	il1ac1.CreateElement(cbcAmount).CreateAttr(cID, MYR).Element().SetText("0.00")

	il1ac2 := il1.CreateElement(cacAllowanceCharge)
	il1ac2.CreateElement(cbcChargeIndicator).SetText("true")
	il1ac2.CreateElement(cbcAllowanceChargeReason).SetText("")
	il1ac2.CreateElement(cbcMultiplierFactorNumeric).SetText("0")
	il1ac2.CreateElement(cbcAmount).CreateAttr(cID, MYR).Element().SetText("0.00")

	il1tt := il1.CreateElement(cacTaxTotal)
	il1tt.CreateElement(cbcTaxAmount).CreateAttr(cID, MYR).Element().SetText("10.00")
	il1ttts := il1tt.CreateElement(cacTaxSubtotal)
	il1ttts.CreateElement(cbcTaxableAmount).CreateAttr(cID, MYR).Element().SetText("1000.00")
	il1ttts.CreateElement(cbcTaxAmount).CreateAttr(cID, MYR).Element().SetText("10.00")
	il1ttts.CreateElement(cbcBaseUnitMeasure).CreateAttr("unitCode", "C62").Element().SetText("1")
	il1ttts.CreateElement(cbcPerUnitAmount).CreateAttr(cID, MYR).Element().SetText("10.00")
	il1tttstc := il1ttts.CreateElement(cacTaxCategory)
	il1tttstc.CreateElement(cbcID).SetText("01")
	il1tttstc.CreateElement(cacTaxScheme).CreateElement(cbcID).
		CreateAttr(sID, UE5).Element().CreateAttr(sAID, "6").Element().SetText(OTH)

	il1i := il1.CreateElement(cacItem)
	il1i.CreateElement(cbcDescription).SetText("Laptop Peripherals")
	il1i.CreateElement(cacOriginCountry).CreateElement(cbcIdentificationCode).SetText(cCountryRev["MALAYSIA"].Code)
	il1i.CreateElement(cacCommodityClassification).CreateElement(cbcItemClassificationCode).
		CreateAttr(lID, "PTC").Element().SetText("9800.00.0010")
	il1i.CreateElement(cacCommodityClassification).CreateElement(cbcItemClassificationCode).
		CreateAttr(lID, "CLASS").Element().SetText("003")

	il1.CreateElement(cacPrice).CreateElement(cbcPriceAmount).
		CreateAttr(cID, MYR).Element().SetText("1000.00")
	il1.CreateElement(cacItemPriceExtension).CreateElement(cbcAmount).
		CreateAttr(cID, MYR).Element().SetText("1000.00")

	il2 := inv.CreateElement(cacInvoiceLine)
	il2.CreateElement(cbcID).SetText("002")
	il2.CreateElement(cbcInvoicedQuantity).CreateAttr("unitCode", "C62").Element().SetText("1")
	il2.CreateElement(cbcLineExtensionAmount).CreateAttr(cID, MYR).Element().SetText("1500.00")

	il2ac1 := il2.CreateElement(cacAllowanceCharge)
	il2ac1.CreateElement(cbcChargeIndicator).SetText("false")
	il2ac1.CreateElement(cbcAllowanceChargeReason).SetText("")
	il2ac1.CreateElement(cbcMultiplierFactorNumeric).SetText("0")
	il2ac1.CreateElement(cbcAmount).CreateAttr(cID, MYR).Element().SetText("0.00")

	il2ac2 := il2.CreateElement(cacAllowanceCharge)
	il2ac2.CreateElement(cbcChargeIndicator).SetText("true")
	il2ac2.CreateElement(cbcAllowanceChargeReason).SetText("")
	il2ac2.CreateElement(cbcMultiplierFactorNumeric).SetText("0")
	il2ac2.CreateElement(cbcAmount).CreateAttr(cID, MYR).Element().SetText("0.00")

	il2tt := il2.CreateElement(cacTaxTotal)
	il2tt.CreateElement(cbcTaxAmount).CreateAttr(cID, MYR).Element().SetText("75.00")
	il2ttts := il2tt.CreateElement(cacTaxSubtotal)
	il2ttts.CreateElement(cbcTaxableAmount).CreateAttr(cID, MYR).Element().SetText("1500.00")
	il2ttts.CreateElement(cbcTaxAmount).CreateAttr(cID, MYR).Element().SetText("75.00")
	il2ttts.CreateElement(cbcPercent).SetText("5.00")
	il2tttstc := il2ttts.CreateElement(cacTaxCategory)
	il2tttstc.CreateElement(cbcID).SetText("02")
	il2tttstc.CreateElement(cacTaxScheme).CreateElement(cbcID).
		CreateAttr(sID, UE5).Element().CreateAttr(sAID, "6").Element().SetText(OTH)

	il2i := il2.CreateElement(cacItem)
	il2i.CreateElement(cbcDescription).SetText("Computer Monitor")
	il2i.CreateElement(cacOriginCountry).CreateElement(cbcIdentificationCode).SetText(cCountryRev["MALAYSIA"].Code)
	il2i.CreateElement(cacCommodityClassification).CreateElement(cbcItemClassificationCode).
		CreateAttr(lID, "PTC").Element().SetText("9800.00.0011")
	il2i.CreateElement(cacCommodityClassification).CreateElement(cbcItemClassificationCode).
		CreateAttr(lID, "CLASS").Element().SetText("003")

	il2.CreateElement(cacPrice).CreateElement(cbcPriceAmount).
		CreateAttr(cID, MYR).Element().SetText("1500.00")
	il2.CreateElement(cacItemPriceExtension).CreateElement(cbcAmount).
		CreateAttr(cID, MYR).Element().SetText("1500.00")

	il3 := inv.CreateElement(cacInvoiceLine)
	il3.CreateElement(cbcID).SetText("003")
	il3.CreateElement(cbcInvoicedQuantity).CreateAttr("unitCode", "C62").Element().SetText("1")
	il3.CreateElement(cbcLineExtensionAmount).CreateAttr(cID, MYR).Element().SetText("2000.00")

	il3ac1 := il3.CreateElement(cacAllowanceCharge)
	il3ac1.CreateElement(cbcChargeIndicator).SetText("false")
	il3ac1.CreateElement(cbcAllowanceChargeReason).SetText("")
	il3ac1.CreateElement(cbcMultiplierFactorNumeric).SetText("0")
	il3ac1.CreateElement(cbcAmount).CreateAttr(cID, MYR).Element().SetText("0.00")

	il3ac2 := il3.CreateElement(cacAllowanceCharge)
	il3ac2.CreateElement(cbcChargeIndicator).SetText("true")
	il3ac2.CreateElement(cbcAllowanceChargeReason).SetText("")
	il3ac2.CreateElement(cbcMultiplierFactorNumeric).SetText("0")
	il3ac2.CreateElement(cbcAmount).CreateAttr(cID, MYR).Element().SetText("0.00")

	il3tt := il3.CreateElement(cacTaxTotal)
	il3tt.CreateElement(cbcTaxAmount).CreateAttr(cID, MYR).Element().SetText("0.00")
	il3ttts1 := il3tt.CreateElement(cacTaxSubtotal)
	il3ttts1.CreateElement(cbcTaxableAmount).CreateAttr(cID, MYR).Element().SetText("0.00")
	il3ttts1.CreateElement(cbcTaxAmount).CreateAttr(cID, MYR).Element().SetText("0.00")
	il3ttts1.CreateElement(cbcPercent).SetText("5.00")
	il3ttts1tc := il3ttts1.CreateElement(cacTaxCategory)
	il3ttts1tc.CreateElement(cbcID).SetText("01")
	il3ttts1tc.CreateElement(cacTaxScheme).CreateElement(cbcID).
		CreateAttr(sID, UE5).Element().CreateAttr(sAID, "6").Element().SetText(OTH)
	il3ttts2 := il3tt.CreateElement(cacTaxSubtotal)
	il3ttts2.CreateElement(cbcTaxableAmount).CreateAttr(cID, MYR).Element().SetText("2000.00")
	il3ttts2.CreateElement(cbcTaxAmount).CreateAttr(cID, MYR).Element().SetText("100.00")
	il3ttts2tc := il3ttts2.CreateElement(cacTaxCategory)
	il3ttts2tc.CreateElement(cbcID).SetText("E")
	il3ttts2tc.CreateElement(cbcTaxExemptionReason).SetText("Special Case")
	il3ttts2tc.CreateElement(cacTaxScheme).CreateElement(cbcID).
		CreateAttr(sID, UE5).Element().CreateAttr(sAID, "6").Element().SetText(OTH)

	il3i := il3.CreateElement(cacItem)
	il3i.CreateElement(cbcDescription).SetText("Wireless Mouse")
	il3i.CreateElement(cacOriginCountry).CreateElement(cbcIdentificationCode).SetText(cCountryRev["MALAYSIA"].Code)
	il3i.CreateElement(cacCommodityClassification).CreateElement(cbcItemClassificationCode).
		CreateAttr(lID, "PTC").Element().SetText("9800.00.0012")
	il3i.CreateElement(cacCommodityClassification).CreateElement(cbcItemClassificationCode).
		CreateAttr(lID, "CLASS").Element().SetText("003")

	il3.CreateElement(cacPrice).CreateElement(cbcPriceAmount).
		CreateAttr(cID, MYR).Element().SetText("2000.00")
	il3.CreateElement(cacItemPriceExtension).CreateElement(cbcAmount).
		CreateAttr(cID, MYR).Element().SetText("2000.00")

	invSav(doc, out)
}

// nolint:unused
func sdkSampleInvoiceConsolidated10(out string) {
	doc := etree.NewDocument()
	inv := doc.CreateElement("Invoice").CreateAttr("xmlns", urnInvoice2).Element().
		CreateAttr("xmlns:cac", urnCommonAggregateComponents2).Element().
		CreateAttr("xmlns:cbc", urnCommonBasicComponents2).Element()

	inv.CreateElement(cbcID).SetText("XML-INV12345")
	inv.CreateElement(cbcIssueDate).SetText("2024-07-23")
	inv.CreateElement(cbcIssueTime).SetText("00:40:00Z")
	inv.CreateElement(cbcInvoiceTypeCode).CreateAttr("listVersionID", "1.0").Element().SetText("01") // 1.1 requires signature!
	inv.CreateElement(cbcDocumentCurrencyCode).SetText(MYR)
	inv.CreateElement(cbcTaxCurrencyCode).SetText(MYR)

	asp := inv.CreateElement(cacAccountingSupplierParty)
	sp := asp.CreateElement(cacParty)

	sp.CreateElement(cbcIndustryClassificationCode).
		CreateAttr("name", cMSICSC[spMSICSC].Description).Element().SetText(spMSICSC)

	sp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, TIN).Element().SetText("Supplier's TIN")
	sp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, BRN).Element().SetText("Supplier's BRN")
	sp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, SST).Element().SetText("NA")
	sp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, TTX).Element().SetText("NA")

	spa := sp.CreateElement(cacPostalAddress)
	spa.CreateElement(cbcCityName).SetText(reSubWP.ReplaceAllString(cState[spCSC].State, ""))
	spa.CreateElement(cbcPostalZone).SetText("50480")
	spa.CreateElement(cbcCountrySubentityCode).SetText(spCSC)
	spa.CreateElement(cacAddressLine).CreateElement(cbcLine).SetText("Lot 66")
	spa.CreateElement(cacAddressLine).CreateElement(cbcLine).SetText("Bangunan Merdeka")
	spa.CreateElement(cacAddressLine).CreateElement(cbcLine).SetText("Persiaran Jaya")
	spa.CreateElement(cacCountry).CreateElement(cbcIdentificationCode).
		CreateAttr(lID, "ISO3166-1").Element().CreateAttr(lAID, "6").
		Element().SetText(cCountryRev[spCIC].Code)
	sp.CreateElement(cacPartyLegalEntity).CreateElement(cbcRegistrationName).SetText("Supplier's Name")
	sc := sp.CreateElement(cacContact)
	sc.CreateElement(cbcTelephone).SetText("+60123456789")
	sc.CreateElement(cbcElectronicMail).SetText("supplier@email.com")

	acp := inv.CreateElement(cacAccountingCustomerParty)
	cp := acp.CreateElement(cacParty)

	cp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, TIN).Element().SetText("EI00000000010")
	cp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, BRN).Element().SetText("NA")
	cp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, SST).Element().SetText("NA")
	cp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, TTX).Element().SetText("NA")

	cpa := cp.CreateElement(cacPostalAddress)
	cpa.CreateElement(cbcCityName).SetText("")
	cpa.CreateElement(cbcPostalZone).SetText("")
	cpa.CreateElement(cbcCountrySubentityCode).SetText("")
	cpa.CreateElement(cacAddressLine).CreateElement(cbcLine).SetText("NA")
	cpa.CreateElement(cacAddressLine).CreateElement(cbcLine).SetText("")
	cpa.CreateElement(cacAddressLine).CreateElement(cbcLine).SetText("")
	cpa.CreateElement(cacCountry).CreateElement(cbcIdentificationCode).
		CreateAttr(lID, "ISO3166-1").Element().CreateAttr(lAID, "6").
		Element().SetText("")
	cp.CreateElement(cacPartyLegalEntity).CreateElement(cbcRegistrationName).SetText("Consolidated Buyers")
	cc := cp.CreateElement(cacContact)
	cc.CreateElement(cbcTelephone).SetText("NA")
	cc.CreateElement(cbcElectronicMail).SetText("NA")

	tt := inv.CreateElement(cacTaxTotal)
	tt.CreateElement(cbcTaxAmount).CreateAttr(cID, MYR).Element().SetText("3000")

	ts1 := tt.CreateElement(cacTaxSubtotal)
	ts1.CreateElement(cbcTaxableAmount).CreateAttr(cID, MYR).Element().SetText("30000")
	ts1.CreateElement(cbcTaxAmount).CreateAttr(cID, MYR).Element().SetText("3000")
	ts1tc := ts1.CreateElement(cacTaxCategory)
	ts1tc.CreateElement(cbcID).SetText("01")
	ts1tc.CreateElement(cacTaxScheme).CreateElement(cbcID).
		CreateAttr(sID, UE5).Element().CreateAttr(sAID, "6").Element().SetText(OTH)

	lmt := inv.CreateElement(cacLegalMonetaryTotal)
	lmt.CreateElement(cbcLineExtensionAmount).CreateAttr(cID, MYR).Element().SetText("30000")
	lmt.CreateElement(cbcTaxExclusiveAmount).CreateAttr(cID, MYR).Element().SetText("30000")
	lmt.CreateElement(cbcTaxInclusiveAmount).CreateAttr(cID, MYR).Element().SetText("33000")
	lmt.CreateElement(cbcAllowanceTotalAmount).CreateAttr(cID, MYR).Element().SetText("0")
	lmt.CreateElement(cbcChargeTotalAmount).CreateAttr(cID, MYR).Element().SetText("0")
	lmt.CreateElement(cbcPayableRoundingAmount).CreateAttr(cID, MYR).Element().SetText("0")
	lmt.CreateElement(cbcPayableAmount).CreateAttr(cID, MYR).Element().SetText("33000")

	il1 := inv.CreateElement(cacInvoiceLine)
	il1.CreateElement(cbcID).SetText("1")
	il1.CreateElement(cbcInvoicedQuantity).CreateAttr("unitCode", "C62").Element().SetText("1")
	il1.CreateElement(cbcLineExtensionAmount).CreateAttr(cID, MYR).Element().SetText("10000")

	il1tt := il1.CreateElement(cacTaxTotal)
	il1tt.CreateElement(cbcTaxAmount).CreateAttr(cID, MYR).Element().SetText("1000")
	il1ttts := il1tt.CreateElement(cacTaxSubtotal)
	il1ttts.CreateElement(cbcTaxableAmount).CreateAttr(cID, MYR).Element().SetText("10000")
	il1ttts.CreateElement(cbcTaxAmount).CreateAttr(cID, MYR).Element().SetText("1000")
	il1ttts.CreateElement(cbcPercent).SetText("10.00")
	il1tttstc := il1ttts.CreateElement(cacTaxCategory)
	il1tttstc.CreateElement(cbcID).SetText("01")
	il1tttstc.CreateElement(cacTaxScheme).CreateElement(cbcID).
		CreateAttr(sID, UE5).Element().CreateAttr(sAID, "6").Element().SetText(OTH)

	il1i := il1.CreateElement(cacItem)
	il1i.CreateElement(cbcDescription).SetText("Receipt 001 - 100")
	il1i.CreateElement(cacOriginCountry).CreateElement(cbcIdentificationCode).SetText(cCountryRev["MALAYSIA"].Code)
	il1i.CreateElement(cacCommodityClassification).CreateElement(cbcItemClassificationCode).
		CreateAttr(lID, "PTC").Element().SetText("")
	il1i.CreateElement(cacCommodityClassification).CreateElement(cbcItemClassificationCode).
		CreateAttr(lID, "CLASS").Element().SetText("004")

	il1.CreateElement(cacPrice).CreateElement(cbcPriceAmount).
		CreateAttr(cID, MYR).Element().SetText("10000")
	il1.CreateElement(cacItemPriceExtension).CreateElement(cbcAmount).
		CreateAttr(cID, MYR).Element().SetText("10000")

	il2 := inv.CreateElement(cacInvoiceLine)
	il2.CreateElement(cbcID).SetText("2")
	il2.CreateElement(cbcInvoicedQuantity).CreateAttr("unitCode", "C62").Element().SetText("1")
	il2.CreateElement(cbcLineExtensionAmount).CreateAttr(cID, MYR).Element().SetText("20000")

	il2tt := il2.CreateElement(cacTaxTotal)
	il2tt.CreateElement(cbcTaxAmount).CreateAttr(cID, MYR).Element().SetText("2000")
	il2ttts := il2tt.CreateElement(cacTaxSubtotal)
	il2ttts.CreateElement(cbcTaxableAmount).CreateAttr(cID, MYR).Element().SetText("20000")
	il2ttts.CreateElement(cbcTaxAmount).CreateAttr(cID, MYR).Element().SetText("2000")
	il2ttts.CreateElement(cbcPercent).SetText("10.00")
	il2tttstc := il2ttts.CreateElement(cacTaxCategory)
	il2tttstc.CreateElement(cbcID).SetText("01")
	il2tttstc.CreateElement(cacTaxScheme).CreateElement(cbcID).
		CreateAttr(sID, UE5).Element().CreateAttr(sAID, "6").Element().SetText(OTH)

	il2i := il2.CreateElement(cacItem)
	il2i.CreateElement(cbcDescription).SetText("Receipt 101 - 200")
	il2i.CreateElement(cacOriginCountry).CreateElement(cbcIdentificationCode).SetText(cCountryRev["MALAYSIA"].Code)
	il2i.CreateElement(cacCommodityClassification).CreateElement(cbcItemClassificationCode).
		CreateAttr(lID, "PTC").Element().SetText("")
	il2i.CreateElement(cacCommodityClassification).CreateElement(cbcItemClassificationCode).
		CreateAttr(lID, "CLASS").Element().SetText("004") // ClassificationCodes.json

	il2.CreateElement(cacPrice).CreateElement(cbcPriceAmount).
		CreateAttr(cID, MYR).Element().SetText("20000")
	il2.CreateElement(cacItemPriceExtension).CreateElement(cbcAmount).
		CreateAttr(cID, MYR).Element().SetText("20000")

	invSav(doc, out)
}

// nolint:unused
func sdkSampleInvoiceForeignCurrency10(out string) {
	doc := etree.NewDocument()
	inv := doc.CreateElement("Invoice").CreateAttr("xmlns", urnInvoice2).Element().
		CreateAttr("xmlns:cac", urnCommonAggregateComponents2).Element().
		CreateAttr("xmlns:cbc", urnCommonBasicComponents2).Element()

	inv.CreateElement(cbcID).SetText("XML-INV12345")
	inv.CreateElement(cbcIssueDate).SetText("2024-07-23")
	inv.CreateElement(cbcIssueTime).SetText("00:40:00Z")
	inv.CreateElement(cbcInvoiceTypeCode).CreateAttr("listVersionID", "1.0").Element().SetText("01") // 1.1 requires signature!
	inv.CreateElement(cbcDocumentCurrencyCode).SetText(USD)
	inv.CreateElement(cbcTaxCurrencyCode).SetText(MYR)

	period := inv.CreateElement(cacInvoicePeriod)
	period.CreateElement(cbcStartDate).SetText("2024-07-01")
	period.CreateElement(cbcEndDate).SetText("2024-07-31")
	period.CreateElement(cbcDescription).SetText("Monthly")

	inv.CreateElement(cacBillingReference).CreateElement(cacAdditionalDocumentReference).
		CreateElement(cbcID).SetText("151891-1981")

	adf1 := inv.CreateElement(cacAdditionalDocumentReference)
	adf1.CreateElement(cbcID).SetText("L1")
	adf1.CreateElement(cbcDocumentType).SetText("CustomsImportForm")

	adf2 := inv.CreateElement(cacAdditionalDocumentReference)
	adf2.CreateElement(cbcID).SetText("FTA")
	adf2.CreateElement(cbcDocumentType).SetText("FreeTradeAgreement")
	adf2.CreateElement(cbcDocumentDescription).SetText("Sample Description")

	adf3 := inv.CreateElement(cacAdditionalDocumentReference)
	adf3.CreateElement(cbcID).SetText("L1")
	adf3.CreateElement(cbcDocumentType).SetText("K2")

	adf4 := inv.CreateElement(cacAdditionalDocumentReference)
	adf4.CreateElement(cbcID).SetText("L1")
	// adf4.CreateElement(cbcDocumentType).SetText("K2")

	asp := inv.CreateElement(cacAccountingSupplierParty)
	asp.CreateElement(cbcAdditionalAccountID).CreateAttr("schemeAgencyName", "CertEX").
		Element().SetText("CPT-CCN-W-211111-KL-000002")

	sp := asp.CreateElement(cacParty)

	sp.CreateElement(cbcIndustryClassificationCode).
		CreateAttr("name", cMSICSC[spMSICSC].Description).Element().SetText(spMSICSC)

	sp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, TIN).Element().SetText("Supplier's TIN")
	sp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, BRN).Element().SetText("Supplier's BRN")
	sp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, SST).Element().SetText("NA")
	sp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, TTX).Element().SetText("NA")

	spa := sp.CreateElement(cacPostalAddress)
	spa.CreateElement(cbcCityName).SetText(reSubWP.ReplaceAllString(cState[spCSC].State, ""))
	spa.CreateElement(cbcPostalZone).SetText("50480")
	spa.CreateElement(cbcCountrySubentityCode).SetText(spCSC)
	spa.CreateElement(cacAddressLine).CreateElement(cbcLine).SetText("Lot 66")
	spa.CreateElement(cacAddressLine).CreateElement(cbcLine).SetText("Bangunan Merdeka")
	spa.CreateElement(cacAddressLine).CreateElement(cbcLine).SetText("Persiaran Jaya")
	spa.CreateElement(cacCountry).CreateElement(cbcIdentificationCode).
		CreateAttr(lID, "ISO3166-1").Element().CreateAttr(lAID, "6").
		Element().SetText(cCountryRev[spCIC].Code)
	sp.CreateElement(cacPartyLegalEntity).CreateElement(cbcRegistrationName).SetText("Supplier's Name")
	sc := sp.CreateElement(cacContact)
	sc.CreateElement(cbcTelephone).SetText("+60123456789")
	sc.CreateElement(cbcElectronicMail).SetText("supplier@email.com")

	acp := inv.CreateElement(cacAccountingCustomerParty)
	cp := acp.CreateElement(cacParty)

	cp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, TIN).Element().SetText("Buyer's TIN")
	cp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, BRN).Element().SetText("Buyer's BRN")
	cp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, SST).Element().SetText("NA")
	cp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, TTX).Element().SetText("NA")

	cpa := cp.CreateElement(cacPostalAddress)
	cpa.CreateElement(cbcCityName).SetText(reSubWP.ReplaceAllString(cState[spCSC].State, ""))
	cpa.CreateElement(cbcPostalZone).SetText("50480")
	cpa.CreateElement(cbcCountrySubentityCode).SetText(spCSC)
	cpa.CreateElement(cacAddressLine).CreateElement(cbcLine).SetText("Lot 66")
	cpa.CreateElement(cacAddressLine).CreateElement(cbcLine).SetText("Bangunan Merdeka")
	cpa.CreateElement(cacAddressLine).CreateElement(cbcLine).SetText("Persiaran Jaya")
	cpa.CreateElement(cacCountry).CreateElement(cbcIdentificationCode).
		CreateAttr(lID, "ISO3166-1").Element().CreateAttr(lAID, "6").
		Element().SetText(cCountryRev["MALAYSIA"].Code)
	cp.CreateElement(cacPartyLegalEntity).CreateElement(cbcRegistrationName).SetText("Buyer's Name")
	cc := cp.CreateElement(cacContact)
	cc.CreateElement(cbcTelephone).SetText("+60123456780")
	cc.CreateElement(cbcElectronicMail).SetText("buyer@email.com")

	d := inv.CreateElement(cacDelivery)
	dp := d.CreateElement(cacDeliveryParty)

	dp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, TIN).Element().SetText("Recipient's TIN")
	dp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, BRN).Element().SetText("Recipient's BRN")

	dpa := dp.CreateElement(cacPostalAddress)
	dpa.CreateElement(cbcCityName).SetText(reSubWP.ReplaceAllString(cState[spCSC].State, ""))
	dpa.CreateElement(cbcPostalZone).SetText("50480")
	dpa.CreateElement(cbcCountrySubentityCode).SetText(spCSC)
	dpa.CreateElement(cacAddressLine).CreateElement(cbcLine).SetText("Lot 66")
	dpa.CreateElement(cacAddressLine).CreateElement(cbcLine).SetText("Bangunan Merdeka")
	dpa.CreateElement(cacAddressLine).CreateElement(cbcLine).SetText("Persiaran Jaya")
	dpa.CreateElement(cacCountry).CreateElement(cbcIdentificationCode).
		CreateAttr(lID, "ISO3166-1").Element().CreateAttr(lAID, "6").
		Element().SetText(cCountryRev["MALAYSIA"].Code)
	dp.CreateElement(cacPartyLegalEntity).CreateElement(cbcRegistrationName).SetText("Recipient's Name")

	ds := d.CreateElement(cacShipment)
	ds.CreateElement(cbcID).SetText("1234")
	dsfac := ds.CreateElement(cacFreightAllowanceCharge)
	dsfac.CreateElement(cbcChargeIndicator).SetText("true")
	dsfac.CreateElement(cbcAllowanceChargeReason).SetText("Service charge")
	dsfac.CreateElement(cbcAmount).CreateAttr(cID, USD).Element().SetText("100")

	pm := inv.CreateElement(cacPaymentMeans)
	pm.CreateElement(cbcPaymentMeansCode).SetText("01") // PaymentMethods.json
	pm.CreateElement(cacPayeeFinancialAccount).CreateElement(cbcID).SetText("1234567890")

	inv.CreateElement(cacPaymentTerms).CreateElement(cbcNote).SetText("Payment method is cash")

	pp := inv.CreateElement(cacPrepaidPayment)
	pp.CreateElement(cbcID).SetText("E12345678912")
	pp.CreateElement(cbcPaidAmount).CreateAttr(cID, USD).Element().SetText("1.00")
	pp.CreateElement(cbcPaidDate).SetText("2024-07-23")
	pp.CreateElement(cbcPaidTime).SetText("00:30:00Z")

	ac1 := inv.CreateElement(cacAllowanceCharge)
	ac1.CreateElement(cbcChargeIndicator).SetText("false")
	ac1.CreateElement(cbcAllowanceChargeReason).SetText("Sample Description")
	ac1.CreateElement(cbcAmount).CreateAttr(cID, USD).Element().SetText("100")

	ac2 := inv.CreateElement(cacAllowanceCharge)
	ac2.CreateElement(cbcChargeIndicator).SetText("true")
	ac2.CreateElement(cbcAllowanceChargeReason).SetText("Service charge")
	ac2.CreateElement(cbcAmount).CreateAttr(cID, USD).Element().SetText("100")

	tt := inv.CreateElement(cacTaxTotal)
	tt.CreateElement(cbcTaxAmount).CreateAttr(cID, MYR).Element().SetText("87.63")

	ts1 := tt.CreateElement(cacTaxSubtotal)
	ts1.CreateElement(cbcTaxableAmount).CreateAttr(cID, USD).Element().SetText("87.63")
	ts1.CreateElement(cbcTaxAmount).CreateAttr(cID, MYR).Element().SetText("87.63")
	ts1tc := ts1.CreateElement(cacTaxCategory)
	ts1tc.CreateElement(cbcID).SetText("01")
	ts1tc.CreateElement(cacTaxScheme).CreateElement(cbcID).
		CreateAttr(sID, UE5).Element().CreateAttr(sAID, "6").Element().SetText(OTH)

	lmt := inv.CreateElement(cacLegalMonetaryTotal)
	lmt.CreateElement(cbcLineExtensionAmount).CreateAttr(cID, USD).Element().SetText("1436.50")
	lmt.CreateElement(cbcTaxExclusiveAmount).CreateAttr(cID, USD).Element().SetText("1436.50")
	lmt.CreateElement(cbcTaxInclusiveAmount).CreateAttr(cID, USD).Element().SetText("1436.50")
	lmt.CreateElement(cbcAllowanceTotalAmount).CreateAttr(cID, USD).Element().SetText("1436.50")
	lmt.CreateElement(cbcChargeTotalAmount).CreateAttr(cID, USD).Element().SetText("1436.50")
	lmt.CreateElement(cbcPayableRoundingAmount).CreateAttr(cID, USD).Element().SetText("0.30")
	lmt.CreateElement(cbcPayableAmount).CreateAttr(cID, USD).Element().SetText("1436.50")

	il1 := inv.CreateElement(cacInvoiceLine)
	il1.CreateElement(cbcID).SetText("1234")
	il1.CreateElement(cbcInvoicedQuantity).CreateAttr("unitCode", "C62").Element().SetText("1")
	il1.CreateElement(cbcLineExtensionAmount).CreateAttr(cID, USD).Element().SetText("1436.50")

	il1ac1 := il1.CreateElement(cacAllowanceCharge)
	il1ac1.CreateElement(cbcChargeIndicator).SetText("false")
	il1ac1.CreateElement(cbcAllowanceChargeReason).SetText("Sample Description")
	il1ac1.CreateElement(cbcMultiplierFactorNumeric).SetText("0.15")
	il1ac1.CreateElement(cbcAmount).CreateAttr(cID, USD).Element().SetText("100")

	il1ac2 := il1.CreateElement(cacAllowanceCharge)
	il1ac2.CreateElement(cbcChargeIndicator).SetText("true")
	il1ac2.CreateElement(cbcAllowanceChargeReason).SetText("Sample Description")
	il1ac2.CreateElement(cbcMultiplierFactorNumeric).SetText("0.10")
	il1ac2.CreateElement(cbcAmount).CreateAttr(cID, USD).Element().SetText("100")

	il1ter := il1.CreateElement(cacTaxExchangeRate)
	il1ter.CreateElement(cbcSourceCurrencyCode).SetText(USD)
	il1ter.CreateElement(cbcTargetCurrencyCode).SetText(MYR)
	il1ter.CreateElement(cbcCalculationRate).SetText("4.72")

	il1tt := il1.CreateElement(cacTaxTotal)
	il1tt.CreateElement(cbcTaxAmount).CreateAttr(cID, MYR).Element().SetText("0")
	il1ttts := il1tt.CreateElement(cacTaxSubtotal)
	il1ttts.CreateElement(cbcTaxableAmount).CreateAttr(cID, USD).Element().SetText("1460.50")
	il1ttts.CreateElement(cbcTaxAmount).CreateAttr(cID, MYR).Element().SetText("0")
	il1ttts.CreateElement(cbcPercent).SetText("6.00")
	il1tttstc := il1ttts.CreateElement(cacTaxCategory)
	il1tttstc.CreateElement(cbcID).SetText("E")
	il1tttstc.CreateElement(cbcTaxExemptionReason).SetText("Exempt New Means of Transport")
	il1tttstc.CreateElement(cacTaxScheme).CreateElement(cbcID).
		CreateAttr(sID, UE5).Element().CreateAttr(sAID, "6").Element().SetText(OTH)

	il1i := il1.CreateElement(cacItem)
	il1i.CreateElement(cbcDescription).SetText("Laptop Peripherals")
	il1i.CreateElement(cacOriginCountry).CreateElement(cbcIdentificationCode).SetText(cCountryRev["MALAYSIA"].Code)
	il1i.CreateElement(cacCommodityClassification).CreateElement(cbcItemClassificationCode).
		CreateAttr(lID, "PTC").Element().SetText("9800.00.0010")
	il1i.CreateElement(cacCommodityClassification).CreateElement(cbcItemClassificationCode).
		CreateAttr(lID, "CLASS").Element().SetText("003")

	il1.CreateElement(cacPrice).CreateElement(cbcPriceAmount).
		CreateAttr(cID, USD).Element().SetText("17")
	il1.CreateElement(cacItemPriceExtension).CreateElement(cbcAmount).
		CreateAttr(cID, USD).Element().SetText("100")

	invSav(doc, out)
}

// nolint: unused
func sdkSampleCreditNote10(out string) {
	doc := etree.NewDocument()
	inv := doc.CreateElement("Invoice").CreateAttr("xmlns", urnInvoice2).Element().
		CreateAttr("xmlns:cac", urnCommonAggregateComponents2).Element().
		CreateAttr("xmlns:cbc", urnCommonBasicComponents2).Element()
	inv.CreateElement(cbcID).SetText("XML-CN12345")
	inv.CreateElement(cbcIssueDate).SetText("2024-07-23")
	inv.CreateElement(cbcIssueTime).SetText("00:40:00Z")
	inv.CreateElement(cbcInvoiceTypeCode).CreateAttr("listVersionID", "1.0").Element().SetText(iTCCreditNote) // 1.1 requires signature!
	inv.CreateElement(cbcDocumentCurrencyCode).SetText(MYR)
	inv.CreateElement(cbcTaxCurrencyCode).SetText(MYR)

	period := inv.CreateElement(cacInvoicePeriod)
	period.CreateElement(cbcStartDate).SetText("2024-07-01")
	period.CreateElement(cbcEndDate).SetText("2024-07-31")
	period.CreateElement(cbcDescription).SetText("Monthly")

	idr := inv.CreateElement(cacBillingReference).CreateElement(cacInvoiceDocumentReference)
	idr.CreateElement(cbcID).SetText("XML-INV12345")
	idr.CreateElement(cbcUUID).SetText("Reference Invoice UUID")

	inv.CreateElement(cacBillingReference).CreateElement(cacAdditionalDocumentReference).CreateElement(cbcID).SetText("151891-1981")

	adr1 := inv.CreateElement(cacAdditionalDocumentReference)
	adr1.CreateElement(cbcID).SetText("L1")
	adr1.CreateElement(cbcDocumentType).SetText("CustomsImportForm")

	adr2 := inv.CreateElement(cacAdditionalDocumentReference)
	adr2.CreateElement(cbcID).SetText("FTA")
	adr2.CreateElement(cbcDocumentType).SetText("FreeTradeAgreement")
	adr2.CreateElement(cbcDocumentDescription).SetText("Sample Description")

	adr3 := inv.CreateElement(cacAdditionalDocumentReference)
	adr3.CreateElement(cbcID).SetText("L1")
	adr3.CreateElement(cbcDocumentType).SetText("K2")

	adr4 := inv.CreateElement(cacAdditionalDocumentReference)
	adr4.CreateElement(cbcID).SetText("L1")

	asp := inv.CreateElement(cacAccountingSupplierParty)
	asp.CreateElement(cbcAdditionalAccountID).CreateAttr("schemeAgencyName", "CertEX").
		Element().SetText("CPT-CCN-W-211111-KL-000002")

	sp := asp.CreateElement(cacParty)
	sp.CreateElement(cbcIndustryClassificationCode).
		CreateAttr("name", cMSICSC[spMSICSC].Description).Element().SetText(spMSICSC)

	sp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, TIN).Element().SetText("Supplier's TIN")
	sp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, BRN).Element().SetText("Supplier's BRN")
	sp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, SST).Element().SetText("NA")
	sp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, TTX).Element().SetText("NA")

	spa := sp.CreateElement(cacPostalAddress)
	spa.CreateElement(cbcCityName).SetText(reSubWP.ReplaceAllString(cState[spCSC].State, ""))
	spa.CreateElement(cbcPostalZone).SetText("50480")
	spa.CreateElement(cbcCountrySubentityCode).SetText(spCSC)
	spa.CreateElement(cacAddressLine).CreateElement(cbcLine).SetText("Lot 66")
	spa.CreateElement(cacAddressLine).CreateElement(cbcLine).SetText("Bangunan Merdeka")
	spa.CreateElement(cacAddressLine).CreateElement(cbcLine).SetText("Persiaran Jaya")
	spa.CreateElement(cacCountry).CreateElement(cbcIdentificationCode).
		CreateAttr(lID, "ISO3166-1").Element().CreateAttr(lAID, "6").
		Element().SetText(cCountryRev[spCIC].Code)
	sp.CreateElement(cacPartyLegalEntity).CreateElement(cbcRegistrationName).SetText("Supplier's Name")
	sc := sp.CreateElement(cacContact)
	sc.CreateElement(cbcTelephone).SetText("+60123456789")
	sc.CreateElement(cbcElectronicMail).SetText("supplier@email.com")

	acp := inv.CreateElement(cacAccountingCustomerParty)
	cp := acp.CreateElement(cacParty)

	cp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, TIN).Element().SetText("Buyer's TIN")
	cp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, BRN).Element().SetText("Buyer's BRN")
	cp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, SST).Element().SetText("NA")
	cp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, TTX).Element().SetText("NA")

	cpa := cp.CreateElement(cacPostalAddress)
	cpa.CreateElement(cbcCityName).SetText(reSubWP.ReplaceAllString(cState[spCSC].State, ""))
	cpa.CreateElement(cbcPostalZone).SetText("50480")
	cpa.CreateElement(cbcCountrySubentityCode).SetText(spCSC)
	cpa.CreateElement(cacAddressLine).CreateElement(cbcLine).SetText("Lot 66")
	cpa.CreateElement(cacAddressLine).CreateElement(cbcLine).SetText("Bangunan Merdeka")
	cpa.CreateElement(cacAddressLine).CreateElement(cbcLine).SetText("Persiaran Jaya")
	cpa.CreateElement(cacCountry).CreateElement(cbcIdentificationCode).
		CreateAttr(lID, "ISO3166-1").Element().CreateAttr(lAID, "6").
		Element().SetText(cCountryRev["MALAYSIA"].Code)
	cp.CreateElement(cacPartyLegalEntity).CreateElement(cbcRegistrationName).SetText("Buyer's Name")
	cc := cp.CreateElement(cacContact)
	cc.CreateElement(cbcTelephone).SetText("+60123456780")
	cc.CreateElement(cbcElectronicMail).SetText("buyer@email.com")

	d := inv.CreateElement(cacDelivery)
	dp := d.CreateElement(cacDeliveryParty)

	dp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, TIN).Element().SetText("Recipient's TIN")
	dp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, BRN).Element().SetText("Recipient's BRN")

	dpa := dp.CreateElement(cacPostalAddress)
	dpa.CreateElement(cbcCityName).SetText(reSubWP.ReplaceAllString(cState[spCSC].State, ""))
	dpa.CreateElement(cbcPostalZone).SetText("50480")
	dpa.CreateElement(cbcCountrySubentityCode).SetText(spCSC)
	dpa.CreateElement(cacAddressLine).CreateElement(cbcLine).SetText("Lot 66")
	dpa.CreateElement(cacAddressLine).CreateElement(cbcLine).SetText("Bangunan Merdeka")
	dpa.CreateElement(cacAddressLine).CreateElement(cbcLine).SetText("Persiaran Jaya")
	dpa.CreateElement(cacCountry).CreateElement(cbcIdentificationCode).
		CreateAttr(lID, "ISO3166-1").Element().CreateAttr(lAID, "6").
		Element().SetText(cCountryRev["MALAYSIA"].Code)
	dp.CreateElement(cacPartyLegalEntity).CreateElement(cbcRegistrationName).SetText("Recipient's Name")

	ds := d.CreateElement(cacShipment)
	ds.CreateElement(cbcID).SetText("1234")
	dsfac := ds.CreateElement(cacFreightAllowanceCharge)
	dsfac.CreateElement(cbcChargeIndicator).SetText("true")
	dsfac.CreateElement(cbcAllowanceChargeReason).SetText("Service charge")
	dsfac.CreateElement(cbcAmount).CreateAttr(cID, MYR).Element().SetText("100")

	pm := inv.CreateElement(cacPaymentMeans)
	pm.CreateElement(cbcPaymentMeansCode).SetText("01") // PaymentMethods.json
	pm.CreateElement(cacPayeeFinancialAccount).CreateElement(cbcID).SetText("1234567890")

	inv.CreateElement(cacPaymentTerms).CreateElement(cbcNote).SetText("Payment method is cash")

	pp := inv.CreateElement(cacPrepaidPayment)
	pp.CreateElement(cbcID).SetText("E12345678912")
	pp.CreateElement(cbcPaidAmount).CreateAttr(cID, MYR).Element().SetText("1.00")
	pp.CreateElement(cbcPaidDate).SetText("2024-07-23")
	pp.CreateElement(cbcPaidTime).SetText("00:30:00Z")

	ac1 := inv.CreateElement(cacAllowanceCharge)
	ac1.CreateElement(cbcChargeIndicator).SetText("false")
	ac1.CreateElement(cbcAllowanceChargeReason).SetText("Sample Description")
	ac1.CreateElement(cbcAmount).CreateAttr(cID, MYR).Element().SetText("100")

	ac2 := inv.CreateElement(cacAllowanceCharge)
	ac2.CreateElement(cbcChargeIndicator).SetText("true")
	ac2.CreateElement(cbcAllowanceChargeReason).SetText("Service charge")
	ac2.CreateElement(cbcAmount).CreateAttr(cID, MYR).Element().SetText("100")

	tt := inv.CreateElement(cacTaxTotal)
	tt.CreateElement(cbcTaxAmount).CreateAttr(cID, MYR).Element().SetText("87.63")

	ts := tt.CreateElement(cacTaxSubtotal)
	ts.CreateElement(cbcTaxableAmount).CreateAttr(cID, MYR).Element().SetText("87.63")
	ts.CreateElement(cbcTaxAmount).CreateAttr(cID, MYR).Element().SetText("87.63")
	tstc := ts.CreateElement(cacTaxCategory)
	tstc.CreateElement(cbcID).SetText("01")
	tstc.CreateElement(cacTaxScheme).CreateElement(cbcID).
		CreateAttr(sID, UE5).Element().CreateAttr(sAID, "6").Element().SetText(OTH)

	lmt := inv.CreateElement(cacLegalMonetaryTotal)
	lmt.CreateElement(cbcLineExtensionAmount).CreateAttr(cID, MYR).Element().SetText("1436.50")
	lmt.CreateElement(cbcTaxExclusiveAmount).CreateAttr(cID, MYR).Element().SetText("1436.50")
	lmt.CreateElement(cbcTaxInclusiveAmount).CreateAttr(cID, MYR).Element().SetText("1436.50")
	lmt.CreateElement(cbcAllowanceTotalAmount).CreateAttr(cID, MYR).Element().SetText("1436.50")
	lmt.CreateElement(cbcChargeTotalAmount).CreateAttr(cID, MYR).Element().SetText("1436.50")
	lmt.CreateElement(cbcPayableRoundingAmount).CreateAttr(cID, MYR).Element().SetText("0.30")
	lmt.CreateElement(cbcPayableAmount).CreateAttr(cID, MYR).Element().SetText("1436.50")

	il := inv.CreateElement(cacInvoiceLine)
	il.CreateElement(cbcID).SetText("1234")
	il.CreateElement(cbcInvoicedQuantity).CreateAttr("unitCode", "C62").Element().SetText("1")
	il.CreateElement(cbcLineExtensionAmount).CreateAttr(cID, MYR).Element().SetText("1436.50")

	ilac1 := il.CreateElement(cacAllowanceCharge)
	ilac1.CreateElement(cbcChargeIndicator).SetText("false")
	ilac1.CreateElement(cbcAllowanceChargeReason).SetText("Sample Description")
	ilac1.CreateElement(cbcMultiplierFactorNumeric).SetText("0.15")
	ilac1.CreateElement(cbcAmount).CreateAttr(cID, MYR).Element().SetText("100")

	ilac2 := il.CreateElement(cacAllowanceCharge)
	ilac2.CreateElement(cbcChargeIndicator).SetText("true")
	ilac2.CreateElement(cbcAllowanceChargeReason).SetText("Sample Description")
	ilac2.CreateElement(cbcMultiplierFactorNumeric).SetText("0.10")
	ilac2.CreateElement(cbcAmount).CreateAttr(cID, MYR).Element().SetText("100")

	iltt := il.CreateElement(cacTaxTotal)
	iltt.CreateElement(cbcTaxAmount).CreateAttr(cID, MYR).Element().SetText("0")
	ilttts := iltt.CreateElement(cacTaxSubtotal)
	ilttts.CreateElement(cbcTaxableAmount).CreateAttr(cID, MYR).Element().SetText("1460.50")
	ilttts.CreateElement(cbcTaxAmount).CreateAttr(cID, MYR).Element().SetText("0")
	ilttts.CreateElement(cbcPercent).SetText("6.00")
	ittttstc := ilttts.CreateElement(cacTaxCategory)
	ittttstc.CreateElement(cbcID).SetText("E")
	ittttstc.CreateElement(cbcTaxExemptionReason).SetText("Exempt New Means of Transport")
	ittttstc.CreateElement(cacTaxScheme).CreateElement(cbcID).
		CreateAttr(sID, UE5).Element().CreateAttr(sAID, "6").Element().SetText(OTH)

	ili := il.CreateElement(cacItem)
	ili.CreateElement(cbcDescription).SetText("Laptop Peripherals")
	ili.CreateElement(cacOriginCountry).CreateElement(cbcIdentificationCode).SetText(cCountryRev["MALAYSIA"].Code)
	ili.CreateElement(cacCommodityClassification).CreateElement(cbcItemClassificationCode).
		CreateAttr(lID, "PTC").Element().SetText("9800.00.0010")
	ili.CreateElement(cacCommodityClassification).CreateElement(cbcItemClassificationCode).
		CreateAttr(lID, "CLASS").Element().SetText("003")

	il.CreateElement(cacPrice).CreateElement(cbcPriceAmount).CreateAttr(cID, MYR).
		Element().SetText("17")
	il.CreateElement(cacItemPriceExtension).CreateElement(cbcAmount).CreateAttr(cID, MYR).
		Element().SetText("100")

	invSav(doc, out)
}

// nolint: unused
func sdkSampleDebitNote10(out string) {
	doc := etree.NewDocument()
	inv := doc.CreateElement("Invoice").CreateAttr("xmlns", urnInvoice2).Element().
		CreateAttr("xmlns:cac", urnCommonAggregateComponents2).Element().
		CreateAttr("xmlns:cbc", urnCommonBasicComponents2).Element()
	inv.CreateElement(cbcID).SetText("XML-DN12345")
	inv.CreateElement(cbcIssueDate).SetText("2024-07-23")
	inv.CreateElement(cbcIssueTime).SetText("00:40:00Z")
	inv.CreateElement(cbcInvoiceTypeCode).CreateAttr("listVersionID", "1.0").Element().SetText(iTCDebitNote) // 1.1 requires signature!
	inv.CreateElement(cbcDocumentCurrencyCode).SetText(MYR)
	inv.CreateElement(cbcTaxCurrencyCode).SetText(MYR)

	period := inv.CreateElement(cacInvoicePeriod)
	period.CreateElement(cbcStartDate).SetText("2024-07-01")
	period.CreateElement(cbcEndDate).SetText("2024-07-31")
	period.CreateElement(cbcDescription).SetText("Monthly")

	idr := inv.CreateElement(cacBillingReference).CreateElement(cacInvoiceDocumentReference)
	idr.CreateElement(cbcID).SetText("XML-INV12345")
	idr.CreateElement(cbcUUID).SetText("Reference Invoice UUID")

	inv.CreateElement(cacBillingReference).CreateElement(cacAdditionalDocumentReference).CreateElement(cbcID).SetText("151891-1981")

	adr1 := inv.CreateElement(cacAdditionalDocumentReference)
	adr1.CreateElement(cbcID).SetText("L1")
	adr1.CreateElement(cbcDocumentType).SetText("CustomsImportForm")

	adr2 := inv.CreateElement(cacAdditionalDocumentReference)
	adr2.CreateElement(cbcID).SetText("FTA")
	adr2.CreateElement(cbcDocumentType).SetText("FreeTradeAgreement")
	adr2.CreateElement(cbcDocumentDescription).SetText("Sample Description")

	adr3 := inv.CreateElement(cacAdditionalDocumentReference)
	adr3.CreateElement(cbcID).SetText("L1")
	adr3.CreateElement(cbcDocumentType).SetText("K2")

	adr4 := inv.CreateElement(cacAdditionalDocumentReference)
	adr4.CreateElement(cbcID).SetText("L1")

	asp := inv.CreateElement(cacAccountingSupplierParty)
	asp.CreateElement(cbcAdditionalAccountID).CreateAttr("schemeAgencyName", "CertEX").
		Element().SetText("CPT-CCN-W-211111-KL-000002")

	sp := asp.CreateElement(cacParty)
	sp.CreateElement(cbcIndustryClassificationCode).
		CreateAttr("name", cMSICSC[spMSICSC].Description).Element().SetText(spMSICSC)

	sp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, TIN).Element().SetText("Supplier's TIN")
	sp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, BRN).Element().SetText("Supplier's BRN")
	sp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, SST).Element().SetText("NA")
	sp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, TTX).Element().SetText("NA")

	spa := sp.CreateElement(cacPostalAddress)
	spa.CreateElement(cbcCityName).SetText(reSubWP.ReplaceAllString(cState[spCSC].State, ""))
	spa.CreateElement(cbcPostalZone).SetText("50480")
	spa.CreateElement(cbcCountrySubentityCode).SetText(spCSC)
	spa.CreateElement(cacAddressLine).CreateElement(cbcLine).SetText("Lot 66")
	spa.CreateElement(cacAddressLine).CreateElement(cbcLine).SetText("Bangunan Merdeka")
	spa.CreateElement(cacAddressLine).CreateElement(cbcLine).SetText("Persiaran Jaya")
	spa.CreateElement(cacCountry).CreateElement(cbcIdentificationCode).
		CreateAttr(lID, "ISO3166-1").Element().CreateAttr(lAID, "6").
		Element().SetText(cCountryRev[spCIC].Code)
	sp.CreateElement(cacPartyLegalEntity).CreateElement(cbcRegistrationName).SetText("Supplier's Name")
	sc := sp.CreateElement(cacContact)
	sc.CreateElement(cbcTelephone).SetText("+60123456789")
	sc.CreateElement(cbcElectronicMail).SetText("supplier@email.com")

	acp := inv.CreateElement(cacAccountingCustomerParty)
	cp := acp.CreateElement(cacParty)

	cp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, TIN).Element().SetText("Buyer's TIN")
	cp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, BRN).Element().SetText("Buyer's BRN")
	cp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, SST).Element().SetText("NA")
	cp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, TTX).Element().SetText("NA")

	cpa := cp.CreateElement(cacPostalAddress)
	cpa.CreateElement(cbcCityName).SetText(reSubWP.ReplaceAllString(cState[spCSC].State, ""))
	cpa.CreateElement(cbcPostalZone).SetText("50480")
	cpa.CreateElement(cbcCountrySubentityCode).SetText(spCSC)
	cpa.CreateElement(cacAddressLine).CreateElement(cbcLine).SetText("Lot 66")
	cpa.CreateElement(cacAddressLine).CreateElement(cbcLine).SetText("Bangunan Merdeka")
	cpa.CreateElement(cacAddressLine).CreateElement(cbcLine).SetText("Persiaran Jaya")
	cpa.CreateElement(cacCountry).CreateElement(cbcIdentificationCode).
		CreateAttr(lID, "ISO3166-1").Element().CreateAttr(lAID, "6").
		Element().SetText(cCountryRev["MALAYSIA"].Code)
	cp.CreateElement(cacPartyLegalEntity).CreateElement(cbcRegistrationName).SetText("Buyer's Name")
	cc := cp.CreateElement(cacContact)
	cc.CreateElement(cbcTelephone).SetText("+60123456780")
	cc.CreateElement(cbcElectronicMail).SetText("buyer@email.com")

	d := inv.CreateElement(cacDelivery)
	dp := d.CreateElement(cacDeliveryParty)

	dp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, TIN).Element().SetText("Recipient's TIN")
	dp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, BRN).Element().SetText("Recipient's BRN")

	dpa := dp.CreateElement(cacPostalAddress)
	dpa.CreateElement(cbcCityName).SetText(reSubWP.ReplaceAllString(cState[spCSC].State, ""))
	dpa.CreateElement(cbcPostalZone).SetText("50480")
	dpa.CreateElement(cbcCountrySubentityCode).SetText(spCSC)
	dpa.CreateElement(cacAddressLine).CreateElement(cbcLine).SetText("Lot 66")
	dpa.CreateElement(cacAddressLine).CreateElement(cbcLine).SetText("Bangunan Merdeka")
	dpa.CreateElement(cacAddressLine).CreateElement(cbcLine).SetText("Persiaran Jaya")
	dpa.CreateElement(cacCountry).CreateElement(cbcIdentificationCode).
		CreateAttr(lID, "ISO3166-1").Element().CreateAttr(lAID, "6").
		Element().SetText(cCountryRev["MALAYSIA"].Code)
	dp.CreateElement(cacPartyLegalEntity).CreateElement(cbcRegistrationName).SetText("Recipient's Name")

	ds := d.CreateElement(cacShipment)
	ds.CreateElement(cbcID).SetText("1234")
	dsfac := ds.CreateElement(cacFreightAllowanceCharge)
	dsfac.CreateElement(cbcChargeIndicator).SetText("true")
	dsfac.CreateElement(cbcAllowanceChargeReason).SetText("Service charge")
	dsfac.CreateElement(cbcAmount).CreateAttr(cID, MYR).Element().SetText("100")

	pm := inv.CreateElement(cacPaymentMeans)
	pm.CreateElement(cbcPaymentMeansCode).SetText("01") // PaymentMethods.json
	pm.CreateElement(cacPayeeFinancialAccount).CreateElement(cbcID).SetText("1234567890")

	inv.CreateElement(cacPaymentTerms).CreateElement(cbcNote).SetText("Payment method is cash")

	pp := inv.CreateElement(cacPrepaidPayment)
	pp.CreateElement(cbcID).SetText("E12345678912")
	pp.CreateElement(cbcPaidAmount).CreateAttr(cID, MYR).Element().SetText("1.00")
	pp.CreateElement(cbcPaidDate).SetText("2024-07-23")
	pp.CreateElement(cbcPaidTime).SetText("00:30:00Z")

	ac1 := inv.CreateElement(cacAllowanceCharge)
	ac1.CreateElement(cbcChargeIndicator).SetText("false")
	ac1.CreateElement(cbcAllowanceChargeReason).SetText("Sample Description")
	ac1.CreateElement(cbcAmount).CreateAttr(cID, MYR).Element().SetText("100")

	ac2 := inv.CreateElement(cacAllowanceCharge)
	ac2.CreateElement(cbcChargeIndicator).SetText("true")
	ac2.CreateElement(cbcAllowanceChargeReason).SetText("Service charge")
	ac2.CreateElement(cbcAmount).CreateAttr(cID, MYR).Element().SetText("100")

	tt := inv.CreateElement(cacTaxTotal)
	tt.CreateElement(cbcTaxAmount).CreateAttr(cID, MYR).Element().SetText("87.63")

	ts := tt.CreateElement(cacTaxSubtotal)
	ts.CreateElement(cbcTaxableAmount).CreateAttr(cID, MYR).Element().SetText("87.63")
	ts.CreateElement(cbcTaxAmount).CreateAttr(cID, MYR).Element().SetText("87.63")
	tstc := ts.CreateElement(cacTaxCategory)
	tstc.CreateElement(cbcID).SetText("01")
	tstc.CreateElement(cacTaxScheme).CreateElement(cbcID).
		CreateAttr(sID, UE5).Element().CreateAttr(sAID, "6").Element().SetText(OTH)

	lmt := inv.CreateElement(cacLegalMonetaryTotal)
	lmt.CreateElement(cbcLineExtensionAmount).CreateAttr(cID, MYR).Element().SetText("1436.50")
	lmt.CreateElement(cbcTaxExclusiveAmount).CreateAttr(cID, MYR).Element().SetText("1436.50")
	lmt.CreateElement(cbcTaxInclusiveAmount).CreateAttr(cID, MYR).Element().SetText("1436.50")
	lmt.CreateElement(cbcAllowanceTotalAmount).CreateAttr(cID, MYR).Element().SetText("1436.50")
	lmt.CreateElement(cbcChargeTotalAmount).CreateAttr(cID, MYR).Element().SetText("1436.50")
	lmt.CreateElement(cbcPayableRoundingAmount).CreateAttr(cID, MYR).Element().SetText("0.30")
	lmt.CreateElement(cbcPayableAmount).CreateAttr(cID, MYR).Element().SetText("1436.50")

	il := inv.CreateElement(cacInvoiceLine)
	il.CreateElement(cbcID).SetText("1234")
	il.CreateElement(cbcInvoicedQuantity).CreateAttr("unitCode", "C62").Element().SetText("1")
	il.CreateElement(cbcLineExtensionAmount).CreateAttr(cID, MYR).Element().SetText("1436.50")

	ilac1 := il.CreateElement(cacAllowanceCharge)
	ilac1.CreateElement(cbcChargeIndicator).SetText("false")
	ilac1.CreateElement(cbcAllowanceChargeReason).SetText("Sample Description")
	ilac1.CreateElement(cbcMultiplierFactorNumeric).SetText("0.15")
	ilac1.CreateElement(cbcAmount).CreateAttr(cID, MYR).Element().SetText("100")

	ilac2 := il.CreateElement(cacAllowanceCharge)
	ilac2.CreateElement(cbcChargeIndicator).SetText("true")
	ilac2.CreateElement(cbcAllowanceChargeReason).SetText("Sample Description")
	ilac2.CreateElement(cbcMultiplierFactorNumeric).SetText("0.10")
	ilac2.CreateElement(cbcAmount).CreateAttr(cID, MYR).Element().SetText("100")

	iltt := il.CreateElement(cacTaxTotal)
	iltt.CreateElement(cbcTaxAmount).CreateAttr(cID, MYR).Element().SetText("0")
	ilttts := iltt.CreateElement(cacTaxSubtotal)
	ilttts.CreateElement(cbcTaxableAmount).CreateAttr(cID, MYR).Element().SetText("1460.50")
	ilttts.CreateElement(cbcTaxAmount).CreateAttr(cID, MYR).Element().SetText("0")
	ilttts.CreateElement(cbcPercent).SetText("6.00")
	ittttstc := ilttts.CreateElement(cacTaxCategory)
	ittttstc.CreateElement(cbcID).SetText("E")
	ittttstc.CreateElement(cbcTaxExemptionReason).SetText("Exempt New Means of Transport")
	ittttstc.CreateElement(cacTaxScheme).CreateElement(cbcID).
		CreateAttr(sID, UE5).Element().CreateAttr(sAID, "6").Element().SetText(OTH)

	ili := il.CreateElement(cacItem)
	ili.CreateElement(cbcDescription).SetText("Laptop Peripherals")
	ili.CreateElement(cacOriginCountry).CreateElement(cbcIdentificationCode).SetText(cCountryRev["MALAYSIA"].Code)
	ili.CreateElement(cacCommodityClassification).CreateElement(cbcItemClassificationCode).
		CreateAttr(lID, "PTC").Element().SetText("9800.00.0010")
	ili.CreateElement(cacCommodityClassification).CreateElement(cbcItemClassificationCode).
		CreateAttr(lID, "CLASS").Element().SetText("003")

	il.CreateElement(cacPrice).CreateElement(cbcPriceAmount).CreateAttr(cID, MYR).
		Element().SetText("17")
	il.CreateElement(cacItemPriceExtension).CreateElement(cbcAmount).CreateAttr(cID, MYR).
		Element().SetText("100")

	invSav(doc, out)
}

// nolint: unused
func sdkSampleRefundNote10(out string) {
	doc := etree.NewDocument()
	inv := doc.CreateElement("Invoice").CreateAttr("xmlns", urnInvoice2).Element().
		CreateAttr("xmlns:cac", urnCommonAggregateComponents2).Element().
		CreateAttr("xmlns:cbc", urnCommonBasicComponents2).Element()
	inv.CreateElement(cbcID).SetText("XML-RN12345")
	inv.CreateElement(cbcIssueDate).SetText("2024-07-23")
	inv.CreateElement(cbcIssueTime).SetText("00:40:00Z")
	inv.CreateElement(cbcInvoiceTypeCode).CreateAttr("listVersionID", "1.0").Element().SetText(iTCRefundNote) // 1.1 requires signature!
	inv.CreateElement(cbcDocumentCurrencyCode).SetText(MYR)
	inv.CreateElement(cbcTaxCurrencyCode).SetText(MYR)

	period := inv.CreateElement(cacInvoicePeriod)
	period.CreateElement(cbcStartDate).SetText("2024-07-01")
	period.CreateElement(cbcEndDate).SetText("2024-07-31")
	period.CreateElement(cbcDescription).SetText("Monthly")

	idr := inv.CreateElement(cacBillingReference).CreateElement(cacInvoiceDocumentReference)
	idr.CreateElement(cbcID).SetText("XML-INV12345")
	idr.CreateElement(cbcUUID).SetText("Reference Invoice UUID")

	inv.CreateElement(cacBillingReference).CreateElement(cacAdditionalDocumentReference).CreateElement(cbcID).SetText("151891-1981")

	adr1 := inv.CreateElement(cacAdditionalDocumentReference)
	adr1.CreateElement(cbcID).SetText("L1")
	adr1.CreateElement(cbcDocumentType).SetText("CustomsImportForm")

	adr2 := inv.CreateElement(cacAdditionalDocumentReference)
	adr2.CreateElement(cbcID).SetText("FTA")
	adr2.CreateElement(cbcDocumentType).SetText("FreeTradeAgreement")
	adr2.CreateElement(cbcDocumentDescription).SetText("Sample Description")

	adr3 := inv.CreateElement(cacAdditionalDocumentReference)
	adr3.CreateElement(cbcID).SetText("L1")
	adr3.CreateElement(cbcDocumentType).SetText("K2")

	adr4 := inv.CreateElement(cacAdditionalDocumentReference)
	adr4.CreateElement(cbcID).SetText("L1")

	asp := inv.CreateElement(cacAccountingSupplierParty)
	asp.CreateElement(cbcAdditionalAccountID).CreateAttr("schemeAgencyName", "CertEX").
		Element().SetText("CPT-CCN-W-211111-KL-000002")

	sp := asp.CreateElement(cacParty)
	sp.CreateElement(cbcIndustryClassificationCode).
		CreateAttr("name", cMSICSC[spMSICSC].Description).Element().SetText(spMSICSC)

	sp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, TIN).Element().SetText("Supplier's TIN")
	sp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, BRN).Element().SetText("Supplier's BRN")
	sp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, SST).Element().SetText("NA")
	sp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, TTX).Element().SetText("NA")

	spa := sp.CreateElement(cacPostalAddress)
	spa.CreateElement(cbcCityName).SetText(reSubWP.ReplaceAllString(cState[spCSC].State, ""))
	spa.CreateElement(cbcPostalZone).SetText("50480")
	spa.CreateElement(cbcCountrySubentityCode).SetText(spCSC)
	spa.CreateElement(cacAddressLine).CreateElement(cbcLine).SetText("Lot 66")
	spa.CreateElement(cacAddressLine).CreateElement(cbcLine).SetText("Bangunan Merdeka")
	spa.CreateElement(cacAddressLine).CreateElement(cbcLine).SetText("Persiaran Jaya")
	spa.CreateElement(cacCountry).CreateElement(cbcIdentificationCode).
		CreateAttr(lID, "ISO3166-1").Element().CreateAttr(lAID, "6").
		Element().SetText(cCountryRev[spCIC].Code)
	sp.CreateElement(cacPartyLegalEntity).CreateElement(cbcRegistrationName).SetText("Supplier's Name")
	sc := sp.CreateElement(cacContact)
	sc.CreateElement(cbcTelephone).SetText("+60123456789")
	sc.CreateElement(cbcElectronicMail).SetText("supplier@email.com")

	acp := inv.CreateElement(cacAccountingCustomerParty)
	cp := acp.CreateElement(cacParty)

	cp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, TIN).Element().SetText("Buyer's TIN")
	cp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, BRN).Element().SetText("Buyer's BRN")
	cp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, SST).Element().SetText("NA")
	cp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, TTX).Element().SetText("NA")

	cpa := cp.CreateElement(cacPostalAddress)
	cpa.CreateElement(cbcCityName).SetText(reSubWP.ReplaceAllString(cState[spCSC].State, ""))
	cpa.CreateElement(cbcPostalZone).SetText("50480")
	cpa.CreateElement(cbcCountrySubentityCode).SetText(spCSC)
	cpa.CreateElement(cacAddressLine).CreateElement(cbcLine).SetText("Lot 66")
	cpa.CreateElement(cacAddressLine).CreateElement(cbcLine).SetText("Bangunan Merdeka")
	cpa.CreateElement(cacAddressLine).CreateElement(cbcLine).SetText("Persiaran Jaya")
	cpa.CreateElement(cacCountry).CreateElement(cbcIdentificationCode).
		CreateAttr(lID, "ISO3166-1").Element().CreateAttr(lAID, "6").
		Element().SetText(cCountryRev["MALAYSIA"].Code)
	cp.CreateElement(cacPartyLegalEntity).CreateElement(cbcRegistrationName).SetText("Buyer's Name")
	cc := cp.CreateElement(cacContact)
	cc.CreateElement(cbcTelephone).SetText("+60123456780")
	cc.CreateElement(cbcElectronicMail).SetText("buyer@email.com")

	d := inv.CreateElement(cacDelivery)
	dp := d.CreateElement(cacDeliveryParty)

	dp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, TIN).Element().SetText("Recipient's TIN")
	dp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, BRN).Element().SetText("Recipient's BRN")

	dpa := dp.CreateElement(cacPostalAddress)
	dpa.CreateElement(cbcCityName).SetText(reSubWP.ReplaceAllString(cState[spCSC].State, ""))
	dpa.CreateElement(cbcPostalZone).SetText("50480")
	dpa.CreateElement(cbcCountrySubentityCode).SetText(spCSC)
	dpa.CreateElement(cacAddressLine).CreateElement(cbcLine).SetText("Lot 66")
	dpa.CreateElement(cacAddressLine).CreateElement(cbcLine).SetText("Bangunan Merdeka")
	dpa.CreateElement(cacAddressLine).CreateElement(cbcLine).SetText("Persiaran Jaya")
	dpa.CreateElement(cacCountry).CreateElement(cbcIdentificationCode).
		CreateAttr(lID, "ISO3166-1").Element().CreateAttr(lAID, "6").
		Element().SetText(cCountryRev["MALAYSIA"].Code)
	dp.CreateElement(cacPartyLegalEntity).CreateElement(cbcRegistrationName).SetText("Recipient's Name")

	ds := d.CreateElement(cacShipment)
	ds.CreateElement(cbcID).SetText("1234")
	dsfac := ds.CreateElement(cacFreightAllowanceCharge)
	dsfac.CreateElement(cbcChargeIndicator).SetText("true")
	dsfac.CreateElement(cbcAllowanceChargeReason).SetText("Service charge")
	dsfac.CreateElement(cbcAmount).CreateAttr(cID, MYR).Element().SetText("100")

	pm := inv.CreateElement(cacPaymentMeans)
	pm.CreateElement(cbcPaymentMeansCode).SetText("01") // PaymentMethods.json
	pm.CreateElement(cacPayeeFinancialAccount).CreateElement(cbcID).SetText("1234567890")

	inv.CreateElement(cacPaymentTerms).CreateElement(cbcNote).SetText("Payment method is cash")

	pp := inv.CreateElement(cacPrepaidPayment)
	pp.CreateElement(cbcID).SetText("E12345678912")
	pp.CreateElement(cbcPaidAmount).CreateAttr(cID, MYR).Element().SetText("1.00")
	pp.CreateElement(cbcPaidDate).SetText("2024-07-23")
	pp.CreateElement(cbcPaidTime).SetText("00:30:00Z")

	ac1 := inv.CreateElement(cacAllowanceCharge)
	ac1.CreateElement(cbcChargeIndicator).SetText("false")
	ac1.CreateElement(cbcAllowanceChargeReason).SetText("Sample Description")
	ac1.CreateElement(cbcAmount).CreateAttr(cID, MYR).Element().SetText("100")

	ac2 := inv.CreateElement(cacAllowanceCharge)
	ac2.CreateElement(cbcChargeIndicator).SetText("true")
	ac2.CreateElement(cbcAllowanceChargeReason).SetText("Service charge")
	ac2.CreateElement(cbcAmount).CreateAttr(cID, MYR).Element().SetText("100")

	tt := inv.CreateElement(cacTaxTotal)
	tt.CreateElement(cbcTaxAmount).CreateAttr(cID, MYR).Element().SetText("87.63")

	ts := tt.CreateElement(cacTaxSubtotal)
	ts.CreateElement(cbcTaxableAmount).CreateAttr(cID, MYR).Element().SetText("87.63")
	ts.CreateElement(cbcTaxAmount).CreateAttr(cID, MYR).Element().SetText("87.63")
	tstc := ts.CreateElement(cacTaxCategory)
	tstc.CreateElement(cbcID).SetText("01")
	tstc.CreateElement(cacTaxScheme).CreateElement(cbcID).
		CreateAttr(sID, UE5).Element().CreateAttr(sAID, "6").Element().SetText(OTH)

	lmt := inv.CreateElement(cacLegalMonetaryTotal)
	lmt.CreateElement(cbcLineExtensionAmount).CreateAttr(cID, MYR).Element().SetText("1436.50")
	lmt.CreateElement(cbcTaxExclusiveAmount).CreateAttr(cID, MYR).Element().SetText("1436.50")
	lmt.CreateElement(cbcTaxInclusiveAmount).CreateAttr(cID, MYR).Element().SetText("1436.50")
	lmt.CreateElement(cbcAllowanceTotalAmount).CreateAttr(cID, MYR).Element().SetText("1436.50")
	lmt.CreateElement(cbcChargeTotalAmount).CreateAttr(cID, MYR).Element().SetText("1436.50")
	lmt.CreateElement(cbcPayableRoundingAmount).CreateAttr(cID, MYR).Element().SetText("0.30")
	lmt.CreateElement(cbcPayableAmount).CreateAttr(cID, MYR).Element().SetText("1436.50")

	il := inv.CreateElement(cacInvoiceLine)
	il.CreateElement(cbcID).SetText("1234")
	il.CreateElement(cbcInvoicedQuantity).CreateAttr("unitCode", "C62").Element().SetText("1")
	il.CreateElement(cbcLineExtensionAmount).CreateAttr(cID, MYR).Element().SetText("1436.50")

	ilac1 := il.CreateElement(cacAllowanceCharge)
	ilac1.CreateElement(cbcChargeIndicator).SetText("false")
	ilac1.CreateElement(cbcAllowanceChargeReason).SetText("Sample Description")
	ilac1.CreateElement(cbcMultiplierFactorNumeric).SetText("0.15")
	ilac1.CreateElement(cbcAmount).CreateAttr(cID, MYR).Element().SetText("100")

	ilac2 := il.CreateElement(cacAllowanceCharge)
	ilac2.CreateElement(cbcChargeIndicator).SetText("true")
	ilac2.CreateElement(cbcAllowanceChargeReason).SetText("Sample Description")
	ilac2.CreateElement(cbcMultiplierFactorNumeric).SetText("0.10")
	ilac2.CreateElement(cbcAmount).CreateAttr(cID, MYR).Element().SetText("100")

	iltt := il.CreateElement(cacTaxTotal)
	iltt.CreateElement(cbcTaxAmount).CreateAttr(cID, MYR).Element().SetText("0")
	ilttts := iltt.CreateElement(cacTaxSubtotal)
	ilttts.CreateElement(cbcTaxableAmount).CreateAttr(cID, MYR).Element().SetText("1460.50")
	ilttts.CreateElement(cbcTaxAmount).CreateAttr(cID, MYR).Element().SetText("0")
	ilttts.CreateElement(cbcPercent).SetText("6.00")
	ittttstc := ilttts.CreateElement(cacTaxCategory)
	ittttstc.CreateElement(cbcID).SetText("E")
	ittttstc.CreateElement(cbcTaxExemptionReason).SetText("Exempt New Means of Transport")
	ittttstc.CreateElement(cacTaxScheme).CreateElement(cbcID).
		CreateAttr(sID, UE5).Element().CreateAttr(sAID, "6").Element().SetText(OTH)

	ili := il.CreateElement(cacItem)
	ili.CreateElement(cbcDescription).SetText("Laptop Peripherals")
	ili.CreateElement(cacOriginCountry).CreateElement(cbcIdentificationCode).SetText(cCountryRev["MALAYSIA"].Code)
	ili.CreateElement(cacCommodityClassification).CreateElement(cbcItemClassificationCode).
		CreateAttr(lID, "PTC").Element().SetText("9800.00.0010")
	ili.CreateElement(cacCommodityClassification).CreateElement(cbcItemClassificationCode).
		CreateAttr(lID, "CLASS").Element().SetText("003")

	il.CreateElement(cacPrice).CreateElement(cbcPriceAmount).CreateAttr(cID, MYR).
		Element().SetText("17")
	il.CreateElement(cacItemPriceExtension).CreateElement(cbcAmount).CreateAttr(cID, MYR).
		Element().SetText("100")

	invSav(doc, out)
}

// nolint: unused
func sdkSampleSelfBilledInvoice10(out string) {
	doc := etree.NewDocument()
	inv := doc.CreateElement("Invoice").CreateAttr("xmlns", urnInvoice2).Element().
		CreateAttr("xmlns:cac", urnCommonAggregateComponents2).Element().
		CreateAttr("xmlns:cbc", urnCommonBasicComponents2).Element()
	inv.CreateElement(cbcID).SetText("XML-SBINV12345")
	inv.CreateElement(cbcIssueDate).SetText("2024-07-23")
	inv.CreateElement(cbcIssueTime).SetText("00:40:00Z")
	inv.CreateElement(cbcInvoiceTypeCode).CreateAttr("listVersionID", "1.0").Element().SetText(iTCSBInvoice) // 1.1 requires signature!
	inv.CreateElement(cbcDocumentCurrencyCode).SetText(MYR)
	inv.CreateElement(cbcTaxCurrencyCode).SetText(MYR)

	period := inv.CreateElement(cacInvoicePeriod)
	period.CreateElement(cbcStartDate).SetText("2024-07-01")
	period.CreateElement(cbcEndDate).SetText("2024-07-31")
	period.CreateElement(cbcDescription).SetText("Monthly")

	inv.CreateElement(cacBillingReference).CreateElement(cacAdditionalDocumentReference).CreateElement(cbcID).SetText("151891-1981")

	adr1 := inv.CreateElement(cacAdditionalDocumentReference)
	adr1.CreateElement(cbcID).SetText("L1")
	adr1.CreateElement(cbcDocumentType).SetText("CustomsImportForm")

	adr2 := inv.CreateElement(cacAdditionalDocumentReference)
	adr2.CreateElement(cbcID).SetText("FTA")
	adr2.CreateElement(cbcDocumentType).SetText("FreeTradeAgreement")
	adr2.CreateElement(cbcDocumentDescription).SetText("Sample Description")

	adr3 := inv.CreateElement(cacAdditionalDocumentReference)
	adr3.CreateElement(cbcID).SetText("L1")
	adr3.CreateElement(cbcDocumentType).SetText("K2")

	adr4 := inv.CreateElement(cacAdditionalDocumentReference)
	adr4.CreateElement(cbcID).SetText("L1")

	asp := inv.CreateElement(cacAccountingSupplierParty)
	asp.CreateElement(cbcAdditionalAccountID).CreateAttr("schemeAgencyName", "CertEX").
		Element().SetText("CPT-CCN-W-211111-KL-000002")

	sp := asp.CreateElement(cacParty)
	sp.CreateElement(cbcIndustryClassificationCode).
		CreateAttr("name", cMSICSC[spMSICSC].Description).Element().SetText(spMSICSC)

	sp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, TIN).Element().SetText("Supplier's TIN")
	sp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, BRN).Element().SetText("Supplier's BRN")
	sp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, SST).Element().SetText("NA")
	sp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, TTX).Element().SetText("NA")

	spa := sp.CreateElement(cacPostalAddress)
	spa.CreateElement(cbcCityName).SetText(reSubWP.ReplaceAllString(cState[spCSC].State, ""))
	spa.CreateElement(cbcPostalZone).SetText("50480")
	spa.CreateElement(cbcCountrySubentityCode).SetText(spCSC)
	spa.CreateElement(cacAddressLine).CreateElement(cbcLine).SetText("Lot 66")
	spa.CreateElement(cacAddressLine).CreateElement(cbcLine).SetText("Bangunan Merdeka")
	spa.CreateElement(cacAddressLine).CreateElement(cbcLine).SetText("Persiaran Jaya")
	spa.CreateElement(cacCountry).CreateElement(cbcIdentificationCode).
		CreateAttr(lID, "ISO3166-1").Element().CreateAttr(lAID, "6").
		Element().SetText(cCountryRev[spCIC].Code)
	sp.CreateElement(cacPartyLegalEntity).CreateElement(cbcRegistrationName).SetText("Supplier's Name")
	sc := sp.CreateElement(cacContact)
	sc.CreateElement(cbcTelephone).SetText("+60123456789")
	sc.CreateElement(cbcElectronicMail).SetText("supplier@email.com")

	acp := inv.CreateElement(cacAccountingCustomerParty)
	cp := acp.CreateElement(cacParty)

	cp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, TIN).Element().SetText("Buyer's TIN")
	cp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, BRN).Element().SetText("Buyer's BRN")
	cp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, SST).Element().SetText("NA")
	cp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, TTX).Element().SetText("NA")

	cpa := cp.CreateElement(cacPostalAddress)
	cpa.CreateElement(cbcCityName).SetText(reSubWP.ReplaceAllString(cState[spCSC].State, ""))
	cpa.CreateElement(cbcPostalZone).SetText("50480")
	cpa.CreateElement(cbcCountrySubentityCode).SetText(spCSC)
	cpa.CreateElement(cacAddressLine).CreateElement(cbcLine).SetText("Lot 66")
	cpa.CreateElement(cacAddressLine).CreateElement(cbcLine).SetText("Bangunan Merdeka")
	cpa.CreateElement(cacAddressLine).CreateElement(cbcLine).SetText("Persiaran Jaya")
	cpa.CreateElement(cacCountry).CreateElement(cbcIdentificationCode).
		CreateAttr(lID, "ISO3166-1").Element().CreateAttr(lAID, "6").
		Element().SetText(cCountryRev["MALAYSIA"].Code)
	cp.CreateElement(cacPartyLegalEntity).CreateElement(cbcRegistrationName).SetText("Buyer's Name")
	cc := cp.CreateElement(cacContact)
	cc.CreateElement(cbcTelephone).SetText("+60123456780")
	cc.CreateElement(cbcElectronicMail).SetText("buyer@email.com")

	d := inv.CreateElement(cacDelivery)
	dp := d.CreateElement(cacDeliveryParty)

	dp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, TIN).Element().SetText("Recipient's TIN")
	dp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, BRN).Element().SetText("Recipient's BRN")

	dpa := dp.CreateElement(cacPostalAddress)
	dpa.CreateElement(cbcCityName).SetText(reSubWP.ReplaceAllString(cState[spCSC].State, ""))
	dpa.CreateElement(cbcPostalZone).SetText("50480")
	dpa.CreateElement(cbcCountrySubentityCode).SetText(spCSC)
	dpa.CreateElement(cacAddressLine).CreateElement(cbcLine).SetText("Lot 66")
	dpa.CreateElement(cacAddressLine).CreateElement(cbcLine).SetText("Bangunan Merdeka")
	dpa.CreateElement(cacAddressLine).CreateElement(cbcLine).SetText("Persiaran Jaya")
	dpa.CreateElement(cacCountry).CreateElement(cbcIdentificationCode).
		CreateAttr(lID, "ISO3166-1").Element().CreateAttr(lAID, "6").
		Element().SetText(cCountryRev["MALAYSIA"].Code)
	dp.CreateElement(cacPartyLegalEntity).CreateElement(cbcRegistrationName).SetText("Recipient's Name")

	ds := d.CreateElement(cacShipment)
	ds.CreateElement(cbcID).SetText("1234")
	dsfac := ds.CreateElement(cacFreightAllowanceCharge)
	dsfac.CreateElement(cbcChargeIndicator).SetText("true")
	dsfac.CreateElement(cbcAllowanceChargeReason).SetText("Service charge")
	dsfac.CreateElement(cbcAmount).CreateAttr(cID, MYR).Element().SetText("100")

	pm := inv.CreateElement(cacPaymentMeans)
	pm.CreateElement(cbcPaymentMeansCode).SetText("01") // PaymentMethods.json
	pm.CreateElement(cacPayeeFinancialAccount).CreateElement(cbcID).SetText("1234567890")

	inv.CreateElement(cacPaymentTerms).CreateElement(cbcNote).SetText("Payment method is cash")

	pp := inv.CreateElement(cacPrepaidPayment)
	pp.CreateElement(cbcID).SetText("E12345678912")
	pp.CreateElement(cbcPaidAmount).CreateAttr(cID, MYR).Element().SetText("1.00")
	pp.CreateElement(cbcPaidDate).SetText("2024-07-23")
	pp.CreateElement(cbcPaidTime).SetText("00:30:00Z")

	ac1 := inv.CreateElement(cacAllowanceCharge)
	ac1.CreateElement(cbcChargeIndicator).SetText("false")
	ac1.CreateElement(cbcAllowanceChargeReason).SetText("Sample Description")
	ac1.CreateElement(cbcAmount).CreateAttr(cID, MYR).Element().SetText("100")

	ac2 := inv.CreateElement(cacAllowanceCharge)
	ac2.CreateElement(cbcChargeIndicator).SetText("true")
	ac2.CreateElement(cbcAllowanceChargeReason).SetText("Service charge")
	ac2.CreateElement(cbcAmount).CreateAttr(cID, MYR).Element().SetText("100")

	tt := inv.CreateElement(cacTaxTotal)
	tt.CreateElement(cbcTaxAmount).CreateAttr(cID, MYR).Element().SetText("87.63")

	ts := tt.CreateElement(cacTaxSubtotal)
	ts.CreateElement(cbcTaxableAmount).CreateAttr(cID, MYR).Element().SetText("87.63")
	ts.CreateElement(cbcTaxAmount).CreateAttr(cID, MYR).Element().SetText("87.63")
	tstc := ts.CreateElement(cacTaxCategory)
	tstc.CreateElement(cbcID).SetText("01")
	tstc.CreateElement(cacTaxScheme).CreateElement(cbcID).
		CreateAttr(sID, UE5).Element().CreateAttr(sAID, "6").Element().SetText(OTH)

	lmt := inv.CreateElement(cacLegalMonetaryTotal)
	lmt.CreateElement(cbcLineExtensionAmount).CreateAttr(cID, MYR).Element().SetText("1436.50")
	lmt.CreateElement(cbcTaxExclusiveAmount).CreateAttr(cID, MYR).Element().SetText("1436.50")
	lmt.CreateElement(cbcTaxInclusiveAmount).CreateAttr(cID, MYR).Element().SetText("1436.50")
	lmt.CreateElement(cbcAllowanceTotalAmount).CreateAttr(cID, MYR).Element().SetText("1436.50")
	lmt.CreateElement(cbcChargeTotalAmount).CreateAttr(cID, MYR).Element().SetText("1436.50")
	lmt.CreateElement(cbcPayableRoundingAmount).CreateAttr(cID, MYR).Element().SetText("0.30")
	lmt.CreateElement(cbcPayableAmount).CreateAttr(cID, MYR).Element().SetText("1436.50")

	il := inv.CreateElement(cacInvoiceLine)
	il.CreateElement(cbcID).SetText("1234")
	il.CreateElement(cbcInvoicedQuantity).CreateAttr("unitCode", "C62").Element().SetText("1")
	il.CreateElement(cbcLineExtensionAmount).CreateAttr(cID, MYR).Element().SetText("1436.50")

	ilac1 := il.CreateElement(cacAllowanceCharge)
	ilac1.CreateElement(cbcChargeIndicator).SetText("false")
	ilac1.CreateElement(cbcAllowanceChargeReason).SetText("Sample Description")
	ilac1.CreateElement(cbcMultiplierFactorNumeric).SetText("0.15")
	ilac1.CreateElement(cbcAmount).CreateAttr(cID, MYR).Element().SetText("100")

	ilac2 := il.CreateElement(cacAllowanceCharge)
	ilac2.CreateElement(cbcChargeIndicator).SetText("true")
	ilac2.CreateElement(cbcAllowanceChargeReason).SetText("Sample Description")
	ilac2.CreateElement(cbcMultiplierFactorNumeric).SetText("0.10")
	ilac2.CreateElement(cbcAmount).CreateAttr(cID, MYR).Element().SetText("100")

	iltt := il.CreateElement(cacTaxTotal)
	iltt.CreateElement(cbcTaxAmount).CreateAttr(cID, MYR).Element().SetText("0")
	ilttts := iltt.CreateElement(cacTaxSubtotal)
	ilttts.CreateElement(cbcTaxableAmount).CreateAttr(cID, MYR).Element().SetText("1460.50")
	ilttts.CreateElement(cbcTaxAmount).CreateAttr(cID, MYR).Element().SetText("0")
	ilttts.CreateElement(cbcPercent).SetText("6.00")
	ittttstc := ilttts.CreateElement(cacTaxCategory)
	ittttstc.CreateElement(cbcID).SetText("E")
	ittttstc.CreateElement(cbcTaxExemptionReason).SetText("Exempt New Means of Transport")
	ittttstc.CreateElement(cacTaxScheme).CreateElement(cbcID).
		CreateAttr(sID, UE5).Element().CreateAttr(sAID, "6").Element().SetText(OTH)

	ili := il.CreateElement(cacItem)
	ili.CreateElement(cbcDescription).SetText("Laptop Peripherals")
	ili.CreateElement(cacOriginCountry).CreateElement(cbcIdentificationCode).SetText(cCountryRev["MALAYSIA"].Code)
	ili.CreateElement(cacCommodityClassification).CreateElement(cbcItemClassificationCode).
		CreateAttr(lID, "PTC").Element().SetText("9800.00.0010")
	ili.CreateElement(cacCommodityClassification).CreateElement(cbcItemClassificationCode).
		CreateAttr(lID, "CLASS").Element().SetText("003")

	il.CreateElement(cacPrice).CreateElement(cbcPriceAmount).CreateAttr(cID, MYR).
		Element().SetText("17")
	il.CreateElement(cacItemPriceExtension).CreateElement(cbcAmount).CreateAttr(cID, MYR).
		Element().SetText("100")

	invSav(doc, out)
}

// nolint: unused
func sdkSampleSelfBilledCreditNote10(out string) {
	doc := etree.NewDocument()
	inv := doc.CreateElement("Invoice").CreateAttr("xmlns", urnInvoice2).Element().
		CreateAttr("xmlns:cac", urnCommonAggregateComponents2).Element().
		CreateAttr("xmlns:cbc", urnCommonBasicComponents2).Element()
	inv.CreateElement(cbcID).SetText("XML-SBCN12345")
	inv.CreateElement(cbcIssueDate).SetText("2024-07-23")
	inv.CreateElement(cbcIssueTime).SetText("00:40:00Z")
	inv.CreateElement(cbcInvoiceTypeCode).CreateAttr("listVersionID", "1.0").Element().
		SetText(iTCSBCreditNote) // 1.1 requires signature!
	inv.CreateElement(cbcDocumentCurrencyCode).SetText(MYR)
	inv.CreateElement(cbcTaxCurrencyCode).SetText(MYR)

	period := inv.CreateElement(cacInvoicePeriod)
	period.CreateElement(cbcStartDate).SetText("2024-07-01")
	period.CreateElement(cbcEndDate).SetText("2024-07-31")
	period.CreateElement(cbcDescription).SetText("Monthly")

	idr := inv.CreateElement(cacBillingReference).CreateElement(cacInvoiceDocumentReference)
	idr.CreateElement(cbcID).SetText("XML-SBINV12345")
	idr.CreateElement(cbcUUID).SetText("Reference Self-billed Invoice UUID")

	inv.CreateElement(cacBillingReference).CreateElement(cacAdditionalDocumentReference).CreateElement(cbcID).SetText("151891-1981")

	adr1 := inv.CreateElement(cacAdditionalDocumentReference)
	adr1.CreateElement(cbcID).SetText("L1")
	adr1.CreateElement(cbcDocumentType).SetText("CustomsImportForm")

	adr2 := inv.CreateElement(cacAdditionalDocumentReference)
	adr2.CreateElement(cbcID).SetText("FTA")
	adr2.CreateElement(cbcDocumentType).SetText("FreeTradeAgreement")
	adr2.CreateElement(cbcDocumentDescription).SetText("Sample Description")

	adr3 := inv.CreateElement(cacAdditionalDocumentReference)
	adr3.CreateElement(cbcID).SetText("L1")
	adr3.CreateElement(cbcDocumentType).SetText("K2")

	adr4 := inv.CreateElement(cacAdditionalDocumentReference)
	adr4.CreateElement(cbcID).SetText("L1")

	asp := inv.CreateElement(cacAccountingSupplierParty)
	asp.CreateElement(cbcAdditionalAccountID).CreateAttr("schemeAgencyName", "CertEX").
		Element().SetText("CPT-CCN-W-211111-KL-000002")

	sp := asp.CreateElement(cacParty)
	sp.CreateElement(cbcIndustryClassificationCode).
		CreateAttr("name", cMSICSC[spMSICSC].Description).Element().SetText(spMSICSC)

	sp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, TIN).Element().SetText("Supplier's TIN")
	sp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, BRN).Element().SetText("Supplier's BRN")
	sp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, SST).Element().SetText("NA")
	sp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, TTX).Element().SetText("NA")

	spa := sp.CreateElement(cacPostalAddress)
	spa.CreateElement(cbcCityName).SetText(reSubWP.ReplaceAllString(cState[spCSC].State, ""))
	spa.CreateElement(cbcPostalZone).SetText("50480")
	spa.CreateElement(cbcCountrySubentityCode).SetText(spCSC)
	spa.CreateElement(cacAddressLine).CreateElement(cbcLine).SetText("Lot 66")
	spa.CreateElement(cacAddressLine).CreateElement(cbcLine).SetText("Bangunan Merdeka")
	spa.CreateElement(cacAddressLine).CreateElement(cbcLine).SetText("Persiaran Jaya")
	spa.CreateElement(cacCountry).CreateElement(cbcIdentificationCode).
		CreateAttr(lID, "ISO3166-1").Element().CreateAttr(lAID, "6").
		Element().SetText(cCountryRev[spCIC].Code)
	sp.CreateElement(cacPartyLegalEntity).CreateElement(cbcRegistrationName).SetText("Supplier's Name")
	sc := sp.CreateElement(cacContact)
	sc.CreateElement(cbcTelephone).SetText("+60123456789")
	sc.CreateElement(cbcElectronicMail).SetText("supplier@email.com")

	acp := inv.CreateElement(cacAccountingCustomerParty)
	cp := acp.CreateElement(cacParty)

	cp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, TIN).Element().SetText("Buyer's TIN")
	cp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, BRN).Element().SetText("Buyer's BRN")
	cp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, SST).Element().SetText("NA")
	cp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, TTX).Element().SetText("NA")

	cpa := cp.CreateElement(cacPostalAddress)
	cpa.CreateElement(cbcCityName).SetText(reSubWP.ReplaceAllString(cState[spCSC].State, ""))
	cpa.CreateElement(cbcPostalZone).SetText("50480")
	cpa.CreateElement(cbcCountrySubentityCode).SetText(spCSC)
	cpa.CreateElement(cacAddressLine).CreateElement(cbcLine).SetText("Lot 66")
	cpa.CreateElement(cacAddressLine).CreateElement(cbcLine).SetText("Bangunan Merdeka")
	cpa.CreateElement(cacAddressLine).CreateElement(cbcLine).SetText("Persiaran Jaya")
	cpa.CreateElement(cacCountry).CreateElement(cbcIdentificationCode).
		CreateAttr(lID, "ISO3166-1").Element().CreateAttr(lAID, "6").
		Element().SetText(cCountryRev["MALAYSIA"].Code)
	cp.CreateElement(cacPartyLegalEntity).CreateElement(cbcRegistrationName).SetText("Buyer's Name")
	cc := cp.CreateElement(cacContact)
	cc.CreateElement(cbcTelephone).SetText("+60123456780")
	cc.CreateElement(cbcElectronicMail).SetText("buyer@email.com")

	d := inv.CreateElement(cacDelivery)
	dp := d.CreateElement(cacDeliveryParty)

	dp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, TIN).Element().SetText("Recipient's TIN")
	dp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, BRN).Element().SetText("Recipient's BRN")

	dpa := dp.CreateElement(cacPostalAddress)
	dpa.CreateElement(cbcCityName).SetText(reSubWP.ReplaceAllString(cState[spCSC].State, ""))
	dpa.CreateElement(cbcPostalZone).SetText("50480")
	dpa.CreateElement(cbcCountrySubentityCode).SetText(spCSC)
	dpa.CreateElement(cacAddressLine).CreateElement(cbcLine).SetText("Lot 66")
	dpa.CreateElement(cacAddressLine).CreateElement(cbcLine).SetText("Bangunan Merdeka")
	dpa.CreateElement(cacAddressLine).CreateElement(cbcLine).SetText("Persiaran Jaya")
	dpa.CreateElement(cacCountry).CreateElement(cbcIdentificationCode).
		CreateAttr(lID, "ISO3166-1").Element().CreateAttr(lAID, "6").
		Element().SetText(cCountryRev["MALAYSIA"].Code)
	dp.CreateElement(cacPartyLegalEntity).CreateElement(cbcRegistrationName).SetText("Recipient's Name")

	ds := d.CreateElement(cacShipment)
	ds.CreateElement(cbcID).SetText("1234")
	dsfac := ds.CreateElement(cacFreightAllowanceCharge)
	dsfac.CreateElement(cbcChargeIndicator).SetText("true")
	dsfac.CreateElement(cbcAllowanceChargeReason).SetText("Service charge")
	dsfac.CreateElement(cbcAmount).CreateAttr(cID, MYR).Element().SetText("100")

	pm := inv.CreateElement(cacPaymentMeans)
	pm.CreateElement(cbcPaymentMeansCode).SetText("01") // PaymentMethods.json
	pm.CreateElement(cacPayeeFinancialAccount).CreateElement(cbcID).SetText("1234567890")

	inv.CreateElement(cacPaymentTerms).CreateElement(cbcNote).SetText("Payment method is cash")

	pp := inv.CreateElement(cacPrepaidPayment)
	pp.CreateElement(cbcID).SetText("E12345678912")
	pp.CreateElement(cbcPaidAmount).CreateAttr(cID, MYR).Element().SetText("1.00")
	pp.CreateElement(cbcPaidDate).SetText("2024-07-23")
	pp.CreateElement(cbcPaidTime).SetText("00:30:00Z")

	ac1 := inv.CreateElement(cacAllowanceCharge)
	ac1.CreateElement(cbcChargeIndicator).SetText("false")
	ac1.CreateElement(cbcAllowanceChargeReason).SetText("Sample Description")
	ac1.CreateElement(cbcAmount).CreateAttr(cID, MYR).Element().SetText("100")

	ac2 := inv.CreateElement(cacAllowanceCharge)
	ac2.CreateElement(cbcChargeIndicator).SetText("true")
	ac2.CreateElement(cbcAllowanceChargeReason).SetText("Service charge")
	ac2.CreateElement(cbcAmount).CreateAttr(cID, MYR).Element().SetText("100")

	tt := inv.CreateElement(cacTaxTotal)
	tt.CreateElement(cbcTaxAmount).CreateAttr(cID, MYR).Element().SetText("87.63")

	ts := tt.CreateElement(cacTaxSubtotal)
	ts.CreateElement(cbcTaxableAmount).CreateAttr(cID, MYR).Element().SetText("87.63")
	ts.CreateElement(cbcTaxAmount).CreateAttr(cID, MYR).Element().SetText("87.63")
	tstc := ts.CreateElement(cacTaxCategory)
	tstc.CreateElement(cbcID).SetText("01")
	tstc.CreateElement(cacTaxScheme).CreateElement(cbcID).
		CreateAttr(sID, UE5).Element().CreateAttr(sAID, "6").Element().SetText(OTH)

	lmt := inv.CreateElement(cacLegalMonetaryTotal)
	lmt.CreateElement(cbcLineExtensionAmount).CreateAttr(cID, MYR).Element().SetText("1436.50")
	lmt.CreateElement(cbcTaxExclusiveAmount).CreateAttr(cID, MYR).Element().SetText("1436.50")
	lmt.CreateElement(cbcTaxInclusiveAmount).CreateAttr(cID, MYR).Element().SetText("1436.50")
	lmt.CreateElement(cbcAllowanceTotalAmount).CreateAttr(cID, MYR).Element().SetText("1436.50")
	lmt.CreateElement(cbcChargeTotalAmount).CreateAttr(cID, MYR).Element().SetText("1436.50")
	lmt.CreateElement(cbcPayableRoundingAmount).CreateAttr(cID, MYR).Element().SetText("0.30")
	lmt.CreateElement(cbcPayableAmount).CreateAttr(cID, MYR).Element().SetText("1436.50")

	il := inv.CreateElement(cacInvoiceLine)
	il.CreateElement(cbcID).SetText("1234")
	il.CreateElement(cbcInvoicedQuantity).CreateAttr("unitCode", "C62").Element().SetText("1")
	il.CreateElement(cbcLineExtensionAmount).CreateAttr(cID, MYR).Element().SetText("1436.50")

	ilac1 := il.CreateElement(cacAllowanceCharge)
	ilac1.CreateElement(cbcChargeIndicator).SetText("false")
	ilac1.CreateElement(cbcAllowanceChargeReason).SetText("Sample Description")
	ilac1.CreateElement(cbcMultiplierFactorNumeric).SetText("0.15")
	ilac1.CreateElement(cbcAmount).CreateAttr(cID, MYR).Element().SetText("100")

	ilac2 := il.CreateElement(cacAllowanceCharge)
	ilac2.CreateElement(cbcChargeIndicator).SetText("true")
	ilac2.CreateElement(cbcAllowanceChargeReason).SetText("Sample Description")
	ilac2.CreateElement(cbcMultiplierFactorNumeric).SetText("0.10")
	ilac2.CreateElement(cbcAmount).CreateAttr(cID, MYR).Element().SetText("100")

	iltt := il.CreateElement(cacTaxTotal)
	iltt.CreateElement(cbcTaxAmount).CreateAttr(cID, MYR).Element().SetText("0")
	ilttts := iltt.CreateElement(cacTaxSubtotal)
	ilttts.CreateElement(cbcTaxableAmount).CreateAttr(cID, MYR).Element().SetText("1460.50")
	ilttts.CreateElement(cbcTaxAmount).CreateAttr(cID, MYR).Element().SetText("0")
	ilttts.CreateElement(cbcPercent).SetText("6.00")
	ittttstc := ilttts.CreateElement(cacTaxCategory)
	ittttstc.CreateElement(cbcID).SetText("E")
	ittttstc.CreateElement(cbcTaxExemptionReason).SetText("Exempt New Means of Transport")
	ittttstc.CreateElement(cacTaxScheme).CreateElement(cbcID).
		CreateAttr(sID, UE5).Element().CreateAttr(sAID, "6").Element().SetText(OTH)

	ili := il.CreateElement(cacItem)
	ili.CreateElement(cbcDescription).SetText("Laptop Peripherals")
	ili.CreateElement(cacOriginCountry).CreateElement(cbcIdentificationCode).SetText(cCountryRev["MALAYSIA"].Code)
	ili.CreateElement(cacCommodityClassification).CreateElement(cbcItemClassificationCode).
		CreateAttr(lID, "PTC").Element().SetText("9800.00.0010")
	ili.CreateElement(cacCommodityClassification).CreateElement(cbcItemClassificationCode).
		CreateAttr(lID, "CLASS").Element().SetText("003")

	il.CreateElement(cacPrice).CreateElement(cbcPriceAmount).CreateAttr(cID, MYR).
		Element().SetText("17")
	il.CreateElement(cacItemPriceExtension).CreateElement(cbcAmount).CreateAttr(cID, MYR).
		Element().SetText("100")

	invSav(doc, out)
}

// nolint: unused
func sdkSampleSelfBilledDebitNote10(out string) {
	doc := etree.NewDocument()
	inv := doc.CreateElement("Invoice").CreateAttr("xmlns", urnInvoice2).Element().
		CreateAttr("xmlns:cac", urnCommonAggregateComponents2).Element().
		CreateAttr("xmlns:cbc", urnCommonBasicComponents2).Element()
	inv.CreateElement(cbcID).SetText("XML-SBDN12345")
	inv.CreateElement(cbcIssueDate).SetText("2024-07-23")
	inv.CreateElement(cbcIssueTime).SetText("00:40:00Z")
	inv.CreateElement(cbcInvoiceTypeCode).CreateAttr("listVersionID", "1.0").Element().SetText(iTCSBDebitNote) // 1.1 requires signature!
	inv.CreateElement(cbcDocumentCurrencyCode).SetText(MYR)
	inv.CreateElement(cbcTaxCurrencyCode).SetText(MYR)

	period := inv.CreateElement(cacInvoicePeriod)
	period.CreateElement(cbcStartDate).SetText("2024-07-01")
	period.CreateElement(cbcEndDate).SetText("2024-07-31")
	period.CreateElement(cbcDescription).SetText("Monthly")

	idr := inv.CreateElement(cacBillingReference).CreateElement(cacInvoiceDocumentReference)
	idr.CreateElement(cbcID).SetText("XML-SBINV12345")
	idr.CreateElement(cbcUUID).SetText("Reference Self-billed Invoice UUID")

	inv.CreateElement(cacBillingReference).CreateElement(cacAdditionalDocumentReference).CreateElement(cbcID).SetText("151891-1981")

	adr1 := inv.CreateElement(cacAdditionalDocumentReference)
	adr1.CreateElement(cbcID).SetText("L1")
	adr1.CreateElement(cbcDocumentType).SetText("CustomsImportForm")

	adr2 := inv.CreateElement(cacAdditionalDocumentReference)
	adr2.CreateElement(cbcID).SetText("FTA")
	adr2.CreateElement(cbcDocumentType).SetText("FreeTradeAgreement")
	adr2.CreateElement(cbcDocumentDescription).SetText("Sample Description")

	adr3 := inv.CreateElement(cacAdditionalDocumentReference)
	adr3.CreateElement(cbcID).SetText("L1")
	adr3.CreateElement(cbcDocumentType).SetText("K2")

	adr4 := inv.CreateElement(cacAdditionalDocumentReference)
	adr4.CreateElement(cbcID).SetText("L1")

	asp := inv.CreateElement(cacAccountingSupplierParty)
	asp.CreateElement(cbcAdditionalAccountID).CreateAttr("schemeAgencyName", "CertEX").
		Element().SetText("CPT-CCN-W-211111-KL-000002")

	sp := asp.CreateElement(cacParty)
	sp.CreateElement(cbcIndustryClassificationCode).
		CreateAttr("name", cMSICSC[spMSICSC].Description).Element().SetText(spMSICSC)

	sp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, TIN).Element().SetText("Supplier's TIN")
	sp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, BRN).Element().SetText("Supplier's BRN")
	sp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, SST).Element().SetText("NA")
	sp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, TTX).Element().SetText("NA")

	spa := sp.CreateElement(cacPostalAddress)
	spa.CreateElement(cbcCityName).SetText(reSubWP.ReplaceAllString(cState[spCSC].State, ""))
	spa.CreateElement(cbcPostalZone).SetText("50480")
	spa.CreateElement(cbcCountrySubentityCode).SetText(spCSC)
	spa.CreateElement(cacAddressLine).CreateElement(cbcLine).SetText("Lot 66")
	spa.CreateElement(cacAddressLine).CreateElement(cbcLine).SetText("Bangunan Merdeka")
	spa.CreateElement(cacAddressLine).CreateElement(cbcLine).SetText("Persiaran Jaya")
	spa.CreateElement(cacCountry).CreateElement(cbcIdentificationCode).
		CreateAttr(lID, "ISO3166-1").Element().CreateAttr(lAID, "6").
		Element().SetText(cCountryRev[spCIC].Code)
	sp.CreateElement(cacPartyLegalEntity).CreateElement(cbcRegistrationName).SetText("Supplier's Name")
	sc := sp.CreateElement(cacContact)
	sc.CreateElement(cbcTelephone).SetText("+60123456789")
	sc.CreateElement(cbcElectronicMail).SetText("supplier@email.com")

	acp := inv.CreateElement(cacAccountingCustomerParty)
	cp := acp.CreateElement(cacParty)

	cp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, TIN).Element().SetText("Buyer's TIN")
	cp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, BRN).Element().SetText("Buyer's BRN")
	cp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, SST).Element().SetText("NA")
	cp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, TTX).Element().SetText("NA")

	cpa := cp.CreateElement(cacPostalAddress)
	cpa.CreateElement(cbcCityName).SetText(reSubWP.ReplaceAllString(cState[spCSC].State, ""))
	cpa.CreateElement(cbcPostalZone).SetText("50480")
	cpa.CreateElement(cbcCountrySubentityCode).SetText(spCSC)
	cpa.CreateElement(cacAddressLine).CreateElement(cbcLine).SetText("Lot 66")
	cpa.CreateElement(cacAddressLine).CreateElement(cbcLine).SetText("Bangunan Merdeka")
	cpa.CreateElement(cacAddressLine).CreateElement(cbcLine).SetText("Persiaran Jaya")
	cpa.CreateElement(cacCountry).CreateElement(cbcIdentificationCode).
		CreateAttr(lID, "ISO3166-1").Element().CreateAttr(lAID, "6").
		Element().SetText(cCountryRev["MALAYSIA"].Code)
	cp.CreateElement(cacPartyLegalEntity).CreateElement(cbcRegistrationName).SetText("Buyer's Name")
	cc := cp.CreateElement(cacContact)
	cc.CreateElement(cbcTelephone).SetText("+60123456780")
	cc.CreateElement(cbcElectronicMail).SetText("buyer@email.com")

	d := inv.CreateElement(cacDelivery)
	dp := d.CreateElement(cacDeliveryParty)

	dp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, TIN).Element().SetText("Recipient's TIN")
	dp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, BRN).Element().SetText("Recipient's BRN")

	dpa := dp.CreateElement(cacPostalAddress)
	dpa.CreateElement(cbcCityName).SetText(reSubWP.ReplaceAllString(cState[spCSC].State, ""))
	dpa.CreateElement(cbcPostalZone).SetText("50480")
	dpa.CreateElement(cbcCountrySubentityCode).SetText(spCSC)
	dpa.CreateElement(cacAddressLine).CreateElement(cbcLine).SetText("Lot 66")
	dpa.CreateElement(cacAddressLine).CreateElement(cbcLine).SetText("Bangunan Merdeka")
	dpa.CreateElement(cacAddressLine).CreateElement(cbcLine).SetText("Persiaran Jaya")
	dpa.CreateElement(cacCountry).CreateElement(cbcIdentificationCode).
		CreateAttr(lID, "ISO3166-1").Element().CreateAttr(lAID, "6").
		Element().SetText(cCountryRev["MALAYSIA"].Code)
	dp.CreateElement(cacPartyLegalEntity).CreateElement(cbcRegistrationName).SetText("Recipient's Name")

	ds := d.CreateElement(cacShipment)
	ds.CreateElement(cbcID).SetText("1234")
	dsfac := ds.CreateElement(cacFreightAllowanceCharge)
	dsfac.CreateElement(cbcChargeIndicator).SetText("true")
	dsfac.CreateElement(cbcAllowanceChargeReason).SetText("Service charge")
	dsfac.CreateElement(cbcAmount).CreateAttr(cID, MYR).Element().SetText("100")

	pm := inv.CreateElement(cacPaymentMeans)
	pm.CreateElement(cbcPaymentMeansCode).SetText("01") // PaymentMethods.json
	pm.CreateElement(cacPayeeFinancialAccount).CreateElement(cbcID).SetText("1234567890")

	inv.CreateElement(cacPaymentTerms).CreateElement(cbcNote).SetText("Payment method is cash")

	pp := inv.CreateElement(cacPrepaidPayment)
	pp.CreateElement(cbcID).SetText("E12345678912")
	pp.CreateElement(cbcPaidAmount).CreateAttr(cID, MYR).Element().SetText("1.00")
	pp.CreateElement(cbcPaidDate).SetText("2024-07-23")
	pp.CreateElement(cbcPaidTime).SetText("00:30:00Z")

	ac1 := inv.CreateElement(cacAllowanceCharge)
	ac1.CreateElement(cbcChargeIndicator).SetText("false")
	ac1.CreateElement(cbcAllowanceChargeReason).SetText("Sample Description")
	ac1.CreateElement(cbcAmount).CreateAttr(cID, MYR).Element().SetText("100")

	ac2 := inv.CreateElement(cacAllowanceCharge)
	ac2.CreateElement(cbcChargeIndicator).SetText("true")
	ac2.CreateElement(cbcAllowanceChargeReason).SetText("Service charge")
	ac2.CreateElement(cbcAmount).CreateAttr(cID, MYR).Element().SetText("100")

	tt := inv.CreateElement(cacTaxTotal)
	tt.CreateElement(cbcTaxAmount).CreateAttr(cID, MYR).Element().SetText("87.63")

	ts := tt.CreateElement(cacTaxSubtotal)
	ts.CreateElement(cbcTaxableAmount).CreateAttr(cID, MYR).Element().SetText("87.63")
	ts.CreateElement(cbcTaxAmount).CreateAttr(cID, MYR).Element().SetText("87.63")
	tstc := ts.CreateElement(cacTaxCategory)
	tstc.CreateElement(cbcID).SetText("01")
	tstc.CreateElement(cacTaxScheme).CreateElement(cbcID).
		CreateAttr(sID, UE5).Element().CreateAttr(sAID, "6").Element().SetText(OTH)

	lmt := inv.CreateElement(cacLegalMonetaryTotal)
	lmt.CreateElement(cbcLineExtensionAmount).CreateAttr(cID, MYR).Element().SetText("1436.50")
	lmt.CreateElement(cbcTaxExclusiveAmount).CreateAttr(cID, MYR).Element().SetText("1436.50")
	lmt.CreateElement(cbcTaxInclusiveAmount).CreateAttr(cID, MYR).Element().SetText("1436.50")
	lmt.CreateElement(cbcAllowanceTotalAmount).CreateAttr(cID, MYR).Element().SetText("1436.50")
	lmt.CreateElement(cbcChargeTotalAmount).CreateAttr(cID, MYR).Element().SetText("1436.50")
	lmt.CreateElement(cbcPayableRoundingAmount).CreateAttr(cID, MYR).Element().SetText("0.30")
	lmt.CreateElement(cbcPayableAmount).CreateAttr(cID, MYR).Element().SetText("1436.50")

	il := inv.CreateElement(cacInvoiceLine)
	il.CreateElement(cbcID).SetText("1234")
	il.CreateElement(cbcInvoicedQuantity).CreateAttr("unitCode", "C62").Element().SetText("1")
	il.CreateElement(cbcLineExtensionAmount).CreateAttr(cID, MYR).Element().SetText("1436.50")

	ilac1 := il.CreateElement(cacAllowanceCharge)
	ilac1.CreateElement(cbcChargeIndicator).SetText("false")
	ilac1.CreateElement(cbcAllowanceChargeReason).SetText("Sample Description")
	ilac1.CreateElement(cbcMultiplierFactorNumeric).SetText("0.15")
	ilac1.CreateElement(cbcAmount).CreateAttr(cID, MYR).Element().SetText("100")

	ilac2 := il.CreateElement(cacAllowanceCharge)
	ilac2.CreateElement(cbcChargeIndicator).SetText("true")
	ilac2.CreateElement(cbcAllowanceChargeReason).SetText("Sample Description")
	ilac2.CreateElement(cbcMultiplierFactorNumeric).SetText("0.10")
	ilac2.CreateElement(cbcAmount).CreateAttr(cID, MYR).Element().SetText("100")

	iltt := il.CreateElement(cacTaxTotal)
	iltt.CreateElement(cbcTaxAmount).CreateAttr(cID, MYR).Element().SetText("0")
	ilttts := iltt.CreateElement(cacTaxSubtotal)
	ilttts.CreateElement(cbcTaxableAmount).CreateAttr(cID, MYR).Element().SetText("1460.50")
	ilttts.CreateElement(cbcTaxAmount).CreateAttr(cID, MYR).Element().SetText("0")
	ilttts.CreateElement(cbcPercent).SetText("6.00")
	ittttstc := ilttts.CreateElement(cacTaxCategory)
	ittttstc.CreateElement(cbcID).SetText("E")
	ittttstc.CreateElement(cbcTaxExemptionReason).SetText("Exempt New Means of Transport")
	ittttstc.CreateElement(cacTaxScheme).CreateElement(cbcID).
		CreateAttr(sID, UE5).Element().CreateAttr(sAID, "6").Element().SetText(OTH)

	ili := il.CreateElement(cacItem)
	ili.CreateElement(cbcDescription).SetText("Laptop Peripherals")
	ili.CreateElement(cacOriginCountry).CreateElement(cbcIdentificationCode).SetText(cCountryRev["MALAYSIA"].Code)
	ili.CreateElement(cacCommodityClassification).CreateElement(cbcItemClassificationCode).
		CreateAttr(lID, "PTC").Element().SetText("9800.00.0010")
	ili.CreateElement(cacCommodityClassification).CreateElement(cbcItemClassificationCode).
		CreateAttr(lID, "CLASS").Element().SetText("003")

	il.CreateElement(cacPrice).CreateElement(cbcPriceAmount).CreateAttr(cID, MYR).
		Element().SetText("17")
	il.CreateElement(cacItemPriceExtension).CreateElement(cbcAmount).CreateAttr(cID, MYR).
		Element().SetText("100")

	invSav(doc, out)
}

// nolint: unused
func sdkSampleSelfBilledRefundNote10(out string) {
	doc := etree.NewDocument()
	inv := doc.CreateElement("Invoice").CreateAttr("xmlns", urnInvoice2).Element().
		CreateAttr("xmlns:cac", urnCommonAggregateComponents2).Element().
		CreateAttr("xmlns:cbc", urnCommonBasicComponents2).Element()
	inv.CreateElement(cbcID).SetText("XML-SBRN12345")
	inv.CreateElement(cbcIssueDate).SetText("2024-07-23")
	inv.CreateElement(cbcIssueTime).SetText("00:40:00Z")
	inv.CreateElement(cbcInvoiceTypeCode).CreateAttr("listVersionID", "1.0").Element().SetText(iTCSBRefundNote) // 1.1 requires signature!
	inv.CreateElement(cbcDocumentCurrencyCode).SetText(MYR)
	inv.CreateElement(cbcTaxCurrencyCode).SetText(MYR)

	period := inv.CreateElement(cacInvoicePeriod)
	period.CreateElement(cbcStartDate).SetText("2024-07-01")
	period.CreateElement(cbcEndDate).SetText("2024-07-31")
	period.CreateElement(cbcDescription).SetText("Monthly")

	idr := inv.CreateElement(cacBillingReference).CreateElement(cacInvoiceDocumentReference)
	idr.CreateElement(cbcID).SetText("XML-SBINV12345")
	idr.CreateElement(cbcUUID).SetText("Reference Self-billed Invoice UUID")

	inv.CreateElement(cacBillingReference).CreateElement(cacAdditionalDocumentReference).CreateElement(cbcID).SetText("151891-1981")

	adr1 := inv.CreateElement(cacAdditionalDocumentReference)
	adr1.CreateElement(cbcID).SetText("L1")
	adr1.CreateElement(cbcDocumentType).SetText("CustomsImportForm")

	adr2 := inv.CreateElement(cacAdditionalDocumentReference)
	adr2.CreateElement(cbcID).SetText("FTA")
	adr2.CreateElement(cbcDocumentType).SetText("FreeTradeAgreement")
	adr2.CreateElement(cbcDocumentDescription).SetText("Sample Description")

	adr3 := inv.CreateElement(cacAdditionalDocumentReference)
	adr3.CreateElement(cbcID).SetText("L1")
	adr3.CreateElement(cbcDocumentType).SetText("K2")

	adr4 := inv.CreateElement(cacAdditionalDocumentReference)
	adr4.CreateElement(cbcID).SetText("L1")

	asp := inv.CreateElement(cacAccountingSupplierParty)
	asp.CreateElement(cbcAdditionalAccountID).CreateAttr("schemeAgencyName", "CertEX").
		Element().SetText("CPT-CCN-W-211111-KL-000002")

	sp := asp.CreateElement(cacParty)
	sp.CreateElement(cbcIndustryClassificationCode).
		CreateAttr("name", cMSICSC[spMSICSC].Description).Element().SetText(spMSICSC)

	sp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, TIN).Element().SetText("Supplier's TIN")
	sp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, BRN).Element().SetText("Supplier's BRN")
	sp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, SST).Element().SetText("NA")
	sp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, TTX).Element().SetText("NA")

	spa := sp.CreateElement(cacPostalAddress)
	spa.CreateElement(cbcCityName).SetText(reSubWP.ReplaceAllString(cState[spCSC].State, ""))
	spa.CreateElement(cbcPostalZone).SetText("50480")
	spa.CreateElement(cbcCountrySubentityCode).SetText(spCSC)
	spa.CreateElement(cacAddressLine).CreateElement(cbcLine).SetText("Lot 66")
	spa.CreateElement(cacAddressLine).CreateElement(cbcLine).SetText("Bangunan Merdeka")
	spa.CreateElement(cacAddressLine).CreateElement(cbcLine).SetText("Persiaran Jaya")
	spa.CreateElement(cacCountry).CreateElement(cbcIdentificationCode).
		CreateAttr(lID, "ISO3166-1").Element().CreateAttr(lAID, "6").
		Element().SetText(cCountryRev[spCIC].Code)
	sp.CreateElement(cacPartyLegalEntity).CreateElement(cbcRegistrationName).SetText("Supplier's Name")
	sc := sp.CreateElement(cacContact)
	sc.CreateElement(cbcTelephone).SetText("+60123456789")
	sc.CreateElement(cbcElectronicMail).SetText("supplier@email.com")

	acp := inv.CreateElement(cacAccountingCustomerParty)
	cp := acp.CreateElement(cacParty)

	cp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, TIN).Element().SetText("Buyer's TIN")
	cp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, BRN).Element().SetText("Buyer's BRN")
	cp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, SST).Element().SetText("NA")
	cp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, TTX).Element().SetText("NA")

	cpa := cp.CreateElement(cacPostalAddress)
	cpa.CreateElement(cbcCityName).SetText(reSubWP.ReplaceAllString(cState[spCSC].State, ""))
	cpa.CreateElement(cbcPostalZone).SetText("50480")
	cpa.CreateElement(cbcCountrySubentityCode).SetText(spCSC)
	cpa.CreateElement(cacAddressLine).CreateElement(cbcLine).SetText("Lot 66")
	cpa.CreateElement(cacAddressLine).CreateElement(cbcLine).SetText("Bangunan Merdeka")
	cpa.CreateElement(cacAddressLine).CreateElement(cbcLine).SetText("Persiaran Jaya")
	cpa.CreateElement(cacCountry).CreateElement(cbcIdentificationCode).
		CreateAttr(lID, "ISO3166-1").Element().CreateAttr(lAID, "6").
		Element().SetText(cCountryRev["MALAYSIA"].Code)
	cp.CreateElement(cacPartyLegalEntity).CreateElement(cbcRegistrationName).SetText("Buyer's Name")
	cc := cp.CreateElement(cacContact)
	cc.CreateElement(cbcTelephone).SetText("+60123456780")
	cc.CreateElement(cbcElectronicMail).SetText("buyer@email.com")

	d := inv.CreateElement(cacDelivery)
	dp := d.CreateElement(cacDeliveryParty)

	dp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, TIN).Element().SetText("Recipient's TIN")
	dp.CreateElement(cacPartyIdentification).CreateElement(cbcID).
		CreateAttr(sID, BRN).Element().SetText("Recipient's BRN")

	dpa := dp.CreateElement(cacPostalAddress)
	dpa.CreateElement(cbcCityName).SetText(reSubWP.ReplaceAllString(cState[spCSC].State, ""))
	dpa.CreateElement(cbcPostalZone).SetText("50480")
	dpa.CreateElement(cbcCountrySubentityCode).SetText(spCSC)
	dpa.CreateElement(cacAddressLine).CreateElement(cbcLine).SetText("Lot 66")
	dpa.CreateElement(cacAddressLine).CreateElement(cbcLine).SetText("Bangunan Merdeka")
	dpa.CreateElement(cacAddressLine).CreateElement(cbcLine).SetText("Persiaran Jaya")
	dpa.CreateElement(cacCountry).CreateElement(cbcIdentificationCode).
		CreateAttr(lID, "ISO3166-1").Element().CreateAttr(lAID, "6").
		Element().SetText(cCountryRev["MALAYSIA"].Code)
	dp.CreateElement(cacPartyLegalEntity).CreateElement(cbcRegistrationName).SetText("Recipient's Name")

	ds := d.CreateElement(cacShipment)
	ds.CreateElement(cbcID).SetText("1234")
	dsfac := ds.CreateElement(cacFreightAllowanceCharge)
	dsfac.CreateElement(cbcChargeIndicator).SetText("true")
	dsfac.CreateElement(cbcAllowanceChargeReason).SetText("Service charge")
	dsfac.CreateElement(cbcAmount).CreateAttr(cID, MYR).Element().SetText("100")

	pm := inv.CreateElement(cacPaymentMeans)
	pm.CreateElement(cbcPaymentMeansCode).SetText("01") // PaymentMethods.json
	pm.CreateElement(cacPayeeFinancialAccount).CreateElement(cbcID).SetText("1234567890")

	inv.CreateElement(cacPaymentTerms).CreateElement(cbcNote).SetText("Payment method is cash")

	pp := inv.CreateElement(cacPrepaidPayment)
	pp.CreateElement(cbcID).SetText("E12345678912")
	pp.CreateElement(cbcPaidAmount).CreateAttr(cID, MYR).Element().SetText("1.00")
	pp.CreateElement(cbcPaidDate).SetText("2024-07-23")
	pp.CreateElement(cbcPaidTime).SetText("00:30:00Z")

	ac1 := inv.CreateElement(cacAllowanceCharge)
	ac1.CreateElement(cbcChargeIndicator).SetText("false")
	ac1.CreateElement(cbcAllowanceChargeReason).SetText("Sample Description")
	ac1.CreateElement(cbcAmount).CreateAttr(cID, MYR).Element().SetText("100")

	ac2 := inv.CreateElement(cacAllowanceCharge)
	ac2.CreateElement(cbcChargeIndicator).SetText("true")
	ac2.CreateElement(cbcAllowanceChargeReason).SetText("Service charge")
	ac2.CreateElement(cbcAmount).CreateAttr(cID, MYR).Element().SetText("100")

	tt := inv.CreateElement(cacTaxTotal)
	tt.CreateElement(cbcTaxAmount).CreateAttr(cID, MYR).Element().SetText("87.63")

	ts := tt.CreateElement(cacTaxSubtotal)
	ts.CreateElement(cbcTaxableAmount).CreateAttr(cID, MYR).Element().SetText("87.63")
	ts.CreateElement(cbcTaxAmount).CreateAttr(cID, MYR).Element().SetText("87.63")
	tstc := ts.CreateElement(cacTaxCategory)
	tstc.CreateElement(cbcID).SetText("01")
	tstc.CreateElement(cacTaxScheme).CreateElement(cbcID).
		CreateAttr(sID, UE5).Element().CreateAttr(sAID, "6").Element().SetText(OTH)

	lmt := inv.CreateElement(cacLegalMonetaryTotal)
	lmt.CreateElement(cbcLineExtensionAmount).CreateAttr(cID, MYR).Element().SetText("1436.50")
	lmt.CreateElement(cbcTaxExclusiveAmount).CreateAttr(cID, MYR).Element().SetText("1436.50")
	lmt.CreateElement(cbcTaxInclusiveAmount).CreateAttr(cID, MYR).Element().SetText("1436.50")
	lmt.CreateElement(cbcAllowanceTotalAmount).CreateAttr(cID, MYR).Element().SetText("1436.50")
	lmt.CreateElement(cbcChargeTotalAmount).CreateAttr(cID, MYR).Element().SetText("1436.50")
	lmt.CreateElement(cbcPayableRoundingAmount).CreateAttr(cID, MYR).Element().SetText("0.30")
	lmt.CreateElement(cbcPayableAmount).CreateAttr(cID, MYR).Element().SetText("1436.50")

	il := inv.CreateElement(cacInvoiceLine)
	il.CreateElement(cbcID).SetText("1234")
	il.CreateElement(cbcInvoicedQuantity).CreateAttr("unitCode", "C62").Element().SetText("1")
	il.CreateElement(cbcLineExtensionAmount).CreateAttr(cID, MYR).Element().SetText("1436.50")

	ilac1 := il.CreateElement(cacAllowanceCharge)
	ilac1.CreateElement(cbcChargeIndicator).SetText("false")
	ilac1.CreateElement(cbcAllowanceChargeReason).SetText("Sample Description")
	ilac1.CreateElement(cbcMultiplierFactorNumeric).SetText("0.15")
	ilac1.CreateElement(cbcAmount).CreateAttr(cID, MYR).Element().SetText("100")

	ilac2 := il.CreateElement(cacAllowanceCharge)
	ilac2.CreateElement(cbcChargeIndicator).SetText("true")
	ilac2.CreateElement(cbcAllowanceChargeReason).SetText("Sample Description")
	ilac2.CreateElement(cbcMultiplierFactorNumeric).SetText("0.10")
	ilac2.CreateElement(cbcAmount).CreateAttr(cID, MYR).Element().SetText("100")

	iltt := il.CreateElement(cacTaxTotal)
	iltt.CreateElement(cbcTaxAmount).CreateAttr(cID, MYR).Element().SetText("0")
	ilttts := iltt.CreateElement(cacTaxSubtotal)
	ilttts.CreateElement(cbcTaxableAmount).CreateAttr(cID, MYR).Element().SetText("1460.50")
	ilttts.CreateElement(cbcTaxAmount).CreateAttr(cID, MYR).Element().SetText("0")
	ilttts.CreateElement(cbcPercent).SetText("6.00")
	ittttstc := ilttts.CreateElement(cacTaxCategory)
	ittttstc.CreateElement(cbcID).SetText("E")
	ittttstc.CreateElement(cbcTaxExemptionReason).SetText("Exempt New Means of Transport")
	ittttstc.CreateElement(cacTaxScheme).CreateElement(cbcID).
		CreateAttr(sID, UE5).Element().CreateAttr(sAID, "6").Element().SetText(OTH)

	ili := il.CreateElement(cacItem)
	ili.CreateElement(cbcDescription).SetText("Laptop Peripherals")
	ili.CreateElement(cacOriginCountry).CreateElement(cbcIdentificationCode).SetText(cCountryRev["MALAYSIA"].Code)
	ili.CreateElement(cacCommodityClassification).CreateElement(cbcItemClassificationCode).
		CreateAttr(lID, "PTC").Element().SetText("9800.00.0010")
	ili.CreateElement(cacCommodityClassification).CreateElement(cbcItemClassificationCode).
		CreateAttr(lID, "CLASS").Element().SetText("003")

	il.CreateElement(cacPrice).CreateElement(cbcPriceAmount).CreateAttr(cID, MYR).
		Element().SetText("17")
	il.CreateElement(cacItemPriceExtension).CreateElement(cbcAmount).CreateAttr(cID, MYR).
		Element().SetText("100")

	invSav(doc, out)
}
