package event

import (
	"airbound/internal/errors"
	"airbound/internal/log"
	"context"
	"strings"
)

type Emitter struct {
	events map[Event][]Listener
}

func NewEventEmitter() *Emitter {
	return &Emitter{events: make(map[Event][]Listener)}
}

func (r *Emitter) Register(evt Event, listener Listener) {
	logger := log.WithField(string(log.LogFieldFunctionName), "<Event>Register")

	if strings.EqualFold(string(evt), "") {
		logger.Error(errors.ErrInvalidEventType)
	}

	if r.events == nil {
		r.events = make(map[Event][]Listener)
	}

	if _, ok := r.events[evt]; !ok {
		r.events[evt] = []Listener{listener}
	} else {
		r.events[evt] = append(r.events[evt], listener)
	}
}

func (r *Emitter) Dispatch(ctx context.Context, evt Event, payload interface{}) {
	logger := log.WithField(string(log.LogFieldFunctionName), "<Event>Dispatch")

	if strings.EqualFold(string(evt), "") {
		logger.Error(errors.ErrInvalidEventType)
	}

	if _, ok := r.events[evt]; !ok {
		logger.Error(errors.ErrUnregisteredEvent)
	}

	for _, listener := range r.events[evt] {
		go func(listener Listener) {
			if err := listener.Handle(ctx, payload); err != nil {
				logger.Error("error occurred while dispatching event - %", err)
			}
		}(listener)
	}
}
