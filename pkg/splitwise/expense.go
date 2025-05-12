package splitwise

type Expense struct {
	ID          int
	GroupID     int
	Amount      float64
	Description string
	PaidBy      *User
	Splits      []*Split
}
