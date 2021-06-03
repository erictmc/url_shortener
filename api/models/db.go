package models

import "gorm.io/gorm"

type AppDB struct {
	DB *gorm.DB
}
