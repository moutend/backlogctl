package migrate

import (
	"bytes"
	"database/sql"
	_ "embed"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/golang-migrate/migrate/v4/source/github"
	_ "github.com/mattn/go-sqlite3"
)

//go:embed 0001_init.up.sql
var data []byte

func Setup(db3Path string) error {
	if _, err := os.Stat(db3Path); err == nil {
		return nil
	}

	migrationsPath := filepath.Join(filepath.Dir(db3Path), "migrations")

	os.Mkdir(migrationsPath, 0755)

	file, err := os.Create(filepath.Join(migrationsPath, "0001_init.up.sql"))

	if err != nil {
		return err
	}
	if _, err := io.Copy(file, bytes.NewBuffer(data)); err != nil {
		return err
	}

	file.Close()

	dsn := fmt.Sprintf("file://%s?cache=shared&mode=rwc", db3Path)

	db, err := sql.Open(`sqlite3`, dsn)

	if err != nil {
		return err
	}

	defer db.Close()

	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})

	m, err := migrate.NewWithDatabaseInstance(
		"file://"+migrationsPath,
		"sqlite3", driver)

	if err != nil {
		return err
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}

	m.Close()

	return nil
}
