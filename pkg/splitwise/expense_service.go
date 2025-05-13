package splitwise

type ExpenseService struct {
	strategy SplitStrategy
}

func NewExpenseService(strategy SplitStrategy) *ExpenseService {
	return &ExpenseService{strategy: strategy}
}

func (s *ExpenseService) AddExpense(groupID int,
	paidBy *User,
	amount float64,
	description string,
	participants []*User,
	meta interface{}) *Expense {
	splits := s.strategy.SplitAmount(paidBy, amount, participants, meta)
	expense := &Expense{
		ID:          GenerateExpenseID(),
		GroupID:     groupID,
		Amount:      amount,
		Description: description,
		PaidBy:      paidBy,
		Splits:      splits,
	}

	for _, split := range splits {
		if split.UserID != paidBy.ID {
			paidBy.Balances[split.UserID] += split.Amount
		}
	}

	return expense
}
