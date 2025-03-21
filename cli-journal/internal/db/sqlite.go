package db

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type SQLite struct {
    Path string
}

func (s *SQLite) Connect() (*gorm.DB, error) {
    db, err := gorm.Open(sqlite.Open(s.Path), &gorm.Config{})
    if err != nil {
        return nil, err
    }
    return db, nil
}