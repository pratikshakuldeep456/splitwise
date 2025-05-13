package splitwise

import "fmt"

type Group struct {
	ID       int
	Name     string
	Members  map[int]*User
	Expenses []*Expense
}

func CreateGroup(name string, user []*User) *Group {

	members := make(map[int]*User)
	for _, u := range user {
		members[u.ID] = u
	}
	return &Group{ID: GenerateExpenseID(),
		Name:     name,
		Members:  members,
		Expenses: []*Expense{}}
}

func (g *Group) AddUser(user *User) {
	if _, exists := g.Members[user.ID]; exists {
		fmt.Printf("User %d already exists in the group.\n", user.ID)
		return
	}
	g.Members[user.ID] = user
}

func (g *Group) AddExpense(Amount float64, Description string, PaidBy *User, Splits []*Split) {
	expense := &Expense{
		ID:          GenerateExpenseID(),
		GroupID:     g.ID,
		Amount:      Amount,
		Description: Description,
		PaidBy:      PaidBy,
		Splits:      Splits,
	}
	g.Expenses = append(g.Expenses, expense)
}
