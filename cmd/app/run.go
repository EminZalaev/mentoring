package app

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"mentoring/internal"
	"os"
	"os/signal"
	"syscall"
)

func Run() error {
	cfg, err := internal.InitConfig()
	if err != nil {
		return err
	}
	app := fiber.New()
	store, err := internal.NewStorage(cfg)
	if err != nil {
		return err
	}
	srv := internal.NewService(store, app)
	srv.InitRoutes()
	if err := app.Listen(":" + cfg.Port); err != nil {
		return err
	}
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)
	go func() {
		internal.UpdateCurrency(store)
	}()

	<-signals
	log.Println("Terminating...")

	_ = srv.Stop()
	log.Println("Terminated!")

	return nil
}
