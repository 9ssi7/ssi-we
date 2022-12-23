package internal

import (
	"encoding/json"

	"github.com/gofiber/websocket/v2"
	"github.com/ssibrahimbas/ssi-we/src/dto"
)

func (h *Handler) listenAnyEvent(c *websocket.Conn) {
	var (
		msg []byte
		err error
	)
	ip := c.Locals("IpAddr").(string)
	for {
		if _, msg, err = c.ReadMessage(); err != nil {
			_ = c.WriteMessage(websocket.TextMessage, []byte("An error occurred while reading the message"))
			continue
		}
		dto := &dto.EventCreateRequest{}
		if err := json.Unmarshal(msg, dto); err != nil {
			_ = c.WriteMessage(websocket.TextMessage, []byte("Invalid message format"))
			continue
		}
		dto.IP = ip
		h.s.Create(dto)
	}
}
