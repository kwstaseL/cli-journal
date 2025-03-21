/*
Copyright © 2025 Konstantinos Eleftheriou <eleftheriou.konst@gmail.com>
*/
package model

import (
	"fmt"

	"gorm.io/gorm"
)

type Note struct {
	gorm.Model
	Title     string  
	Content   string 
	Tags      []string
	Category  string
}

func (n Note) String() string {
	return fmt.Sprintf("Note title: %s\n Note content: %s", n.Title, n.Content)
}

func NewNote(title, content, category string, tags []string) *Note {
	return &Note{
		Title:     title,
		Content:   content,
		Tags:      tags,
		Category:  category,
	}
}