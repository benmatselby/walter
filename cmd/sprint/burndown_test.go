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
	"github.com/spf13/viper"
)

func TestNewBurndownCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	client := jira.NewMockAPI(ctrl)

	cmd := sprint.NewBurndownCommand(client)

	use := "burndown"
	short := "Display the burndown for the sprint"

	if cmd.Use != use {
		t.Fatalf("expected use: %s; got %s", use, cmd.Use)
	}

	if cmd.Short != short {
		t.Fatalf("expected use: %s; got %s", short, cmd.Short)
	}
}

func TestShowBurndown(t *testing.T) {
	tt := []struct {
		name            string
		customFieldKey  string
		customFieldKeys []string
		storyPointField string
		output          string
		err             error
	}{
		{name: "can return a list of issues", customFieldKey: "field_1", customFieldKeys: []string{"field_2", "field_3"}, storyPointField: "field_1", output: `Column Items Points
------ ----- ------
Todo   1     1
------ ----- ------
Total  1     1
------ ----- ------
`, err: nil},
		{name: "can use a secondary story point field", customFieldKey: "field_1", customFieldKeys: []string{"field_2", "field_3"}, storyPointField: "field_2", output: `Column Items Points
------ ----- ------
Todo   1     1
------ ----- ------
Total  1     1
------ ----- ------
`, err: nil},
		{name: "can use a third story point field", customFieldKey: "field_1", customFieldKeys: []string{"field_2", "field_3"}, storyPointField: "field_3", output: `Column Items Points
------ ----- ------
Todo   1     1
------ ----- ------
Total  1     1
------ ----- ------
`, err: nil},
		{name: "returns error if we cannot get list of issues", output: "", err: errors.New("something")},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			client := jira.NewMockAPI(ctrl)

			viper.Set("boards.board.story_point_field", tc.customFieldKey)
			viper.Set("boards.board.story_point_fields", tc.customFieldKeys)

			unknowns := make(map[string]interface{})
			unknowns[tc.storyPointField] = float64(1)

			fields := goJira.IssueFields{
				Summary:  "Issue 1",
				Status:   &goJira.Status{Name: "Todo"},
				Unknowns: unknowns,
			}

			jiraIssues := goJira.Issue{
				Key:    "KEY-1",
				Fields: &fields,
			}

			issues := make([]goJira.Issue, 0)
			issues = append(issues, jiraIssues)

			client.
				EXPECT().
				IssueSearch(gomock.Eq("sprint = 'sprint' and type = 'Story'"), gomock.Eq(&goJira.SearchOptions{MaxResults: 7})).
				Return(issues, tc.err).
				AnyTimes()

			client.
				EXPECT().
				GetBoardLayout(gomock.Eq("board")).
				Return([]string{"Todo"}, tc.err).
				AnyTimes()

			client.
				EXPECT().
				GetStoryPoint(gomock.Any(), gomock.Eq("board")).
				Return(1, nil).
				AnyTimes()

			var b bytes.Buffer
			writer := bufio.NewWriter(&b)

			opts := sprint.BurndownOptions{
				Args:       []string{"board", "sprint"},
				FilterType: "Story",
				MaxResults: 7,
			}

			err := sprint.ShowBurndown(client, opts, writer)
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
