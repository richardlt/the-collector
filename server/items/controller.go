package items

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/richardlt/the-collector/server/types"
)

// HandleGetAllForCollection : handler get all for collection
func HandleGetAllForCollection(c echo.Context) error {
	collection := c.Get("collection").(*types.Collection)
	items, err := GetAllByCollectionID(collection.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}
	return c.JSON(http.StatusOK, items)
}

// HandlePost : handler for post
func HandlePost(c echo.Context) error {
	collection := c.Get("collection").(*types.Collection)
	var data *types.Item
	c.Bind(&data)
	item := &types.Item{Name: data.Name, CollectionID: collection.ID}
	err := Create(item)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}
	return c.JSON(http.StatusOK, item)
}
