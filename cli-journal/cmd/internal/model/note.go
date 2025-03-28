/*
Copyright © 2025 Konstantinos Eleftheriou <eleftheriou.konst@gmail.com>
*/
package model

import (
	"fmt"
	"strings"

	"gorm.io/gorm"
)

type Note struct {
	gorm.Model
	Title     string   `json:"title" gorm:"not null"`
	Content   string   `json:"content" gorm:"not null"`
	Tags      []string `json:"tags" gorm:"type:text[]"`
	Category  string   `json:"category" gorm:"default:'general'"`
}

func (n Note) String() string {
	return fmt.Sprintf("Note ID: %d\nTitle: %s\nContent: %s\nCategory: %s\nTags: %s", 
		n.ID, n.Title, n.Content, n.Category, strings.Join(n.Tags, ", "))
}

func NewNote(title, content, category string, tags []string) *Note {
	return &Note{
		Title:     title,
		Content:   content,
		Tags:      tags,
		Category:  category,
	}
}