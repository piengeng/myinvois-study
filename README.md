# Study of MYInvois UBL-Invoice-2.1.xsd sdksamples

Purpose:

- to understand UBL-Invoice-2.1.xsd a bit more.
- to build e-invoice document block via functions, instead of xml-like-approach (eg. samples.go) document structure unlikely to change much due to UBL-2.1 anyway.
- as baseline for next steps.

```yaml
files:
  - .justfile # setup directory structure, download sdk/ubl # adjust path accordingly, same goes to *.go files
  - samples.go # to reproduce IRBM's sdksamples/*.xml
  - sampling.go # the functions/reduced approach
  - const_ubl.go # `just go-gen` make ubl's namespaces as const
```

Notes:

- Sequence is important when building UBL-Invoice-2.1.xsd's XML, so XSD validation before submission is required.
- UBL-CreditNote-2.1.xsd and UBL-DebitNote-2.1.xsd should not be confused with MyInvois's Credit/Debit/Refund Note, as these UBL's specs are not used in MyInvois, there's no mention of this in the SDK site.
- you might find 1 of their sample XML failing XSD validation, I've emailed them about this, it's not my role to force them to improve their own documentation.
- When IRBM force businesses to use digital signature (1.0 vs 1.1) on e-invoice, expect minimum RM1500/year/certificate on self-generated e-invoice. Still looking for low-cost option, haven't found any.
- misleading wording "using algorithms like SHA256" for documentHash, but must be SHA256, if allow other algorithms, the param will include options for other algorithms.
- some fields/tags like contact>telephone|electronicMail is also a MUST when generating documents, except consolidated, then 'NA' is used.
- PDK 2022 / Harmonized-System (HS) coding are not included. It was refered from [WCO](https://www.wcoomd.org/en/topics/nomenclature/instrument-and-tools.aspx) then localized/modified to Malaysia.
- maybe a better [LHDN MyInvois SDK.postman_collection.json](https://github.com/user-attachments/files/20915041/LHDN.MyInvois.SDK.postman_collection.json)
- for QR validation link: portalUrl + docUuid + /share/ + docLongId eg. `https://preprod.myinvois.hasil.gov.my/PP2...J10/share/9P8...557`. /api/v1.0/documents/:uuid/details key=longId but /api/v1.0/documents/:uuid/raw key=longID 🤷
- `just go-gen` to generate various go files, python3 and sqlc are required.
- go.mod for dependencies, valkey for simple caching, sqlite3 for simple submissions/sql storage.
- UBL2.1 is just the base schema, tax information/amount is mandatory, IRBM use-case, do'h.
- to achieve these few basic processes, build e-invoice xml documents, validation, submissions, storage, check/update status, generate validation link for QR in pdf.
- `config.yml.example` for reference, if you're testing this out.
- automation with cron for token renewal, to consider multiple api client/secret pairs, possible workaround on rate-limit.

### end of study, sorry if it's hard for non-techie to understand the code behind.
