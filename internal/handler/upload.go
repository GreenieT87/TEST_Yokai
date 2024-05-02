package handler

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/ankorstore/yokai/config"
	"github.com/labstack/echo/v4"
)

// UploadHandler is an example HTTP handler.
type UploadHandler struct {
	config *config.Config
}

// NewUploadHandler returns a new [UploadHandler].
func NewUploadHandler(config *config.Config) *UploadHandler {
	return &UploadHandler{
		config: config,
	}
}

// Handle handles HTTP requests.
func (h *UploadHandler) Handle() echo.HandlerFunc {
	return func(c echo.Context) error {
		// c.Request().ParseMultipartForm(10 << 20)
		file, multipartFileHeader, err := c.Request().FormFile("file")
		if err != nil {
			return err
		}
		defer file.Close()

		localFileName := "./" + multipartFileHeader.Filename
		out, err := os.Create(localFileName)
		if err != nil {
			fmt.Printf("failed to open the file %s for writing", localFileName)
			return err
		}
		defer out.Close()
		_, err = io.Copy(out, file)
		if err != nil {
			return err
		}
		
		return c.Render(http.StatusOK, "modal.htm", map[string]interface{}{
			"name": multipartFileHeader.Filename,
		})
		// return c.String(http.StatusOK, fmt.Sprintf("<p>File: %s successfully uploaded</p>", multipartFileHeader.Filename))
	}
}
