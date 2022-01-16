package identitydetails

import (
	"airbound/internal/ent"
	"airbound/internal/ent/enums"
	"airbound/internal/ent/transaction"
	"context"
	"time"

	"github.com/google/uuid"
)

type AddressRepository interface {
	transaction.TxRunner
	CreateAddressWithTx(ctx context.Context, txClient *ent.Client, a *Address) (*ent.Address, error)
}

type AccountRepository interface {
	transaction.TxRunner
	CreateAccountWithTx(ctx context.Context, txClient *ent.Client, a *Account) (*ent.Account, error)
}

type Address struct {
	ID        uuid.UUID `json:"id"`
	Street    string    `json:"street"`
	City      string    `json:"city"`
	State     string    `json:"state"`
	Zipcode   string    `json:"zipcode"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ParseToAddress(model *ent.Address) *Address {
	return &Address{
		ID:        model.ID,
		Street:    model.Street,
		City:      model.City,
		State:     model.State,
		Zipcode:   model.Zipcode,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}
}

type Account struct {
	ID            uuid.UUID           `json:"id"`
	Password      []byte              `json:"password"`
	Salt          []byte              `json:"salt"`
	AccountStatus enums.AccountStatus `json:"account_status"`
	CreatedAt     time.Time           `json:"created_at"`
	UpdatedAt     time.Time           `json:"updated_at"`
}

func ParseToAccount(model *ent.Account) *Account {
	return &Account{
		ID:            model.ID,
		Password:      model.Password,
		Salt:          model.Salt,
		AccountStatus: model.AccountStatus,
		CreatedAt:     model.CreatedAt,
		UpdatedAt:     model.UpdatedAt,
	}
}
