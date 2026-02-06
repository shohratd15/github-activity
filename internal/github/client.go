package github

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

func FetchEvents(ctx context.Context, username string) ([]byte, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s/events", username)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	// GitHub требует User-Agent
	req.Header.Set("User-Agent", "github-activity-cli")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
		// ok
	case http.StatusNotFound:
		return nil, fmt.Errorf("user %q not found", username)
	case http.StatusForbidden:
		return nil, fmt.Errorf("github api rate limit exceeded")
	default:
		return nil, fmt.Errorf("github api returned status %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
