package collections

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/richardlt/the-collector/server/types"
)

// HandleGet : handler for get
func HandleGet(c echo.Context) error {
	collection := c.Get("collection").(*types.Collection)
	return c.JSON(http.StatusOK, collection)
}

// HandleGetAll : handler for get all
func HandleGetAll(c echo.Context) error {
	collections, err := GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}
	return c.JSON(http.StatusOK, collections)
}

// HandlePost : handler for post
func HandlePost(c echo.Context) error {
	var data types.Collection
	c.Bind(&data)
	collection := &types.Collection{Name: data.Name}
	err := Create(collection)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}
	return c.JSON(http.StatusOK, collection)
}
