package board_test

import (
	"testing"

	"github.com/benmatselby/walter/cmd/board"

	"github.com/benmatselby/walter/jira"
	"github.com/golang/mock/gomock"
)

func TestNewBoardCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	client := jira.NewMockAPI(ctrl)

	cmd := board.NewBoardCommand(client)

	use := "board"
	short := "Board related commands"

	if cmd.Use != use {
		t.Fatalf("expected use: %s; got %s", use, cmd.Use)
	}

	if cmd.Short != short {
		t.Fatalf("expected use: %s; got %s", short, cmd.Short)
	}
}
