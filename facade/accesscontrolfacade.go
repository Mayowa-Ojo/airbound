package facade

import (
	"airbound/handlers/accesscontrol"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AccessControlFacade struct {
	ach accesscontrol.AccessControlHandler
}

func NewAccessControlFacade(ach accesscontrol.AccessControlHandler) *AccessControlFacade {
	return &AccessControlFacade{ach}
}

func (f *AccessControlFacade) GetRole(ctx *gin.Context) {
	var roleNameOrID string
	roleID := ctx.Param("role-id")
	roleName := ctx.Query("name")

	if roleName == "" {
		roleNameOrID = roleID
	} else {
		roleNameOrID = roleName
	}

	role, err := f.ach.GetRole(ctx, roleNameOrID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": fmt.Sprint(err),
			"data":    nil,
		})

		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status":  "ok",
		"message": "role retrieved",
		"data":    role,
	})
}

func (f *AccessControlFacade) CreateRole(ctx *gin.Context) {
	var vm accesscontrol.CreateRoleVM

	if err := ctx.ShouldBindJSON(&vm); err != nil {
		ctx.JSON(http.StatusPreconditionFailed, gin.H{
			"status":  "error",
			"message": fmt.Sprint(err),
			"data":    nil,
		})

		return
	}

	role, err := f.ach.CreateRole(ctx, vm)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": fmt.Sprint(err),
			"data":    nil,
		})

		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status":  "ok",
		"message": "role created",
		"data":    role,
	})
}

func (f *AccessControlFacade) CreatePermission(ctx *gin.Context) {
	var vm accesscontrol.CreatePermissionVM

	if err := ctx.ShouldBindJSON(&vm); err != nil {
		ctx.JSON(http.StatusPreconditionFailed, gin.H{
			"status":  "error",
			"message": fmt.Sprint(err),
			"data":    nil,
		})

		return
	}

	role, err := f.ach.CreatePermission(ctx, vm)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": fmt.Sprint(err),
			"data":    nil,
		})

		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status":  "ok",
		"message": "permission created",
		"data":    role,
	})
}

func (f *AccessControlFacade) GrantPermission(ctx *gin.Context) {
	var vm accesscontrol.GrantPermissionVM

	if err := ctx.ShouldBindJSON(&vm); err != nil {
		ctx.JSON(http.StatusPreconditionFailed, gin.H{
			"status":  "error",
			"message": fmt.Sprint(err),
			"data":    nil,
		})

		return
	}

	vm.RoleID = ctx.Param("role-id")

	role, err := f.ach.GrantPermission(ctx, vm)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": fmt.Sprint(err),
			"data":    nil,
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "permission granted",
		"data":    role,
	})
}
