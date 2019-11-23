package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	repo := parseFlags()
	fmt.Printf("GitHub repository '%s'", repo)
	println()

	events := githubEvents(repo)

	eventsByAuthor := make(map[actor]int)
	for _, event := range events {
		eventsByAuthor[event.Actor]++
	}
	for author, eventNumber := range eventsByAuthor {
		fmt.Printf("  %v events created by %s", eventNumber, author.Login)
		println()
	}
}

// Parse command line arguments to retrieve repository name
func parseFlags() string {
	repo := flag.String("repo", "", "a GitHub repository (i.e. 'nicokosi/pullpigo')")
	flag.Parse()
	return *repo
}

// Retrieve events from GitHub API
func githubEvents(repo string) []rawEvent {
	url := fmt.Sprintf("https://api.github.com/repos/%v/events?access_token=&page=1", repo)
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
