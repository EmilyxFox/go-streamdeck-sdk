package streamdeck

type Action interface {
	GetUUID() string
}

type ActionConfig struct {
	UUID string
}

func (a *ActionConfig) GetUUID() string {
	return a.UUID
}

// StreamDeckInfo holds the parsed JSON data from the -info flag
type StreamDeckInfo struct {
	Application struct {
		Font            string `json:"font"`
		Platform        string `json:"platform"`
		PlatformVersion string `json:"platformVersion"`
		Version         string `json:"version"`
	} `json:"application"`
	Plugin struct {
		UUID    string `json:"uuid"`
		Version string `json:"version"`
	} `json:"plugin"`
	DevicePixelRatio int `json:"devicePixelRatio"`
	Colors           struct {
		ButtonPressedBackgroundColor string `json:"buttonPressedBackgroundColor"`
		ButtonPressedBorderColor     string `json:"buttonPressedBorderColor"`
		ButtonPressedTextColor       string `json:"buttonPressedTextColor"`
		DisabledColor                string `json:"disabledColor"`
		HighlightColor               string `json:"highlightColor"`
		MouseDownColor               string `json:"mouseDownColor"`
	} `json:"colors"`
	Device []struct {
		Id   string `json:"id"`
		Name string `json:"name"`
		Size struct {
			Columns int    `json:"columns"`
			Rows    string `json:"rows"`
		} `json:"device"`
		Type int `json:"type"`
	} `json:"devices"`
}
