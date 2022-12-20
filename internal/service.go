package internal

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type Service struct {
	Store    Store
	app      *fiber.App
	stopChan chan struct{}
}

func NewService(store Store, app *fiber.App) *Service {
	return &Service{
		stopChan: make(chan struct{}),
		Store:    store,
		app:      app,
	}
}

func (s *Service) Stop() error {
	if httpSrv := s.app.Server(); httpSrv != nil {
		err := httpSrv.Shutdown()
		if err != nil {
			log.Println("Terminate error: ", err.Error())
			return err
		}

		close(s.stopChan)

		s.app = nil
	}

	return nil
}

func (s *Service) InitRoutes() {
	group := s.app.Group("/api/currency")
	group.Get("/", getCurrency(s.Store))
	group.Post("/", postCurrency(s.Store))
	group.Put("/", putCurrency(s.Store))
}
