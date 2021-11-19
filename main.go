package main

import (
	"encoding/csv"
	"flag"
	"os"
	"strconv"
	"strings"

	"github.com/olekukonko/tablewriter"
)

var (
	Target string
)

/*
* To generate all tables - go run main.go
* To generate phase 2 table - go run main.go -t phase2
* To generate phase 3 table - go run main.go -t phase3
 */

func init() {
	flag.StringVar(&Target, "t", "all", "Target Table")
	flag.Parse()
}

func main() {
	if Target == "phase2" || Target == "all" {
		generateFromCsv("./phase2-data.csv", "./phase2-output.txt", []string{"Organisation", "Deployment", "Consentimento API", "Dados Cadastrais (PF) API", "Dados Cadastrais (PJ) API", "Resources API", "Contas API", "Cartão de Crédito API", "Operações de Crédito - Empréstimos API", "Operações de Crédito - Financiamentos API", "Operações de Crédito - Adiantamento a Depositantes API", "Operações de Crédito - Direitos Creditórios Descontados API"})
	}
	if Target == "phase3" || Target == "all" {
		generateFromCsv("./phase3-data.csv", "./phase3-output.txt", []string{"Organisation", "Deployment", "Payments"})
	}
}

func generateFromCsv(inputFile string, outputFile string, headers []string) {
	f, _ := os.Open(inputFile)
	defer f.Close()

	//Read lines from file
	lines, _ := csv.NewReader(f).ReadAll()
	lines = lines[1:]

	//Set the table to output as a string
	tableOutput := &strings.Builder{}
	table := tablewriter.NewWriter(tableOutput)

	var indexHeaders []string

	for index := range headers {
		indexHeaders = append(indexHeaders, strconv.Itoa(index))
	}

	//Configure table
	table.SetHeader(indexHeaders)
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetAutoWrapText(false)
	table.SetCenterSeparator("|")
	table.AppendBulk(lines)
	table.Render()

	//Open output file
	output, _ := os.OpenFile(outputFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0775)

	toWrite := tableOutput.String()

	//Replace headersxs
	for index, value := range headers {
		toWrite = strings.Replace(toWrite, strconv.Itoa(index), value, 1)
	}

	//Write result of table to file
	output.Write([]byte(toWrite))
	output.Close()
}
