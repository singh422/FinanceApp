package format

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/singh422/FinanceApp/internal/category"
	expensetype "github.com/singh422/FinanceApp/internal/expense_type"
	"github.com/singh422/FinanceApp/internal/source"
	"github.com/singh422/FinanceApp/internal/util"
)

type Metadata struct {
	BankCategory string
	BankType     string
	FileName     string
}

type Expense struct {
	Type        expensetype.ExpenseType
	Source      source.Source
	Description string
	Amount      float64
	Date        time.Time
	Category    category.Category
	Metadata    *Metadata
}

func ConvertRecord(user string, src source.Source, record []string) (*Expense, error) {
	switch src {
	case source.Chase:
		return convertChaseRecord(record)
	case source.ChaseShared:
		return convertChaseSharedRecord(record)
	case source.AMEX:
		return convertAmexRecord(user, record)
	case source.Apple:
		return convertAppleRecord(record)
	case source.Discover:
		return convertDiscoverRecord(record)
	case source.Paypal:
		return convertPaypalRecord(record)
	case source.Venmo:
		return convertVenmoRecord(user, record)
	}
	return nil, fmt.Errorf("ERROR: Unrecognized source, cannot convert record %v", record)
}

func convertAmexRecord(user string, record []string) (*Expense, error) {
	offset := 0
	if user == "Avantika Bagri" {
		offset = 2
	}

	amt, err := strconv.ParseFloat(record[2+offset], 64)
	if err != nil {
		return nil, err
	}

	dateString := record[0]
	parsedTime, err := time.Parse("01/02/2006", dateString)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	expense := Expense{
		Source:      source.AMEX,
		Amount:      amt,
		Description: record[1],
		Date:        parsedTime,
		Metadata: &Metadata{
			BankCategory: record[10+offset],
		},
	}

	if amt > 0 {
		expense.Type = expensetype.Expense
	} else if util.CaseInsensitiveSubstring(expense.Description, "Online Payment - Thank You") {
		expense.Type = expensetype.Payment
	} else {
		expense.Type = expensetype.Credit
	}

	return &expense, nil
}

func convertAppleRecord(record []string) (*Expense, error) {
	amt, err := strconv.ParseFloat(record[6], 64)
	if err != nil {
		return nil, err
	}

	dateString := record[0]
	parsedTime, err := time.Parse("01/02/2006", dateString)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	expense := Expense{
		Source:      source.Apple,
		Amount:      amt,
		Description: record[2],
		Date:        parsedTime,
		Metadata: &Metadata{
			BankCategory: record[4],
			BankType:     record[5],
		},
	}

	if record[5] == "Purchase" {
		expense.Type = expensetype.Expense
	} else if record[5] == "Debit" {
		expense.Type = expensetype.Credit
	} else if record[5] == "Payment" {
		expense.Type = expensetype.Payment
	} else if record[5] == "Credit" {
		expense.Type = expensetype.Credit
	} else {
		log.Fatalf("FATAL: Unrecognized original BankType: [%s] for Apple please add correct Type.", record[5])
	}

	return &expense, nil
}

func convertChaseRecord(record []string) (*Expense, error) {
	amt, err := strconv.ParseFloat(record[5], 64)
	if err != nil {
		return nil, err
	}

	dateString := record[0]
	parsedTime, err := time.Parse("01/02/2006", dateString)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	bankType := record[4]
	expenseType := expensetype.Expense

	if bankType == "Sale" {
		expenseType = expensetype.Expense
	} else if bankType == "Payment" {
		expenseType = expensetype.Payment
	} else if bankType == "Return" {
		expenseType = expensetype.Credit
	} else if bankType == "Adjustment" {
		expenseType = expensetype.Credit
	} else {
		log.Fatalf("FATAL: Unrecognized original BankType for Chase please add correct Type.")
	}

	expense := Expense{
		Type:        expenseType,
		Source:      source.Chase,
		Amount:      -1 * amt,
		Description: record[2],
		Date:        parsedTime,
		Metadata: &Metadata{
			BankCategory: record[3],
			BankType:     bankType,
		},
	}

	return &expense, nil
}

