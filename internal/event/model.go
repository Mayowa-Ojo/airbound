package event

import "context"

type Event string

const (
	CONFIRM_FLIGHT_RESERVATION = Event("CONFIRM_FLIGHT_RESERVATION")
)

type Listener interface {
	Handle(ctx context.Context, payload interface{}) error
}
