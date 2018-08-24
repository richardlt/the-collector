package items

import (
	"context"

	"github.com/labstack/echo"
	"github.com/richardlt/the-collector/server/api/errors"
	"github.com/richardlt/the-collector/server/types"
)

// Middleware .
func Middleware(next echo.HandlerFunc) echo.HandlerFunc {
	return echo.HandlerFunc(func(c echo.Context) (err error) {
		co := c.Get("collection").(*types.Collection)

		uuid := c.Param("itemUUID")
		if uuid == "" {
			return errors.NewNotFound()
		}

		i, err := Get(context.Background(), newCriteria().CollectionID(co.ID).UUID(uuid))
		if err != nil {
			return err
		}
		if i == nil {
			return errors.NewNotFound()
		}

		c.Set("item", i)
		return next(c)
	})
}
