package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type Server struct {
	logger *logrus.Logger
}

func NewServer(l *logrus.Logger) *Server {
	return &Server{
		logger: l,
	}
}

func (s *Server) Run() {
	app := fiber.New(
		fiber.Config{
			DisableStartupMessage: true,
		},
	)

	router := app.Group("/api")

	router.Get("/login", s.login)
	router.Get("/logout", s.logout)

	// TODO to config
	port := 8080

	s.logger.Infof("start service port %d", port)

	app.Listen(fmt.Sprintf(":%d", port))
}

func (s *Server) login(ctx *fiber.Ctx) error {
	s.logger.Info("login handle")
	return nil
}

func (s *Server) logout(ctx *fiber.Ctx) error {
	s.logger.Info("logout handle")
	return nil
}
