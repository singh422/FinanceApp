package main

import (
	"fmt"
	"log"
	"os"

	"github.com/singh422/FinanceApp/internal/format"
	"github.com/singh422/FinanceApp/internal/processor"
	"github.com/singh422/FinanceApp/internal/reader"
)

func main() {

	arg := os.Args[1:]

	user := "Avinash Singh"

	if len(arg) > 0 && arg[0] == "AB" {
		user = "Avantika Bagri"
	}

	directoryPath := fmt.Sprintf("/Users/Avinash/Desktop/Expenses/%s", user)

	// Read
	reader := reader.Reader{
		DirectoryPath: directoryPath,
	}
	files, err := reader.ReadCSVFiles(directoryPath)
	if err != nil {
		fmt.Println("Error:", err)
	}

	var unifiedExpenseRecords []format.Expense

	// Convert
	for _, fileInfo := range files {
		fileExpenses := unifyFileRecords(user, *fileInfo)
		unifiedExpenseRecords = append(unifiedExpenseRecords, fileExpenses...)
	}

	// Processor
	expenseReport := processor.CreateSummaryReport(unifiedExpenseRecords)

	// Output
	expenseReport.OutputToCSV(fmt.Sprintf("%s/Output Reports", directoryPath))
	expenseReport.OutputSummary()
}

func unifyFileRecords(user string, fileInfo reader.FileInfo) (expenses []format.Expense) {
	for i, record := range fileInfo.Records {

		// Skip Header row
		if i == 0 {
			continue
		}

		convertedExpense, err := format.ConvertRecord(user, fileInfo.Source, record)
		if err != nil {
			log.Printf("WARNING: Failed to parse record %d for FileName: %s, Source: %s\nError: %v", i, fileInfo.FileName, fileInfo.Source, err)
			continue
		}

		if convertedExpense.Metadata != nil {
			convertedExpense.Metadata.FileName = fileInfo.FileName
		}
		expenses = append(expenses, *convertedExpense)
	}

	return
}
