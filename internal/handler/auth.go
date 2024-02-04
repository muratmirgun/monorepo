package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/muratmirgun/monorepo/internal/user/auth"
)

type AuthHandler struct {
	AuthService *auth.Service
}

func (h *AuthHandler) Register(e *echo.Group) {
	e.POST("/sign-in", h.SignIn)
	e.POST("/sign-up", h.SignUp)
}

func (h *AuthHandler) SignIn(c echo.Context) error {
	return nil
}

func (h *AuthHandler) SignUp(c echo.Context) error {
	return nil
}
