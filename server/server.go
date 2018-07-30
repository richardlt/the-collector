package server

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"github.com/richardlt/the-collector/server/collections"
	"github.com/richardlt/the-collector/server/files"
	"github.com/richardlt/the-collector/server/items"
	"gopkg.in/mgo.v2"
)

// Start : start server
func Start(databaseURI string, databaseName string) {

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

	e.File("*", "./client/dist/index.html")

	e.Static("/static", "./client/dist/static")

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
