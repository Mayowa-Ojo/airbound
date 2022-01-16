package accesscontrol

import (
	"airbound/internal/ent"
	"airbound/internal/ent/enums"
	"airbound/internal/ent/permission"
	"airbound/internal/ent/role"
	"context"

	"github.com/google/uuid"
)

type Datastore struct {
	c *ent.Client
}

func NewRoleRepository(client *ent.Client) RoleRepository {
	return &Datastore{c: client}
}

func (d *Datastore) CreateRole(ctx context.Context, r *Role) (*ent.Role, error) {
	entity, err := d.c.Role.
		Create().
		SetName(r.Name).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *Datastore) FindRoleByName(ctx context.Context, name enums.Role) (*ent.Role, error) {
	entity, err := d.c.Role.Query().Where(role.NameEQ(name)).WithPermissions().Only(ctx)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *Datastore) FindRoleByID(ctx context.Context, roleID uuid.UUID) (*ent.Role, error) {
	entity, err := d.c.Role.Query().Where(role.IDEQ(roleID)).WithPermissions().Only(ctx)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *Datastore) UpdateRoleAddPermissions(ctx context.Context, roleID uuid.UUID, permissionIDs []uuid.UUID) (*ent.Role, error) {
	entity, err := d.c.Role.
		UpdateOneID(roleID).
		AddPermissionIDs(permissionIDs...).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *Datastore) UpdateRoleRemovePermissions(ctx context.Context, roleID uuid.UUID, permissionIDs []uuid.UUID) (*ent.Role, error) {
	entity, err := d.c.Role.
		UpdateOneID(roleID).
		RemovePermissionIDs(permissionIDs...).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func NewPermissionRepository(client *ent.Client) PermissionRepository {
	return &Datastore{c: client}
}

func (d *Datastore) CreatePermission(ctx context.Context, p *Permission, roleIDs []uuid.UUID) (*ent.Permission, error) {
	entity, err := d.c.Permission.
		Create().
		SetPermission(p.Permission).
		AddRoleIDs(roleIDs...).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *Datastore) FindPermissionByID(ctx context.Context, permissionID uuid.UUID) (*ent.Permission, error) {
	entity, err := d.c.Permission.
		Query().
		Where(permission.IDEQ(permissionID)).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *Datastore) FindListPermissionsByID(ctx context.Context, permissionIDs []uuid.UUID) ([]*ent.Permission, error) {
	entities, err := d.c.Permission.
		Query().
		Where(permission.IDIn(permissionIDs...)).
		All(ctx)
	if err != nil {
		return nil, err
	}

	return entities, nil
}
