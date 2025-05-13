package splitwise

type User struct {
	ID          int
	Name        string
	Email       string
	PhoneNo     string
	Balances    map[int]float64
	Transaction map[int]*Transaction
}

func CreateAccount(name string, email string, phoneno string) *User {
	return &User{
		ID:          GenerateExpenseID(),
		Name:        name,
		Email:       email,
		PhoneNo:     phoneno,
		Balances:    map[int]float64{},
		Transaction: map[int]*Transaction{},
	}
}
