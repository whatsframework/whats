package database

import (
	"whats/app/models"

	"gorm.io/gorm"
)

// migrate 数据库自动迁移
func migrate(db *gorm.DB) {
	db.AutoMigrate(&models.UserAuth{})
	db.AutoMigrate(&models.UserBase{})
}
