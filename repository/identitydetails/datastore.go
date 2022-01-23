package identitydetails

import (
	"airbound/internal/ent"
	"airbound/internal/ent/enums"
	"context"

	"github.com/google/uuid"
)

type Datastore struct {
	c *ent.Client
}

func (d *Datastore) GetClient() *ent.Client {
	return d.c
}

func NewAddressRepository(client *ent.Client) AddressRepository {
	return &Datastore{c: client}
}

func (d *Datastore) CreateAddressWithTx(ctx context.Context, txClient *ent.Client, a *Address) (*ent.Address, error) {
	entity, err := d.c.Address.
		Create().
		SetStreet(a.Street).
		SetCity(a.City).
		SetState(a.State).
		SetZipcode(a.Zipcode).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func NewAccountRepository(client *ent.Client) AccountRepository {
	return &Datastore{c: client}
}

func (d *Datastore) CreateAccountWithTx(ctx context.Context, txClient *ent.Client, a *Account) (*ent.Account, error) {
	entity, err := txClient.Account.
		Create().
		SetPassword(a.Password).
		SetSalt(a.Salt).
		SetAccountStatus(a.AccountStatus).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *Datastore) UpdateAccountStatus(ctx context.Context, accountID uuid.UUID, status enums.AccountStatus) (*ent.Account, error) {
	entity, err := d.c.Account.
		UpdateOneID(accountID).
		SetAccountStatus(status).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *Datastore) UpdateAccountVerificationToken(ctx context.Context, accountID uuid.UUID, token string) (*ent.Account, error) {
	entity, err := d.c.Account.
		UpdateOneID(accountID).
		SetVerificationToken(token).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *Datastore) UpdateAccount2FaSecret(ctx context.Context, accountID uuid.UUID, secret string) (*ent.Account, error) {
	entity, err := d.c.Account.
		UpdateOneID(accountID).
		SetTwoFaSecret(secret).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *Datastore) UpdateAccount2FaStatus(ctx context.Context, accountID uuid.UUID, status bool) (*ent.Account, error) {
	entity, err := d.c.Account.
		UpdateOneID(accountID).
		SetTwoFaCompleted(status).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return entity, nil
}
