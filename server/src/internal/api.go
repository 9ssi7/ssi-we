package internal

import (
	"log"

	"github.com/gofiber/websocket/v2"
)

func (h *Handler) listenAnyEvent(c *websocket.Conn) {
	var (
		mt  int
		msg []byte
		err error
	)
	for {
		if mt, msg, err = c.ReadMessage(); err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", msg)
		log.Printf("msg type: %d", mt)
	}
}
