package handler

import (
	"net/http"

	"github.com/ankorstore/yokai/config"
	"github.com/labstack/echo/v4"
)

type TemplateHandler struct {
	config *config.Config
}

func NewTemplateHandler(cfg *config.Config) *TemplateHandler {
	return &TemplateHandler{
		config: cfg,
	}
}

func (h *TemplateHandler) Handle() echo.HandlerFunc {
	return func(c echo.Context) error {
		// will render: "<html><body><h1>App name is app</h1></body></html>"
		return c.Render(http.StatusOK, "app.htm", map[string]interface{}{
			"name": h.config.AppName(),
		})
	}
}
