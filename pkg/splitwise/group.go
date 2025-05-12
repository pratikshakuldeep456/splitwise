package splitwise

type Group struct {
	ID       int
	Name     string
	Members  map[int]*User
	Expenses []*Expense
}
