package main

import (
	"fmt"
	"log"

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

// func OpenWebsocketAndRegisterPlugin() {

// }

func SendEventToStreamDeck(response interface{}) error {
	if WsClient == nil {
		return fmt.Errorf("WebSocket client is not initialised")
	}
	log.Println("Sending response:", response)
	return WsClient.WriteJSON(response)
}
