package control

import (
	"github.com/gofiber/fiber/v2"
)

type Handlers interface {
	FetchGraph() fiber.Handler
	NewCommunication() fiber.Handler
}