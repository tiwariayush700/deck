package util

import (
	`fmt`
	`math/rand`
	`time`

	`deck/model`
)

func GetCode(suit model.Suit, rank model.Rank) string {
	if rank == model.Ace || rank == model.Jack || rank == model.Queen || rank == model.King {
		return fmt.Sprintf("%s%s", rank.String()[0:1], suit.String()[0:1])
	}

	return fmt.Sprintf("%d%s", rank, suit.String()[0:1])
}

func Shuffle(cards []model.Card) []model.Card {
	ret := make([]model.Card, len(cards))
	r := rand.New(rand.NewSource(time.Now().Unix()))
	perms := r.Perm(len(cards))
	//perms = { 4, 1, 3, 2, 0}
	for idx, perm := range perms {
		ret[idx] = cards[perm]
	}

	return ret
}
