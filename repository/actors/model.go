package actors

import (
	"airbound/internal/ent"
	"airbound/internal/ent/transaction"
	"airbound/repository/accesscontrol"
	"airbound/repository/identitydetails"
	"context"
	"time"

	"github.com/google/uuid"
)

type UserRepository interface {
	transaction.TxRunner
	CreateUser(ctx context.Context, u *User) (*ent.User, error)
	CreateUserWithTx(ctx context.Context, txClient *ent.Client, u *User) (*ent.User, error)
	FindUserByEmail(ctx context.Context, email string, edges ...string) (*ent.User, error)
	FetchUserIncludingDetails(ctx context.Context, userID uuid.UUID) (*ent.User, error)
	LinkAccountWithTx(ctx context.Context, txClient *ent.Client, userID, accountID uuid.UUID) (*ent.User, error)
	LinkAddressWithTx(ctx context.Context, txClient *ent.Client, userID, addressID uuid.UUID) (*ent.User, error)
	SetUserRoleWithTx(ctx context.Context, txClient *ent.Client, userID, roleID uuid.UUID) (*ent.User, error)
}

type User struct {
	ID        uuid.UUID                `json:"id"`
	Firstname string                   `json:"firstname"`
	Lastname  string                   `json:"lastname"`
	Email     string                   `json:"email"`
	Phone     string                   `json:"phone"`
	Token     string                   `json:"token"`
	Account   *identitydetails.Account `json:"account"`
	Address   *identitydetails.Address `json:"address"`
	Role      *accesscontrol.Role      `json:"role"`
	CreatedAt time.Time                `json:"created_at"`
	UpdatedAt time.Time                `json:"updated_at"`
}

func ParseToUser(model *ent.User, token string) *User {
	var a *identitydetails.Account
	var ad *identitydetails.Address
	var r *accesscontrol.Role

	if model.Edges.Account != nil {
		a = identitydetails.ParseToAccount(model.Edges.Account)
	}

	if model.Edges.Address != nil {
		ad = identitydetails.ParseToAddress(model.Edges.Address)
	}

	if model.Edges.Role != nil {
		r = accesscontrol.ParseToRole(model.Edges.Role)
	}

	u := &User{
		ID:        model.ID,
		Firstname: model.Firstname,
		Lastname:  model.Lastname,
		Email:     model.Email,
		Phone:     model.Phone,
		Token:     token,
		Account:   a,
		Address:   ad,
		Role:      r,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}

	return sanitizeUser(u)
}

func sanitizeUser(u *User) *User {
	u.Account.Password = nil
	u.Account.Salt = nil

	return u
}

type AdminRepository interface {
	transaction.TxRunner
	CreateAdmin(ctx context.Context, a *Admin, userID uuid.UUID) (*ent.Admin, error)
}

type Admin struct {
	ID             uuid.UUID `json:"id"`
	TwoFaSecret    string    `json:"two_fa_secret"`
	TwoFaCompleted bool      `json:"two_fa_completed"`
	Token          string    `json:"token"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type CustomerRepository interface {
	transaction.TxRunner
	CreateCustomer(ctx context.Context, c *Customer, userID uuid.UUID) (*ent.Customer, error)
}

type Customer struct {
	ID                  uuid.UUID `json:"id"`
	FrequentFlyerNumber string    `json:"frequent_flyer_number"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
}

type PilotRepository interface {
	transaction.TxRunner
	CreatePilot(ctx context.Context, p *Pilot, userID uuid.UUID) (*ent.Pilot, error)
}

type Pilot struct {
	ID            uuid.UUID `json:"id"`
	EmployeeID    string    `json:"employee_id"`
	LicenceNumber string    `json:"licence_number"`
	FlightHours   int       `json:"flight_hours"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type CrewRepository interface {
	transaction.TxRunner
	CreateCrew(ctx context.Context, c *Crew, userID uuid.UUID) (*ent.Crew, error)
}

type Crew struct {
	ID         uuid.UUID `json:"id"`
	EmployeeID string    `json:"employee_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type FrontDeskRepository interface {
	transaction.TxRunner
	CreateFrontDesk(ctx context.Context, f *FrontDesk, userID uuid.UUID) (*ent.FrontDesk, error)
}

type FrontDesk struct {
	ID         uuid.UUID `json:"id"`
	EmployeeID string    `json:"employee_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
