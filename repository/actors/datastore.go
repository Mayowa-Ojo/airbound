package actors

import (
	"airbound/internal/ent"
	"airbound/internal/ent/account"
	"airbound/internal/ent/address"
	"airbound/internal/ent/admin"
	"airbound/internal/ent/permission"
	"airbound/internal/ent/role"
	"airbound/internal/ent/user"
	"context"

	"github.com/google/uuid"
)

type Datastore struct {
	c *ent.Client
}

func NewUserRepository(client *ent.Client) UserRepository {
	return &Datastore{c: client}
}

func (d *Datastore) GetClient() *ent.Client {
	return d.c
}

func (d *Datastore) queryEdges(ctx context.Context, entity *ent.User, edges ...string) error {
	for idx, edge := range edges {
		switch edge {
		case address.Table:
			if entity.Edges.Address != nil {
				break
			}
			address, err := entity.QueryAddress().Only(ctx)
			if err != nil {
				return err
			}

			entity.Edges.Address = address

		case account.Table:
			if entity.Edges.Account != nil {
				break
			}
			account, err := entity.QueryAccount().Only(ctx)
			if err != nil {
				return err
			}

			entity.Edges.Account = account

		case admin.Table:
			if entity.Edges.Admin != nil {
				break
			}
			admin, err := entity.QueryAdmin().Only(ctx)
			if err != nil {
				return err
			}

			entity.Edges.Admin = admin

		case role.Table:
			if entity.Edges.Role != nil {
				break
			}
			role, err := entity.QueryRole().Only(ctx)
			if err != nil {
				return err
			}

			entity.Edges.Role = role

		case permission.Table:
			if entity.Edges.Role.Edges.Permissions != nil {
				break
			}
			if entity.Edges.Role == nil {
				return d.queryEdges(ctx, entity, append([]string{role.Table}, edges[idx:]...)...)
			}
			permissions, err := entity.Edges.Role.QueryPermissions().All(ctx)
			if err != nil {
				return err
			}

			entity.Edges.Role.Edges.Permissions = permissions
		}
	}

	return nil
}

func (d *Datastore) CreateUser(ctx context.Context, u *User) (*ent.User, error) {
	entity, err := d.c.User.
		Create().
		SetFirstname(u.Firstname).
		SetLastname(u.Lastname).
		SetEmail(u.Email).
		SetPhone(u.Phone).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *Datastore) CreateUserWithTx(ctx context.Context, txClient *ent.Client, u *User) (*ent.User, error) {
	entity, err := txClient.User.
		Create().
		SetFirstname(u.Firstname).
		SetLastname(u.Lastname).
		SetEmail(u.Email).
		SetPhone(u.Phone).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *Datastore) FindUserByEmail(ctx context.Context, email string, edges ...string) (*ent.User, error) {
	entity, err := d.c.User.
		Query().
		Where(user.EmailEQ(email)).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	if err := d.queryEdges(ctx, entity, edges...); err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *Datastore) FetchUserIncludingDetails(ctx context.Context, userID uuid.UUID) (*ent.User, error) {
	entity, err := d.c.User.
		Query().
		Where(user.IDEQ(userID)).
		WithAccount().
		WithAddress().
		Only(ctx)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *Datastore) LinkAccountWithTx(ctx context.Context, txClient *ent.Client, userID, accountID uuid.UUID) (*ent.User, error) {
	entity, err := txClient.User.
		UpdateOneID(userID).
		SetAccountID(accountID).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *Datastore) LinkAddressWithTx(ctx context.Context, txClient *ent.Client, userID, addressID uuid.UUID) (*ent.User, error) {
	entity, err := txClient.User.
		UpdateOneID(userID).
		SetAddressID(addressID).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *Datastore) SetUserRoleWithTx(ctx context.Context, txClient *ent.Client, userID, roleID uuid.UUID) (*ent.User, error) {
	entity, err := txClient.User.
		UpdateOneID(userID).
		SetRoleID(roleID).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func NewAdminRepository(client *ent.Client) AdminRepository {
	return &Datastore{c: client}
}

func (d *Datastore) CreateAdmin(ctx context.Context, a *Admin, userID uuid.UUID) (*ent.Admin, error) {
	entity, err := d.c.Admin.
		Create().
		SetUserID(userID).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func NewCustomerRepository(client *ent.Client) CustomerRepository {
	return &Datastore{c: client}
}

func (d *Datastore) CreateCustomer(ctx context.Context, c *Customer, userID uuid.UUID) (*ent.Customer, error) {
	entity, err := d.c.Customer.
		Create().
		SetFrequentFlyerNumber(c.FrequentFlyerNumber).
		SetUserID(userID).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func NewPilotRepository(client *ent.Client) PilotRepository {
	return &Datastore{c: client}
}

func (d *Datastore) CreatePilot(ctx context.Context, p *Pilot, userID uuid.UUID) (*ent.Pilot, error) {
	entity, err := d.c.Pilot.
		Create().
		SetEmployeeID(p.EmployeeID).
		SetLicenceNumber(p.LicenceNumber).
		SetUserID(userID).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func NewCrewRepository(client *ent.Client) CrewRepository {
	return &Datastore{c: client}
}

func (d *Datastore) CreateCrew(ctx context.Context, c *Crew, userID uuid.UUID) (*ent.Crew, error) {
	entity, err := d.c.Crew.
		Create().
		SetEmployeeID(c.EmployeeID).
		SetUserID(userID).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func NewFrontDeskRepository(client *ent.Client) FrontDeskRepository {
	return &Datastore{c: client}
}

func (d *Datastore) CreateFrontDesk(ctx context.Context, f *FrontDesk, userID uuid.UUID) (*ent.FrontDesk, error) {
	entity, err := d.c.FrontDesk.
		Create().
		SetEmployeeID(f.EmployeeID).
		SetUserID(userID).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return entity, nil
}
