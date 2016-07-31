package items

import (
	"net/http"

	"github.com/labstack/echo"
)

// Middleware : collection middleware
func Middleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return echo.HandlerFunc(func(c echo.Context) (err error) {
			itemUUID := c.Param("itemUUID")
			if itemUUID == "" {
				return c.JSON(http.StatusBadRequest, nil)
			}
			item, err := GetByUUID(itemUUID)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, nil)
			}
			if item == nil {
				return c.JSON(http.StatusNotFound, nil)
			}
			c.Set("item", item)
			return next(c)
		})
	}
}
