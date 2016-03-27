package server

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"github.com/richardlt/the-collector/server/collections"
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
	collections.DatabaseCollection = db.C("collections")
	items.DatabaseCollection = db.C("items")

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	api := e.Group("/api")

	co := api.Group("/collections")
	collections.DeclareRoutes(co)

	it := api.Group("/items")
	it.Use(collections.Middleware())
	items.DeclareRoutes(it)

	e.Run(standard.New(":8080"))
}
