package export

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/kwstaseL/cli-journal/pkg/model"
)

type TxtExporter struct {
	Path string
}

func (t *TxtExporter) Export(note model.Note) error {
	filename := fmt.Sprintf("%d.txt", note.ID)
	fullPath := filepath.Join(ExportPath, filename)

	file, err := os.Create(fullPath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	content := fmt.Sprintf("Title: %s\n\n%s\n\nTags: %s\nCategory: %s\nCreated At: %s\n",
		note.Header, note.Body, note.Tags, note.Category, note.CreatedAt.Format("2006-01-02 15:04"))

	_, err = file.WriteString(content)
	if err != nil {
		return fmt.Errorf("failed to write to file: %w", err)
	}

	fmt.Println("Note exported to", fullPath)
	return nil
}