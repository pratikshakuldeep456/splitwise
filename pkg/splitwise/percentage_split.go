package splitwise

import "fmt"

type PercentageSplit struct {
}

func (p *PercentageSplit) SplitAmount(paidby *User, amount float64, participants []*User, meta interface{}) []*Split {
	cents, ok := meta.([]float64)
	if !ok || len(cents) != len(participants) {
		fmt.Printf("Invalid meta data for exact split")
	}
	Splits := []*Split{}
	for i, user := range participants {
		Splits = append(Splits, &Split{UserID: user.ID, Amount: (cents[i] / 100.0) * amount})
	}
	return Splits

}
