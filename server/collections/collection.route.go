package collections

import "github.com/labstack/echo"

// DeclareRoutes : declare routes for collections
func DeclareRoutes(p *echo.Group) {
	p.Get("/collections", HandleGetAll())
	p.Post("/collections", HandlePost())
	p.Get("/collections/:collectionSlug", HandleGet())
}
