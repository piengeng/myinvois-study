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
- misleading wording "using algorithms like SHA256" for documentHash, it must be SHA256, if allow other algorithms, the param will include options for other algorithms. 
- some fields/tags like contact>telephone|electronicMail is also a MUST when generating documents, except consolidated, then 'NA' is used.
