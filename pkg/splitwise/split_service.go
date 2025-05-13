package splitwise

import (
	"fmt"
	"sync"
)

type SplitService struct {
	Users  map[int]*User
	Groups map[int]*Group

	Mu             *sync.Mutex
	ExpenseService *ExpenseService
}

var SplitServiceInstance *SplitService
var Once sync.Once

func NewSplitService(expenseService *ExpenseService) *SplitService {
	return &SplitService{
		Users:  make(map[int]*User),
		Groups: make(map[int]*Group),

		Mu:             &sync.Mutex{},
		ExpenseService: expenseService,
	}
}

func InitSplitService(expenseService *ExpenseService) *SplitService {
	Once.Do(func() {
		SplitServiceInstance = NewSplitService(expenseService)
	})
	return SplitServiceInstance
}

func (ss *SplitService) AddUser(u *User) {
	ss.Users[u.ID] = u

}
func (ss *SplitService) AddGroup(g *Group) {
	ss.Groups[g.ID] = g
}

// func (ss *SplitService) SplitExpense(groupID int,
// 	paidBy *User,
// 	amount float64,
// 	description string,
// 	participants []*User,
// 	meta interface{}) {
// 	ss.Mu.Lock()
// 	defer ss.Mu.Unlock()

// 	group, exists := ss.Groups[groupID]
// 	if !exists {
// 		fmt.Println("Group does not exist")
// 		return
// 	}

// 	splits := ss.ExpenseService.strategy.SplitAmount(paidBy, amount, participants, meta)
// 	group.AddExpense(amount, description, paidBy, splits)

// }
func (ss *SplitService) AddExpense(groupID int,
	strategy SplitStrategy,
	paidBy *User,
	amount float64,
	description string,
	participants []*User,
	meta interface{}) {
	group, ok := ss.Groups[groupID]
	if !ok {
		fmt.Printf("Group with ID %d does not exist\n", groupID)
		return
	}
	expenseService := NewExpenseService(strategy)
	expense := expenseService.AddExpense(groupID, paidBy, amount, description, participants, meta)
	group.AddExpense(amount, description, paidBy, expense.Splits)
	ss.UpdateBalance(expense)

}

func (ss *SplitService) UpdateBalance(e *Expense) {

	// splits := e.Splits

	for _, exists := range e.Splits {
		payer := e.PaidBy.ID
		rec := exists.UserID
		if payer == rec {
			continue
		}
		e.PaidBy.Balances[rec] += exists.Amount
		ss.Users[rec].Balances[payer] -= exists.Amount
	}

}
func (ss *SplitService) AddUserToGroup(gId int, user *User) {
	group, exists := ss.Groups[gId]
	if !exists {
		fmt.Printf("Group with ID %d does not exist\n", gId)
		return
	}
	group.AddUser(user)
	// ss.Groups[gId].AddUser(user)

}

func (ss *SplitService) Settleup(sID, recId int, amount float64) {
	sender := ss.Users[sID]
	rec := ss.Users[recId]
	if sender.Balances[recId] < amount {
		fmt.Println("Insufficient balance to settle")
		return
	}
	sender.Balances[recId] -= amount
	rec.Balances[sID] += amount

	txn := &Transaction{Id: GenerateExpenseID(),
		User1:  sender,
		User2:  rec,
		Amount: amount,
	}
	sender.Transaction[txn.Id] = txn
	rec.Transaction[txn.Id] = txn
}

func (ss *SplitService) PrintBalance(userID int) {
	user, ok := ss.Users[userID]
	if !ok {
		fmt.Println("User not found")
		return
	}
	for otherID, balance := range user.Balances {
		if balance > 0 {
			fmt.Printf("%s owes %s: %.2f\n", ss.Users[otherID].Name, user.Name, balance)
		} else if balance < 0 {
			fmt.Printf("%s owes %s: %.2f\n", user.Name, ss.Users[otherID].Name, -balance)
		}
	}
}
