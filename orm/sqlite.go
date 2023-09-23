package orm

import (
	"errors"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type SqliteConnection struct {
	ConnString string
	db         gorm.DB
}

func (con *SqliteConnection) Connect(connString ...string) error {
	if len(connString) == 0 {
		db, err := gorm.Open(sqlite.Open(con.ConnString), &gorm.Config{})
		con.db = *db
		return err

	} else if len(connString) == 1 {
		db, err := gorm.Open(sqlite.Open(connString[0]), &gorm.Config{})
		con.db = *db
		return err

	} else {
		return errors.New("too many parameters")
	}
}
