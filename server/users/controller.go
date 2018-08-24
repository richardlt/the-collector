package users

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/richardlt/the-collector/server/types"
)

// HandleGetMe .
func HandleGetMe(c echo.Context) error {
	return c.JSON(http.StatusOK, c.Get("me").(*types.User))
}
