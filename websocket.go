package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/url"

	"github.com/gorilla/websocket"
)

// PluginConfig holds the configuration passed via flags
type PluginConfigType struct {
	Port          string
	PluginUUID    string
	RegisterEvent string
	Info          StreamDeckInfo
}

var PluginConfig PluginConfigType

var WsClient *websocket.Conn

func OpenWebsocketAndRegisterPlugin() {
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

	registerMessage := map[string]string{
		"event": PluginConfig.RegisterEvent,
		"uuid":  PluginConfig.PluginUUID,
	}

	if err := c.WriteJSON(registerMessage); err != nil {
		log.Fatalf("Error sending register message: %v", err)
	}
}

func SendEventToStreamDeck(response interface{}) error {
	if WsClient == nil {
		return fmt.Errorf("WebSocket client is not initialised")
	}
	log.Println("Sending response:", response)
	return WsClient.WriteJSON(response)
}
