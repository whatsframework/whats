package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

// UserAuth 用户授权表
type UserAuth struct {
	ID          uint           `gorm:"primarykey" json:"-"`
	UID         string         `gorm:"type:char(36);uniqueIndex" json:"uid" validate:"required"`                                            // 用户ID
	Type        string         `gorm:"type:char(36);uniqueIndex:idx_user_auth_type_identity" json:"type" validate:"required,ValidatorType"` // 身份标识类型 phone email account
	Identity    string         `gorm:"type:varchar(128);uniqueIndex:idx_user_auth_type_identity" json:"identity" validate:"required"`       // 身份标识 上述对应的值
	Certificate string         `gorm:"type:char(36)" json:"certificate" validate:"required"`                                                // 密码
	CreatedAt   time.Time      `json:"-"`
	UpdatedAt   time.Time      `json:"-"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

// ValidatorType ValidatorType
func ValidatorType(fl validator.FieldLevel) bool {
	switch fl.Field().String() {
	case "account", "phone", "email":
	default:
		return false
	}
	return true
}

// Validate the fields.
func (u *UserAuth) Validate() error {
	validate := validator.New()
	_ = validate.RegisterValidation("ValidatorType", ValidatorType)
	return validate.Struct(u)
}

// UserBase 用户基础资料表
type UserBase struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	UID       string         `gorm:"type:char(36);uniqueIndex" json:"uid" validate:"required"` // 用户UID
	Nickname  string         `gorm:"type:char(16)" json:"nickname" validate:"required"`        // 用户昵称
	Avatar    string         `gorm:"type:char(128)" json:"avatar"`                             // 头像
	AuthName  string         `gorm:"type:char(16)" json:"auth_name"`                           // 授权名称
	Wechat    string         `gorm:"type:char(32);uniqueIndex" json:"wechat"`                  // 微信号
	Phone     string         `gorm:"type:char(16);uniqueIndex" json:"phone"`                   // 手机号
	Address   string         `gorm:"type:char(128)" json:"address"`                            // 地址
	Gender    int            `gorm:"type:tinyint(1);not null;default:0" json:"gender"`         // 状态 0未知/保密，1男，2女
	Status    int            `gorm:"type:tinyint(1);not null;default:0" json:"status"`         // 状态 0暂存，1使用
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// Validate the fields.
func (u *UserBase) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}
