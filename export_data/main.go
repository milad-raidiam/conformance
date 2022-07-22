package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func exportData(apiFamilyTypes []string, fileName string) {
	// Creating the header for the table
	tableHeader := []string{"Brand Name"}
	for _, familyType := range apiFamilyTypes {
		tableHeader = append(tableHeader, translateName(familyType))
	}

	// Requesting data from the API participants endpoint
	data, err := importData("https://data.directory.openbankingbrasil.org.br/participants")
	if err != nil {
		log.Fatal("Failed to request data from the API:", err)
	}

	// Creating the csv file
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal("Failed to open file:", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Writing header to the file
	if err := writer.Write(tableHeader); err != nil {
		log.Fatal("Failed to write to file:", err)
	}

	// Writing data to the file
	for _, participant := range data {
		for _, server := range participant.AuthorisationServers {
			row_elements := make(map[string]string)
			row_elements["Brand name"] = server.CustomerFriendlyName

			for _, resource := range server.APIResources {
				if contains(apiFamilyTypes, resource.APIFamilyType) && len(resource.APIDiscoveryEndpoints) > 0 {
					row_elements[resource.APIFamilyType] = fmt.Sprintf(
						"[API URI] (%s), [Certification Plan] (%s)",
						resource.APIDiscoveryEndpoints[0].APIEndpoint,
						resource.APICertificationURI,
					)
				}
			}

			row := make([]string, len(apiFamilyTypes) + 1)
			row[0] = row_elements["Brand name"]
			for i, familyType := range apiFamilyTypes {
				row[i + 1] = row_elements[familyType]
			}

			if err := writer.Write(row); err != nil {
				log.Fatal("Failed to write to file:", err)
			}
		}
	}
}

func main() {
	// Family types that are going to be used once we have v4 data
	// apiFamilyTypes := []string{
	// 	"CustomerFriendlyName",
	// 	"opendata-investments_funds",
	// 	"opendata-investments_bank-fixed-incomes",
	// 	"opendata-investments_credit-fixed-incomes",
	// 	"opendata-investments_variable-incomes",
	// 	"opendata-investments_treasure-titles",
	// 	"opendata-capitalization_bonds",
	// 	"opendata-exchange_online-rates",
	// 	"opendata-exchange_vet-values",
	// 	"opendata-acquiring-services_personals",
	// 	"opendata-acquiring-services_businesses",
	// 	"opendata-pension_risk-coverages",
	// 	"opendata-pension_survival-coverages",
	// 	"opendata-insurance_automotives",
	// 	"opendata-insurance_homes",
	// 	"opendata-insurance_personals",
	// }

	// Provisory family types
	apiFamilyTypes := []string{
		"accounts",
		"admin",
		"channels",
		"consents",
		"credit-cards-accounts",
		"customers-business",
		"customers-personal",
		"discovery",
		"financings",
		"invoice-financings",
		"loans",
		"payments-consents",
		"payments-pix",
		"products-services",
		"resources",
		"unarranged-accounts-overdraft",
	}

	exportData(apiFamilyTypes, "../phase4-data.csv")
}



