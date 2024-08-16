package main

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

// SetTitle sends a "setTitle" event to the Stream Deck software to update the title
// displayed on the Stream Deck button.
//
// Parameters:
//   - title: The new title to be displayed on the button.
//   - options: Optional integers where the first value is the target and the second
//     is the state. Both default to 0 if not provided.
//
// Usage:
//
//	e.SetTitle("New Title")                // Default target and state
//	e.SetTitle("New Title", 1)             // Specified target, default state
//	e.SetTitle("New Title", 1, 2)          // Specified target and state
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
