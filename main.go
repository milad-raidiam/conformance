package main

import (
	"encoding/csv"
	"os"
	"strconv"
	"strings"

	"github.com/olekukonko/tablewriter"
)

func main() {

	//Open input file
	f, _ := os.Open("./data.csv")
	defer f.Close()

	//Read lines from file
	lines, _ := csv.NewReader(f).ReadAll()
	lines = lines[1:]

	//Set the table to output as a string
	tableOutput := &strings.Builder{}
	table := tablewriter.NewWriter(tableOutput)

	headers := []string{"Organisation", "Deployment", "Consentimento API", "Dados Cadastrais (PF) API", "Dados Cadastrais (PJ) API", "Resources API", "Contas API", "Cartão de Crédito API", "Operações de Crédito - Empréstimos API", "Operações de Crédito - Financiamentos API", "Operações de Crédito - Adiantamento a Depositantes API", "Operações de Crédito - Direitos Creditórios Descontados API"}

	//Configure table
	table.SetHeader([]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11"})
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetAutoWrapText(false)
	table.SetCenterSeparator("|")
	table.AppendBulk(lines)
	table.Render()

	//Open output file
	output, _ := os.OpenFile("./output.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0775)

	toWrite := tableOutput.String()

	//Replace headersxs
	for index, value := range headers {
		toWrite = strings.Replace(toWrite, strconv.Itoa(index), value, 1)
	}

	//Write result of table to file
	output.Write([]byte(toWrite))
	output.Close()
}
