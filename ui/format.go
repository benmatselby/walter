package ui

import (
	"fmt"
	"io"
	"strings"

	goJira "github.com/andygrunwald/go-jira"
	"github.com/benmatselby/walter/jira"
)

// RenderTitle will produce an underlined title.
func RenderTitle(w io.Writer, title string) {
	fmt.Fprintf(w, "\n%s\n", title)
	fmt.Fprintf(w, strings.Repeat("=", len(title))+"\n")
}

// RenderItem understands how to render a single work item.
func RenderItem(w io.Writer, item goJira.Issue) {
	points, err := jira.GetStoryPoint(item, "")

	if err != nil {
		fmt.Fprintf(w, "* %s - %s\n", item.Key, item.Fields.Summary)
	} else {
		fmt.Fprintf(w, "* (%v) %s - %s\n", points, item.Key, item.Fields.Summary)
	}
}
