package collections

import "github.com/labstack/echo"

// Middleware : collection middleware
func Middleware() echo.MiddlewareFunc {
	return func(next echo.Handler) echo.Handler {
		return echo.HandlerFunc(func(c echo.Context) (err error) {
			return next.Handle(c)
		})
	}
}
