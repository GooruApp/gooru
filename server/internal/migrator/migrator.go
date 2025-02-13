package migrator

import (
	"embed"
	"fmt"
	"strings"

	"github.com/GooruApp/gooru/server/migrations"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

type Flavor string

const (
	postgres Flavor = "postgres"
	sqlite   Flavor = "sqlite"
)

type migrator struct {
	flavor Flavor
	url    string
	fs     embed.FS
}

func New(url string) (*migrator, error) {
	before, _, found := strings.Cut(url, "://")
	if !found {
		return nil, fmt.Errorf("could not parse db flavor from provided url")
	}

	var flavor Flavor
	var fs embed.FS

	switch before {
	case "sqlite":
		flavor = sqlite
		fs = migrations.SQLiteMigrations
	case "postgres":
		flavor = postgres
		fs = migrations.PGMigrations
	default:
		return nil, fmt.Errorf("%s is not supported", before)
	}

	migrator := &migrator{
		flavor: flavor,
		url:    url,
		fs:     fs,
	}

	return migrator, nil
}

func (m *migrator) Up() error {
	source, err := iofs.New(m.fs, string(m.flavor))
	if err != nil {
		return err
	}

	migrations, err := migrate.NewWithSourceInstance("iofs", source, m.url)
	if err != nil {
		return err
	}

	err = migrations.Up()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}

func (m *migrator) Flavor() Flavor { return m.flavor }
