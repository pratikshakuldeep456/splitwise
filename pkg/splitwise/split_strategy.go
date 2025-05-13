package splitwise

type SplitStrategy interface {
	SplitAmount(paidby *User, amount float64, participants []*User, meta interface{}) []*Split
}
