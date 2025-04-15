package utils

import (
	"fmt"
	"strings"
	"time"

	"github.com/kwstaseL/cli-journal/pkg/model"
)

type NoteFormatOptions struct {
	IncludeTags     bool
	IncludeCategory bool
	IncludeCreated  bool
}

func FormatNoteAsMarkdown(note model.Note, opts NoteFormatOptions) string {
	var builder strings.Builder

	builder.WriteString(fmt.Sprintf("# %s\n", note.Header))

	if opts.IncludeCategory && strings.TrimSpace(note.Category) != "" {
		builder.WriteString(fmt.Sprintf("**Category**: %s  \n", note.Category))
	}
	if opts.IncludeTags && strings.TrimSpace(note.Tags) != "" {
		builder.WriteString(fmt.Sprintf("**Tags**: %s  \n", note.Tags))
	}

	builder.WriteString("\n---\n\n")
	builder.WriteString(note.Body)

	if opts.IncludeCreated {
		builder.WriteString(fmt.Sprintf("\n\n_Created at: %s_\n", note.CreatedAt.Format(time.RFC1123)))
	}

	return builder.String()
}

func FormatNoteAsTxt(note model.Note, opts NoteFormatOptions) string {
	var builder strings.Builder

	builder.WriteString(fmt.Sprintf("Title: %s\n\n", note.Header))
	builder.WriteString(fmt.Sprintf("%s\n\n", note.Body))

	if opts.IncludeTags && strings.TrimSpace(note.Tags) != "" {
		builder.WriteString(fmt.Sprintf("Tags: %s\n", note.Tags))
	}

	if opts.IncludeCategory && strings.TrimSpace(note.Category) != "" {
		builder.WriteString(fmt.Sprintf("Category: %s\n", note.Category))
	}

	if opts.IncludeCreated {
		builder.WriteString(fmt.Sprintf("Created At: %s\n", note.CreatedAt.Format(time.RFC1123)))
	}

	return builder.String()
}
