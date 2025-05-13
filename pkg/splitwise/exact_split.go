package splitwise

import "fmt"

type ExactSplit struct {
}

func (e *ExactSplit) SplitAmount(paidby *User, amount float64, participants []*User, meta interface{}) []*Split {
	exactAmounts, ok := meta.([]float64)
	if !ok || len(exactAmounts) != len(participants) {
		fmt.Printf("Invalid meta data for exact split")
	}

	Splits := []*Split{}
	for i, user := range participants {
		Splits = append(Splits, &Split{UserID: user.ID, Amount: exactAmounts[i]})
	}
	return Splits
}
