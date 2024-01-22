package expensetype

type ExpenseType string

const (
	Credit  ExpenseType = "Credit"
	Expense ExpenseType = "Expense"
	Payment ExpenseType = "Payment"
)

func (e ExpenseType) String() string {
	switch e {
	case Credit:
		return "Credit"
	case Expense:
		return "Expense"
	case Payment:
		return "Payment"
	}
	return ""
}
