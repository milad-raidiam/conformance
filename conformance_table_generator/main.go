package main

import (
	"flag"
	"log"
	"strconv"

	"github.com/OpenBanking-Brasil/conformance/tree/main/conformance_table_generator/utils"
)

var (
	Target  string
	Version string
)

func init() {
	flag.StringVar(&Target, "t", "all", "Target Table")
	flag.StringVar(&Version, "v", "2", "API Version")
	flag.Parse()
}

// go run main.go -t <phaseNo> -v <versionNo>
// example
// go run main.go -t phase2 -v 2

// Add a .env file containing your github access token to the root of this project
// GITHUB_AT=ghp_...
func main() {
	var apis []string

	if Target == "phase2" || Target == "all" {
		apis = []string{
			"accounts", "business",
			"consents", "credit-card",
			"financings", "invoice-financings",
			"loans", "personal",
			"resources", "unarranged-overdraft",
		}

		utils.GenerateTable(apis, "phase2", Version)
	}

	if _, err := strconv.Atoi(Version); err != nil {
		log.Fatalf("Invalid version entered: %s. Error: %s", Version, err)
	}
}