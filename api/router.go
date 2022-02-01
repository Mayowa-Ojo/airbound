package api

import (
	"airbound/facade"
	"airbound/handlers/accesscontrol"
	"airbound/handlers/actors"
	"airbound/handlers/air"
	"airbound/handlers/flights"
	"airbound/internal/config"
	"airbound/internal/email"
	"airbound/internal/ent"
	"airbound/internal/event"
	"airbound/internal/log"
	"airbound/internal/middleware"
	"airbound/internal/queue"
	"airbound/internal/rbac"
	acRepository "airbound/repository/accesscontrol"
	actorRepository "airbound/repository/actors"
	airRepository "airbound/repository/air"
	flightRepository "airbound/repository/flights"
	idRepository "airbound/repository/identitydetails"
	"net/http"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Run(client *ent.Client, cfg *config.Config, sess, sqss *session.Session) *gin.Engine {
	router := gin.New()

	router.Use(log.Logger(logrus.New()), gin.Recovery())

	registerRoutes(router, client, cfg, sess, sqss)

	return router
}

func registerRoutes(router *gin.Engine, client *ent.Client, cfg *config.Config, sess, sqss *session.Session) {
	emitter := event.NewEventEmitter()
	emailService := email.NewEmailService(sess, cfg)
	queueService := queue.NewQueueService(sqss, cfg, emitter)

	// queueService.SendFlightReservationMessage(context.Background(), &flightRepository.FlightReservation{
	// 	ID:                uuid.MustParse("6cc46f7a-b0be-4dbf-8068-0dafa5769d7c"),
	// 	ReservationNumber: "1F6nvRT5mYvi1KxKe63fuRxAZ8zu3L8vDq",
	// 	ReservationStatus: enums.ReservationStatusRequested,
	// 	FlightInstance: &flightRepository.FlightInstance{
	// 		ID: uuid.MustParse("e353e222-1cce-4bca-aa8e-ec4a334fa906"),
	// 	},
	// }, uuid.MustParse("86659acf-770d-4cb3-80b3-cced65280495"))

	flightRepo := flightRepository.NewFlightRepository(client)
	flightInstanceRepo := flightRepository.NewFlightInstanceRepository(client)
	flightScheduleRepo := flightRepository.NewFlightScheduleRepository(client)
	flightReservationRepo := flightRepository.NewFlightReservationRepository(client)
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
	airlineRepo := airRepository.NewAirlineRepository(client)
	airportRepo := airRepository.NewAirportRepository(client)
	aircraftRepo := airRepository.NewAircraftRepository(client)

	flightHandler := flights.NewFlightHandler(
		flightRepo, flightInstanceRepo, flightScheduleRepo, flightReservationRepo, cfg, emailService, queueService,
	)
	actorHandler := actors.NewActorHandler(
		userRepo, adminRepo, customerRepo, pilotRepo, crewRepo,
		frontDeskRepo, accountRepo, addressRepo, roleRepo, cfg, emailService,
	)
	accessControlHandler := accesscontrol.NewAccessControlHandler(roleRepo, permissionRepo, cfg, emailService)
	airHandler := air.NewAirHandler(airlineRepo, airportRepo, aircraftRepo, addressRepo, cfg)

	flightFacade := facade.NewFlightFacade(flightHandler)
	actorFacade := facade.NewActorFacade(*actorHandler)
	accessControlFacade := facade.NewAccessControlFacade(*accessControlHandler)
	airFacade := facade.NewAirFacade(airHandler)

	emitter.Register(event.CONFIRM_FLIGHT_RESERVATION, &flights.ConfirmReservationEvent{Fh: flightHandler})

	router.LoadHTMLGlob("templates/*")

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

		v1.GET("/flights", flightFacade.GetFlights)
		v1.GET("/users/verify-account", actorFacade.VerifyAccount)

		v1.POST("/users/register-admin", actorFacade.RegisterUserAsAdmin)
		v1.POST("/users/register-customer", actorFacade.RegisterUserAsCustomer)
		v1.POST("/users/register-pilot", actorFacade.RegisterUserAsPilot)
		v1.POST("/users/register-crew", actorFacade.RegisterUserAsCrew)
		v1.POST("/users/register-frontdesk", actorFacade.RegisterUserAsFrontDesk)
		v1.POST("/users/login", actorFacade.LoginUser)
		v1.POST("/users/begin-2fa-setup", actorFacade.BeginTwoFaSetup)
		v1.POST("/users/complete-2fa-setup", actorFacade.CompleteTwoFaSetup)

		v1.Use(middleware.AuthenticateUser(cfg))
		v1.Use(middleware.EnsureTwoFa())
		v1.Use(middleware.CheckAccountStatus())

		v1.GET("/roles/:role-id", middleware.AuthorizeUser(rbac.AccessControl.GetRole), accessControlFacade.GetRole)

		v1.POST("/roles", middleware.AuthorizeUser(rbac.AccessControl.CreateRole), accessControlFacade.CreateRole)
		v1.POST("/permissions", middleware.AuthorizeUser(rbac.AccessControl.CreatePermission), accessControlFacade.CreatePermission)
		v1.POST("/airlines/register", middleware.AuthorizeUser(rbac.Air.RegisterAirline), airFacade.RegisterAirline)
		v1.POST("/airports/register", middleware.AuthorizeUser(rbac.Air.RegisterAirport), airFacade.RegisterAirport)
		v1.POST("/aircrafts/register", middleware.AuthorizeUser(rbac.Air.RegisterAircraft), airFacade.RegisterAircraft)

		v1.PATCH("/roles/:role-id/grant", middleware.AuthorizeUser(rbac.AccessControl.GrantPermission), accessControlFacade.GrantPermission)
	}
}
