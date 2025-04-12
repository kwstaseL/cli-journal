package share

import "github.com/kwstaseL/cli-journal/pkg/model"

type NoteSharer interface {
	Share(note model.Note) error
}
