package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/gorilla/websocket"
)

var actionRegistry = make(map[string]Action)

func main() {
	// Open log file
	logFile, err := os.OpenFile("streamdeck.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}

	// Set log output to the file
	log.SetOutput(logFile)

	defer logFile.Close()

	OpenWebsocketAndRegisterPlugin()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	exampleAction := &ExampleAction{
		BaseAction: BaseAction{UUID: "com.emilyxfox.counter.counter"},
	}
	RegisterAction(exampleAction)

	// Listen for messages from WebSocket
	go func() {
		for {
			_, message, err := WsClient.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}

			HandleEvent(message)
		}
	}()

	// Keep running until interrupted
	for {
		<-interrupt
		log.Println("Interrupt received, shutting down...")
		WsClient.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
		return
	}

}

func RegisterAction(action Action) {
	actionRegistry[action.GetUUID()] = action
}
