package export

import (
	"fmt"
	"strings"

	"github.com/kwstaseL/cli-journal/pkg/config"
	"github.com/kwstaseL/cli-journal/pkg/export/components"
)

func GetExporter(platform string) (NoteExporter, error) {
	switch strings.ToLower(platform) {
	case "txt":
		return &components.TxtExporter{Path: config.Config.ExportPath}, nil
	case "md": 
		return &components.MdExporter{Path: config.Config.ExportPath}, nil
	case "json":
		return nil, fmt.Errorf("exporting with json is not yet implemented")
	default:
		return nil, fmt.Errorf("exporting with %s is not yet implemented", strings.ToLower(platform))
	}
	
}