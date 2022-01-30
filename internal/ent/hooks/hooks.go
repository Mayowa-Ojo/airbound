package hooks

import (
	"airbound/internal/ent"
	"airbound/internal/ent/hook"
	"airbound/internal/log"
	"context"
	"strings"
)

func EnsureUppercaseField() ent.Hook {
	return hook.On(func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			logger := log.WithField(string(log.LogFieldFunctionName), "<Ent-Hook>EnsureUppercaseField")

			a, ok := m.(*ent.AirlineMutation)
			if !ok {
				logger.Error("unable to get concrete type of mutation")
				return next.Mutate(ctx, m)
			}

			callSign, exists := a.CallSign()
			if !exists {
				logger.Warn("call_sign fied not set")
				return next.Mutate(ctx, m)
			}

			a.SetCallSign(strings.ToUpper(callSign))
			return next.Mutate(ctx, m)
		})
	}, ent.OpCreate)
}
