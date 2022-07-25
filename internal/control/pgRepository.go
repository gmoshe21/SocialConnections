package control

import (
	"SocialConnections/internal/models"
	"context"
)

type Repository interface {
	NewCommunication(ctx context.Context, params models.Connections) error
	FetchConnections(ctx context.Context) (result []byte, err error)
}