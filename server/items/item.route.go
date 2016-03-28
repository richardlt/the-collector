package items

import "github.com/labstack/echo"

// DeclareRoutes : declare routes for items
func DeclareRoutes(p *echo.Group) {
	p.Get("", HandleGetAll())
	p.Post("", HandlePost())
	p.Get("/:itemSlug", HandleGet())
}
