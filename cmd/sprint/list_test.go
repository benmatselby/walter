package sprint_test

import (
	"bufio"
	"bytes"
	"errors"
	"testing"
	"time"

	"github.com/benmatselby/walter/cmd/sprint"

	goJira "github.com/andygrunwald/go-jira"
	"github.com/benmatselby/walter/jira"
	"github.com/golang/mock/gomock"
)

func TestNewListCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	client := jira.NewMockAPI(ctrl)

	cmd := sprint.NewListCommand(client)

	use := "list"
	short := "List all the sprints"

	if cmd.Use != use {
		t.Fatalf("expected use: %s; got %s", use, cmd.Use)
	}

	if cmd.Short != short {
		t.Fatalf("expected use: %s; got %s", short, cmd.Short)
	}
}

func TestDisplayBoards(t *testing.T) {

	tt := []struct {
		name    string
		verbose bool
		args    []string
		output  string
		err     error
	}{
		{name: "can return a list of sprints", output: "2019.1\n", verbose: false, args: []string{"board"}, err: nil},
		{name: "can be verbose when asked", output: `2019.1
 ID: 717
 Start date: 2019-08-01 09:00:00 +0000 UTC
 End date: 2019-08-14 17:00:00 +0000 UTC
 Completed date: 2019-08-14 16:00:00 +0000 UTC

`, verbose: true, args: []string{"board"}, err: nil},
		{name: "returns error if we cannot get list of sprints", output: "", verbose: false, args: []string{"board"}, err: errors.New("something")},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			client := jira.NewMockAPI(ctrl)

			start, _ := time.Parse(time.RFC3339, "2019-08-01T09:00:00Z")
			completed, _ := time.Parse(time.RFC3339, "2019-08-14T16:00:00Z")
			end, _ := time.Parse(time.RFC3339, "2019-08-14T17:00:00Z")
			name := "2019.1"
			jiraSprint := goJira.Sprint{
				Name:         name,
				ID:           717,
				StartDate:    &start,
				CompleteDate: &completed,
				EndDate:      &end,
			}

			sprints := make([]goJira.Sprint, 0)
			sprints = append(sprints, jiraSprint)

			client.
				EXPECT().
				GetSprints("board").
				Return(sprints, tc.err).
				AnyTimes()

			var b bytes.Buffer
			writer := bufio.NewWriter(&b)

			opts := sprint.ListSprintOptions{
				Verbose: tc.verbose,
				Args:    tc.args,
			}

			sprint.ListSprints(client, opts, writer)
			writer.Flush()

			if b.String() != tc.output {
				t.Fatalf("expected '%s'; got '%s'", tc.output, b.String())
			}
		})
	}
}
