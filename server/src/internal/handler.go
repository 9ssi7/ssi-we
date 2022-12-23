package internal

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/websocket/v2"
	"github.com/ssibrahimbas/ssi-we/pkg/http"
	"github.com/ssibrahimbas/ssi-we/src/config"
)

type Handler struct {
	s *Srv
	c *config.App
	h *http.Client
}

type HandlerParams struct {
	Srv  *Srv
	Http *http.Client
	Cnf  *config.App
}

func NewHandler(p *HandlerParams) *Handler {
	return &Handler{
		s: p.Srv,
		h: p.Http,
		c: p.Cnf,
	}
}

func (h *Handler) Init() {
	gr := h.h.App.Group("/ws")
	gr.Use(h.cors())
	gr.Use(h.wsHandshake())
	gr.Get("/", websocket.New(h.listenAnyEvent))
}

func (h *Handler) cors() fiber.Handler {
	return cors.New(cors.Config{
		AllowOrigins:     h.c.CorsAllowOrigins,
		AllowMethods:     h.c.CorsAllowMethods,
		AllowHeaders:     h.c.CorsAllowHeaders,
		AllowCredentials: h.c.CorsAllowCredentials,
	})
}

func (h *Handler) wsHandshake() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			c.Locals("IpAddr", c.IP())
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	}
}
