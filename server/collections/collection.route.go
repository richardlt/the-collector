package collections

import "github.com/labstack/echo"

// DeclareRoutes : declare routes for collections
func DeclareRoutes(p *echo.Group) {
	p.Get("", HandleGetAll())
	p.Post("", HandlePost())
	p.Get("/:collectionSlug", HandleGet())
}
