package export

import (
	"fmt"
	"strings"

	"github.com/kwstaseL/cli-journal/pkg/export/components"
)

// TODO: Load Export path from Config
func GetExporter(platform string) (NoteExporter, error) {
	switch strings.ToLower(platform) {
	case "txt":
		return &components.TxtExporter{}, nil
	case "md": 
		return &components.MdExporter{}, nil
	case "json":
		return nil, fmt.Errorf("exporting with json is not yet implemented")
	default:
		return nil, fmt.Errorf("exporting with %s is not yet implemented", strings.ToLower(platform))
	}
	
}