package collections

import (
	"context"

	"github.com/labstack/echo"
	"github.com/richardlt/the-collector/server/api/errors"
)

// Middleware .
func Middleware(next echo.HandlerFunc) echo.HandlerFunc {
	return echo.HandlerFunc(func(c echo.Context) (err error) {
		slugOrUUID := c.Param("collectionSlugOrUUID")
		if slugOrUUID == "" {
			return errors.NewNotFound()
		}

		co, err := Get(context.Background(), newCriteria().SlugOrUUID(slugOrUUID))
		if err != nil {
			return err
		}
		if co == nil {
			return errors.NewNotFound()
		}

		c.Set("collection", co)
		return next(c)
	})
}
