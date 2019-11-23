package main

import (
	"testing"
	"time"
)

func Test_GithubEvents_can_return_empty_array(t *testing.T) {
	events := decodeEvents([]byte("{}"))
	if len(events) != 0 {
		t.Errorf("Expecting empty array")
	}
}

func Test_GithubEvents_can_return_a_single_event_PullRequestEvent(t *testing.T) {
	events := decodeEvents([]byte(
		`[{
			"id":"1",
			"type":"PullRequestEvent",
			"actor":{
				"login":"alice",
				"display_login":"Alice"
			},
			"repo":{
				"id":2,
				"name":"softwarevidal/fake-repo"
			},
			"payload":{
				"action":"opened"
			},
			"public":false,
			"created_at":"2016-12-01T16:26:43Z"
		}]`))
	if len(events) != 1 {
		t.Errorf("Expecting single event")
	}
	expectedCreatedAt, _ := time.Parse(time.RFC3339, "2016-12-01T16:26:43Z")
	expected := rawEvent{
		Actor:     actor{Login: "alice"},
		Payload:   payload{Action: "opened"},
		EventType: "PullRequestEvent",
		CreatedAt: expectedCreatedAt,
	}
	if events[0] != expected {
		t.Errorf("Unexpected event")
	}

}
