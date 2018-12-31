package server

import (
	"context"
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/mongodb/mongo-go-driver/mongo"
	errorsP "github.com/pkg/errors"

	"github.com/richardlt/the-collector/server/api"
	"github.com/richardlt/the-collector/server/api/errors"
	"github.com/richardlt/the-collector/server/collections"
	"github.com/richardlt/the-collector/server/facebook"
	"github.com/richardlt/the-collector/server/files"
	"github.com/richardlt/the-collector/server/items"
	"github.com/richardlt/the-collector/server/users"
)

// Start .
func Start(
	appURI, jwtSecret, secret string, debug bool,
	databaseURI, databaseName string,
	facebookAppID, facebookSecret string,
	minioURI, minioAccessKey, minioSecretKey, minioBucket string, minioSSL bool,
) error {
	client, err := mongo.NewClient(fmt.Sprintf("mongodb://%s", databaseURI))
	if err != nil {
		return errorsP.WithStack(err)
	}
	if err := client.Connect(context.Background()); err != nil {
		return errorsP.WithStack(err)
	}

	db := client.Database(databaseName)
	collections.InitDatabase(context.Background(), db)
	items.InitDatabase(context.Background(), db)
	files.InitDatabase(context.Background(), db)
	users.InitDatabase(context.Background(), db)

	files.InitStorage(minioURI, minioAccessKey, minioSecretKey,
		minioBucket, minioSSL)

	items.Init(jwtSecret)
	files.Init(jwtSecret)
	users.Init(jwtSecret, secret, appURI)
	facebook.Init(facebookAppID, facebookSecret)

	e := echo.New()
	e.Debug = debug
	e.Logger = api.NewLoggerConverter()
	e.Use(
		middleware.Secure(),
		middleware.CSRFWithConfig(middleware.CSRFConfig{
			CookiePath: "/",
		}),
		api.RequestLogger(debug),
		middleware.Recover(),
		errors.Middleware,
	)

	e.File("*", "./client/dist/index.html")

	e.Static("/static", "./client/dist/static")

	api := e.Group("/api")
	{ // routes for /api
		csg := api.Group("/collections", users.MiddlewareAuth()...)
		{ // routes for /api/collections
			csg.GET("", collections.HandleGetAll)
			csg.POST("", collections.HandlePost)
			cg := csg.Group("/:collectionSlugOrUUID", collections.Middleware)
			{ // routes for /api/collections/:collectionSlugOrUUID
				cg.GET("", collections.HandleGet)
				cg.DELETE("", collections.HandleDelete)
				isg := cg.Group("/items")
				{ // routes for /api/collections/:collectionSlugOrUUID/items
					isg.GET("", items.HandleGetAllForCollection)
					isg.POST("", items.HandlePost)
					ig := isg.Group("/:itemUUID", items.Middleware)
					{ // routes for /api/collections/:collectionSlugOrUUID/items/:itemUUID
						ig.GET("", items.HandleGet)
						ig.DELETE("", items.HandleDelete)
						ig.POST("/file", items.HandlePostFile)
					}
				}
			}
		}
		usg := api.Group("/users", users.MiddlewareAuth()...)
		{ // routes for /api/users
			usg.GET("/me", users.HandleGetMe)
		}
		fsg := api.Group("/files")
		{ // routes for /api/files
			fsg.GET("/:fileToken/:filename", files.HandleGet)
		}
		ag := api.Group("/auth")
		{ // routes for /api/auth
			ag.GET("/login", users.HandleLogin)
			ag.GET("/callback", users.HandleCallback)
		}
	}

	return errorsP.WithStack(e.Start(":8080"))
}
