package api

type DeckRequest struct {
	Shuffle bool     `json:"shuffle"`
	Partial bool     `json:"partial"`
	Cards   []string `json:"cards"`
}
