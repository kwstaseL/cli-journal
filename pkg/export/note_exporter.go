package export

import "github.com/kwstaseL/cli-journal/pkg/model"

type NoteExporter interface {
	Export(note model.Note) error
}
