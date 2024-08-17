package main

import (
	"log"
	"strconv"

	"github.com/emilyxfox/go-streamdeck-sdk/streamdeck"
)

type MyCounterAction struct {
	streamdeck.ActionConfig
}

var CounterAction = &MyCounterAction{
	ActionConfig: streamdeck.ActionConfig{UUID: "tld.domain.counterplugin.counteraction"},
}

var counter uint32

func (a *MyCounterAction) HandleKeyDown(event *streamdeck.KeyDownEvent) {
	log.Printf("Handling KeyDownEvent for UUID: %s, Event: %+v", a.UUID, event)
	counter++
	event.SetTitle(strconv.FormatUint(uint64(counter), 10))
}

func (a *MyCounterAction) HandleWillAppear(event *streamdeck.WillAppearEvent) {
	event.SetTitle(strconv.FormatUint(uint64(counter), 10))
}

// More actions can be defined similarly...
