package streamdeck

type StreamDeckEvent interface {
	GetEventType() string
	IsActionAssociated() bool
}

type ActionAssociatedEvent struct {
	Event   string `json:"event"`
	Action  string `json:"action"`
	Context string `json:"context"`
	Device  string `json:"device,omitempty"`
}

type GlobalEvent struct {
	Event  string `json:"event"`
	Device string `json:"device,omitempty"`
}

func (e *ActionAssociatedEvent) GetEventType() string {
	return e.Event
}

func (e *GlobalEvent) GetEventType() string {
	return e.Event
}

func (e *ActionAssociatedEvent) IsActionAssociated() bool {
	return true
}

func (e *GlobalEvent) IsActionAssociated() bool {
	return false
}

func (e *ActionAssociatedEvent) GetAction() (string, bool) {
	if e.Action != "" {
		return e.Action, true
	} else {
		return "", false
	}
}

// Update settings associated with action
//
// Parameters:
// - settings: A map[string]any which is persistently saved as a json for the action's instance.
//
// Usage:
//
//	newSettings := map[string]any{
//		"apikey": "mX8ulcBHYmMniSshmB59"
//	}
//	e.SetSettings(newSettings)
//
// Docs:
// https://docs.elgato.com/sdk/plugins/events-sent#setsettings
func (e *ActionAssociatedEvent) SetSettings(settings map[string]any) error {
	response := map[string]interface{}{
		"event":   "setSettings",
		"context": e.Context,
		"payload": settings,
	}
	return SendEventToStreamDeck(response)
}

// Request the persistent data stored for the action's instance:
//
// Usage:
//
//	e.GetSettings()
//
// Docs:
// https://docs.elgato.com/sdk/plugins/events-sent#getsettings
func (e *ActionAssociatedEvent) GetSettings() error {
	response := map[string]string{
		"event":   "getSettings",
		"context": e.Context,
	}
	return SendEventToStreamDeck(response)
}

// The plugin and Property Inspector can save persistent data globally. The data will be saved securely
// to the Keychain on macOS and the Credential Store on Windows. This API can be used to save tokens
// that should be available to every action in the plugin.
//
// Parameters:
// - settings: A map[string]any which is persistently saved  persistently saved globally.
//
// Usage:
//
//	newSettings := map[string]any{
//		"apikey": "mX8ulcBHYmMniSshmB59"
//	}
//	e.setGlobalSettings(newSettings)
//
// Docs:
// https://docs.elgato.com/sdk/plugins/events-sent#setglobalsettings
func (e *ActionAssociatedEvent) SetGlobalSettings(settings map[string]any) error {
	response := map[string]interface{}{
		"event":   "setGlobalSettings",
		"context": PluginConfig.PluginUUID,
		"payload": settings,
	}
	return SendEventToStreamDeck(response)
}

// Request the persistent global data:
//
// Usage:
//
//	e.GetGlobalSettings()
//
// Docs: https://docs.elgato.com/sdk/plugins/events-sent#getglobalsettings
func (e *ActionAssociatedEvent) GetGlobalSettings() error {
	response := map[string]string{
		"event":   "getGlobalSettings",
		"context": PluginConfig.PluginUUID,
	}
	return SendEventToStreamDeck(response)
}

// Tell the Stream Deck application to open an URL in the default browser:
//
// Usage:
//
//	e.OpenUrl("https://www.example.com")
//
// Docs:
// https://docs.elgato.com/sdk/plugins/events-sent#openurl
func (e *ActionAssociatedEvent) OpenUrl(url string) error {
	response := map[string]interface{}{
		"event": "openUrl",
		"payload": map[string]interface{}{
			"url": url,
		},
	}
	return SendEventToStreamDeck(response)
}

// Write a debug message to the logs file:
//
// Usage:
//
//	e.LogMessage("Button was clicked!")
//
// Docs:
// https://docs.elgato.com/sdk/plugins/events-sent#logmessage
func (e *ActionAssociatedEvent) LogMessage(message string) error {
	response := map[string]interface{}{
		"event": "logMessage",
		"payload": map[string]interface{}{
			"message": message,
		},
	}
	return SendEventToStreamDeck(response)
}

// Update the title displayed on the Stream Deck button.
//
// Parameters:
//   - title:   The new title to be displayed on the button.
//   - options: Optional integers where the first value is the target and the second
//     is the state. Both default to 0 if not provided.
//
// Usage:
//
//	e.SetTitle("New Title")                // Default target and state
//	e.SetTitle("New Title", 1)             // Specified target, default state
//	e.SetTitle("New Title", 1, 2)          // Specified target and state
//
// Docs:
// https://docs.elgato.com/sdk/plugins/events-sent#settitle
func (e *ActionAssociatedEvent) SetTitle(title string, options ...uint8) error {
	var target, state uint8 = 0, 0

	if len(options) > 0 {
		target = options[0]
	}
	if len(options) > 1 {
		state = options[1]
	}

	response := map[string]interface{}{
		"event":   "setTitle",
		"context": e.Context,
		"payload": map[string]interface{}{
			"title":  title,
			"target": target,
			"state":  state,
		},
	}
	return SendEventToStreamDeck(response)
}

