package streamdeck

import (
	"encoding/json"
	"fmt"
)

type eventParser func(data []byte) (StreamDeckEvent, error)

var eventParsers = map[string]eventParser{
	"didReceiveSettings":            parseDidReceiveSettings,
	"didReceiveGlobalSettings":      parseDidReceiveGlobalSettings,
	"didReceiveDeepLink":            parseDidReceiveDeepLink,
	"touchTap":                      parseTouchTap,
	"dialDown":                      parseDialDown,
	"dialUp":                        parseDialUp,
	"dialRotate":                    parseDialRotate,
	"keyDown":                       parseKeyDown,
	"keyUp":                         parseKeyUp,
	"willAppear":                    parseWillAppear,
	"willDisappear":                 parseWillDisappear,
	"titleParametersDidChange":      parseTitleParametersDidChange,
	"deviceDidConnect":              parseDeviceDidConnect,
	"deviceDidDisconnect":           parseDeviceDidDisconnect,
	"applicationDidLaunch":          parseApplicationDidLaunch,
	"applicationDidTerminate":       parseApplicationDidTerminate,
	"systemDidWakeUp":               parseSystemDidWakeUp,
	"propertyInspectorDidAppear":    parsePropertyInspectorDidAppear,
	"propertyInspectorDidDisappear": parsePropertyInspectorDidDisappear,
	"sendToPlugin":                  parseSendToPlugin,
	"sendToPropertyInspector":       parseSendToPropertyInspector,
}

func parseDidReceiveSettings(data []byte) (StreamDeckEvent, error) {
	var event DidReceiveSettingsEvent
	err := json.Unmarshal(data, &event)
	return &event, err
}
func parseDidReceiveGlobalSettings(data []byte) (StreamDeckEvent, error) {
	var event DidReceiveGlobalSettingsEvent
	err := json.Unmarshal(data, &event)
	return &event, err
}
func parseDidReceiveDeepLink(data []byte) (StreamDeckEvent, error) {
	var event DidReceiveDeepLinkEvent
	err := json.Unmarshal(data, &event)
	return &event, err
}
func parseTouchTap(data []byte) (StreamDeckEvent, error) {
	var event TouchTapEvent
	err := json.Unmarshal(data, &event)
	return &event, err
}
func parseDialDown(data []byte) (StreamDeckEvent, error) {
	var event DialDownEvent
	err := json.Unmarshal(data, &event)
	return &event, err
}
func parseDialUp(data []byte) (StreamDeckEvent, error) {
	var event DialUpEvent
	err := json.Unmarshal(data, &event)
	return &event, err
}
func parseDialRotate(data []byte) (StreamDeckEvent, error) {
	var event DialRotateEvent
	err := json.Unmarshal(data, &event)
	return &event, err
}
func parseKeyDown(data []byte) (StreamDeckEvent, error) {
	var event KeyDownEvent
	err := json.Unmarshal(data, &event)
	return &event, err
}
func parseKeyUp(data []byte) (StreamDeckEvent, error) {
	var event KeyUpEvent
	err := json.Unmarshal(data, &event)
	return &event, err
}
func parseWillAppear(data []byte) (StreamDeckEvent, error) {
	var event WillAppearEvent
	err := json.Unmarshal(data, &event)
	return &event, err
}
func parseWillDisappear(data []byte) (StreamDeckEvent, error) {
	var event WillDisappearEvent
	err := json.Unmarshal(data, &event)
	return &event, err
}
func parseTitleParametersDidChange(data []byte) (StreamDeckEvent, error) {
	var event TitleParametersDidChangeEvent
	err := json.Unmarshal(data, &event)
	return &event, err
}
func parseDeviceDidConnect(data []byte) (StreamDeckEvent, error) {
	var event DeviceDidConnectEvent
	err := json.Unmarshal(data, &event)
	return &event, err
}
func parseDeviceDidDisconnect(data []byte) (StreamDeckEvent, error) {
	var event DeviceDidDisconnectEvent
	err := json.Unmarshal(data, &event)
	return &event, err
}
func parseApplicationDidLaunch(data []byte) (StreamDeckEvent, error) {
	var event ApplicationDidLaunchEvent
	err := json.Unmarshal(data, &event)
	return &event, err
}
func parseApplicationDidTerminate(data []byte) (StreamDeckEvent, error) {
	var event ApplicationDidTerminateEvent
	err := json.Unmarshal(data, &event)
	return &event, err
}
func parseSystemDidWakeUp(data []byte) (StreamDeckEvent, error) {
	var event SystemDidWakeUpEvent
	err := json.Unmarshal(data, &event)
	return &event, err
}
func parsePropertyInspectorDidAppear(data []byte) (StreamDeckEvent, error) {
	var event PropertyInspectorDidAppearEvent
	err := json.Unmarshal(data, &event)
	return &event, err
}
func parsePropertyInspectorDidDisappear(data []byte) (StreamDeckEvent, error) {
	var event PropertyInspectorDidDisappearEvent
	err := json.Unmarshal(data, &event)
	return &event, err
}
func parseSendToPlugin(data []byte) (StreamDeckEvent, error) {
	var event SendToPluginEvent
	err := json.Unmarshal(data, &event)
	return &event, err
}
func parseSendToPropertyInspector(data []byte) (StreamDeckEvent, error) {
	var event SendToPropertyInspectorEvent
	err := json.Unmarshal(data, &event)
	return &event, err
}

func ParseEvent(data []byte) (StreamDeckEvent, error) {
	var temp struct {
		Event string `json:"event"`
	}

	if err := json.Unmarshal(data, &temp); err != nil {
		return nil, fmt.Errorf("error unmarshaling event type: %w", err)
	}

	// Look up the parser function for the event type
	parser, exists := eventParsers[temp.Event]
	if !exists {
		return nil, fmt.Errorf("unknown event type: %s", temp.Event)
	}

	return parser(data)
}
