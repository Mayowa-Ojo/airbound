package flights

type FlightRepository interface {
	FindListFlightInstances()
}

type Flight struct{}
