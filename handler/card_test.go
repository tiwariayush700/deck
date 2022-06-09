package handler

import (
	`context`
	`reflect`
	`testing`

	`github.com/gin-gonic/gin`

	`deck/core/config`
	`deck/service`
)

func TestCardHandler_CreateDecks(t *testing.T) {
	type fields struct {
		CardService   service.Card
		Configuration *config.Config
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   gin.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &CardHandler{
				CardService:   tt.fields.CardService,
				Configuration: tt.fields.Configuration,
			}
			if got := h.CreateDecks(tt.args.ctx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateDecks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCardHandler_DrawCards(t *testing.T) {
	type fields struct {
		CardService   service.Card
		Configuration *config.Config
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   gin.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &CardHandler{
				CardService:   tt.fields.CardService,
				Configuration: tt.fields.Configuration,
			}
			if got := h.DrawCards(tt.args.ctx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DrawCards() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCardHandler_GetDeck(t *testing.T) {
	type fields struct {
		CardService   service.Card
		Configuration *config.Config
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   gin.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &CardHandler{
				CardService:   tt.fields.CardService,
				Configuration: tt.fields.Configuration,
			}
			if got := h.GetDeck(tt.args.ctx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDeck() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCardHandler_RegisterRoutes(t *testing.T) {
	type fields struct {
		CardService   service.Card
		Configuration *config.Config
	}
	type args struct {
		ctx    context.Context
		router *gin.Engine
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &CardHandler{
				CardService:   tt.fields.CardService,
				Configuration: tt.fields.Configuration,
			}
		})
	}
}

func TestNewCardHandler(t *testing.T) {
	type args struct {
		cardService   service.Card
		configuration *config.Config
	}
	tests := []struct {
		name string
		args args
		want *CardHandler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCardHandler(tt.args.cardService, tt.args.configuration); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCardHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}