// Change the image displayed by an instance of an action:
//
// Usage:
//
//	e.SetImage("Button was clicked!")
//
// Docs:
// https://docs.elgato.com/sdk/plugins/events-sent#setimage
func (e *ActionAssociatedEvent) SetImage(base64image string, options ...uint8) error {
	var target, state uint8 = 0, 0

	if len(options) > 0 {
		target = options[0]
	}
	if len(options) > 1 {
		state = options[1]
	}

	response := map[string]interface{}{
		"event":   "setImage",
		"context": e.Context,
		"payload": map[string]interface{}{
			"image":  base64image,
			"target": target,
			"state":  state,
		},
	}
	return SendEventToStreamDeck(response)
}

// !! SetFeedback

// !! SetFeedbackLayout

// !! SetTriggerDescription

// Temporarily show an alert icon on the image displayed by an instance of an action:
//
// Usage:
//
//	e.ShowAlert()
//
// Docs:
// https://docs.elgato.com/sdk/plugins/events-sent#showalert
func (e *ActionAssociatedEvent) ShowAlert() error {
	response := map[string]interface{}{
		"event":   "showAlert",
		"context": e.Context,
	}
	return SendEventToStreamDeck(response)
}

// Temporarily show an OK checkmark icon on the image displayed by an instance of an action:
//
// Usage:
//
//	e.ShowOk()
//
// Docs:
// https://docs.elgato.com/sdk/plugins/events-sent#showok
func (e *ActionAssociatedEvent) ShowOk() error {
	response := map[string]interface{}{
		"event":   "showOk",
		"context": e.Context,
	}
	return SendEventToStreamDeck(response)
}

// Change the state of an action supporting multiple states:
//
// Usage:
//
//	e.SetState(1)
//
// Docs:
// https://docs.elgato.com/sdk/plugins/events-sent#setstate
func (e *ActionAssociatedEvent) SetState(state uint8) error {
	response := map[string]interface{}{
		"event":   "setState",
		"context": e.Context,
		"payload": map[string]interface{}{
			"state": state,
		},
	}
	return SendEventToStreamDeck(response)
}

// Switch to a preconfigured read-only profile:
//
// Usage:
//
//	e.SwitchToProfile("MyPluginProfile")
//	e.SwitchToProfile("MyPluginProfile", 1)
//
// Docs:
// https://docs.elgato.com/sdk/plugins/events-sent#switchtoprofile
// func (e *ActionAssociatedEvent) SwitchToProfile(profile string, page ...uint8) error {
// 	var pageIndex uint8 = 0

// 	if len(page) > 0 {
// 		pageIndex = page[0]
// 	}
// 	response := map[string]interface{}{
// 		"event":   "switchToProfile",
// 		"context": e.Context, //PLUGINUUID
// 		"device":  e.Device,
// 		"payload": map[string]interface{}{
// 			"profile": profile,
// 			"page":    pageIndex,
// 		},
// 	}
// 	return SendEventToStreamDeck(response)
// }

// Send a payload to the Property Inspector:
//
// Usage:
//
//	e.SetState(map[string]any{}{
//		"arbitrary": "value",
//	})
//
// Docs:
// https://docs.elgato.com/sdk/plugins/events-sent#sendtopropertyinspector
func (e *ActionAssociatedEvent) SendToPropertyInspector(payload map[string]any) error {
	response := map[string]interface{}{
		"event":   "setState",
		"action":  e.Action,
		"context": e.Context,
		"payload": payload,
	}
	return SendEventToStreamDeck(response)
}

type ActionCoordinates struct {
	Column int `json:"column"`
	Row    int `json:"row"`
}

type ActionSettings map[string]any

type GlobalSettings map[string]any

type TapPosition [2]int

type DidReceiveSettingsEvent struct {
	ActionAssociatedEvent
	Payload struct {
		Settings        ActionSettings    `json:"settings"`
		Coordinates     ActionCoordinates `json:"coordinates"`
		IsInMultiAction bool              `json:"isInMultiAction"`
	} `json:"payload"`
}

type DidReceiveGlobalSettingsEvent struct {
	GlobalEvent
	Payload struct {
		Settings GlobalSettings `json:"settings"`
	} `json:"payload"`
}

type DidReceiveDeepLinkEvent struct {
	GlobalEvent
	Payload struct {
		Url string `json:"url"`
	} `json:"payload"`
}

