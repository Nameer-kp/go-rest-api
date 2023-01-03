package database

import (
	"fmt"
	"github.com/golang-migrate/migrate"
	migrate2 "github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func (d *Database) MigrateDB() error {
	fmt.Println("Migrating Databse")
	driver, err := postgres.WithInstance(d.Client.DB, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("could not create the postgres driver: %w", err)
	}
	m, err := migrate2.NewWithDatabaseInstance("file:///migrations", "postgres", driver)
	if err != nil || err != migrate.ErrNoChange {
		fmt.Println(err)
		return err
	}
	if err := m.Up(); err != nil {
		return fmt.Errorf("migration failed:%w", err)
	}
	fmt.Println("successfully migrated the database")
	return nil
}
