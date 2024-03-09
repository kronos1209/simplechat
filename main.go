package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"simplechat/model/message"
	"simplechat/model/room"

	"golang.org/x/net/websocket"
)

var (
	r = room.New()
)

func entry(ws *websocket.Conn) {
	r.Add(ws)
	defer r.Del(ws)
	for {
		if _, err := ws.Read(nil); err != nil {
			return
		}
	}
}
func main() {
	http.HandleFunc("/index/", func(w http.ResponseWriter, req *http.Request) {
		msgList := r.GetMessageList()
		tmpl := template.Must(template.ParseFiles("views/index.html", "views/parts/message.html"))
		tmpl.ExecuteTemplate(w, "index.html", msgList)
	})
	http.HandleFunc("/post", func(w http.ResponseWriter, req *http.Request) {
		user := req.PostFormValue("Name")
		m := req.PostFormValue("Message")
		msg := message.New(user, m)
		if msg == nil {
			return
		}
		r.AddMessage(msg)
		if err := r.Broadcast(msg); err != nil {
			fmt.Println(err)
		}
	})
	http.Handle("/entry", websocket.Handler(entry))
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
