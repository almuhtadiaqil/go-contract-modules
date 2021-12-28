package go_contract_modules

import (
	"github.com/jinzhu/gorm"
)

type RolePermission struct {
	gorm.Model
	RoleId       uint       `gorm:"primaryKey"`
	PermissionId uint       `gorm:"primaryKey"`
	Role         Role       `gorm:"foreignKey:RoleId"`
	Permission   Permission `gorm:"foreignKey:PermissionId"`
}
