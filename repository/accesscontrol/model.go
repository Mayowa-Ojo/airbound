package accesscontrol

import (
	"airbound/internal/ent"
	"airbound/internal/ent/enums"
	"context"
	"time"

	"github.com/google/uuid"
)

type RoleRepository interface {
	CreateRole(ctx context.Context, r *Role) (*ent.Role, error)
	FindRoleByName(ctx context.Context, name enums.Role) (*ent.Role, error)
	FindRoleByID(ctx context.Context, roleID uuid.UUID) (*ent.Role, error)
	UpdateRoleAddPermissions(ctx context.Context, roleID uuid.UUID, permissionIDs []uuid.UUID) (*ent.Role, error)
	UpdateRoleRemovePermissions(ctx context.Context, roleID uuid.UUID, permissionIDs []uuid.UUID) (*ent.Role, error)
}

type PermissionRepository interface {
	CreatePermission(ctx context.Context, p *Permission, roleIDs []uuid.UUID) (*ent.Permission, error)
	FindPermissionByID(ctx context.Context, permissionID uuid.UUID) (*ent.Permission, error)
	FindListPermissionsByID(ctx context.Context, permissionIDs []uuid.UUID) ([]*ent.Permission, error)
}

type Role struct {
	ID          uuid.UUID     `json:"id"`
	Name        enums.Role    `json:"name"`
	Permissions []*Permission `json:"permissions"`
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at"`
}

func ParseToRole(model *ent.Role) *Role {
	var ps []*Permission

	for _, p := range model.Edges.Permissions {
		ps = append(ps, ParseToPermission(p))
	}

	return &Role{
		ID:          model.ID,
		Name:        model.Name,
		Permissions: ps,
		CreatedAt:   model.CreatedAt,
		UpdatedAt:   model.UpdatedAt,
	}
}

type Permission struct {
	ID         uuid.UUID `json:"id"`
	Permission string    `json:"permission"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func ParseToPermission(model *ent.Permission) *Permission {
	return &Permission{
		ID:         model.ID,
		Permission: model.Permission,
		CreatedAt:  model.CreatedAt,
		UpdatedAt:  model.UpdatedAt,
	}
}
