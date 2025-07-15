package database

import (
    "database/sql"
    "log"

    _ "github.com/lib/pq"
    "github.com/golang-migrate/migrate/v4"
    "github.com/golang-migrate/migrate/v4/database/postgres"
    _ "github.com/golang-migrate/migrate/v4/source/file"
)

func New(databaseURL string) (*sql.DB, error) {
    db, err := sql.Open("postgres", databaseURL)
    if err != nil {
        return nil, err
    }

    if err := db.Ping(); err != nil {
        return nil, err
    }

    log.Println("Database connected successfully")
    return db, nil
}

func RunMigrations(databaseURL string) error {
    db, err := sql.Open("postgres", databaseURL)
    if err != nil {
        return err
    }
    defer db.Close()

    driver, err := postgres.WithInstance(db, &postgres.Config{})
    if err != nil {
        return err
    }

    m, err := migrate.NewWithDatabaseInstance(
        "file://migrations",
        "postgres", driver)
    if err != nil {
        return err
    }

    if err := m.Up(); err != nil && err != migrate.ErrNoChange {
        return err
    }

    log.Println("Migrations completed successfully")
    return nil
}