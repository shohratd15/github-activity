package usecase

import (
	"context"
	"encoding/json"
	"fmt"

	"github-activity/internal/github"
	"github-activity/internal/models"
)

func GetUserActivity(ctx context.Context, username string) ([]string, error) {
	data, err := github.FetchEvents(ctx, username)
	if err != nil {
		return nil, err
	}

	var events []models.Event
	if err := json.Unmarshal(data, &events); err != nil {
		return nil, err
	}

	lines := make([]string, 0, len(events))
	for _, e := range events {
		line := formatEvent(e)
		if line != "" {
			lines = append(lines, line)
		}
	}

	return lines, nil
}

func formatEvent(e models.Event) string {
	switch e.Type {
	case "PushEvent":
		return fmt.Sprintf(
			"Pushed %d commits to %s",
			len(e.Payload.Commits),
			e.Repo.Name,
		)

	case "IssuesEvent":
		return fmt.Sprintf(
			"Issues activity in %s",
			e.Repo.Name,
		)

	case "WatchEvent":
		return fmt.Sprintf(
			"Starred %s",
			e.Repo.Name,
		)

	default:
		return ""
	}
}
