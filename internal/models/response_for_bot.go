package models

type Response struct {
	ID           uint        `json:"id"`
	ChatID       int         `json:"chat_id"`
	Text         string      `json:"text"`
	Atachments   []Atachment `json:"attachments"`
	KeyboardFlag bool        `json:"keyboard_flag"`
	Keyboard     Keyboard    `json:"keyboard"`
}

type Atachment struct {
	Type string `json:"type"`
	Url  string `json:"url"`
}

type Keyboard struct {
	Type    string   `json:"type"`
	Buttons []Button `json:"buttons"`
}

type Button struct {
	Caption string `json:"caption"`
	Data    string `json:"data"`
	Row     uint8  `json:"row"`
	Order   uint8  `json:"order"`
}

type ResponseState struct {
	ID           uint64
	ChatID       int64
	Text         string
	Atachments   []Atachment
	KeyboardFlag bool
	Keyboard     Keyboard
	Delay        uint8
	Conditions   []Condition
	UpdateInfo   map[string]string
	Trigger      Trigger
}

type Trigger struct {
	Type string
}

type Condition struct {
	Variables map[string]string
	State     UserState
}
