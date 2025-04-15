package export

import (
	"fmt"
	"strings"
)

// TODO: Load Export path from Config
func GetExporter(platform string) (NoteExporter, error) {
	switch strings.ToLower(platform) {
	case "txt":
		return &TxtExporter{}, nil
	case "md": 
		return &MdExporter{}, nil
	case "json":
		return nil, fmt.Errorf("exporting with json is not yet implemented")
	default:
		return nil, fmt.Errorf("exporting with %s is not yet implemented", strings.ToLower(platform))
	}
	
}