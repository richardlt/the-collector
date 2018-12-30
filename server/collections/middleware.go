package collections

import (
	"context"

	"github.com/labstack/echo"

	"github.com/richardlt/the-collector/server/api/errors"
	"github.com/richardlt/the-collector/server/types"
)

// Middleware .
func Middleware(next echo.HandlerFunc) echo.HandlerFunc {
	return echo.HandlerFunc(func(c echo.Context) (err error) {
		me := c.Get("me").(*types.User)

		slugOrUUID := c.Param("collectionSlugOrUUID")
		if slugOrUUID == "" {
			return errors.NewNotFound()
		}

		co, err := Get(context.Background(), newCriteria().
			UserID(me.ID).SlugOrUUID(slugOrUUID))
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
