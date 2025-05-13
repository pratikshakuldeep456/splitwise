package splitwise

type EqualSplit struct {
}

func (e *EqualSplit) SplitAmount(paidby *User, amount float64, participants []*User, meta interface{}) []*Split {
	length := len(participants)
	share := amount / float64(length)

	Splits := []*Split{}
	for _, user := range participants {
		Splits = append(Splits, &Split{UserID: user.ID, Amount: share})
	}
	return Splits
}
