package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github-activity/internal/usecase"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	username, err := parseArgs(os.Args)
	if err != nil {
		log.Fatal(err)
	}

	lines, err := usecase.GetUserActivity(ctx, username)
	if err != nil {
		log.Fatal(err)
	}

	if len(lines) == 0 {
		fmt.Println("No activity found for user", username)
		return
	}

	for _, line := range lines {
		fmt.Println(line)
	}
}

func parseArgs(args []string) (string, error) {
	if len(args) < 2 {
		return "", fmt.Errorf("usage: github-activity <username>")
	}
	return args[1], nil
}
