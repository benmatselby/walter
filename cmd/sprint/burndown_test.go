package sprint_test

import (
	"testing"

	"github.com/benmatselby/walter/cmd/sprint"

	"github.com/benmatselby/walter/jira"
	"github.com/golang/mock/gomock"
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
