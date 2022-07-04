package search_test

import (
	"bufio"
	"bytes"
	"errors"
	"testing"

	goJira "github.com/andygrunwald/go-jira"
	"github.com/benmatselby/walter/cmd/search"
	"github.com/benmatselby/walter/jira"
	"github.com/golang/mock/gomock"
	"github.com/spf13/viper"
)

func TestCommandOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	client := jira.NewMockAPI(ctrl)

	cmd := search.NewSearchCommand(client)

	use := "search"
	short := "Search for issues"

	if cmd.Use != use {
		t.Fatalf("expected use: %s; got %s", use, cmd.Use)
	}

	if cmd.Short != short {
		t.Fatalf("expected use: %s; got %s", short, cmd.Short)
	}
}

func TestQueryIssues(t *testing.T) {
	tt := []struct {
		name               string
		format             string
		query              string
		template           string
		maxResults         int
		expectedQuery      string
		expectedMaxResults int
		output             string
		err                error
	}{
		{name: "can return a list of issues via query option", query: "status != Completed", maxResults: 10, expectedMaxResults: 10, expectedQuery: "status != Completed", output: "* 101 - Issue 1\n", err: nil},
		{name: "can return a list of issues via template option", template: "closed-issues", expectedMaxResults: 10, expectedQuery: "status = Completed", output: "* 101 - Issue 1\n", err: nil},
		{name: "can handle an error from the search", query: "status != Completed", maxResults: 10, expectedMaxResults: 10, expectedQuery: "status != Completed", output: "* 101 - Issue 1\n", err: errors.New("boom")},
		{name: "can handle the fact we dont pass template or query", query: "", expectedMaxResults: 43, err: errors.New("please use --query or --template to search")},
		{name: "can handle the fact a template may not be defined", template: "undefined-flim-flam", expectedMaxResults: 43, err: errors.New("undefined-flim-flam is not defined")},
		{name: "can return a table of issues via query option", query: "status != Completed", format: "table", maxResults: 10, expectedMaxResults: 10, expectedQuery: "status != Completed", output: `Metric Count
------ -----
Issues 1
------ -----
`, err: nil},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			client := jira.NewMockAPI(ctrl)

			// Setup the configuration via Viper
			viper.Set("templates.closed-issues.query", "status = Completed")
			viper.Set("templates.closed-issues.count", 43)

			// Define the mock jira issues
			fields := goJira.IssueFields{
				Summary: "Issue 1",
				Status:  &goJira.Status{Name: "Todo"},
			}

			jiraIssues := goJira.Issue{
				Key:    "101",
				Fields: &fields,
			}

			issues := make([]goJira.Issue, 0)
			issues = append(issues, jiraIssues)

			opts := search.CommandOptions{
				Format:     tc.format,
				MaxResults: tc.maxResults,
				Query:      tc.query,
				Template:   tc.template,
			}

			searchOpts := goJira.SearchOptions{
				MaxResults: tc.maxResults,
			}

			client.
				EXPECT().
				IssueSearch(gomock.Eq(tc.expectedQuery), gomock.Eq(&searchOpts)).
				Return(issues, tc.err).
				AnyTimes()

			var b bytes.Buffer
			writer := bufio.NewWriter(&b)

			err := search.QueryIssues(client, opts, writer)
			writer.Flush()

			if tc.err == nil {
				if b.String() != tc.output {
					t.Fatalf("expected '%s'; got '%s'", tc.output, b.String())
				}
			} else {
				if tc.err.Error() != err.Error() {
					t.Fatalf("expected '%v'; got '%v'", tc.err, err)
				}
			}
		})
	}
}

func TestQueryIssuesCanHandleStoryPoints(t *testing.T) {
	tt := []struct {
		name              string
		format            string
		query             string
		storyPointDefined bool
		points            float64
		output            string
	}{
		{name: "can handle the happy path of the story point defined", query: "status != Completed", storyPointDefined: true, points: 15, output: "* (15) 101 - Issue 1\n"},
		{name: "can handle the story point field defined but not value", query: "status != Completed", storyPointDefined: false, points: 0, output: "* 101 - Issue 1\n"},
		{name: "can return a table of issues via query option", query: "status != Completed", storyPointDefined: true, points: 15, format: "table", output: `Metric      Count
------      -----
Issues      1
Points      15
Not pointed 0
------      -----
`},
		{name: "can return a table of issues via query option no story points", query: "status != Completed", storyPointDefined: false, format: "table", output: `Metric Count
------ -----
Issues 1
------ -----
`},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			client := jira.NewMockAPI(ctrl)

			// Setup the configuration via Viper
			viper.Set("templates.closed-issues.query", "status = Completed")
			viper.Set("templates.closed-issues.count", 43)

			// Define the mock jira issues
			fields := goJira.IssueFields{
				Summary: "Issue 1",
				Status:  &goJira.Status{Name: "Todo"},
			}

			if tc.storyPointDefined {
				viper.Set("fields.story_point_field", "story_point_one")
				fields.Unknowns = map[string]interface{}{"story_point_one": tc.points}
			}

			jiraIssues := goJira.Issue{
				Key:    "101",
				Fields: &fields,
			}

			issues := make([]goJira.Issue, 0)
			issues = append(issues, jiraIssues)

			opts := search.CommandOptions{
				Format:     tc.format,
				MaxResults: 41,
				Query:      tc.query,
			}

			searchOpts := goJira.SearchOptions{
				MaxResults: 41,
			}

			client.
				EXPECT().
				IssueSearch(gomock.Eq(tc.query), gomock.Eq(&searchOpts)).
				Return(issues, nil).
				AnyTimes()

			var b bytes.Buffer
			writer := bufio.NewWriter(&b)

			_ = search.QueryIssues(client, opts, writer)
			writer.Flush()

			if b.String() != tc.output {
				t.Fatalf("expected '%s'; got '%s'", tc.output, b.String())
			}
		})
	}
}
