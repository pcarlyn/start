package models

type Response struct {
	Ok     bool     `json:"ok"`
	Result []Result `json:"result"`
}

type Result struct {
	UpdateID int     `json:"update_id"`
	Message  Message `json:"message"`
}

type Message struct {
	MessageID int      `json:"message_id"`
	From      User     `json:"from"`
	Chat      Chat     `json:"chat"`
	Date      int64    `json:"date"`
	Text      string   `json:"text"`
	Entities  []Entity `json:"entities"`
}

type User struct {
	ID           int    `json:"id"`
	IsBot        bool   `json:"is_bot"`
	FirstName    string `json:"first_name"`
	Username     string `json:"username"`
	LanguageCode string `json:"language_code"`
	IsPremium    bool   `json:"is_premium"`
}

type Chat struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	Username  string `json:"username"`
	Type      string `json:"type"`
}

type Entity struct {
	Offset int    `json:"offset"`
	Length int    `json:"length"`
	Type   string `json:"type"`
}

type Auth struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
