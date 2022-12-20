package app

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"mentoring/internal"

	"github.com/gofiber/fiber/v2"
)

func Run() (err error) {
	cfg, err := internal.InitConfig()
	if err != nil {
		return fmt.Errorf("error init config: %w", err)
	}

	app := fiber.New()

	store, err := internal.NewStorage(cfg)
	if err != nil {
		return fmt.Errorf("error storage: %w", err)
	}
	defer func() {
		if err := store.CloseDBConnection(); err != nil {
			log.Println(err)
		}
	}()

	srv := internal.NewService(store, app)
	log.Println("server start on port:", cfg.Port)

	go func() {
		ticker := time.NewTicker(1 * time.Hour)
		for range ticker.C {
			if err := internal.UpdateCurrency(cfg.CurrencyApiKey, store); err != nil {
				log.Println(err)
			}
		}
	}()

	srv.InitRoutes()
	go func() {
		if err := app.Listen(":" + cfg.Port); err != nil {
			log.Fatal(err)
		}
	}()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)

	<-signals
	log.Println("Terminating...")

	_ = srv.Stop()
	log.Println("Terminated!")

	return err
}