type TouchTapEvent struct {
	ActionAssociatedEvent
	Payload struct {
		Settings    ActionSettings    `json:"settings"`
		Controller  string            `json:"controller"`
		Coordinates ActionCoordinates `json:"coordinates"`
		TapPos      TapPosition       `json:"tapPos"`
		Hold        bool              `json:"hold"`
	} `json:"payload"`
}

type DialDownEvent struct {
	ActionAssociatedEvent
	Payload struct {
		Controller  string            `json:"controller"`
		Settings    ActionSettings    `json:"settings"`
		Coordinates ActionCoordinates `json:"coordinates"`
	} `json:"payload"`
}

type DialUpEvent struct {
	ActionAssociatedEvent
	Payload struct {
		Controller  string            `json:"controller"`
		Settings    ActionSettings    `json:"settings"`
		Coordinates ActionCoordinates `json:"coordinates"`
	} `json:"payload"`
}

type DialRotateEvent struct {
	ActionAssociatedEvent
	Payload struct {
		Settings    ActionSettings    `json:"settings"`
		Controller  string            `json:"controller"`
		Coordinates ActionCoordinates `json:"coordinates"`
		Ticks       int               `json:"ticks"`
		Pressed     bool              `json:"pressed"`
	} `json:"payload"`
}

type KeyDownEvent struct {
	ActionAssociatedEvent
	Payload struct {
		Settings         ActionSettings    `json:"settings"`
		Coordinates      ActionCoordinates `json:"coordinates"`
		State            int               `json:"states"`
		UserDesiredState int               `json:"userDesiredState"`
		IsInMultiAction  bool              `json:"isInMultiAction"`
	} `json:"payload"`
}

type KeyUpEvent struct {
	ActionAssociatedEvent
	Payload struct {
		Settings         ActionSettings    `json:"settings"`
		Coordinates      ActionCoordinates `json:"coordinates"`
		State            int               `json:"state"`
		UserDesiredState int               `json:"userDesiredState"`
		IsInMultiAction  bool              `json:"isInMultiAction"`
	} `json:"payload"`
}

type WillAppearEvent struct {
	ActionAssociatedEvent
	Payload struct {
		Settings        ActionSettings    `json:"settings"`
		Coordinates     ActionCoordinates `json:"coordinates"`
		Controller      string            `json:"controller"`
		State           int               `json:"state"`
		IsInMultiAction bool              `json:"isInMultiAction"`
	} `json:"payload"`
}

type WillDisappearEvent struct {
	ActionAssociatedEvent
	Payload struct {
		Settings        ActionSettings    `json:"settings"`
		Coordinates     ActionCoordinates `json:"coordinates"`
		Controller      string            `json:"controller"`
		State           int               `json:"state"`
		IsInMultiAction bool              `json:"isInMultiAction"`
	} `json:"payload"`
}

type TitleParametersDidChangeEvent struct {
	ActionAssociatedEvent
	Payload struct {
		Coordinates     ActionCoordinates `json:"coordinates"`
		Settings        ActionSettings    `json:"settings"`
		State           int               `json:"state"`
		Title           string            `json:"title"`
		TitleParameters struct {
			FontFamily     string `json:"fontFamily"`
			FontSize       int    `json:"fontSize"`
			FontStyle      string `json:"fontStyle"`
			FontUnderline  bool   `json:"fontUnderline"`
			ShowTitle      bool   `json:"showTitle"`
			TitleAlignment string `json:"titleAlignment"`
			TitleColor     string `json:"titleColor"`
		} `json:"titleParameters"`
	} `json:"payload"`
}

type DeviceDidConnectEvent struct {
	GlobalEvent
	DeviceInfo struct {
		Name string `json:"name"`
		Type int    `json:"type"`
		Size struct {
			Columns int `json:"columns"`
			Rows    int `json:"rows"`
		} `json:"size"`
	} `json:"deviceInfo"`
}

type DeviceDidDisconnectEvent struct {
	GlobalEvent
}

type ApplicationDidLaunchEvent struct {
	GlobalEvent
	Payload struct {
		Application string `json:"application"`
	} `json:"payload"`
}

type ApplicationDidTerminateEvent struct {
	GlobalEvent
	Payload struct {
		Application string `json:"application"`
	} `json:"payload"`
}

type SystemDidWakeUpEvent struct {
	GlobalEvent
}

type PropertyInspectorDidAppearEvent struct {
	ActionAssociatedEvent
}

type PropertyInspectorDidDisappearEvent struct {
	ActionAssociatedEvent
}

type SendToPluginEvent struct {
	ActionAssociatedEvent
	Payload map[string]any `json:"payload"`
}

type SendToPropertyInspectorEvent struct {
	ActionAssociatedEvent
	Payload map[string]any `json:"payload"`
}
