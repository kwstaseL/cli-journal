package share

import (
	"fmt"
	"strings"

	"github.com/kwstaseL/cli-journal/pkg/share/components"
)

func GetSharer(platform string) (NoteSharer, error) {
	switch strings.ToLower(platform) {
	case "x", "twitter":
		return &components.XSharer{}, nil
	case "email":
		return nil, fmt.Errorf("sharing with email is not supported yet")
	default:
		return nil, fmt.Errorf("unsupported platform: %s", platform)
	}
}
