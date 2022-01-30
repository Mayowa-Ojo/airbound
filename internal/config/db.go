package config

import (
	"airbound/internal/ent"
	"airbound/internal/ent/hooks"
	"airbound/internal/ent/migrate"
	"airbound/internal/log"
	"context"
	"fmt"
	"os"

	_ "airbound/internal/ent/runtime"

	"entgo.io/ent/dialect/sql/schema"
	_ "github.com/lib/pq"
)

func NewEntClient(cfg *Config) *ent.Client {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName,
	)

	client, err := ent.Open(cfg.DBDriver, dsn)
	if err != nil {
		log.Fatal("[ENT]: error connecting to db %s", err)
	}

	client.Airline.Use(hooks.EnsureUppercaseField())

	log.Info("[ENT]: connected to database")

	return client
}

func SchemaMigrateUp(ctx context.Context, client *ent.Client) {
	if err := client.Schema.Create(ctx, schema.WithAtlas(true)); err != nil {
		log.Fatal("[ENT]: error running migration %s", err)
	}
}

func SchemaMigrateDown(ctx context.Context, client *ent.Client) {
	if err := client.Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	); err != nil {
		log.Fatal("[ENT]: error running migration %s", err)
	}
}

func DumpMigrations(ctx context.Context, client *ent.Client) {
	file, err := os.Create("migrations.sql")
	if err != nil {
		log.Fatal("[ENT]: error reading/creating migrations file %s", err)
	}

	if err := client.Schema.WriteTo(ctx, file); err != nil {
		log.Fatal("[ENT]: error writing migrations to file %s", err)
	}
}
