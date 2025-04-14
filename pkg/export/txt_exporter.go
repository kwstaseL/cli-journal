package export

import (
	"fmt"

	"github.com/kwstaseL/cli-journal/pkg/model"
)

type TxtExporter struct {
	Path string
}

func (t *TxtExporter) Export(note model.Note) error {
	fmt.Println("Exporting as txt at path ", t.Path)
	return nil
}