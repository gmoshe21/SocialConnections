package repository

import (
	"SocialConnections/internal/control"
	"SocialConnections/internal/models"
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type controlRepo struct {
	db     *sqlx.DB
}

func NewControlRepository(db *sqlx.DB) control.Repository {
	return &controlRepo{db: db}
}

func (c *controlRepo) NewCommunication(ctx context.Context, params models.Connections) error {
	_, err := c.db.ExecContext(ctx, params.Sender, params.Recipient)
	if err != nil {
		errors.Wrapf(err, "controlRepo.NewCommunication.ExecContext(Sender : %s, Recipient: %s)",
			params.Sender,
			params.Recipient,
		)
	}
	return nil
}

func (c *controlRepo) FetchConnections(ctx context.Context) (result []byte, err error) {
	err = c.db.GetContext(ctx, &result, queryFetchConnections)
	if err != nil {
		return nil, errors.Wrap(err, "controlRepo.FetchConnections.GetContext()")
	}
	return result, nil
}

