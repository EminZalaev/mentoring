package internal

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func getCurrency(s Store) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		writeCors(ctx)
		currency, err := s.GetCurrency()
		if err != nil {
			writeError(ctx, "internal error", http.StatusInternalServerError)
			return err
		}

		body, err := json.Marshal(currency)
		if err != nil {
			writeError(ctx, "internal error", http.StatusInternalServerError)
			return err
		}

		ctx.Response().SetBody(body)
		return nil
	}
}

func postCurrency(s Store) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		writeCors(ctx)
		currency := &currencyRequest{}
		if err := json.Unmarshal(ctx.Request().Body(), &currency); err != nil {
			writeError(ctx, "wrong json", http.StatusBadRequest)
			return err
		}

		err := s.PostCurrency(currency)
		if err != nil {
			writeError(ctx, "internal error", http.StatusInternalServerError)
			return err
		}

		setStatusSuccess(ctx)

		return nil
	}
}

func putCurrency(s Store) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		writeCors(ctx)
		currency := &currencyRequest{}
		if err := json.Unmarshal(ctx.Request().Body(), &currency); err != nil {
			writeError(ctx, "wrong json", http.StatusBadRequest)
			return err
		}

		err := s.PutCurrency(currency)
		if err != nil {
			writeError(ctx, "internal error", http.StatusInternalServerError)
			return err
		}

		setStatusSuccess(ctx)

		return nil
	}
}

func writeCors(ctx *fiber.Ctx) {
	ctx.Response().Header.Set("Access-Control-Allow-Credentials", "*")
	ctx.Response().Header.Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	ctx.Response().Header.SetBytesV("Access-Control-Allow-Origin", []byte("Accept, Content-Type, Content-Length, Accept-Encoding"))
}

func writeError(ctx *fiber.Ctx, msg string, statusCode int) {
	ctx.Request().SetBodyString(msg)
	ctx.Status(statusCode)
}

func setStatusSuccess(ctx *fiber.Ctx) {
	status := status{"success"}
	response, err := json.Marshal(status)
	if err != nil {
		writeError(ctx, "internal error", http.StatusInternalServerError)
		return
	}
	ctx.Response().SetBody(response)
}
