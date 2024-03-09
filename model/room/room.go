package room

import (
	"errors"
	"fmt"
	"simplechat/model/message"
	"sync"

	"golang.org/x/net/websocket"
)

type Room struct {
	subscribers map[*websocket.Conn]bool
	MessageList []*message.Message
	sync.RWMutex
}

func New() *Room {
	return &Room{
		subscribers: make(map[*websocket.Conn]bool),
	}
}

func (r *Room) Add(ws *websocket.Conn) {
	r.subscribers[ws] = true
}

func (r *Room) Del(ws *websocket.Conn) {
	delete(r.subscribers, ws)
}

func (r *Room) AddMessage(msg *message.Message) {
	r.Lock()
	r.MessageList = append(r.MessageList, msg)
	r.Unlock()
}

func (r *Room) GetMessageList() []*message.Message {
	r.RLock()
	defer r.RUnlock()
	return r.MessageList
}

func (r *Room) Broadcast(msg *message.Message) error {
	fmt.Println(r)
	var errs []error
	for ws := range r.subscribers {
		if err := websocket.Message.Send(ws, msg.ToHtml()); err != nil {
			errs = append(errs, err)
		}
	}
	if len(errs) != 0 {
		res := fmt.Errorf("on broadcast : there are %v errors", len(errs))
		return errors.Join(res, errors.Join(errs...))
	}
	return nil
}
