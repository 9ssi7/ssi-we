package internal

import (
	"encoding/json"
	"log"

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
			log.Println("read:", err)
			break
		}
		dto := &dto.EventCreateRequest{}
		if err := json.Unmarshal(msg, dto); err != nil {
			log.Println("json unmarshal:", err)
			break
		}
		dto.IP = ip
		h.s.Create(dto)
	}
}
