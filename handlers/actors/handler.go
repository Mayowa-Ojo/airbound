package actors

import (
	"airbound/internal"
	"airbound/internal/auth"
	"airbound/internal/config"
	"airbound/internal/email"
	"airbound/internal/ent"
	"airbound/internal/ent/account"
	"airbound/internal/ent/address"
	"airbound/internal/ent/enums"
	"airbound/internal/ent/permission"
	"airbound/internal/ent/role"
	"airbound/internal/ent/transaction"
	"airbound/internal/errors"
	"airbound/internal/log"
	"airbound/internal/twofa"
	"airbound/repository/accesscontrol"
	"airbound/repository/actors"
	"airbound/repository/identitydetails"
	"context"
	"strings"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type ActorHandler struct {
	ur    actors.UserRepository
	ar    actors.AdminRepository
	cr    actors.CustomerRepository
	pr    actors.PilotRepository
	crr   actors.CrewRepository
	fr    actors.FrontDeskRepository
	acr   identitydetails.AccountRepository
	adr   identitydetails.AddressRepository
	rr    accesscontrol.RoleRepository
	cfg   *config.Config
	email email.EmailService
}

func NewActorHandler(
	ur actors.UserRepository,
	ar actors.AdminRepository,
	cr actors.CustomerRepository,
	pr actors.PilotRepository,
	crr actors.CrewRepository,
	fr actors.FrontDeskRepository,
	acr identitydetails.AccountRepository,
	adr identitydetails.AddressRepository,
	rr accesscontrol.RoleRepository,
	cfg *config.Config,
	e email.EmailService,
) *ActorHandler {
	return &ActorHandler{ur, ar, cr, pr, crr, fr, acr, adr, rr, cfg, e}
}

func (h *ActorHandler) createUserAndLinkAccount(ctx context.Context, vm UserRegistrationVM, roleName enums.Role) (*ent.User, error) {
	logger := log.WithField("FunctionName", "createUserAndLinkAccount")

	if roleName != enums.RoleCustomer {
		if isValid := internal.CheckInternalEmail(vm.Email); !isValid {
			return nil, errors.ErrInvalidInternalEmail
		}
	}

	user, _ := h.ur.FindUserByEmail(ctx, vm.Email)
	if user != nil {
		return nil, errors.ErrUserAlreadyExists
	}

	if err := auth.ValidatePassword(vm.Password); err != nil {
		return nil, err
	}

	hash, salt, err := auth.HashPassword(vm.Password)
	if err != nil {
		return nil, err
	}

	sanitizedPhone := internal.SanitizePhone(vm.Phone)

	u := &actors.User{
		Firstname: vm.Firstname,
		Lastname:  vm.Lastname,
		Email:     vm.Email,
		Phone:     sanitizedPhone,
	}

	ac := &identitydetails.Account{
		Password:      hash,
		Salt:          salt,
		AccountStatus: enums.AccountStatusNone,
	}

	ad := &identitydetails.Address{
		Street:  vm.Street,
		City:    vm.City,
		State:   vm.State,
		Zipcode: vm.Zipcode,
	}

	userRole, err := h.rr.FindRoleByName(ctx, roleName)
	if err != nil {
		return nil, err
	}

	entClient := h.ur.GetClient()

	if err := transaction.WithTransaction(ctx, entClient, func(tx *ent.Tx) error {
		client := tx.Client()

		user, err = h.ur.CreateUserWithTx(ctx, client, u)
		if err != nil {
			return err
		}

		account, err := h.acr.CreateAccountWithTx(ctx, client, ac)
		if err != nil {
			return err
		}

		address, err := h.adr.CreateAddressWithTx(ctx, client, ad)
		if err != nil {
			return err
		}

		_, err = h.ur.SetUserRoleWithTx(ctx, client, user.ID, userRole.ID)
		if err != nil {
			return err
		}

		user, err = h.ur.LinkAccountWithTx(ctx, client, user.ID, account.ID)
		if err != nil {
			logger.Error("error occured while linking account: %s", err)
			return err
		}

		user, err = h.ur.LinkAddressWithTx(ctx, client, user.ID, address.ID)
		if err != nil {
			logger.Error("error occured while linking address: %s", err)
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}

	user, err = h.ur.FindUserByEmail(ctx, user.Email, []string{account.Table, address.Table, role.Table, permission.Table}...)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (h *ActorHandler) CreateUserAdmin(ctx context.Context, vm UserRegistrationVM) (*actors.User, error) {
	user, err := h.createUserAndLinkAccount(ctx, vm, enums.RoleAdmin)
	if err != nil {
		return nil, err
	}

	a := &actors.Admin{}

	_, err = h.ar.CreateAdmin(ctx, a, user.ID)
	if err != nil {
		return nil, err
	}

	token, err := auth.GenerateToken(auth.TokenData{User: actors.ParseToUser(user, ""), Role: enums.RoleAdmin}, h.cfg)
	if err != nil {
		return nil, err
	}

	user, err = h.ur.FetchUserIncludingDetails(ctx, user.ID)
	if err != nil {
		return nil, err
	}

	return actors.ParseToUser(user, token), nil
}

func (h *ActorHandler) CreateUserCustomer(ctx context.Context, vm UserRegistrationVM) (*actors.User, error) {
	logger := log.WithField("FunctionName", "CreateUserCustomer")

	user, err := h.createUserAndLinkAccount(ctx, vm, enums.RoleCustomer)
	if err != nil {
		return nil, err
	}

	c := &actors.Customer{}

	_, err = h.cr.CreateCustomer(ctx, c, user.ID)
	if err != nil {
		logger.Error("error creating customer entity - %s", err)
		return nil, err
	}

	token, err := auth.GenerateToken(auth.TokenData{User: actors.ParseToUser(user, ""), Role: enums.RoleCustomer}, h.cfg)
	if err != nil {
		return nil, err
	}

	user, err = h.ur.FetchUserIncludingDetails(ctx, user.ID)
	if err != nil {
		return nil, err
	}

	verificationLink, err := auth.GenerateAccountVerificationLink(user.ID, h.cfg)
	if err != nil {
		return nil, err
	}

	go h.email.SendAccountVerificationEmail("mayowaojo.e@gmail.com", verificationLink)

	return actors.ParseToUser(user, token), nil
}

func (h *ActorHandler) CreateUserPilot(ctx context.Context, vm UserRegistrationVM) (*actors.User, error) {
	user, err := h.createUserAndLinkAccount(ctx, vm, enums.RolePilot)
	if err != nil {
		return nil, err
	}

	p := &actors.Pilot{}

	_, err = h.pr.CreatePilot(ctx, p, user.ID)
	if err != nil {
		return nil, err
	}

	token, err := auth.GenerateToken(auth.TokenData{User: actors.ParseToUser(user, ""), Role: enums.RolePilot}, h.cfg)
	if err != nil {
		return nil, err
	}

	user, err = h.ur.FetchUserIncludingDetails(ctx, user.ID)
	if err != nil {
		return nil, err
	}

	return actors.ParseToUser(user, token), nil
}

func (h *ActorHandler) CreateUserCrew(ctx context.Context, vm UserRegistrationVM) (*actors.User, error) {
	user, err := h.createUserAndLinkAccount(ctx, vm, enums.RoleCrew)
	if err != nil {
		return nil, err
	}

	c := &actors.Crew{}

	_, err = h.crr.CreateCrew(ctx, c, user.ID)
	if err != nil {
		return nil, err
	}

	token, err := auth.GenerateToken(auth.TokenData{User: actors.ParseToUser(user, ""), Role: enums.RoleCrew}, h.cfg)
	if err != nil {
		return nil, err
	}

	user, err = h.ur.FetchUserIncludingDetails(ctx, user.ID)
	if err != nil {
		return nil, err
	}

	return actors.ParseToUser(user, token), nil
}

func (h *ActorHandler) CreateUserFrontDesk(ctx context.Context, vm UserRegistrationVM) (*actors.User, error) {
	user, err := h.createUserAndLinkAccount(ctx, vm, enums.RoleFrontDesk)
	if err != nil {
		return nil, err
	}

	f := &actors.FrontDesk{}

	_, err = h.fr.CreateFrontDesk(ctx, f, user.ID)
	if err != nil {
		return nil, err
	}

	token, err := auth.GenerateToken(auth.TokenData{User: actors.ParseToUser(user, ""), Role: enums.RoleFrontDesk}, h.cfg)
	if err != nil {
		return nil, err
	}

	user, err = h.ur.FetchUserIncludingDetails(ctx, user.ID)
	if err != nil {
		return nil, err
	}

	return actors.ParseToUser(user, token), nil
}

func (h *ActorHandler) LoginUser(ctx context.Context, vm UserLoginVM) (*actors.User, error) {
	logger := log.WithField("FunctionName", "LoginUser")

	user, err := h.ur.FindUserByEmail(ctx, vm.Email, []string{account.Table, address.Table, role.Table, permission.Table}...)
	if err != nil {
		if strings.Contains(err.Error(), "user not found") {
			return nil, errors.ErrUserDoesNotExist
		}
		// TODO: implement better error wrappers to include more specific status codes
		logger.Error("error occured while fetching user: %s", err)
		return nil, err
	}

	if user.Edges.Account.AccountStatus == enums.AccountStatusBlacklisted {
		return nil, errors.ErrAccountBlacklisted
	}
	if user.Edges.Account.AccountStatus == enums.AccountStatusBlocked {
		return nil, errors.ErrAccountBlocked
	}
	if user.Edges.Account.AccountStatus == enums.AccountStatusClosed {
		return nil, errors.ErrAccountClosed
	}

	hash := user.Edges.Account.Password
	salt := user.Edges.Account.Salt

	if err := auth.VerifyPassword(vm.Password, hash, salt); err != nil {
		return nil, err
	}

	// check two-fa
	if (user.Edges.Role.Name != enums.RoleCustomer) || (user.Edges.Role.Name == enums.RoleCustomer && user.Edges.Account.TwoFaCompleted) {
		if isValid := twofa.Validate2FaCode(vm.TwoFaCode, user.Edges.Account.TwoFaSecret); !isValid {
			logger.Info("invalid two-fa code")
			return nil, errors.ErrInvalidTwoFaCode
		}
	}

	token, err := auth.GenerateToken(auth.TokenData{User: actors.ParseToUser(user, ""), Role: user.Edges.Role.Name}, h.cfg)
	if err != nil {
		return nil, err
	}

	return actors.ParseToUser(user, token), nil
}

func (h *ActorHandler) VerifyAccount(ctx context.Context, vm AccountVerificationVM) error {
	logger := log.WithField("FunctionName", "VerifyAccount")

	userID, err := auth.ValidateAccountVerificationToken(vm.Token, h.cfg)
	if err != nil {
		logger.Error("error validating account verification token - %s", err)
		return err
	}

	user, err := h.ur.FetchUserIncludingDetails(ctx, userID)
	if err != nil {
		return err
	}

	if user.Edges.Account.VerificationToken != "" {
		return errors.ErrAccountAlreadyVerified
	}

	if _, err := h.acr.UpdateAccountVerificationToken(ctx, user.Edges.Account.ID, vm.Token); err != nil {
		return err
	}

	if _, err := h.acr.UpdateAccountStatus(ctx, user.Edges.Account.ID, enums.AccountStatusActive); err != nil {
		return err
	}

	return nil
}

func (h *ActorHandler) BeginTwoFaSetup(ctx context.Context, vm TwoFaSetupVM) (*string, error) {
	logger := log.WithField("FunctionName", "BeginTwoFaSetup")
	userID := uuid.MustParse(vm.UserID)

	user, err := h.ur.FetchUserIncludingDetails(ctx, userID)
	if err != nil {
		logger.Error("error occurred while fetching user - %s", err)
		return nil, err
	}

	if user.Edges.Account.TwoFaCompleted {
		return nil, errors.ErrTwoFaSetupCompleted
	}

	kp, err := twofa.GenerateKeyPair(user.Email, h.cfg)
	if err != nil {
		logger.Error("error occurred while generating 2fa key pair - %s", err)
		return nil, err
	}

	if _, err := h.acr.UpdateAccount2FaSecret(ctx, user.Edges.Account.ID, kp.Secret); err != nil {
		logger.Error("error occured while updating account - %s", err)
		return nil, err
	}

	return &kp.QRCode, nil
}

func (h *ActorHandler) CompleteTwoFaSetup(ctx context.Context, vm TwoFaSetupVM) error {
	logger := log.WithField("FunctionName", "CompleteTwoFaSetup")
	userID := uuid.MustParse(vm.UserID)

	user, err := h.ur.FetchUserIncludingDetails(ctx, userID)
	if err != nil {
		logger.Error("error occurred while fetching user - %s", err)
	}

	if user.Edges.Account.TwoFaCompleted {
		return errors.ErrTwoFaSetupCompleted
	}

	if isValid := twofa.Validate2FaCode(vm.TwoFaCode, user.Edges.Account.TwoFaSecret); !isValid {
		logger.
			WithFields(logrus.Fields{"Code": vm.TwoFaCode, "Secret": user.Edges.Account.TwoFaSecret, "isValid": isValid}).
			Info("two-fa verification failed - invalid code")
		return errors.ErrInvalidTwoFaCode
	}

	if _, err := h.acr.UpdateAccount2FaStatus(ctx, user.Edges.Account.ID, true); err != nil {
		logger.Error("error occurred while updating account - %s", err)
		return err
	}

	return nil
}
