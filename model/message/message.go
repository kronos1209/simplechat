package message

import (
	"bytes"
	"html/template"
	"simplechat/lib"
	"time"
)

type Message struct {
	ID        string
	User      string
	Contents  string
	CreatedAt time.Time
}

func New(user, contents string) *Message {
	if user == "" || contents == "" {
		return nil
	}
	return &Message{
		ID:        lib.GenerateRandomString(32),
		User:      user,
		Contents:  contents,
		CreatedAt: time.Now(),
	}
}

func (m *Message) ToHtml() string {
	tmpl := template.Must(template.ParseFiles("views/parts/message.html"))
	tmpl.Parse(`<div hx-swap-oob="beforeend:#chat-room">
		{{ template "message" . }}
	</div>`)
	var buf bytes.Buffer
	tmpl.Execute(&buf, m)
	return buf.String()
}
