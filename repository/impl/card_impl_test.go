package impl

import (
	`context`
	`reflect`
	`testing`

	`github.com/hashicorp/go-uuid`
	`gorm.io/gorm`

	`deck/core/config`
	`deck/core/constant`
	`deck/core/database`
	`deck/core/test`
	`deck/model`
	`deck/repository`
)

const (
	testInvalidDeckId           = "testInvalidDeckId"
	testDeleteCardsFromDeck     = "TestDeleteCardsFromDeck"
	testGetCardView             = "TestCardView"
	testQueryCardsInvalidFilter = "TestQueryCardsInvalidFilter"
	testQueryCards              = "TestQueryCards"
)

var db *gorm.DB

func init() {
	db, _ = test.DB()
}

func TestNewCardRepositoryImpl(t *testing.T) {
	type args struct {
		db *gorm.DB
	}
	tests := []struct {
		name string
		args args
		want repository.CardRepository
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCardRepositoryImpl(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCardRepositoryImpl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewRepositoryImpl(t *testing.T) {
	type args struct {
		pgConfig config.PGConfig
	}
	tests := []struct {
		name    string
		args    args
		want    repository.Repository
		want1   *gorm.DB
		wantErr bool
	}{
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := NewRepositoryImpl(tt.args.pgConfig)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewRepositoryImpl() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRepositoryImpl() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("NewRepositoryImpl() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_cardRepositoryImpl_CreateDeckCards(t *testing.T) {
	type fields struct {
		repositoryImpl repositoryImpl
		db             *gorm.DB
	}
	type args struct {
		ctx   context.Context
		deck  *model.Deck
		cards []model.Card
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "TestInvalidQuery",
			fields: fields{
				repositoryImpl: repositoryImpl{
					db: db,
				},
				db: db,
			},
			args: args{
				ctx: context.Background(),
				deck: &model.Deck{
					Shuffled:  false,
					Remaining: 1,
				},
				cards: []model.Card{
					{
						Rank: model.Ace,
						Suit: model.Spade,
						Code: "AS",
					},
				},
			},
			wantErr: true,
		},
		{
			name: "TestCreateDeckCards",
			fields: fields{
				repositoryImpl: repositoryImpl{
					db: db,
				},
				db: db,
			},
			args: args{
				ctx: context.Background(),
				deck: &model.Deck{
					Shuffled:  false,
					Remaining: 1,
				},
				cards: []model.Card{
					{
						BaseModel: database.BaseModel{ID: "AS"},
						Suit:      model.Spade,
						Rank:      model.Ace,
						Code:      "AS",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &cardRepositoryImpl{
				repositoryImpl: tt.fields.repositoryImpl,
				db:             tt.fields.db,
			}
			err := c.CreateDeckCards(tt.args.ctx, tt.args.deck, tt.args.cards)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateDeckCards() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_cardRepositoryImpl_DeleteCardsFromDeck(t *testing.T) {
	type fields struct {
		repositoryImpl repositoryImpl
		db             *gorm.DB
	}
	deckID1, _ := uuid.GenerateUUID()
	deckID2, _ := uuid.GenerateUUID()
	type createDeckCardArgs struct {
		ctx   context.Context
		deck  *model.Deck
		cards []model.Card
	}
	tests := []struct {
		name               string
		fields             fields
		createDeckCardArgs createDeckCardArgs
		err1Want           bool
		err2Want           bool
	}{
		{
			name: testInvalidDeckId,
			fields: fields{
				repositoryImpl: repositoryImpl{
					db: db,
				},
				db: db,
			},
			createDeckCardArgs: createDeckCardArgs{
				ctx: context.Background(),
				deck: &model.Deck{
					BaseModel: database.BaseModel{
						ID: deckID1,
					},
					Shuffled:  false,
					Remaining: 1,
				},
				cards: []model.Card{
					{
						BaseModel: database.BaseModel{ID: "AS"},
						Suit:      model.Spade,
						Rank:      model.Ace,
						Code:      "AS",
					},
				},
			},
			err1Want: false,
			err2Want: true,
		},
		{
			name: testDeleteCardsFromDeck,
			fields: fields{
				repositoryImpl: repositoryImpl{
					db: db,
				},
				db: db,
			},
			createDeckCardArgs: createDeckCardArgs{
				ctx: context.Background(),
				deck: &model.Deck{
					BaseModel: database.BaseModel{
						ID: deckID2,
					},
					Shuffled:  false,
					Remaining: 1,
				},
				cards: []model.Card{
					{
						BaseModel: database.BaseModel{ID: "AS"},
						Suit:      model.Spade,
						Rank:      model.Ace,
						Code:      "AS",
					},
				},
			},
			err1Want: false,
			err2Want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &cardRepositoryImpl{
				repositoryImpl: tt.fields.repositoryImpl,
				db:             tt.fields.db,
			}

			err := c.CreateDeckCards(tt.createDeckCardArgs.ctx, tt.createDeckCardArgs.deck, tt.createDeckCardArgs.cards)
			if (err != nil) != tt.err1Want {
				t.Errorf("CreateDeckCards() error = %v, wantErr %v", err, tt.err1Want)
				return
			}

			cardIDs := make([]string, 0)
			for _, card := range tt.createDeckCardArgs.cards {
				cardIDs = append(cardIDs, card.ID)
			}

			switch tt.name {
			case testInvalidDeckId:
				cardIDs = nil
			}

			err = c.DeleteCardsFromDeck(tt.createDeckCardArgs.ctx, deckID2, cardIDs)
			if (err != nil) != tt.err2Want {
				t.Errorf("DeleteCardsFromDeck() error = %v, wantErr %v", err, tt.err2Want)
				return
			}
		})
	}
}

func Test_cardRepositoryImpl_GetDeckView(t *testing.T) {
	type fields struct {
		repositoryImpl repositoryImpl
		db             *gorm.DB
	}
	deckID, _ := uuid.GenerateUUID()
	type args struct {
		ctx   context.Context
		deck  *model.Deck
		cards []model.Card
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *model.DeckView
		err    bool
	}{
		{
			name: testInvalidDeckId,
			fields: fields{
				repositoryImpl: repositoryImpl{
					db: db,
				},
				db: db,
			},
			args: args{
				ctx: context.Background(),
				deck: &model.Deck{
					BaseModel: database.BaseModel{
						ID: "invalidDeckID",
					},
					Shuffled:  false,
					Remaining: 1,
				},
				cards: []model.Card{
					{
						BaseModel: database.BaseModel{ID: "AS"},
						Suit:      model.Spade,
						Rank:      model.Ace,
						Code:      "AS",
					},
				},
			},
			want: nil,
			err:  true,
		},
		{
			name: testGetCardView,
			fields: fields{
				repositoryImpl: repositoryImpl{
					db: db,
				},
				db: db,
			},
			args: args{
				ctx: context.Background(),
				deck: &model.Deck{
					BaseModel: database.BaseModel{
						ID: deckID,
					},
					Shuffled:  false,
					Remaining: 1,
				},
				cards: []model.Card{
					{
						BaseModel: database.BaseModel{ID: "AS"},
						Suit:      model.Spade,
						Rank:      model.Ace,
						Code:      "AS",
					},
				},
			},
			want: &model.DeckView{
				DeckID:    deckID,
				Shuffled:  false,
				Remaining: 1,
				Cards: []model.CardView{
					{
						Suit: model.Spade.String(),
						Rank: model.Ace.String(),
						Code: "AS",
					},
				},
			},
			err: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &cardRepositoryImpl{
				repositoryImpl: tt.fields.repositoryImpl,
				db:             tt.fields.db,
			}

			switch tt.name {
			case testGetCardView:
				err := c.CreateDeckCards(tt.args.ctx, tt.args.deck, tt.args.cards)
				if (err != nil) != tt.err {
					t.Errorf("CreateDeckCards() error = %v, wantErr %v", err, tt.err)
					return
				}
			}

			got, err := c.GetDeckView(tt.args.ctx, tt.args.deck.BaseModel.ID, constant.DefaultCardCount)
			if (err != nil) != tt.err {
				t.Errorf("GetDeckView() error = %v, wantErr %v", err, tt.err)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDeckView() got1 = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_cardRepositoryImpl_QueryCards(t *testing.T) {
	type fields struct {
		repositoryImpl repositoryImpl
		db             *gorm.DB
	}
	type args struct {
		ctx    context.Context
		params map[string]interface{}
		filter string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []model.Card
		err    bool
	}{
		{
			name: testQueryCardsInvalidFilter,
			fields: fields{
				repositoryImpl: repositoryImpl{
					db: db,
				},
				db: db,
			},
			args: args{
				ctx:    context.Background(),
				params: map[string]interface{}{"code": "INVALIDCODE"},
				filter: "",
			},
			want: nil,
			err:  true,
		},
		{
			name: testQueryCards,
			fields: fields{
				repositoryImpl: repositoryImpl{
					db: db,
				},
				db: db,
			},
			args: args{
				ctx:    context.Background(),
				params: map[string]interface{}{"code": "AS"},
				filter: "",
			},
			want: []model.Card{
				{
					BaseModel: database.BaseModel{ID: "AS"},
					Suit:      model.Spade,
					Rank:      model.Ace,
					Code:      "AS",
				},
			},
			err: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &cardRepositoryImpl{
				repositoryImpl: tt.fields.repositoryImpl,
				db:             tt.fields.db,
			}
			got, err := c.QueryCards(tt.args.ctx, tt.args.params, tt.args.filter)
			if err == nil {
				if !reflect.DeepEqual(got[0].Code, tt.want[0].Code) {
					t.Errorf("QueryCards() got = %v, want %v", got, tt.want)
				}
			}
			if (err != nil) != tt.err {
				t.Errorf("QueryCards() error = %v, wantErr %v", err, tt.err)
				return
			}
		})
	}
}

func Test_cardRepositoryImpl_QueryInCards(t *testing.T) {
	type fields struct {
		repositoryImpl repositoryImpl
		db             *gorm.DB
	}
	type args struct {
		ctx context.Context
		v   []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []model.Card
		wantErr bool
	}{
		{
			name: testQueryCardsInvalidFilter,
			fields: fields{
				repositoryImpl: repositoryImpl{
					db: db,
				},
				db: db,
			},
			args: args{
				ctx: context.Background(),
				v:   []string{"INVALID_CARD_ID"},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: testQueryCards,
			fields: fields{
				repositoryImpl: repositoryImpl{
					db: db,
				},
				db: db,
			},
			args: args{
				ctx: context.Background(),
				v:   []string{"ASS"},
			},
			want: []model.Card{
				{
					BaseModel: database.BaseModel{ID: "AS"},
					Suit:      model.Spade,
					Rank:      model.Ace,
					Code:      "AS",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &cardRepositoryImpl{
				repositoryImpl: tt.fields.repositoryImpl,
				db:             tt.fields.db,
			}
			got, err := c.QueryInCards(tt.args.ctx, tt.args.v)
			if err == nil {
				if !reflect.DeepEqual(got[0].Code, tt.want[0].Code) {
					t.Errorf("QueryInCards() got = %v, want %v", got, tt.want)
				}
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("QueryCards() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_postgresURI(t *testing.T) {
	type args struct {
		pgConfig config.PGConfig
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "PostgresURI",
			args: args{
				pgConfig: config.PGConfig{
					PostgresUser:     "deck_user",
					PostgresPassword: "Vm28xMykxKM",
					PostgresServer:   "localhost",
					PostgresPort:     "5431",
					PostgresDB:       "deck_db",
					TestPostgresDB:   "deck_test_db",
				},
			},
			want: "host=localhost user=deck_user password=Vm28xMykxKM dbname=deck_db port=5431 sslmode=disable",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := postgresURI(tt.args.pgConfig); got != tt.want {
				t.Errorf("postgresURI() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repositoryImpl_Create(t *testing.T) {
	id, _ := uuid.GenerateUUID()

	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx context.Context
		out interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   database.Error
	}{
		{
			name: "TestCreate",
			fields: fields{
				db: db,
			},
			args: args{
				ctx: context.Background(),
				out: &model.Deck{
					BaseModel: database.BaseModel{
						ID: id,
					},
					Shuffled:  false,
					Remaining: 1,
				},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			impl := &repositoryImpl{
				db: tt.fields.db,
			}
			if got := impl.Create(tt.args.ctx, tt.args.out); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repositoryImpl_Delete(t *testing.T) {

	id, _ := uuid.GenerateUUID()
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx context.Context
		out interface{}
		id  interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   database.Error
	}{
		{
			name: "TestSoftDelete",
			fields: fields{
				db: db,
			},
			args: args{
				ctx: context.Background(),
				out: &model.Deck{
					BaseModel: database.BaseModel{
						ID: id,
					},
					Shuffled:  false,
					Remaining: 1,
				},
				id: id,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			impl := &repositoryImpl{
				db: tt.fields.db,
			}
			if err := impl.Create(tt.args.ctx, tt.args.out); !reflect.DeepEqual(err, tt.want) {
				t.Errorf("Create() = %v, want %v", err, tt.want)
			}
			if got := impl.Delete(tt.args.ctx, tt.args.out, tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repositoryImpl_Get(t *testing.T) {
	card := &model.Card{}
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx context.Context
		out interface{}
		id  interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   database.Error
	}{
		{
			name: "GetCard",
			fields: fields{
				db: db,
			},
			args: args{
				ctx: context.Background(),
				out: card,
				id:  "AS",
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			impl := &repositoryImpl{
				db: tt.fields.db,
			}
			if got := impl.Get(tt.args.ctx, tt.args.out, tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repositoryImpl_Update(t *testing.T) {
	id, _ := uuid.GenerateUUID()
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx context.Context
		out interface{}
		id  interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   database.Error
	}{
		{
			name: "TestUpdate",
			fields: fields{
				db: db,
			},
			args: args{
				ctx: context.Background(),
				out: &model.Deck{
					BaseModel: database.BaseModel{
						ID: id,
					},
					Shuffled:  false,
					Remaining: 1,
				},
				id: id,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			impl := &repositoryImpl{
				db: tt.fields.db,
			}
			if err := impl.Create(tt.args.ctx, tt.args.out); !reflect.DeepEqual(err, tt.want) {
				t.Errorf("Create() = %v, want %v", err, tt.want)
			}
			if got := impl.Update(tt.args.ctx, tt.args.out, tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Update() = %v, want %v", got, tt.want)
			}
		})
	}
}
