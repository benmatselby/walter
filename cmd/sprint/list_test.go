package sprint_test

import (
	"bufio"
	"bytes"
	"errors"
	"testing"

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
		name   string
		output string
		err    error
	}{
		{name: "can return a list of sprints", output: "2019.1\n", err: nil},
		{name: "returns error if we cannot get list of sprints", output: "", err: errors.New("something")},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			client := jira.NewMockAPI(ctrl)

			name := "2019.1"
			jiraSprint := goJira.Sprint{
				Name: name,
			}

			sprints := make([]goJira.Sprint, 0)
			sprints = append(sprints, jiraSprint)

			client.
				EXPECT().
				GetSprints("").
				Return(sprints, tc.err).
				AnyTimes()

			var b bytes.Buffer
			writer := bufio.NewWriter(&b)

			sprint.ListSprints(client, "", writer)
			writer.Flush()

			if b.String() != tc.output {
				t.Fatalf("expected '%s'; got '%s'", tc.output, b.String())
			}
		})
	}
}
