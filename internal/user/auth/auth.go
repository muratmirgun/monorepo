package auth

import "gorm.io/gorm"

type Service struct {
	db *gorm.DB
}

func NewAuthService(db *gorm.DB) *Service {
	return &Service{
		db: db,
	}
}
