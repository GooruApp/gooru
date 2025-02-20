package start

import (
	"context"
	"database/sql"
	"embed"
	"fmt"

	"github.com/GooruApp/gooru/server/internal/api"
	"github.com/GooruApp/gooru/server/internal/config"
	"github.com/GooruApp/gooru/server/migrations"
	"github.com/Masterminds/squirrel"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"go.uber.org/zap"
)

func Server(ctx context.Context) error {
	logger, err := newLogger("server", config.Settings.Mode.Get())
	if err != nil {
		return fmt.Errorf("couldn't create a new logger: %v", err)
	}

	if err := migrateUp(config.Settings.DBBackend.Get(), config.Settings.DBConnStr.Get()); err != nil {
		return fmt.Errorf("error occured when running migrations: %v", err)
	}

	db, err := newDB(config.Settings.DBBackend.Get(), config.Settings.DBConnStr.Get())
	if err != nil {
		return fmt.Errorf("error occured when setting up db: %v", err)
	}

	api := api.New(ctx, logger, db)
	srv := api.Server(config.Settings.Port.Get())

	go func() { _ = srv.ListenAndServe() }()

	fmt.Printf("Started server on port: %d\n", config.Settings.Port.Get())

	// Blocks until a value is passed on the done ch
	<-ctx.Done()

	_ = srv.Shutdown(ctx)

	return nil
}

func newLogger(service string, appEnv string) (*zap.Logger, error) {
	var logger *zap.Logger
	var err error

	switch appEnv {
	case "production":
		logger, err = zap.NewProduction(zap.Fields(
			zap.String("env", appEnv),
			zap.String("service", service),
		))
	default:
		logger, err = zap.NewDevelopment()
	}

	if err != nil {
		return nil, err
	}

	return logger, nil
}

func migrateUp(backend string, url string) error {
	var fs embed.FS

	switch backend {
	case "sqlite":
		fs = migrations.SQLiteMigrations
	case "postgres":
		fs = migrations.PGMigrations
	default:
		return fmt.Errorf("%s is not supported", backend)
	}

	source, err := iofs.New(fs, backend)
	if err != nil {
		return err
	}

	migrations, err := migrate.NewWithSourceInstance("iofs", source, url)
	if err != nil {
		return err
	}
	defer migrations.Close()

	err = migrations.Up()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}

func newDB(backend string, url string) (squirrel.StatementBuilderType, error) {
	pool, err := sql.Open(backend, url)
	if err != nil {
		return squirrel.StatementBuilderType{}, err
	}

	stmtCache := squirrel.NewStmtCache(pool)
	db := squirrel.StatementBuilder.RunWith(stmtCache)

	if backend == "postgres" {
		db = db.PlaceholderFormat(squirrel.Dollar)
	}

	return db, nil
}
