package splitwise

type User struct {
	ID       int
	Email    string
	PhoneNo  string
	Balances map[int]float64
}

func CreateAccount() *User {
	return &User{}
}
