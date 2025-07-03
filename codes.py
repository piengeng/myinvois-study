#!/usr/bin/env python3

import json
import re
import unicodedata
from pathlib import Path
from pprint import pprint

from caseconverter import pascalcase


def read_json(file_path):
    try:
        with open(file_path, "r") as file:
            data = json.load(file)
            return data
    except FileNotFoundError:
        print(f"Error: The file '{file_path}' was not found.")
    except json.JSONDecodeError:
        print(f"Error: Could not decode JSON from '{file_path}'. Check the file format.")
    except Exception as e:
        print(f"An unexpected error occurred: {e}")


# const -> code mapping
def main():
    lines = []
    lines.append("\t// gen'd by codes.py for Classification")
    lines.extend(doKeyKey(f_classification, "Description", "Code", "cCls"))
    lines.append("\t// gen'd by codes.py for Country")
    lines.extend(doKeyKey(f_country, "Country", "Code", "cCtry"))
    lines.append("\t// gen'd by codes.py for Currency")
    lines.extend(doKeyKey(f_currency, "Currency", "Code", "cCur"))
    lines.append("\t// gen'd by codes.py for EInvoiceTypes")
    lines.extend(doKeyKey(f_e_invoice_types, "Description", "Code", "cInv"))
    lines.append("\t// gen'd by codes.py for MSICSubCategories")
    lines.extend(doKeyKey(f_msic_sub_categories, "Description", "Code", "cMSIC"))
    lines.append("\t// gen'd by codes.py for PaymentMethods")
    lines.extend(doKeyKey(f_payment_methods, "Payment Method", "Code", "cPM"))
    lines.append("\t// gen'd by codes.py for State")
    lines.extend(doKeyKey(f_state, "State", "Code", "cSt"))
    lines.append("\t// gen'd by codes.py for TaxTypes")
    lines.extend(doKeyKey(f_tax_types, "Description", "Code", "cTax"))
    lines.append("\t// gen'd by codes.py for UnitTypes")
    lines.extend(doKeyKey(f_unit_types, "Name", "Code", "cU"))

    with open("const_codes.go", "w") as file:
        for item in wrap(lines):
            file.write(str(item) + "\n")


def doKeyKey(fp, k1, k2, prefix):
    c12n = read_json(fp)
    lines = []

    for i in c12n:
        # cCurPa’anga = "TOP"
        cleaned = re.sub(r"[\n\t]", " ", i[k1])
        line = unicodedata.normalize("NFKD", cleaned).encode("ascii", "ignore").decode("utf-8")
        name = f"{prefix  + pascalcase(line)}"
        # ignoring cCurLeone = "SLL"
        if i[k2] == "SLL" and name == "cCurLeone":
            name = name + "L"
        if i[k2] == "VED" and name == "cCurBolivarSoberano":
            name = name + "D"
        if i[k2] == "16211" and "cMSICManufactureOfVeneerSheetsAndPlywood" in lines[-1]:
            continue  # duplicated on prev.lines, not doing full array search
        if i[k2] == "52231" and name == "cMSICOperationOfTerminalFacilities":
            name = name + "_"
        if name in [
            "cUDenier",
            "cUBall",
            "cUCard",
            "cUKit",
            "cUPiece",
            "cUSet",
            "cUTablet",
            "cUTyre",
            "cUMutuallyDefined",
        ]:
            name = name + i[k2]
        lines.append(f'\t{name} = "{i[k2]}"')
    return lines


def wrap(lines, pkg="main"):
    head = [f"package {pkg}", "", "// nolint: unused", "const ("]
    return head + lines + [")"]


if __name__ == "__main__":
    # code -> const
    d_sdk_files = Path(__file__).parent / "refs/sdk.myinvois.hasil.gov.my/files"

    f_classification = d_sdk_files / "ClassificationCodes.json"
    f_country = d_sdk_files / "CountryCodes.json"
    f_currency = d_sdk_files / "CurrencyCodes.json"
    f_e_invoice_types = d_sdk_files / "EInvoiceTypes.json"
    f_msic_sub_categories = d_sdk_files / "MSICSubCategoryCodes.json"
    f_payment_methods = d_sdk_files / "PaymentMethods.json"
    f_state = d_sdk_files / "StateCodes.json"
    f_tax_types = d_sdk_files / "TaxTypes.json"
    f_unit_types = d_sdk_files / "UnitTypes.json"

    # PDK 2022 Perintah Duti Kastam / Customs Duty Order seen from InvoiceLine
    # https://lom.agc.gov.my/act-view.php?type=pua&no=P.U.%20(A)%20114/2022&lang=BI&language=BI
    # P.U. (A) 114_2022 (DUTI KASTAM).pdf' # refer manually then codify in to const
    # https://ezhs.customs.gov.my/public-home
    # https://github.com/datasets/harmonized-system
    # https://www.wcoomd.org/en/topics/nomenclature/instrument-and-tools.aspx
    f_product_tariff = d_sdk_files / "NotFound.json"  # Harmonized System ProductTariffCode
    main()
