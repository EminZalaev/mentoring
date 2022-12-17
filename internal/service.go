package internal

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

type Service struct {
	Store    Store
	app      *fiber.App
	stopChan chan struct{}
}

func NewService(store Store, app *fiber.App) *Service {
	return &Service{
		Store: store,
		app:   app,
	}
}

func (s *Service) Stop() error {
	if httpSrv := s.app.Server(); httpSrv != nil {
		err := httpSrv.Shutdown()
		if err != nil {
			log.Print("Terminate error: ", err.Error())
			return err
		}
		close(s.stopChan)
		s.app = nil
	}
	return nil
}

func (s *Service) InitRoutes() {
	s.app.Group("/api/currency")
	s.app.Get("/", getCurrency(s.Store))
	s.app.Post("/", postCurrency(s.Store))
	s.app.Put("/", putCurrency(s.Store))
}
