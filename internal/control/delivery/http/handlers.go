package http

import (
	"SocialConnections/config"
	"SocialConnections/internal/control"
	"SocialConnections/internal/models"
	"SocialConnections/pkg/reqvalidator"
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
)

type controlHandlers struct {
	cfg       *config.Config
	controlUC control.UseCase
}

func NewControlHandlers(cfg *config.Config, controlUC control.UseCase) control.Handlers {
	return &controlHandlers{cfg: cfg, controlUC: controlUC}
}

func (ctrl *controlHandlers) FetchGraph() fiber.Handler {
	return func(c *fiber.Ctx) error {
		result, err := ctrl.controlUC.FetchConnections(context.Background())
		if err != nil {
			return err
		}
		
		return c.Send(result)
	}
}

func (ctrl *controlHandlers) NewCommunication() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var params models.Connections
		if err := reqvalidator.ReadRequest(c, &params); err != nil {
			log.Println("controlHandlers.NewCommunication.ReadRequest", err)
			return err
		}

		err := ctrl.controlUC.NewCommunication(context.Background(), params)
		if err != nil {
			return err
		}

		return c.SendStatus(fiber.StatusOK)
	}
}