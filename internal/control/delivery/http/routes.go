package http

import(
	"SocialConnections/internal/control"

	"github.com/gofiber/fiber/v2"
)

func MapAPIRoutes(group fiber.Router, h control.Handlers) {
	group.Get("/fetch_graph", h.FetchGraph())
	group.Post("/new_communication", h.NewCommunication())
}