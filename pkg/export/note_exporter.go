package export

import "github.com/kwstaseL/cli-journal/pkg/model"

// TODO: Move this to Config and then .env
const ExportPath = "./data" 

type NoteExporter interface {
	Export(note model.Note) error
}
