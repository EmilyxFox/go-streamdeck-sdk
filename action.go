package main

import (
	"log"
	"strconv"
)

type ExampleAction struct {
	ActionConfig
}

var CounterAction = &ExampleAction{
	ActionConfig: ActionConfig{UUID: "com.emilyxfox.counter.counter"},
}

var counter uint32

func (a *ExampleAction) HandleKeyDown(event *KeyDownEvent) {
	log.Printf("Handling KeyDownEvent for UUID: %s, Event: %+v", a.UUID, event)
	counter++
	event.SetTitle(strconv.FormatUint(uint64(counter), 10))
}

func (a *ExampleAction) HandleWillAppear(event *WillAppearEvent) {
	event.SetTitle(strconv.FormatUint(uint64(counter), 10))
}

// More actions can be defined similarly...
