# GitHub Activity CLI

A simple command-line interface (CLI) application written in Go that fetches and displays the recent activity of any GitHub user using the GitHub API.

Project URL: https://roadmap.sh/projects/github-user-activity

## Features

- Fetch recent public events for any GitHub user
- Display activity in a clean, readable format
- Support for multiple event types:
  - **Push Events**: Shows number of commits pushed to repositories
  - **Issues Events**: Displays issue-related activities
  - **Watch Events**: Shows starred repositories
- Fast response with 3-second timeout
- No authentication required for public data

## Requirements

- Go 1.25.6 or higher

## Installation

1. Clone the repository:
```bash
git clone https://github.com/shohratd15/github-activity.git
cd github-activity
```

2. Build the application:
```bash
go build -o github-activity ./cmd/github-activity
```

## Usage

Run the CLI with a GitHub username:

```bash
./github-activity <username>
```

### Example

```bash
./github-activity shohratd15
```

### Sample Output

```
Pushed 3 commits to shohratd15/my-project
Starred torvalds/linux
Issues activity in golang/go
Pushed 1 commits to shohratd15/github-activity
```

## Project Structure

```
github-activity/
├── cmd/
│   └── github-activity/
│       └── main.go          # Entry point
├── internal/
│   ├── github/
│   │   └── client.go        # GitHub API client
│   ├── models/
│   │   └── models.go        # Data models
│   └── usecase/
│       └── activity.go      # Business logic
├── go.mod
└── README.md
```

## How It Works

1. The CLI accepts a GitHub username as a command-line argument
2. It fetches the user's recent public events from the GitHub API (`https://api.github.com/users/{username}/events`)
3. Parses the JSON response and formats supported event types
4. Displays the activity in a human-readable format

## Error Handling

- If no username is provided, the CLI displays usage instructions
- If the user is not found or the API is unavailable, an error message is shown
- Requests timeout after 3 seconds to prevent hanging

## License

This project is open source and available for educational purposes.
