package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	config := parseFlags()
	fmt.Printf("GitHub repository '%s'", config.repo)
	if len(config.token) > 0 {
		print(" (token provided)")
	}
	println()
	events := githubEvents(config)

	eventsByAuthor := make(map[actor]int)
	for _, event := range events {
		eventsByAuthor[event.Actor]++
	}
	for author, eventNumber := range eventsByAuthor {
		fmt.Printf("  %v events created by %s", eventNumber, author.Login)
		println()
	}
}

type config struct {
	repo, token string
}

// Parse command line arguments to retrieve the configuration
func parseFlags() config {
	repo := flag.String("repo", "", "a GitHub repository (i.e. 'nicokosi/pullpigo')")
	token := flag.String("token", "", "an optional GitHub token")
	flag.Parse()
	if len(*repo) == 0 {
		println("The 'repo' option is mandatory")
		os.Exit(1)
	}
	return config{*repo, *token}
}

// Retrieve events from GitHub API
func githubEvents(config config) []rawEvent {
	url := fmt.Sprintf("https://api.github.com/repos/%v/events?access_token=%v&page=1", config.repo, config.token)
	resp, getErr := http.Get(url)
	if getErr != nil {
		panic(getErr)
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return decodeEvents(bodyBytes)
}

// Decode events from GitHub API
func decodeEvents(jsonBytes []byte) []rawEvent {
	var events []rawEvent
	json.Unmarshal(jsonBytes, &events)
	return events
}

type actor struct {
	Login string `json:"login"`
}
type payload struct {
	Action string `json:"action"`
}
type rawEvent struct {
	Actor     actor     `json:"actor"`
	Payload   payload   `json:"payload"`
	EventType string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
}
