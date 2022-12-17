package internal

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
)

func getCurrency(s Store) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		currency, err := s.GetCurrency()
		if err != nil {
			return err
		}
		body, err := json.Marshal(currency)
		if err != nil {
			return err
		}
		ctx.Response().SetBody(body)
		return nil
	}
}

func postCurrency(s Store) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		currency := &currency{}
		if err := json.Unmarshal(ctx.Request().Body(), &currency); err != nil {
			return err
		}
		err := s.PostCurrency(currency)
		if err != nil {
			return err
		}
		return nil
	}
}

func putCurrency(s Store) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		currency := &currency{}
		if err := json.Unmarshal(ctx.Request().Body(), &currency); err != nil {
			return err
		}
		err := s.PutCurrency(currency)
		if err != nil {
			return err
		}
		return nil
	}
}
