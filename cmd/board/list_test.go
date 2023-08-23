package board_test

import (
	"bufio"
	"bytes"
	"errors"
	"testing"

	"github.com/benmatselby/walter/cmd/board"

	goJira "github.com/andygrunwald/go-jira"
	"github.com/benmatselby/walter/jira"
	"github.com/golang/mock/gomock"
)

func TestNewListCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	client := jira.NewMockAPI(ctrl)

	cmd := board.NewListCommand(client)

	use := "list"
	short := "List all the boards"

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
		{name: "can return a list of boards", output: "Magical board\n", err: nil},
		{name: "returns error if we cannot get list of boards", output: "", err: errors.New("something")},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			client := jira.NewMockAPI(ctrl)

			name := "Magical board"
			jiraBoard := goJira.Board{
				Name: name,
			}

			boards := make([]goJira.Board, 0)
			boards = append(boards, jiraBoard)

			client.
				EXPECT().
				GetBoards().
				Return(boards, tc.err).
				AnyTimes()

			var b bytes.Buffer
			writer := bufio.NewWriter(&b)

			_ = board.DisplayBoards(client, writer)
			writer.Flush()

			if b.String() != tc.output {
				t.Fatalf("expected '%s'; got '%s'", tc.output, b.String())
			}
		})
	}
}
