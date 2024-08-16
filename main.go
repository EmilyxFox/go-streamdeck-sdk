package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"

	"github.com/gorilla/websocket"
)

// PluginConfig holds the configuration passed via flags
type PluginConfigType struct {
	Port          string
	PluginUUID    string
	RegisterEvent string
	Info          StreamDeckInfo
}

var actionRegistry = make(map[string]Action)

var WsClient *websocket.Conn

var PluginConfig PluginConfigType

func main() {
	// Open log file
	logFile, err := os.OpenFile("streamdeck.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}

	// Set log output to the file
	log.SetOutput(logFile)

	defer logFile.Close()

	// Define flags
	port := flag.String("port", "", "WebSocket port")
	pluginUUID := flag.String("pluginUUID", "", "Plugin UUID")
	registerEvent := flag.String("registerEvent", "", "Event to register")
	info := flag.String("info", "", "Stream Deck information")

	flag.Parse()

	// Parse the info JSON
	var sdInfo StreamDeckInfo
	if err := json.Unmarshal([]byte(*info), &sdInfo); err != nil {
		log.Fatalf("Error parsing info JSON: %v", err)
	}

	PluginConfig = PluginConfigType{
		Port:          *port,
		PluginUUID:    *pluginUUID,
		RegisterEvent: *registerEvent,
		Info:          sdInfo,
	}

	u := url.URL{Scheme: "ws", Host: "127.0.0.1:" + PluginConfig.Port, Path: "/"}
	log.Printf("Connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatalf("Error connecting to WebSocket: %v", err)
	}
	defer c.Close()

	WsClient = c

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	registerMessage := map[string]string{
		"event": PluginConfig.RegisterEvent,
		"uuid":  PluginConfig.PluginUUID,
	}

	println(registerMessage)

	if err := c.WriteJSON(registerMessage); err != nil {
		log.Fatalf("Error sending register message: %v", err)
	}

	exampleAction := &ExampleAction{
		BaseAction: BaseAction{UUID: "com.emilyxfox.counter.counter"},
	}
	RegisterAction(exampleAction)

	// Listen for messages from WebSocket
	go func() {
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}

			handleEvent(message)
		}
	}()

	// Keep running until interrupted
	for {
		<-interrupt
		log.Println("Interrupt received, shutting down...")
		c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
		return
	}

}

func SendEventToStreamDeck(response interface{}) error {
	if WsClient == nil {
		return fmt.Errorf("WebSocket client is not initialised")
	}
	log.Println("Sending response:", response)
	return WsClient.WriteJSON(response)
}

func handleEvent(data []byte) {
	event, err := ParseEvent(data)
	if err != nil {
		log.Printf("Error parsing event: %v", err)
		return
	}

	e := event.GetEventType()
	log.Printf("Received event: %s", e)

	DispatchEvent(event)
}

func RegisterAction(action Action) {
	actionRegistry[action.GetUUID()] = action
}
