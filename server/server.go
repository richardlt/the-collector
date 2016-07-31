package server

import (
	"html/template"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"github.com/richardlt/the-collector/server/collections"
	"github.com/richardlt/the-collector/server/files"
	"github.com/richardlt/the-collector/server/items"
	"gopkg.in/mgo.v2"
)

func handleRoot(mode string) echo.HandlerFunc {
	return func(c echo.Context) error {
		data := make(map[string]interface{})
		if mode == "dev" {
			data["Title"] = "The Collector - DEV"
			data["Bundle"] = "http://localhost:8081/js/bundle.js"
			return c.Render(http.StatusOK, "index", data)
		}
		data["Title"] = "The Collector"
		data["Bundle"] = "/js/bundle.js"
		return c.Render(http.StatusOK, "index", data)
	}
}

// Start : start server
func Start(databaseURI string, databaseName string, mode string) {

	log.Info("[server][Start] Trying connect to database at ", databaseURI)
	s, err := mgo.Dial(databaseURI)
	if err != nil {
		panic(err)
	}
	defer s.Close()
	log.Info("[server][Start] Successfully connected to database")

	db := s.DB(databaseName)
	collections.Collection = db.C("collections")
	items.Collection = db.C("items")

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	t := &Template{Templates: template.Must(template.ParseFiles("./client/dist/templates/index.html"))}
	e.SetRenderer(t)
	e.Any("/*", handleRoot(mode))
	e.Static("/js", "./client/dist/js")

	api := e.Group("/api")
	{ // routes for /api
		cog := api.Group("/collections")
		{ // routes for /api/collections
			cog.Get("", collections.HandleGetAll)
			cog.Post("", collections.HandlePost)
			coug := cog.Group("/:collectionUUID", collections.Middleware())
			{
				couitg := coug.Group("/items")
				{ // routes for /api/items
					couitg.Get("", items.HandleGetAllForCollection)
					couitg.Post("", items.HandlePost)
					couitug := couitg.Group("/:itemUUID", items.Middleware())
					{ // routes for /api/items/:itemUUID
						couitug.Post("/files", files.HandlePost)
					}
				}
			}
		}
	}

	e.Run(standard.New(":8080"))
}
