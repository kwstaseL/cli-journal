package export

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/kwstaseL/cli-journal/pkg/model"
)

type MdExporter struct {
	Path string
}

func (m *MdExporter) Export(note model.Note) error {
	filename := fmt.Sprintf("%d.md", note.ID)
	fullPath := filepath.Join(ExportPath, filename)

	content := FormatNoteAsMarkdown(note, NoteFormatOptions{
		IncludeTags:     true,
		IncludeCategory: true,
		IncludeCreated:  true,
	})

	err := os.WriteFile(fullPath, []byte(content), 0644)
	if err != nil {
		return fmt.Errorf("failed to write markdown note to file: %w", err)
	}

	return nil
}
