package processor

import (
	"fmt"
	"log"
	"sort"

	expensetype "github.com/singh422/FinanceApp/internal/expense_type"
	"github.com/singh422/FinanceApp/internal/format"
	"github.com/singh422/FinanceApp/internal/writer"
)

type Summary struct {
	totalCredit  float64
	totalExpense float64
	totalPayment float64
}

type Expenses []format.Expense

type MonthlyReport struct {
	Expenses Expenses
	Summary  Summary
}

type SummaryReport map[string]*MonthlyReport

func CreateSummaryReport(allExpenses []format.Expense) SummaryReport {
	summary := make(SummaryReport)

	for _, expense := range allExpenses {
		expenseKey := expense.GetMapKey()

		// If it exists, then add to that months expenses
		if month, exists := summary[expenseKey]; exists {
			month.Expenses = append(month.Expenses, expense)
		} else {
			summary[expenseKey] = &MonthlyReport{
				Expenses: []format.Expense{expense},
			}
		}
	}

	summary.process()

	return summary
}

func (sr SummaryReport) process() {
	for _, expenses := range sr {
		expenses.sortExpenses()
		expenses.summarize()
	}
}

func (sr SummaryReport) OutputToCSV(dirPath string) {
	for month, monthlyReport := range sr {
		fileName := fmt.Sprintf("Expense Report-%s", month)
		err := writer.WriteExpensesToCSV(dirPath, fileName, monthlyReport.Expenses)
		if err != nil {
			log.Printf("ERROR: Failed to write report for %s to CSV %v", fileName, err)
		} else {
			log.Printf("INFO: Created summary report for %s in file name: %s", month, fileName)
		}
	}
}

func (sr SummaryReport) OutputSummary() {
	for month, monthlyReport := range sr {

		log.Printf("\n")
		log.Printf("\n")
		log.Printf("Monthly Report for:%s", month)
		log.Printf("\t%s", monthlyReport.Summary.String())
		log.Printf("\n")
		log.Printf("\n")
	}
}

func (m *MonthlyReport) sortExpenses() {
	sort.Sort(m.Expenses)
}

func (m *MonthlyReport) summarize() {
	var summary Summary
	for _, expense := range m.Expenses {
		if expense.Type == expensetype.Expense {
			summary.totalExpense += expense.Amount
		} else if expense.Type == expensetype.Credit {
			summary.totalCredit += expense.Amount
		} else if expense.Type == expensetype.Payment {
			summary.totalPayment += expense.Amount
		}
	}
	m.Summary = summary
}

func (s *Summary) String() string {
	return fmt.Sprintf("Total Expenses: [%f], Total Credits: [%f], Total Payments: [%f]", s.totalExpense, s.totalCredit, s.totalPayment)
}

func (e Expenses) Len() int {
	return len(e)
}

// Less defines the custom sorting order based on age
func (e Expenses) Less(i, j int) bool {
	// Sort Based on Source and Date
	if e[i].Source != e[j].Source {
		return e[i].Source.String() < e[j].Source.String()
	}
	return e[i].Date.Before(e[j].Date)
}

// Swap swaps the elements at positions i and j
func (e Expenses) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}
