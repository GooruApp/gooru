package migrator

import (
	"embed"
	"fmt"
	"strings"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

//go:embed migrations/*/*.sql
var migrationsFS embed.FS

type MigratorFlavor string

const (
	postgres MigratorFlavor = "postgres"
	sqlite   MigratorFlavor = "sqlite"
)

type migrator struct {
	flavor MigratorFlavor
	url    string
}

func New(url string) (*migrator, error) {
	before, _, found := strings.Cut(url, "://")
	if !found {
		return nil, fmt.Errorf("could not parse db flavor from provided url")
	}

	var flavor MigratorFlavor

	switch before {
	case "sqlite":
		flavor = sqlite
	case "postgres":
		flavor = postgres
	default:
		return nil, fmt.Errorf("%s is not supported", before)
	}

	migrator := &migrator{
		flavor: flavor,
		url:    url,
	}

	return migrator, nil
}

func (m *migrator) Up() error {
	path := fmt.Sprintf("migrations/%s", m.flavor)

	source, err := iofs.New(migrationsFS, path)
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

func (m *migrator) Flavor() MigratorFlavor { return m.flavor }
