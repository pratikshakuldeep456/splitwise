package splitwise

var nextID int = 1

func GenerateExpenseID() int {
	id := nextID
	nextID++
	return id
}
