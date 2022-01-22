package accesscontrol

type CreateRoleVM struct {
	Name string `json:"name" binding:"required"`
}

type CreatePermissionVM struct {
	Permission string   `json:"permission" binding:"required"`
	RoleIDs    []string `json:"roleIds" binding:"required"`
}

type GrantPermissionVM struct {
	RoleID        string   `json:"roleId"`
	PermissionIDs []string `json:"permissionIds" binding:"required"`
}
