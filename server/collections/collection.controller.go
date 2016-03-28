package collections

import (
	"net/http"
	"unicode"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"

	"github.com/labstack/echo"
	"github.com/mrvdot/golang-utils"
)

// HandleGet : handler for get
func HandleGet() echo.HandlerFunc {
	return func(c echo.Context) error {
		collection, err := GetBySlug(c.Param("collectionSlug"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, nil)
		}
		if collection == nil {
			return c.JSON(http.StatusNotFound, nil)
		}
		return c.JSON(http.StatusOK, collection.Restify())
	}
}

// HandleGetAll : handler for get all
func HandleGetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		collections, err := GetAll()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, nil)
		}
		collectionsRest := make([]CollectionRest, 0)
		for _, collection := range collections {
			collectionsRest = append(collectionsRest, collection.Restify())
		}
		return c.JSON(http.StatusOK, collectionsRest)
	}
}

func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
}

// HandlePost: handler for post
func HandlePost() echo.HandlerFunc {
	return func(c echo.Context) error {
		var data CollectionRest
		c.Bind(&data)
		if data.Name == "" {
			return c.JSON(http.StatusNotAcceptable, nil)
		}
		t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
		name, _, _ := transform.String(t, data.Name)
		data.Slug = utils.GenerateSlug(name)
		exist, err := ExistSlug(data.Slug)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, data)
		}
		if exist {
			return c.JSON(http.StatusNotAcceptable, nil)
		}
		collection := data.ToCollection()
		err = collection.Create()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, data)
		}
		return c.JSON(http.StatusOK, collection.Restify())
	}
}
