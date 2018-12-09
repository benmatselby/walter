package cmd_test

import (
	"testing"

	"github.com/benmatselby/walter/cmd"
	"github.com/benmatselby/walter/jira"
	"github.com/golang/mock/gomock"
)

func TestNewRootCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	client := jira.NewMockAPI(ctrl)

	cmd := cmd.NewRootCommand(client)

	use := "walter"
	short := "CLI application for retrieving data from Jira"

	if cmd.Use != use {
		t.Fatalf("expected use: %s; got %s", use, cmd.Use)
	}

	if cmd.Short != short {
		t.Fatalf("expected use: %s; got %s", short, cmd.Short)
	}
}
