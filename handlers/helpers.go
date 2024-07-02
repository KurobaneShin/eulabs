package handlers

import (
	"log/slog"
	"net/http"

	echo "github.com/labstack/echo/v4"
)

type APIFunc func(c echo.Context) error

func Make(h APIFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		slog.Info("req", "method", c.Request().Method, "endpoint", c.Request().URL.String())
		if err := h(c); err != nil {
			slog.Error("unhandled HTTP API error", "err", err.Error())
			return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		}
		return nil
	}
}
