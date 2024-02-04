package app

import (
	"github.com/muratmirgun/monorepo/config"
	"github.com/muratmirgun/monorepo/internal/user"
	"github.com/muratmirgun/monorepo/internal/user/auth"
	"github.com/muratmirgun/monorepo/pkg/storage/database"
)

type Service struct {
	Auth *auth.Service
	User *user.Service
}

func NewServices(cfg *config.Config) (*Service, error) {
	var srv Service

	pg, err := database.NewPostgres(&cfg.Database)
	if err != nil {
		return nil, err
	}

	srv.User = user.NewUserService(pg.DB)
	srv.Auth = auth.NewAuthService(pg.DB)
	return &srv, nil
}
