package service

import (
	"github.com/gofiber/fiber/v2"
	"mentoring/internal/storage"
)

type Service struct {
	Store *storage.Storage
	app   *fiber.App
}

func NewService(store *storage.Storage, app *fiber.App) *Service {
	return &Service{
		Store: store,
		app:   app,
	}
}

func (s *Service) GetCurrency() {
	s.app.Get("/")
}
