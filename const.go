package main

// nolint: unused
const (
	iTCInvoice      = "01"
	iTCCreditNote   = "02"
	iTCDebitNote    = "03"
	iTCRefundNote   = "04"
	iTCSBInvoice    = "11"
	iTCSBCreditNote = "12"
	iTCSBDebitNote  = "13"
	iTCSBRefundNote = "14"

	XML  = "XML"
	vE   = `version="1.0" encoding="UTF-8"`
	cID  = "currencyID"
	sID  = "schemeID"
	sAID = "schemeAgencyID"
	sAN  = "schemeAgencyName"
	lID  = "listID"
	lVID = "listVersionID"
	lAID = "listAgencyID"
	nAME = "name"
	UE5  = "UN/ECE 5153"
	NUL  = "NUL"
	OTH  = "OTH"
	uC   = "unitCode"
	_10  = "1.0"
	_11  = "1.1"

	CLASS = "CLASS"

	// irbm-e-invoice-specific-guideline.pdf pg.123 appendix 1 - List of general TIN
	gpTIN = "EI00000000010" // General Public’s TIN
	fbTIN = "EI00000000020" // Foreign Buyer’s / Foreign Shipping Recipient’s TIN
	fsTIN = "EI00000000030" // Foreign Supplier’s TIN
	bTIN  = "EI00000000040" // Buyer’s TIN // government/special

	TIN      = "TIN"
	BRN      = "BRN" // https://sdk.myinvois.hasil.gov.my/faq/ for idType
	NRIC     = "NRIC"
	PASSPORT = "PASSPORT"
	ARMY     = "ARMY"
	SST      = "SST"
	TTX      = "TTX"

	MYR = "MYR"
	USD = "USD"

	D_BASE          = "/home/user/studies/myinvois/"
	D_XSDRT_MAINDOC = "refs/UBL-2.1/xsdrt/maindoc/"
	F_XSD_INVOICE   = "UBL-Invoice-2.1.xsd" // doesn't use UBL-2.1's Credit/Debit Notes.
	D_XML           = "invoices/xml"
	D_PDF           = "invoices/pdf"

	D_C                 = "refs/sdk.myinvois.hasil.gov.my/files/" // coding
	F_C_Classification  = "ClassificationCodes.json"
	F_C_Country         = "CountryCodes.json"
	F_C_Currency        = "CurrencyCodes.json"
	F_C_MSICSubCategory = "MSICSubCategoryCodes.json"
	F_C_State           = "StateCodes.json"

	D_EG = "refs/sdk.myinvois.hasil.gov.my/files/sdksamples/" // orig.samples

	// sp:SupplierParty
	spMSICSC = "46510"
	spCSC    = "14"
	spCIC    = "MALAYSIA"

	invSep = "-" // so that aid#-bid#
)
