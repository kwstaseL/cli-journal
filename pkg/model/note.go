/*
Copyright © 2025 Konstantinos Eleftheriou <eleftheriou.konst@gmail.com>
*/
package model

import (
	"fmt"

	"gorm.io/gorm"
)

type NoteFilters struct {
	Tags      string
	Category  string 
	SearchTerm string
}

type Note struct {
    gorm.Model
    Header    string `json:"header" gorm:"not null"`
    Body  	  string `json:"body" gorm:"not null"`
    Tags     string `json:"tags" gorm:"type:text"`  // (comma-separated tags)
    Category string `json:"category" gorm:"default:'general'"`
}

func (n Note) String() string {
	return fmt.Sprintf("Note ID: %d\nHeader: %s\nBody: %s\nCategory: %s\nTags: %s", 
		n.ID, n.Header, n.Body, n.Category, n.Tags)
}

func NewNote(header, body, category, tags string) *Note {
	return &Note{
		Header:     header,
		Body:   body,
		Tags:      tags,
		Category:  category,
	}
}
