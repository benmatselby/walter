package jira

import (
	"fmt"

	jira "github.com/andygrunwald/go-jira"
	"github.com/spf13/viper"
)

// GetStoryPoint will figure out the story points assigned to a work item,
// based on what is defined in the config. The field to store story points
// can be configured in a few places, which this application tries to
// understand.
func GetStoryPoint(issue jira.Issue, boardName string) (int, error) {
	storyField := viper.GetString("fields.story_point_field")
	boardStoryField := viper.GetString(fmt.Sprintf("boards.%s.story_point_field", boardName))
	boardStoryFields := viper.GetStringSlice(fmt.Sprintf("boards.%s.story_point_fields", boardName))

	if storyField == "" && boardStoryField == "" && len(boardStoryFields) == 0 {
		return 0, fmt.Errorf("no story point field defined")
	}

	storyFieldValue := issue.Fields.Unknowns[storyField]
	if storyFieldValue != nil {
		return int(storyFieldValue.(float64)), nil
	}

	boardStoryFieldValue := issue.Fields.Unknowns[boardStoryField]

	if boardStoryFieldValue != nil {
		return int(boardStoryFieldValue.(float64)), nil
	}

	for _, key := range boardStoryFields {
		value := issue.Fields.Unknowns[key]
		if value != nil {
			return int(value.(float64)), nil
		}
	}

	return 0, fmt.Errorf("could not get story point value")
}
