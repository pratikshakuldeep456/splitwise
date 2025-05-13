package splitwise

type Transaction struct {
	Id     int
	User1  *User
	User2  *User
	Amount float64
}

func CreateTransaction(t Transaction) *Transaction {
	return &Transaction{}
}

// func (t *Transaction) Settleup(user1, user2 User) {

// }
