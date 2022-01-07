package transaction

import (
	"airbound/internal/ent"
	"context"
	"fmt"
)

type TxHandler func(*ent.Tx) error

type TxRunner interface {
	GetClient() *ent.Client
}

func WithTransaction(ctx context.Context, client *ent.Client, fn TxHandler) error {
	tx, err := client.Tx(ctx)
	if err != nil {
		return err
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
	}()

	if err := fn(tx); err != nil {
		return rollback(tx, fmt.Errorf("rolling back transaction: %s", err))
	}

	if err := tx.Commit(); err != nil {
		return rollback(tx, fmt.Errorf("commiting transaction: %s", err))
	}

	return nil
}

func rollback(tx *ent.Tx, err error) error {
	if e := tx.Rollback(); e != nil {
		err = fmt.Errorf("%w: %v", err, e)
	}

	return err
}
