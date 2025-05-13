package main

import "pratikshakuldeep456/splitwise/pkg/splitwise"

func main() {
	expenseService := splitwise.ExpenseService{}
	splitService := splitwise.NewSplitService(&expenseService)
	user1 := splitwise.CreateAccount("Alice", "alice@example.com", "1234x")
	user2 := splitwise.CreateAccount("Bob", "bob@example.com", "1234x")
	user3 := splitwise.CreateAccount("Charlie", "charlie@example.com", "1222")
	splitService.AddUser(user1)
	splitService.AddUser(user2)
	splitService.AddUser(user3)
	group := splitwise.CreateGroup("trip", []*splitwise.User{user1, user2})
	splitService.AddGroup(group)
	splitService.AddUserToGroup(group.ID, user3)

	splitService.AddExpense(group.ID, &splitwise.EqualSplit{}, user1, 1000, "goa food", []*splitwise.User{user1, user2}, nil)
	splitService.PrintBalance(user1.ID)
	splitService.Settleup(user1.ID, user2.ID, 50)
	splitService.PrintBalance(user1.ID)
}
