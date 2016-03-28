package server

import (
	"html/template"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"github.com/richardlt/the-collector/server/collections"
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
	collections.DatabaseCollection = db.C("collections")
	items.DatabaseCollection = db.C("items")

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	t := &Template{Templates: template.Must(template.ParseFiles("./client/dist/templates/index.html"))}
	e.SetRenderer(t)
	e.Any("/*", handleRoot(mode))
	e.Static("/js", "./client/dist/js")

	api := e.Group("/api")

	co := api.Group("/collections")
	collections.DeclareRoutes(co)

	it := api.Group("/items")
	it.Use(collections.Middleware())
	items.DeclareRoutes(it)

	e.Run(standard.New(":8080"))
}
