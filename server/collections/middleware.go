package collections

import (
	"net/http"

	"github.com/labstack/echo"
)

// Middleware : collection middleware
func Middleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return echo.HandlerFunc(func(c echo.Context) (err error) {
			collectionUUID := c.Param("collectionUUID")
			if collectionUUID == "" {
				return c.JSON(http.StatusBadRequest, nil)
			}
			collection, err := GetByUUID(collectionUUID)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, nil)
			}
			if collection == nil {
				return c.JSON(http.StatusNotFound, nil)
			}
			c.Set("collection", collection)
			return next(c)
		})
	}
}
