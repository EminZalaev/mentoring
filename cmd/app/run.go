package app

import (
	"github.com/gofiber/fiber/v2"
	"mentoring/internal/config"
	"mentoring/internal/service"
	"mentoring/internal/storage"
)

func Run() error {
	cfg, err := config.InitConfig()
	if err != nil {
		return err
	}
	app := fiber.New()
	store, err := storage.NewStorage(cfg)
	if err != nil {
		return err
	}
	service.NewService(store, app)
}
