package sprint_test

import (
	"bufio"
	"bytes"
	"errors"
	"testing"

	goJira "github.com/andygrunwald/go-jira"
	"github.com/benmatselby/walter/cmd/sprint"
	"github.com/benmatselby/walter/jira"
	"github.com/golang/mock/gomock"
)

func TestNewIssuesCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	client := jira.NewMockAPI(ctrl)

	cmd := sprint.NewIssueCommand(client)

	use := "issues"
	short := "List all the issues for a given project sprint"

	if cmd.Use != use {
		t.Fatalf("expected use: %s; got %s", use, cmd.Use)
	}

	if cmd.Short != short {
		t.Fatalf("expected use: %s; got %s", short, cmd.Short)
	}
}

func TestNewIssuesCommandReturnsOutput(t *testing.T) {
	tt := []struct {
		name   string
		args   []string
		query  string
		output string
		err    error
	}{
		{
			name: "can return a list of issues",
			args: []string{"board", "sprint"},
			query: "sprint = 'sprint' and type = 'Story'",
			output: "\nTodo\n====\n* KEY-1 - Issue 1\n",
			err: nil,
		},
		{
			name: "returns error if we cannot get list of issues",
			args: []string{"board", "sprint"},
			query: "sprint = 'sprint' and type = 'Story'",
			output: "", 
			err: errors.New("something"),
		},
		{
			name: "returns a list of issue given a project name",
			args: []string{"board", "sprint", "project"},
			query: "sprint = 'sprint' and project = 'project' and type = 'Story'",
			output: "\nTodo\n====\n* KEY-1 - Issue 1\n",
			err: nil,
		},
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
				Key:    "KEY-1",
				Fields: &fields,
			}

			issues := make([]goJira.Issue, 0)
			issues = append(issues, jiraIssues)

			client.
				EXPECT().
				IssueSearch(gomock.Eq(tc.query), gomock.Eq(&goJira.SearchOptions{MaxResults: 7})).
				Return(issues, tc.err).
				AnyTimes()

			client.
				EXPECT().
				GetBoardLayout(gomock.Eq("board")).
				Return([]string{"Todo"}, tc.err).
				AnyTimes()

			var b bytes.Buffer
			writer := bufio.NewWriter(&b)

			opts := sprint.IssueOptions{
				Args:       tc.args,
				MaxResults: 7,
				FilterType: "Story",
			}

			err := sprint.ListIssues(client, opts, writer)
			writer.Flush()

			if b.String() != tc.output {
				t.Fatalf("expected '%s'; got '%s'", tc.output, b.String())
			}

			if err != nil && err != tc.err {
				t.Fatalf("expected err '%s'; got '%s'", tc.err, err)
			}
		})
	}
}
