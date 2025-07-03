package main

import "time"

type CodeMSICSubCategory struct {
	Code                  string `json:"Code"`
	Description           string `json:"Description"`
	MSICCategoryReference string `json:"MSIC Category Reference"`
}

type CodeClassification struct {
	Code        string `json:"Code"`
	Description string `json:"Description"`
}

type CodeCountry struct {
	Code    string `json:"Code"`
	Country string `json:"Country"`
}

type CodeState struct {
	Code  string `json:"Code"`
	State string `json:"State"`
}

type CodeCurrency struct {
	Code     string `json:"Code"`
	Currency string `json:"Currency"`
}

// nolint: unused
type valAttrs struct {
	Val   string
	Attrs [][2]string
}

// nolint: unused
type valAttr struct {
	Val  string
	Attr [2]string
}

type Client struct {
	ID     string `koanf:"id"`
	Secret string `koanf:"secret"`
}

type PartyId struct {
	Sid string `koanf:"sid"`
	Val string `koanf:"val"`
}

type respPostLogin struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
}

// {
//     "submissionUid": "1NWJD9A4Y9VQR367C9D4W1ZJ10",
//     "acceptedDocuments": [
//         {
//             "uuid": "X1DH8MC00J1RP2MDC9D4W1ZJ10",
//             "invoiceCodeNumber": "8888-8888-20250701"
//         }
//     ],
//     "rejectedDocuments": []
// }

type respPostSubmit struct {
	SubmissionUid     string             `json:"submissionUid"`
	AcceptedDocuments []AcceptedDocument `json:"acceptedDocuments"`
	RejectedDocuments []RejectedDocument `json:"rejectedDocuments"`
}

type AcceptedDocument struct {
	UUID              string `json:"uuid"`
	InvoiceCodeNumber string `json:"invoiceCodeNumber"`
}

type RejectedDocument struct {
	InvoiceCodeNumber string   `json:"invoiceCodeNumber"`
	Error             SubError `json:"error"`
}

type SubError struct {
	Code         any        `json:"code"`
	Message      string     `json:"message"`
	Target       string     `json:"target"`
	PropertyPath any        `json:"propertyPath"`
	Details      []SubError `json:"details"`
}

type respGetDocumentDetails struct {
	UUID                  string            `json:"uuid"`
	SubmissionUid         string            `json:"submissionUid"`
	LongID                string            `json:"longId"`
	TypeName              string            `json:"typeName"`
	TypeVersionName       string            `json:"typeVersionName"`
	IssuerTin             string            `json:"issuerTin"`
	IssuerName            string            `json:"issuerName"`
	ReceiverID            string            `json:"receiverId"`
	ReceiverName          string            `json:"receiverName"`
	DateTimeReceived      time.Time         `json:"dateTimeReceived"`
	DateTimeValidated     time.Time         `json:"dateTimeValidated"`
	TotalExcludingTax     float64           `json:"totalExcludingTax"`
	TotalDiscount         float64           `json:"totalDiscount"`
	TotalNetAmount        float64           `json:"totalNetAmount"`
	TotalPayableAmount    float64           `json:"totalPayableAmount"`
	Status                string            `json:"status"`
	CreatedByUserID       string            `json:"createdByUserId"`
	DocumentStatusReason  interface{}       `json:"documentStatusReason"`
	CancelDateTime        interface{}       `json:"cancelDateTime"`
	RejectRequestDateTime interface{}       `json:"rejectRequestDateTime"`
	ValidationResults     ValidationResults `json:"validationResults"`
	InternalID            string            `json:"internalId"`
	DateTimeIssued        time.Time         `json:"dateTimeIssued"`
}

type ValidationResults struct {
	Status          string           `json:"status"`
	ValidationSteps []ValidationStep `json:"validationSteps"`
}

type ValidationStep struct {
	Status string `json:"status"`
	Name   string `json:"name"`
}

type respGetTaxpayersQRCodeInfo struct {
	Name                          string `json:"name"`
	Tin                           string `json:"tin"`
	IDType                        string `json:"idType"`
	IDNumber                      string `json:"idNumber"`
	Sst                           string `json:"sst"`
	Email                         string `json:"email"`
	ContactNumber                 string `json:"contactNumber"`
	Ttx                           string `json:"ttx"`
	BusinessActivityDescriptionBM string `json:"businessActivityDescriptionBM"`
	BusinessActivityDescriptionEN string `json:"businessActivityDescriptionEN"`
	Msic                          string `json:"msic"`
	AddressLine0                  string `json:"addressLine0"`
	AddressLine1                  string `json:"addressLine1"`
	AddressLine2                  string `json:"addressLine2"`
	PostalZone                    string `json:"postalZone"`
	City                          string `json:"city"`
	State                         string `json:"state"`
	Country                       string `json:"country"`
	GeneratedTimestamp            string `json:"generatedTimestamp"` // 2025-06-13T03:17:11 non-std-fmt
}

type Docs struct {
	Documents []Doc `json:"documents"`
}

type Doc struct {
	Format       string `json:"format"`
	DocumentHash string `json:"documentHash"`
	CodeNumber   string `json:"codeNumber"`
	Document     string `json:"document"`
}

type InvLine struct {
	Quantity       string
	UnitCode       string
	Currency       string
	UnitPrice      string
	TotalPrice     string
	Description    string
	OriginCountry  string
	ProductTariff  *valAttr
	Classification *valAttr
	TaxType        string
	TaxAmount      string // taxable * precentage
	TaxTaxable     string
	TaxPercentage  string
}
