package db

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type SQLite struct {
    Path string
}

type SQLiteConfig struct {

}

func (s *SQLite) Connect(config SQLiteConfig) (*gorm.DB, error) {
    db, err := gorm.Open(sqlite.Open(s.Path), &gorm.Config{})
    if err != nil {
        return nil, err
    }
    return db, nil
}