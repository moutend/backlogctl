package cache

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/moutend/go-backlog/pkg/client"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

var (
	BacklogSpace string
	db           *sql.DB
	backlog      *client.Client
)

func Setup(space, token, db3Path string) (err error) {
	backlog, err = client.New(space, token)

	if err != nil {
		return err
	}
	if err := setupDB(db3Path); err != nil {
		return err
	}

	BacklogSpace = space

	return nil
}

func setupDB(db3Path string) (err error) {
	if _, err := os.Stat(db3Path); err != nil {
		return err
	}

	dsn := fmt.Sprintf("file://%s?cache=shared&mode=rwc", db3Path)

	db, err = sql.Open(`sqlite3`, dsn)

	if err != nil {
		return err
	}

	boil.SetDB(db)

	return nil
}

func teardownDB() error {
	db.Close()

	return nil
}
