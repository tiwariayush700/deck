package util

import (
	`fmt`

	`deck/model`
)

func GetCode(suit model.Suit, rank model.Rank) string {
	if rank == model.Ace || rank == model.Jack || rank == model.Queen || rank == model.King {
		return fmt.Sprintf("%s%s", rank.String()[0:1], suit.String()[0:1])
	}

	return fmt.Sprintf("%d%s", rank, suit.String()[0:1])
}
