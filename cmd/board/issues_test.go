package board_test

import (
	"bufio"
	"bytes"
	"errors"
	"testing"

	goJira "github.com/andygrunwald/go-jira"
	"github.com/benmatselby/walter/cmd/board"

	"github.com/benmatselby/walter/jira"
	"github.com/golang/mock/gomock"
)

func TestNewIssueCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	client := jira.NewMockAPI(ctrl)

	cmd := board.NewIssueCommand(client)

	use := "issues"
	short := "List all the issues for the board"

	if cmd.Use != use {
		t.Fatalf("expected use: %s; got %s", use, cmd.Use)
	}

	if cmd.Short != short {
		t.Fatalf("expected use: %s; got %s", short, cmd.Short)
	}
}

func TestListIssues(t *testing.T) {
	tt := []struct {
		name   string
		output string
		err    error
	}{
		{name: "can return a list of issues", output: "\nTodo\n====\n* Issue 1\n", err: nil},
		{name: "returns error if we cannot get list of issues", output: "", err: errors.New("something")},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			client := jira.NewMockAPI(ctrl)

			fields := goJira.IssueFields{
				Summary: "Issue 1",
				Status:  &goJira.Status{Name: "Todo"},
			}

			jiraIssues := goJira.Issue{
				Fields: &fields,
			}

			issues := make([]goJira.Issue, 0)
			issues = append(issues, jiraIssues)

			client.
				EXPECT().
				GetIssuesForBoard(gomock.Eq("board")).
				Return(issues, tc.err).
				AnyTimes()

			client.
				EXPECT().
				GetBoardLayout(gomock.Eq("board")).
				Return([]string{"Todo"}, tc.err).
				AnyTimes()

			var b bytes.Buffer
			writer := bufio.NewWriter(&b)

			_ = board.ListIssues(client, "board", writer)
			writer.Flush()

			if b.String() != tc.output {
				t.Fatalf("expected '%s'; got '%s'", tc.output, b.String())
			}
		})
	}
}
