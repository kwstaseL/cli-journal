package components

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/kwstaseL/cli-journal/pkg/export/utils"
	"github.com/kwstaseL/cli-journal/pkg/model"
)

type TxtExporter struct {
	Path string
}

func (t *TxtExporter) Export(note model.Note) error {
	filename := fmt.Sprintf("%d.txt", note.ID)
	fullPath := filepath.Join("./data", filename)

	content := utils.FormatNoteAsTxt(note, utils.NoteFormatOptions{
		IncludeTags:     true,
		IncludeCategory: true,
		IncludeCreated:  true,
	})

	err := os.WriteFile(fullPath, []byte(content), 0644)
	if err != nil {
		return fmt.Errorf("failed to write txt note to file: %w", err)
	}

	fmt.Println("Note exported to", fullPath)
	return nil
}
