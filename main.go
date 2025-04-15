/*
Copyright © 2025 Konstantinos Eleftheriou <eleftheriou.konst@gmail.com>
*/
package main

import (
	"github.com/kwstaseL/cli-journal/cmd"
	"github.com/kwstaseL/cli-journal/pkg/db"
	"github.com/kwstaseL/cli-journal/pkg/db/utils"
	logger "github.com/kwstaseL/cli-journal/pkg/logging"
	"github.com/kwstaseL/cli-journal/pkg/repository"
	"github.com/kwstaseL/cli-journal/pkg/service"
	"gorm.io/gorm"
)

func initializeDB(dbPath string) (*gorm.DB, error) {
    sqliteDB := db.NewSQLite(dbPath)
    return sqliteDB.Connect(utils.SQLiteConfig)
}

func main() {
	database, err := initializeDB("data/notes.db")
	if err != nil {
		logger.LogError("Failed to connect to the database: ", err)
	}
	logger.LogDebug("test")

	noteRepo := repository.NewNoteRepository(database)
	noteService := service.NewNoteService(noteRepo)

	cmd.SetNoteService(noteService)
	cmd.Execute()
}
