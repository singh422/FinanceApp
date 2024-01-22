package writer

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/singh422/FinanceApp/internal/format"
)

const dirPath = "/Users/Avinash/Desktop/Expenses/Output Reports"

func WriteExpensesToCSV(dirPath string, fileName string, data []format.Expense) error {

	// Create the directory along with any necessary parent directories
	err := os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		return err
	}

	// Create a new CSV file
	file, err := os.Create(fmt.Sprintf("%s/%s.csv", dirPath, fileName))
	if err != nil {
		return err
	}
	defer file.Close()

	// Create a CSV writer
	csvWriter := csv.NewWriter(file)

	// Write the header row with field names
	headerRow := []string{"Type", "Source", "Description", "Amount", "Date", "Category"}
	if err := csvWriter.Write(headerRow); err != nil {
		return err
	}

	// Write data rows with field values
	for _, exp := range data {
		dataRow := []string{
			exp.Type.String(),
			exp.Source.String(),
			exp.Description,
			strconv.FormatFloat(exp.Amount, 'f', 2, 64),
			exp.Date.Format("2006-01-02"),
			exp.Category.String(),
		}
		if err := csvWriter.Write(dataRow); err != nil {
			return err
		}
	}

	// Flush the CSV writer to ensure all data is written
	csvWriter.Flush()

	// Check for errors during flushing
	if err := csvWriter.Error(); err != nil {
		return err
	}

	return nil
}
