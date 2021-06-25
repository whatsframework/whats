package database

import (
	"whats/app/models"

	"gorm.io/gorm"
)

// migrate 数据库自动迁移
func migrate(db *gorm.DB) {
	_ = db.AutoMigrate(&models.UserAuth{})
	_ = db.AutoMigrate(&models.UserBase{})
}
