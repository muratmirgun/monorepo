package database

import (
	"fmt"
	"github.com/muratmirgun/monorepo/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	DB *gorm.DB
	//TODO: add custom options
}

func NewPostgres(cfg *config.Database) (*Postgres, error) {
	pg := &Postgres{}

	db, err := gorm.Open(postgres.Open(fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", cfg.Host, cfg.User, cfg.Password, cfg.Name, cfg.Port)),
		&gorm.Config{})
	if err != nil {
		return nil, err
	}

	pg.DB = db
	return pg, nil
}
