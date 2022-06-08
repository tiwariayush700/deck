package impl

import (
	`context`
	`reflect`
	`testing`

	`deck/core/api`
	`deck/model`
	`deck/repository`
	`deck/service`
)

func TestNewCardServiceImpl(t *testing.T) {
	type args struct {
		cardRepository repository.CardRepository
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCardServiceImpl(tt.args.cardRepository); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCardServiceImpl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_cardServiceImpl_CreateDeck(t *testing.T) {
	type fields struct {
		CardRepository repository.CardRepository
	}
	type args struct {
		ctx     context.Context
		request api.DeckRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Deck
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &cardServiceImpl{
				CardRepository: tt.fields.CardRepository,
			}
			got, err := c.CreateDeck(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateDeck() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateDeck() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_cardServiceImpl_DrawCards(t *testing.T) {
	type fields struct {
		CardRepository repository.CardRepository
	}
	type args struct {
		ctx     context.Context
		deckID  string
		cardIDs []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &cardServiceImpl{
				CardRepository: tt.fields.CardRepository,
			}
			if err := c.DrawCards(tt.args.ctx, tt.args.deckID, tt.args.cardIDs); (err != nil) != tt.wantErr {
				t.Errorf("DrawCards() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_cardServiceImpl_GetDeckView(t *testing.T) {
	type fields struct {
		CardRepository repository.CardRepository
	}
	type args struct {
		ctx    context.Context
		count  int
		deckID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.DeckView
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &cardServiceImpl{
				CardRepository: tt.fields.CardRepository,
			}
			got, err := c.GetDeckView(tt.args.ctx, tt.args.count, tt.args.deckID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetDeckView() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDeckView() got = %v, want %v", got, tt.want)
			}
		})
	}
}
