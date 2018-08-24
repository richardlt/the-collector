package collections

import (
	"context"
	"net/http"

	"github.com/labstack/echo"
	"github.com/richardlt/the-collector/server/api"
	"github.com/richardlt/the-collector/server/api/errors"
	"github.com/richardlt/the-collector/server/database"
	"github.com/richardlt/the-collector/server/types"
)

// HandleGet .
func HandleGet(c echo.Context) error {
	co := c.Get("collection").(*types.Collection)
	return c.JSON(http.StatusOK, co)
}

// HandleGetAll .
func HandleGetAll(c echo.Context) error {
	cs, err := GetAll(context.Background(), &database.BasicCriteria{})
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, cs)
}

// HandlePost .
func HandlePost(c echo.Context) error {
	var data types.Collection
	c.Bind(&data)

	data.Slug = api.Slugify(data.Name)
	res, err := Get(context.Background(), newCriteria().Slug(data.Slug))
	if err != nil {
		return err
	}
	if res != nil {
		return errors.NewData("a collection already exists with the same name")
	}

	if err := Create(context.Background(), &data); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, data)
}

// HandleDelete .
func HandleDelete(c echo.Context) error {
	co := c.Get("collection").(*types.Collection)

	if err := Delete(context.Background(), co); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, nil)
}
