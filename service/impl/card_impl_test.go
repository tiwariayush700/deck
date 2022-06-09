package impl

import (
	`context`
	`reflect`
	`testing`

	`github.com/stretchr/testify/assert`
	`gorm.io/gorm`

	`deck/core/api`
	`deck/core/constant`
	`deck/core/database`
	`deck/mocks`
	`deck/model`
)

const (
	testShufflePartialTrue = "TestShufflePartialTrue"
	testNoInputCards       = "TestNoInputCards"
	testIncorrectCardCode  = "TestIncorrectCardCode"
	testPartial            = "TestPartial"
)

const (
	testInvalidDeckId = "TestInvalidDeckId"
	testDrawCards     = "TestDrawCards"
)

func Test_cardServiceImpl_CreateDeck(t *testing.T) {
	type fields struct {
		CardRepository *mocks.CardRepository
	}
	type args struct {
		ctx     context.Context
		request api.DeckRequest
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		want       *model.Deck
		wantErr    bool
		wantErrVal error
	}{
		{
			name: testShufflePartialTrue,
			fields: fields{
				CardRepository: &mocks.CardRepository{},
			},
			args: args{
				ctx: context.Background(),
				request: api.DeckRequest{
					Shuffle: true,
					Partial: true,
					Cards:   []string{"AS", "5S"},
				},
			},
			want:       nil,
			wantErr:    true,
			wantErrVal: api.NewHTTPError(api.ErrorCodeInvalidRequestPayload, "You cannot have partial and shuffle both"),
		},
		{
			name: testNoInputCards,
			fields: fields{
				CardRepository: &mocks.CardRepository{},
			},
			args: args{
				ctx: context.Background(),
				request: api.DeckRequest{
					Partial: true,
				},
			},
			want:       nil,
			wantErr:    true,
			wantErrVal: api.NewHTTPError(api.ErrorCodeInvalidRequestPayload, "Please provide cards if you want a partial deck"),
		},
		{
			name: testIncorrectCardCode,
			fields: fields{
				CardRepository: &mocks.CardRepository{},
			},
			args: args{
				ctx: context.Background(),
				request: api.DeckRequest{
					Shuffle: false,
					Partial: true,
					Cards:   []string{"ASAP", "5S"},
				},
			},
			want:       nil,
			wantErr:    true,
			wantErrVal: api.NewHTTPError(api.ErrorCodeResourceNotFound, "No cards found for the provided codes"),
		},
		{
			name: testPartial,
			fields: fields{
				CardRepository: &mocks.CardRepository{},
			},
			args: args{
				ctx: context.Background(),
				request: api.DeckRequest{
					Shuffle: false,
					Partial: true,
					Cards:   []string{"AS", "5S"},
				},
			},
			want: &model.Deck{
				BaseModel: database.BaseModel{},
				Shuffled:  false,
				Remaining: 2,
			},
			wantErr:    false,
			wantErrVal: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewCardServiceImpl(tt.fields.CardRepository)

			switch tt.name {
			case testIncorrectCardCode:
				tt.fields.CardRepository.On("QueryInCards", tt.args.ctx, tt.args.request.Cards).Return(nil, database.NewError(gorm.ErrRecordNotFound))
			case testPartial:
				tt.fields.CardRepository.
					On("QueryInCards",
						tt.args.ctx,
						tt.args.request.Cards).
					Return(getPartialCardsData(), nil)
				tt.fields.CardRepository.
					On("CreateDeckCards",
						tt.args.ctx, getTestDeck(tt.args.request.Shuffle, 2),
						getPartialCardsData()).
					Return(nil)
			}

			got, err := c.CreateDeck(tt.args.ctx, tt.args.request)
			d := getTestDeck(tt.args.request.Shuffle, 2)
			if got != nil {
				got.BaseModel = d.BaseModel
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateDeck() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.ErrorIs(t, err, tt.wantErrVal)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateDeck() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_cardServiceImpl_DrawCards(t *testing.T) {
	type fields struct {
		CardRepository *mocks.CardRepository
	}
	type args struct {
		ctx     context.Context
		deckID  string
		cardIDs []string
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantErr    bool
		wantErrVal error
	}{
		{
			name: testInvalidDeckId,
			fields: fields{
				CardRepository: &mocks.CardRepository{},
			},
			args: args{
				ctx:     context.Background(),
				deckID:  "invalidDeckId",
				cardIDs: []string{"AS"},
			},
			wantErr:    true,
			wantErrVal: api.NewHTTPError(api.ErrorCodeResourceNotFound, "No deck found for the provided id"),
		},
		{
			name: testDrawCards,
			fields: fields{
				CardRepository: &mocks.CardRepository{},
			},
			args: args{
				ctx:     context.Background(),
				deckID:  "b74a19f0-e412-6756-7693-c3dc20ad7653",
				cardIDs: []string{"AS"},
			},
			wantErr:    false,
			wantErrVal: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewCardServiceImpl(tt.fields.CardRepository)
			switch tt.name {
			case testInvalidDeckId:
				tt.fields.CardRepository.On("DeleteCardsFromDeck", context.Background(), tt.args.deckID, tt.args.cardIDs).Return(database.NewError(gorm.ErrRecordNotFound))
			case testDrawCards:
				tt.fields.CardRepository.On("DeleteCardsFromDeck", context.Background(), tt.args.deckID, tt.args.cardIDs).Return(nil)
			}
			if err := c.DrawCards(tt.args.ctx, tt.args.deckID, tt.args.cardIDs); (err != nil) != tt.wantErr {
				assert.ErrorAs(t, err, tt.wantErrVal)
				t.Errorf("DrawCards() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_cardServiceImpl_GetDeckView(t *testing.T) {
	type fields struct {
		CardRepository *mocks.CardRepository
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
		{
			name: testInvalidDeckId,
			fields: fields{
				CardRepository: &mocks.CardRepository{},
			},
			args: args{
				ctx:    context.Background(),
				count:  constant.DefaultCardCount,
				deckID: "invalidDeckID",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &cardServiceImpl{
				CardRepository: tt.fields.CardRepository,
			}
			tt.fields.CardRepository.On("GetDeckView", context.Background(), tt.args.deckID, tt.args.count).Return(nil, database.NewError(gorm.ErrRecordNotFound))
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

func getPartialCardsData() []model.Card {
	return []model.Card{
		{
			Rank: model.Ace,
			Suit: model.Spade,
			Code: "AS",
		},
		{
			Rank: model.Five,
			Suit: model.Spade,
			Code: "5S",
		},
	}
}

func getTestDeck(shuffled bool, remaining int) *model.Deck {
	//deckID := ""

	return &model.Deck{
		//BaseModel: database.BaseModel{ID: deckID},
		Shuffled:  shuffled,
		Remaining: remaining,
	}
}
