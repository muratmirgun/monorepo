package user

import "gorm.io/gorm"

type Service struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *Service {
	return &Service{
		db: db,
	}
}
