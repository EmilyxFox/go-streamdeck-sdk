package streamdeck

type SetSettingsCommand struct {
	Event   string         `json:"event"`
	Context string         `json:"context"`
	Payload map[string]any `json:"payload"`
}

type GetSettingsCommand struct {
	Event   string `json:"event"`
	Context string `json:"context"`
}

type SetGlobalSettingsCommand struct {
	Event   string         `json:"event"`
	Context string         `json:"context"`
	Payload map[string]any `json:"payload"`
}

type GetGlobalSettingsCommand struct {
	Event   string `json:"event"`
	Context string `json:"context"`
}

type OpenUrlCommand struct {
	Event   string `json:"event"`
	Payload struct {
		Url string `json:"url"`
	} `json:"payload"`
}

type LogMessageCommand struct {
	Event   string `json:"event"`
	Payload struct {
		Message string `json:"message"`
	} `json:"payload"`
}

type SetTitleCommand struct {
	Event   string `json:"event"`
	Context string `json:"context"`
	Payload struct {
		Title  string `json:"title"`
		Target uint8  `json:"target"`
		State  uint8  `json:"state"`
	} `json:"payload"`
}

type SetImageCommand struct {
	Event   string `json:"event"`
	Context string `json:"context"`
	Payload struct {
		Image  string `json:"image"` //this type might be wrong
		Target uint8  `json:"target"`
		State  uint8  `json:"state"`
	} `json:"payload"`
}

type SetFeedbackEvent struct {
	Event   string         `json:"event"`
	Context string         `json:"context"`
	Payload map[string]any `json:"payload"`
} //not really sure about this one

type SetFeedbackLayoutEvent struct {
	Event   string            `json:"event"`
	Context string            `json:"context"`
	Payload map[string]string `json:"payload"`
} //not really sure about this one

type SetTriggerDescriptionEvent struct {
	Event   string `json:"event"`
	Context string `json:"context"`
	Payload struct {
		Rotate    string `json:"rotate"`
		Push      string `json:"push"`
		Touch     string `json:"touch"`
		LongTouch string `json:"longTouch"`
	} `json:"payload"`
}

type ShowAlertCommand struct {
	Event   string `json:"event"`
	Context string `json:"context"`
}

type ShowOkCommand struct {
	Event   string `json:"event"`
	Context string `json:"context"`
}

type SetStateCommand struct {
	Event   string `json:"event"`
	Context string `json:"context"`
	Payload struct {
		State uint8 `json:"state"`
	} `json:"payload"`
}

type SwitchToProfileCommand struct {
	Event   string `json:"event"`
	Context string `json:"context"`
	Device  string `json:"device"`
	Payload struct {
		Profile string `json:"profile"`
		Page    uint8  `json:"page"` //Maybe its possible to have over 256 pages????
	} `json:"payload"`
}

// Both of these events might not even be necessary
type SendToPropertyInspectorFromPluginCommand struct { //Idk if this rename is a good idea
	Action  string         `json:"action"`
	Event   string         `json:"event"`
	Context string         `json:"context"`
	Payload map[string]any `json:"payload"`
}

type SendToPluginFromPluginEvent struct { //Idk if this rename is a good idea
	Action  string         `json:"action"`
	Event   string         `json:"event"`
	Context string         `json:"context"`
	Payload map[string]any `json:"payload"`
}
