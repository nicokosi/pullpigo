package main

import (
	"math/rand"
	"reflect"
	"testing"
	"testing/quick"
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

func TestCounterMessageShouldBePrintedIfAtLeastAPullRequestEventOccured(t *testing.T) {
	f := func(title string, eventsByAuthor map[actor][]rawEvent) bool {
		message := eventMessage(title, func(rawEvent) bool { return true }, eventsByAuthor)
		if len(eventsByAuthor) > 0 {
			return len(message) > len(title)
		}
		return true
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}
func TestCounterMessageShouldNotBePrintedIfNoPullRequestEventOccured(t *testing.T) {
	f := func(title string, eventsByAuthor map[actor][]rawEvent) bool {
		message := eventMessage(title, func(rawEvent) bool { return true }, eventsByAuthor)
		if len(eventsByAuthor) == 0 {
			return len(message) == len(title)
		}
		return true
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func (rawEvent) Generate(r *rand.Rand, size int) reflect.Value {
	event := rawEvent{
		Actor:     actor{Login: generateString(r, int(r.Int31()%100+1))},
		Payload:   payload{Action: generateString(r, int(r.Int31()%100+1))},
		EventType: "PullRequestEvent",
		CreatedAt: time.Unix(r.Int63(), r.Int63()),
	}
	return reflect.ValueOf(event)
}

// Create a random string
func generateString(r *rand.Rand, size int) string {
	res := make([]byte, size)
	for i := range res {
		res[i] = byte(r.Int())
	}
	return string(res)
}