func convertChaseSharedRecord(record []string) (*Expense, error) {
	amt, err := strconv.ParseFloat(record[5], 64)
	if err != nil {
		return nil, err
	}

	dateString := record[0]
	parsedTime, err := time.Parse("01/02/2006", dateString)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	bankType := record[4]
	expenseType := expensetype.Expense

	if bankType == "Sale" {
		expenseType = expensetype.Expense
	} else if bankType == "Payment" {
		expenseType = expensetype.Payment
	} else if bankType == "Return" {
		expenseType = expensetype.Credit
	} else if bankType == "Adjustment" {
		expenseType = expensetype.Credit
	} else {
		log.Fatalf("FATAL: Unrecognized original BankType for Chase Shared please add correct Type.")
	}

	expense := Expense{
		Type:        expenseType,
		Source:      source.ChaseShared,
		Amount:      (-1 * amt) / 2,
		Description: record[2],
		Date:        parsedTime,
		Metadata: &Metadata{
			BankCategory: record[3],
			BankType:     bankType,
		},
	}

	return &expense, nil
}

func convertDiscoverRecord(record []string) (*Expense, error) {
	amt, err := strconv.ParseFloat(record[3], 64)
	if err != nil {
		return nil, err
	}

	dateString := record[0]
	parsedTime, err := time.Parse("01/02/2006", dateString)
	if err != nil {
		return nil, err
	}

	expense := Expense{
		Source:      source.Discover,
		Amount:      amt,
		Description: record[2],
		Date:        parsedTime,
		Metadata: &Metadata{
			BankCategory: record[4],
		},
	}

	if record[4] == "Payments and Credits" {
		expense.Type = expensetype.Payment
	} else if amt < 0 {
		expense.Type = expensetype.Credit
	} else {
		expense.Type = expensetype.Expense
	}

	return &expense, nil
}

func convertPaypalRecord(record []string) (*Expense, error) {
	amt, err := strconv.ParseFloat(util.ConvertStringAmountToFloat(record[2]), 64)
	if err != nil {
		return nil, err
	}

	dateString := record[0]
	parsedTime, err := time.Parse("1/2/2006", dateString)
	if err != nil {
		return nil, err
	}

	expense := Expense{
		Source:      source.Paypal,
		Amount:      amt,
		Description: record[3],
		Date:        parsedTime,
		Metadata:    &Metadata{},
	}

	if record[3] == "Payment" {
		expense.Type = expensetype.Payment
	} else if amt < 0 {
		expense.Type = expensetype.Credit
	} else {
		expense.Type = expensetype.Expense
	}

	return &expense, nil
}

func convertVenmoRecord(user string, record []string) (*Expense, error) {
	amt, err := strconv.ParseFloat(util.ConvertStringAmountToFloat(record[8]), 64)
	if err != nil {
		return nil, err
	}

	dateString := record[2]
	parsedTime, err := time.Parse("2006-01-02T15:04:05", dateString)
	if err != nil {
		return nil, err
	}

	bankType := record[3]
	fromPerson := record[6]
	toPerson := record[7]

	description := fmt.Sprintf("%s - %s From: %s, To: %s", record[5], bankType, fromPerson, toPerson)

	expense := Expense{
		Source:      source.Venmo,
		Description: description,
		Date:        parsedTime,
		Metadata: &Metadata{
			BankType: bankType,
		},
	}

	if amt < 0 {
		amt = amt * -1
	}

	if (bankType == "Charge" && fromPerson == user) || (bankType == "Payment" && toPerson == user) {
		expense.Type = expensetype.Credit
		expense.Amount = -1 * amt
	} else if (bankType == "Charge" && toPerson == user) || (bankType == "Payment" && fromPerson == user) {
		expense.Type = expensetype.Expense
		expense.Amount = amt
	} else {
		expense.Type = expensetype.Payment
	}

	return &expense, nil
}

func (e *Expense) GetMapKey() string {
	transactionDate := e.Date
	return transactionDate.Format("Jan 2006")
}
