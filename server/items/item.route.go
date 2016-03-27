package items

import "github.com/labstack/echo"

// DeclareRoutes : declare routes for items
func DeclareRoutes(p *echo.Group) {
	p.Get("/items", HandleGetAll())
	p.Post("/items", HandlePost())
	p.Get("/items/:itemSlug", HandleGet())
}
