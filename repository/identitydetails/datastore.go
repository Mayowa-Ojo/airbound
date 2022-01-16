package identitydetails

import (
	"airbound/internal/ent"
	"context"
)

type Datastore struct {
	c *ent.Client
}

func NewAddressRepository(client *ent.Client) AddressRepository {
	return &Datastore{c: client}
}

func NewAccountRepository(client *ent.Client) AccountRepository {
	return &Datastore{c: client}
}

func (d *Datastore) GetClient() *ent.Client {
	return d.c
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
