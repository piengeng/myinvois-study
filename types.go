package main

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
