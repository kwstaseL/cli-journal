package share

import "github.com/kwstaseL/cli-journal/pkg/model"

type EmailSharer struct {}

type Email struct {

}

func (e *EmailSharer) Share(note model.Note) error {
	return nil
}