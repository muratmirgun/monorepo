package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/muratmirgun/monorepo/internal/user"
)

type UsersHandler struct {
	UserService *user.Service
}

func (h *UsersHandler) Register(e *echo.Group) {
	e.POST("/social", h.SocialLogin)
	e.POST("/callback/:provider", h.Callback)
}

func (h *UsersHandler) SocialLogin(c echo.Context) error {
	return nil
}

func (h *UsersHandler) Callback(c echo.Context) error {
	return nil
}
