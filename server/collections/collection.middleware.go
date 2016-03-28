package collections

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

// Middleware : collection middleware
func Middleware() echo.MiddlewareFunc {
	return func(next echo.Handler) echo.Handler {
		return echo.HandlerFunc(func(c echo.Context) (err error) {
			collectionSlug := c.Request().Header().Get("X-COLLECTION-SLUG")
			if collectionSlug == "" {
				return c.JSON(http.StatusForbidden, nil)
			}
			fmt.Println(collectionSlug)
			collection, err := GetBySlug(collectionSlug)
			fmt.Println(collection)
			if err != nil {
				return c.JSON(http.StatusForbidden, nil)
			}
			if collection == nil {
				return c.JSON(http.StatusForbidden, nil)
			}
			c.Request().Header().Add("collectionID", collection.ID.Hex())
			return next.Handle(c)
		})
	}
}
