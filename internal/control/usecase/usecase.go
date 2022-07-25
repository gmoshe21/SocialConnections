package usecase

import (
	"SocialConnections/config"
	"SocialConnections/internal/control"
	"SocialConnections/internal/models"
	"context"
)

type controlUC struct {
	cfg                    *config.Config
	controlRepo            control.Repository
}

func NewControlUseCase( cfg *config.Config, controlRepo control.Repository) control.UseCase {
	return &controlUC{cfg: cfg, controlRepo: controlRepo}
}

func (c *controlUC) NewCommunication(ctx context.Context, params models.Connections) error {
	return c.controlRepo.NewCommunication(ctx, params)
}

func (c *controlUC) FetchConnections(ctx context.Context) (result []byte, err error) {
	return c.controlRepo.FetchConnections(ctx)
}