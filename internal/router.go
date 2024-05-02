package internal

import (
	"github.com/ankorstore/yokai/fxhttpserver"
	"github.com/greeniet87/test_yokai/internal/handler"
	"go.uber.org/fx"
)

// Router is used to register the application HTTP middlewares and handlers.
func Router() fx.Option {
	return fx.Options(
		fxhttpserver.AsHandler("GET", "", handler.NewTemplateHandler),
		fxhttpserver.AsHandler("POST", "upload", handler.NewUploadHandler),
		fxhttpserver.AsHandler("GET", "customerid", handler.NewExampleHandler),
	)
}
