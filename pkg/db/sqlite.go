package db

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type SQLite struct {
	Path string
}

func (s *SQLite) Connect(config DBConfig) (*gorm.DB, error) {
	dir := filepath.Dir(s.Path)
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return nil, fmt.Errorf("failed to create database directory: %w", err)
	}

	var gormConfig *gorm.Config
	if config.EnableLogging {
		gormConfig = &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		}
	} else {
		gormConfig = &gorm.Config{}
	}

	db, err := gorm.Open(sqlite.Open(s.Path), gormConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get database instance: %w", err)
	}

	if config.MaxIdleConns > 0 {
		sqlDB.SetMaxIdleConns(config.MaxIdleConns)
	}
	if config.MaxOpenConns > 0 {
		sqlDB.SetMaxOpenConns(config.MaxOpenConns)
	}

	return db, nil
}

func NewSQLite(path string) *SQLite {
	return &SQLite{Path: path}
}