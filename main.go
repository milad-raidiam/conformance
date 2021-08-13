package main

import (
	"encoding/csv"
	"os"
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

	//Configure table
	table.SetHeader([]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12"})
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetAutoWrapText(false)
	table.SetCenterSeparator("|")
	table.AppendBulk(lines)
	table.Render()

	//Open output file
	output, _ := os.OpenFile("./output.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0775)
	defer output.Close()

	//Write result of table to file
	output.Write([]byte(tableOutput.String()))
}
