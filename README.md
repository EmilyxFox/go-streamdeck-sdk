# StreamDeck SDK for Go

Welcome to a scuffed StreamDeck SDK written for Go.

## Features include

- Potential to work
- Incomplete documentation
- Doubtful WebSocket handling

## Installation

```bash
go get github.com/emilyxfox/go-streamdeck-sdk/streamdeck
```

## Instructions

### Step 1:

Create a new struct that embeds `streamdeck.BaseAction` and implements the event handler methods you need.

```go
package main

import (
	"log"
	"strconv"

	"github.com/emilyxfox/streamdeck-sdk/streamdeck"
)

type MyCounterAction struct {
	streamdeck.BaseAction
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
```

### Step 2:

In your `main` function, register the action, connect to the WebSocket, and start listening for events.

```go
package main

import (
	"log"

	"github.com/emilyxfox/streamdeck-sdk/streamdeck"
)

func main() {
	// Create and register your custom action
	action := &MyCounterAction{
		BaseAction: streamdeck.BaseAction{UUID: "com.emilyxfox.counter.counter"},
	}
	streamdeck.RegisterAction(action)

	// Start listening for events
	streamdeck.StartPlugin()
}
```

## Documentation
For more information and instructions on how to structure your plugin please refer to [the official Stream Deck docs](https://docs.elgato.com/sdk).