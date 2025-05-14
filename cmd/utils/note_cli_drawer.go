package utils

import (
	"fmt"
	"strings"

	"github.com/kwstaseL/cli-journal/pkg/model"
)

type NoteCLIDrawer struct{}

func NewNoteCLIDrawer() *NoteCLIDrawer {
	return &NoteCLIDrawer{}
}

func (d *NoteCLIDrawer) DrawSaveSuccess(note model.Note) {
	fmt.Println("Note saved successfully!")
	d.drawNoteBox(note)
}

func (d *NoteCLIDrawer) DrawListNotes(notes []model.Note) {
	if len(notes) == 0 {
		fmt.Println("No notes found.")
		return
	}

	fmt.Println("Recently added notes:")
	for i, note := range notes {
		fmt.Printf("\nNote #%d\n", i+1)
		d.drawNoteBox(note)
	}
}

func (d *NoteCLIDrawer) DrawError(context string, err error) {
	fmt.Printf("%s: %v\n", context, err)
}

// --- Internal Box Drawing Logic ---
func (d *NoteCLIDrawer) drawNoteBox(note model.Note) {
	const boxWidth = 50

	printLine(boxWidth)
	fmt.Printf("| %-10s: %-33s |\n", "Title", truncate(note.Header, 33))
	fmt.Printf("| %-10s: %-33s |\n", "Body", truncate(note.Body, 33))
	fmt.Printf("| %-10s: %-33s |\n", "Category", truncate(note.Category, 33))
	fmt.Printf("| %-10s: %-33s |\n", "Tags", truncate(note.Tags, 33))
	printLine(boxWidth)
}

func printLine(width int) {
	fmt.Println("+" + strings.Repeat("-", width-2) + "+")
}

func truncate(s string, max int) string {
	if len(s) > max {
		return s[:max-3] + "..."
	}
	return s
}
