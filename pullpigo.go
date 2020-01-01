package main

import (
	"bytes"
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

	eventsByAuthor := make(map[actor][]rawEvent)

	for _, event := range events {
		eventsByAuthor[event.Actor] = append(eventsByAuthor[event.Actor], event)
	}
	println("Pull requests")
	println(
		eventMessage(
			"\topened per author",
			func(event rawEvent) bool {
				return event.pullRequestOpened()
			},
			eventsByAuthor))
	println(
		eventMessage(
			"\tcommented per author",
			func(event rawEvent) bool {
				return event.pullRequestComment()
			},
			eventsByAuthor))
	println(
		eventMessage(
			"\tclosed per author",
			func(event rawEvent) bool {
				return event.pullRequestClosed()
			},
			eventsByAuthor))

}

type eventPredicate = func(event rawEvent) bool

func eventMessage(message string, fn eventPredicate, eventsByAuthor map[actor][]rawEvent) string {
	var buffer bytes.Buffer
	buffer.WriteString(message)
	for actor, events := range eventsByAuthor {
		count := 0
		for _, event := range events {
			if fn(event) {
				count++
			}
		}
		if count > 0 {
			fmt.Fprintf(&buffer, "\n\t\t%s: %v", actor.Login, count)
		}
	}
	return buffer.String()
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
	var events []rawEvent
	page := 1
	for ; page <= 10; page++ {
		url := fmt.Sprintf("https://api.github.com/repos/%v/events?access_token=%v&page=%v", config.repo, config.token, page)
		resp, getErr := http.Get(url)
		if getErr != nil {
			panic(getErr)
		}
		defer resp.Body.Close()
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		pageEvents := decodeEvents(bodyBytes)
		if len(pageEvents) == 0 {
			break
		}
		events = append(events, pageEvents...)
	}
	return events
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

func (event rawEvent) pullRequestClosed() bool {
	return event.EventType == "PullRequestEvent" && event.Payload.Action == "closed"
}

func (event rawEvent) pullRequestOpened() bool {
	return event.EventType == "PullRequestEvent" && event.Payload.Action == "opened"
}

func (event rawEvent) pullRequestComment() bool {
	return event.EventType == "PullRequestReviewCommentEvent" && event.Payload.Action == "created"
}
