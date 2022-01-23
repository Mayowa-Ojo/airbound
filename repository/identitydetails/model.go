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
	UpdateAccountStatus(ctx context.Context, accountID uuid.UUID, status enums.AccountStatus) (*ent.Account, error)
	UpdateAccountVerificationToken(ctx context.Context, accountID uuid.UUID, token string) (*ent.Account, error)
	UpdateAccount2FaSecret(ctx context.Context, accountID uuid.UUID, secret string) (*ent.Account, error)
	UpdateAccount2FaStatus(ctx context.Context, accountID uuid.UUID, status bool) (*ent.Account, error)
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
	ID                  uuid.UUID           `json:"id"`
	Password            []byte              `json:"password"`
	Salt                []byte              `json:"salt"`
	AccountStatus       enums.AccountStatus `json:"account_status"`
	VerificationToken   string              `json:"verification_token"`
	ForgotPasswordToken string              `json:"forgot_password_token"`
	TwoFaSecret         string              `json:"two_fa_secret"`
	TwoFaCompleted      bool                `json:"two_fa_completed"`
	CreatedAt           time.Time           `json:"created_at"`
	UpdatedAt           time.Time           `json:"updated_at"`
}

func ParseToAccount(model *ent.Account) *Account {
	return &Account{
		ID:                model.ID,
		Password:          model.Password,
		Salt:              model.Salt,
		AccountStatus:     model.AccountStatus,
		VerificationToken: model.VerificationToken,
		TwoFaSecret:       model.TwoFaSecret,
		TwoFaCompleted:    model.TwoFaCompleted,
		CreatedAt:         model.CreatedAt,
		UpdatedAt:         model.UpdatedAt,
	}
}
