package control

import (
	"SocialConnections/internal/models"
	"context"
)

type UseCase interface {
	NewCommunication(ctx context.Context, params models.Connections) error
	FetchConnections(ctx context.Context) (result []byte, err error) 

}