package api

import (
	"airbound/facade"
	"airbound/handlers/accesscontrol"
	"airbound/handlers/actors"
	"airbound/handlers/flights"
	"airbound/internal/config"
	"airbound/internal/ent"
	"airbound/internal/ent/user"
	"airbound/internal/log"
	"airbound/internal/middleware"
	"airbound/internal/rbac"
	acRepository "airbound/repository/accesscontrol"
	actorRepository "airbound/repository/actors"
	flightRepository "airbound/repository/flights"
	idRepository "airbound/repository/identitydetails"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Run(client *ent.Client, cfg *config.Config) *gin.Engine {
	router := gin.New()

	router.Use(log.Logger(logrus.New()), gin.Recovery())

	registerRoutes(router, client, cfg)

	// doRapidTest(client)

	return router
}

func registerRoutes(router *gin.Engine, client *ent.Client, cfg *config.Config) {
	flightRepo := flightRepository.NewFlightRepository(client)
	userRepo := actorRepository.NewUserRepository(client)
	adminRepo := actorRepository.NewAdminRepository(client)
	customerRepo := actorRepository.NewCustomerRepository(client)
	pilotRepo := actorRepository.NewPilotRepository(client)
	crewRepo := actorRepository.NewCrewRepository(client)
	frontDeskRepo := actorRepository.NewFrontDeskRepository(client)
	accountRepo := idRepository.NewAccountRepository(client)
	addressRepo := idRepository.NewAddressRepository(client)
	roleRepo := acRepository.NewRoleRepository(client)
	permissionRepo := acRepository.NewPermissionRepository(client)

	flightHandler := flights.NewFlightHandler(flightRepo, cfg)
	actorHandler := actors.NewActorHandler(
		userRepo, adminRepo, customerRepo, pilotRepo, crewRepo,
		frontDeskRepo, accountRepo, addressRepo, roleRepo, cfg,
	)
	accessControlHandler := accesscontrol.NewAccessControlHandler(roleRepo, permissionRepo, cfg)

	flightFacade := facade.NewFlightFacade(flightHandler)
	actorFacade := facade.NewActorFacade(*actorHandler)
	accessControlFacade := facade.NewAccessControlFacade(*accessControlHandler)

	// health-check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"message": "up and running...",
			"data":    nil,
		})
	})

	{
		v1 := router.Group("/api/v1")

		v1.POST("/users/register-admin", actorFacade.RegisterUserAsAdmin)
		v1.POST("/users/register-customer", actorFacade.RegisterUserAsCustomer)
		v1.POST("/users/register-pilot", actorFacade.RegisterUserAsPilot)
		v1.POST("/users/register-crew", actorFacade.RegisterUserAsCrew)
		v1.POST("/users/register-frontdesk", actorFacade.RegisterUserAsFrontDesk)
		v1.POST("/users/login", actorFacade.LoginUser)

		v1.Use(middleware.AuthenticateUser(cfg))
		v1.GET("/flights", middleware.AuthorizeUser(rbac.Flights.GetFlights), flightFacade.GetFlights)

		v1.GET("/roles/:role-id", middleware.AuthorizeUser(rbac.AccessControl.GetRole), accessControlFacade.GetRole)

		v1.POST("/roles", middleware.AuthorizeUser(rbac.AccessControl.CreateRole), accessControlFacade.CreateRole)
		v1.POST("/permissions", middleware.AuthorizeUser(rbac.AccessControl.CreatePermission), accessControlFacade.CreatePermission)

		v1.PATCH("/roles/:role-id/grant", middleware.AuthorizeUser(rbac.AccessControl.GrantPermission), accessControlFacade.GrantPermission)
	}
}

func doRapidTest(c *ent.Client) {
	ctx := context.Background()
	user, err := c.User.
		Query().
		Where(user.LastnameEQ("Brady")).
		Only(ctx)
	if err != nil {
		log.Fatal("[Rapid-Test]: %s", err)
	}

	account, err := user.QueryAccount().Only(ctx)
	if err != nil {
		log.Fatal("[Rapid-Test]: %s", err)
	}
	user.Edges.Account = account

	log.Info("[DEBUG]: User -- %+v", user.Edges.Account)
}
