package splitwise

type SplitStrategy interface {
	SplitAmount(paidby *User, amount float64, participants []*User)
}

type ExpenseService struct {
	strategy SplitStrategy
}

func NewExpenseService(strategy SplitStrategy) *ExpenseService {
	return &ExpenseService{strategy: strategy}
}

func (s *ExpenseService) AddExpense() {

}
