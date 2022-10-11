from os import listdir
import re


def main():
    api_list = ["accounts", "business", "consents", "credit-card", "financings", "invoice-financings",
        "loans", "personal", "resources", "unarranged-overdraft"]
    version = "2.0.1"
    directories = [f"./submissions/functional/{api}/{version}" for api in api_list]

    zips = []
    for directory in directories:
        zips.extend(listdir(directory))

    zips = [element for element in zips if element != ".DS_Store"]

    pattern = re.compile(r"^\d{8}_.+_[a-z-]+_v[12]_(0[1-9]|[12]\d|3[01])-(0[1-9]|1[012])-(20\d\d).(zip|ZIP)$")
    wrong_zips = [file for file in zips if pattern.match(file) is None]

    if len(wrong_zips):
        print("The following zip names are wrong: " + str(wrong_zips))
        return 1
        
    return 0


if __name__ == '__main__':
    raise SystemExit(main())