package items

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/mrvdot/golang-utils"
	"gopkg.in/mgo.v2/bson"
)

// HandleGet : handler for get
func HandleGet() echo.HandlerFunc {
	return func(c echo.Context) error {
		item, err := GetBySlug(c.Param("itemSlug"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, nil)
		}
		if item == nil {
			return c.JSON(http.StatusNotFound, nil)
		}
		return c.JSON(http.StatusOK, item.Restify())
	}
}

// HandleGetAll : handler for get all
func HandleGetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		items, err := GetAllInCollection(bson.ObjectIdHex(c.Request().Header().Get("collectionID")))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, nil)
		}
		itemsRest := make([]ItemRest, 0)
		for _, item := range items {
			itemsRest = append(itemsRest, item.Restify())
		}
		return c.JSON(http.StatusOK, itemsRest)
	}
}

// HandlePost: handler for post
func HandlePost() echo.HandlerFunc {
	return func(c echo.Context) error {
		var data ItemRest
		c.Bind(&data)
		if data.Name == "" {
			return c.JSON(http.StatusNotAcceptable, nil)
		}
		data.Slug = utils.GenerateSlug(data.Name)
		exist, err := ExistSlug(data.Slug)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, data)
		}
		if exist {
			return c.JSON(http.StatusNotAcceptable, nil)
		}
		item := data.ToItem()
		item.CollectionID = bson.ObjectIdHex(c.Request().Header().Get("collectionID"))
		err = item.Create()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, data)
		}
		return c.JSON(http.StatusOK, item.Restify())
	}
}
