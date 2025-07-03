#!/usr/bin/env python3

import re
import sys

pattern = r'"(?:urn|cbc|cac|ext):[^"]*"'
namespaces = []


def read_file(filename):
    try:
        with open(filename, "r") as f:
            text = f.read()
            matches = re.findall(r'"(?:urn|cbc|cac|ext):[^"]*"', text)
            uniques = list(set(matches))
            namespaces.extend(uniques)
            # print(f"Contents of {filename}:\n{content}\n")
    except FileNotFoundError:
        print(f"File not found: {filename}")
    except Exception as e:
        print(f"Error reading {filename}: {e}")


# to extract namespaces as const in go
def main():
    if len(sys.argv) < 2:
        print("Usage: ./ublns.py <file1> <file2> ...")
        return

    for file in sys.argv[1:]:
        read_file(file)

    lines = [f"\t{firstLast(l)} = {l}" for l in sorted(list(set(namespaces)))]
    # pprint(wrap(lines))
    with open("const_ubl.go", "w") as file:
        for item in wrap(lines):
            file.write(str(item) + "\n")


def firstLast(s):
    lst = s.replace('"', "").replace("-", "").split(":")
    return lst[0] + lst[-1]


def wrap(lines, pkg="main"):
    head = [f"package {pkg}", "", "// nolint: unused", "const ("]
    return head + lines + [")"]


if __name__ == "__main__":
    if len(sys.argv) == 1:
        sys.argv += [
            "refs/UBL-2.1/xsdrt/common/UBL-CommonAggregateComponents-2.1.xsd",
            "refs/UBL-2.1/xsdrt/common/UBL-CommonBasicComponents-2.1.xsd",
            "refs/UBL-2.1/xsdrt/common/UBL-CommonExtensionComponents-2.1.xsd",
            # "refs/UBL-2.1/xsdrt/common/UBL-CommonSignatureComponents-2.1.xsd",
            # "refs/UBL-2.1/xsdrt/maindoc/UBL-FreightInvoice-2.1.xsd",
            "refs/UBL-2.1/xsdrt/maindoc/UBL-Invoice-2.1.xsd",  # WHOLE THING USES THIS! NOT OTHERS!
            # "refs/UBL-2.1/xsdrt/maindoc/UBL-SelfBilledInvoice-2.1.xsd",
            # "refs/UBL-2.1/xsdrt/maindoc/UBL-CreditNote-2.1.xsd",
            # "refs/UBL-2.1/xsdrt/maindoc/UBL-SelfBilledCreditNote-2.1.xsd",
            # "refs/UBL-2.1/xsdrt/maindoc/UBL-DebitNote-2.1.xsd",
        ]
    main()
