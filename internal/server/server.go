package server

import (
	controlHttp "SocialConnections/internal/control/delivery/http"
	controlRepository "SocialConnections/internal/control/repository"
	controlUseCase "SocialConnections/internal/control/usecase"
	"SocialConnections/config"
	"context"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	fiber *fiber.App
	cfg   *config.Config
	pgDB  *sqlx.DB
}

func NewServer(cfg *config.Config, database *sqlx.DB) *Server {
	return &Server{
		fiber: fiber.New(fiber.Config{}),
		cfg:   cfg,
		pgDB:  database,
	}
}

func (s *Server) MapHandlers(ctx context.Context) {
	controlRepo := controlRepository.NewControlRepository(s.pgDB)
	controlUC := controlUseCase.NewControlUseCase( s.cfg, controlRepo)
	controlHandlers := controlHttp.NewControlHandlers(s.cfg, controlUC)
	controlGroup := s.fiber.Group("control")

	controlHttp.MapAPIRoutes(controlGroup, controlHandlers)

}

func (s *Server) Run(ctx context.Context) {
	s.MapHandlers(ctx)

	if err := s.fiber.Listen(fmt.Sprintf("%s:%s", s.cfg.Server.Host, s.cfg.Server.Port)); err != nil {
		log.Fatalf("Error starting Server: ", err)
	}
}