package app

import (
	"github.com/labstack/echo/v4"
	"github.com/muratmirgun/monorepo/internal/handler"
)

func RegisterRoutes(e *echo.Echo, di *Service) {
	e.GET("/health", func(c echo.Context) error {
		return c.String(200, "OK")
	})
	v1 := e.Group("/api/v1")

	{
		uh := &handler.UsersHandler{
			UserService: di.User,
		}
		uh.Register(v1.Group("/users"))
	}
	{
		ah := &handler.AuthHandler{
			AuthService: di.Auth,
		}
		ah.Register(v1.Group("/auth"))
	}
}
