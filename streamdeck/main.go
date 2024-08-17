package streamdeck

import (
	"encoding/json"
	"flag"
	"log"
	"net/url"
	"os"
	"os/signal"

	"github.com/gorilla/websocket"
)

var actionRegistry = make(map[string]Action)

func StartPlugin() {
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

	log.Printf("%+v", PluginConfig)

	// OpenWebsocketAndRegisterPlugin()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	// RegisterAction(CounterAction)

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
