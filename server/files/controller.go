package files

import (
	"net/http"

	"github.com/labstack/echo"
)

// HandlePost : handlerfor post
func HandlePost(c echo.Context) error {
	return c.JSON(http.StatusOK, nil)
}
