package facade

import (
	"airbound/handlers/flights"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FlightFacade struct {
	fh *flights.FlightHandler
}

func NewFlightFacade(fh *flights.FlightHandler) *FlightFacade {
	return &FlightFacade{fh}
}

func (f *FlightFacade) GetFlights(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
}
