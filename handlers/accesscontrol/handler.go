package accesscontrol

import (
	"airbound/internal/config"
	"airbound/internal/email"
	"airbound/internal/ent"
	"airbound/internal/ent/enums"
	"airbound/internal/errors"
	"airbound/internal/rbac"
	"airbound/repository/accesscontrol"
	"context"

	"github.com/google/uuid"
)

type AccessControlHandler struct {
	rr    accesscontrol.RoleRepository
	pr    accesscontrol.PermissionRepository
	cfg   *config.Config
	email email.EmailService
}

func NewAccessControlHandler(
	rr accesscontrol.RoleRepository, pr accesscontrol.PermissionRepository,
	cfg *config.Config, e email.EmailService,
) *AccessControlHandler {
	return &AccessControlHandler{rr, pr, cfg, e}
}

func (h *AccessControlHandler) GetRole(ctx context.Context, roleNameOrID string) (*accesscontrol.Role, error) {
	var role *ent.Role
	nameOrID, err := uuid.Parse(roleNameOrID)

	if err != nil {
		role, err = h.rr.FindRoleByName(ctx, enums.Role(roleNameOrID))
		if err != nil {
			return nil, err
		}
	} else {
		role, err = h.rr.FindRoleByID(ctx, nameOrID)
		if err != nil {
			return nil, err
		}
	}

	return accesscontrol.ParseToRole(role), nil
}

func (h *AccessControlHandler) CreateRole(ctx context.Context, vm CreateRoleVM) (*accesscontrol.Role, error) {
	r := &accesscontrol.Role{
		Name: enums.Role(vm.Name),
	}

	role, err := h.rr.CreateRole(ctx, r)
	if err != nil {
		return nil, err
	}

	return accesscontrol.ParseToRole(role), nil
}

func (h *AccessControlHandler) CreatePermission(ctx context.Context, vm CreatePermissionVM) (*accesscontrol.Permission, error) {
	p := &accesscontrol.Permission{
		Permission: vm.Permission,
	}

	var roleIDs []uuid.UUID

	for _, v := range vm.RoleIDs {
		id := uuid.MustParse(v)
		roleIDs = append(roleIDs, id)
	}

	isValid := rbac.ValidatePermission(vm.Permission)
	if !isValid {
		return nil, errors.ErrInvalidPermission
	}

	permission, err := h.pr.CreatePermission(ctx, p, roleIDs)
	if err != nil {
		return nil, err
	}

	return accesscontrol.ParseToPermission(permission), nil
}

func (h *AccessControlHandler) GrantPermission(ctx context.Context, vm GrantPermissionVM) (*accesscontrol.Role, error) {
	roleID := uuid.MustParse(vm.RoleID)
	var permissionIDs []uuid.UUID

	for _, v := range vm.PermissionIDs {
		id := uuid.MustParse(v)
		permissionIDs = append(permissionIDs, id)
	}

	if _, err := h.pr.FindListPermissionsByID(ctx, permissionIDs); err != nil {
		return nil, err
	}

	role, err := h.rr.UpdateRoleAddPermissions(ctx, roleID, permissionIDs)
	if err != nil {
		return nil, err
	}

	return accesscontrol.ParseToRole(role), nil
}
