package facade

import (
	"airbound/handlers/actors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ActorFacade struct {
	ah actors.ActorHandler
}

func NewActorFacade(ah actors.ActorHandler) *ActorFacade {
	return &ActorFacade{ah}
}

func (f *ActorFacade) RegisterUserAsAdmin(ctx *gin.Context) {
	var vm actors.UserRegistrationVM

	if err := ctx.ShouldBindJSON(&vm); err != nil {
		ctx.JSON(http.StatusPreconditionFailed, gin.H{
			"status":  "error",
			"message": fmt.Sprint(err),
			"data":    nil,
		})

		return
	}

	user, err := f.ah.CreateUserAdmin(ctx, vm)
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
		"message": "success",
		"data":    user,
	})
}

func (f *ActorFacade) RegisterUserAsCustomer(ctx *gin.Context) {
	var vm actors.UserRegistrationVM

	if err := ctx.ShouldBindJSON(&vm); err != nil {
		ctx.JSON(http.StatusPreconditionFailed, gin.H{
			"status":  "error",
			"message": fmt.Sprint(err),
			"data":    nil,
		})

		return
	}

	user, err := f.ah.CreateUserCustomer(ctx, vm)
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
		"message": "success",
		"data":    user,
	})
}

func (f *ActorFacade) RegisterUserAsPilot(ctx *gin.Context) {
	var vm actors.UserRegistrationVM

	if err := ctx.ShouldBindJSON(&vm); err != nil {
		ctx.JSON(http.StatusPreconditionFailed, gin.H{
			"status":  "error",
			"message": fmt.Sprint(err),
			"data":    nil,
		})

		return
	}

	user, err := f.ah.CreateUserPilot(ctx, vm)
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
		"message": "success",
		"data":    user,
	})
}

func (f *ActorFacade) RegisterUserAsCrew(ctx *gin.Context) {
	var vm actors.UserRegistrationVM

	if err := ctx.ShouldBindJSON(&vm); err != nil {
		ctx.JSON(http.StatusPreconditionFailed, gin.H{
			"status":  "error",
			"message": fmt.Sprint(err),
			"data":    nil,
		})

		return
	}

	user, err := f.ah.CreateUserCrew(ctx, vm)
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
		"message": "success",
		"data":    user,
	})
}

func (f *ActorFacade) RegisterUserAsFrontDesk(ctx *gin.Context) {
	var vm actors.UserRegistrationVM

	if err := ctx.ShouldBindJSON(&vm); err != nil {
		ctx.JSON(http.StatusPreconditionFailed, gin.H{
			"status":  "error",
			"message": fmt.Sprint(err),
			"data":    nil,
		})

		return
	}

	user, err := f.ah.CreateUserFrontDesk(ctx, vm)
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
		"message": "success",
		"data":    user,
	})
}

func (f *ActorFacade) LoginUser(ctx *gin.Context) {
	var vm actors.UserLoginVM

	if err := ctx.ShouldBindJSON(&vm); err != nil {
		ctx.JSON(http.StatusPreconditionFailed, gin.H{
			"status":  "error",
			"message": fmt.Sprint(err),
			"data":    nil,
		})

		return
	}

	user, err := f.ah.LoginUser(ctx, vm)
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
		"message": "success",
		"data":    user,
	})
}

func (f *ActorFacade) VerifyAccount(ctx *gin.Context) {
	var vm actors.AccountVerificationVM

	if err := ctx.ShouldBindQuery(&vm); err != nil {
		ctx.JSON(http.StatusPreconditionFailed, gin.H{
			"status":  "error",
			"message": fmt.Sprint(err),
			"data":    nil,
		})

		return
	}

	if err := f.ah.VerifyAccount(ctx, vm); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": fmt.Sprint(err),
			"data":    nil,
		})

		return
	}

	ctx.HTML(http.StatusOK, "account_verified.tmpl", gin.H{})
}

func (f *ActorFacade) BeginTwoFaSetup(ctx *gin.Context) {
	var vm actors.TwoFaSetupVM

	if err := ctx.ShouldBindJSON(&vm); err != nil {
		ctx.JSON(http.StatusPreconditionFailed, gin.H{
			"status":  "error",
			"message": fmt.Sprint(err),
			"data":    nil,
		})

		return
	}

	qrCode, err := f.ah.BeginTwoFaSetup(ctx, vm)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": fmt.Sprint(err),
			"data":    nil,
		})

		return
	}

	ctx.HTML(http.StatusOK, "show_2fa_qrcode.tmpl", gin.H{"qrCodeLink": qrCode})
}

func (f *ActorFacade) CompleteTwoFaSetup(ctx *gin.Context) {
	var vm actors.TwoFaSetupVM

	if err := ctx.ShouldBindJSON(&vm); err != nil {
		ctx.JSON(http.StatusPreconditionFailed, gin.H{
			"status":  "error",
			"message": fmt.Sprint(err),
			"data":    nil,
		})

		return
	}

	if vm.TwoFaCode == "" {
		ctx.JSON(http.StatusPreconditionFailed, gin.H{
			"status":  "error",
			"message": "missing/invalid field in request body",
			"data":    nil,
		})

		return
	}

	if err := f.ah.CompleteTwoFaSetup(ctx, vm); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": fmt.Sprint(err),
			"data":    nil,
		})

		return
	}

	ctx.HTML(http.StatusOK, "2fa_completed.tmpl", gin.H{})
}
