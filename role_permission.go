package go_contract_modules

import (
	"github.com/jinzhu/gorm"
)

type RolePermission struct {
	gorm.Model
	RoleId       uint
	PermissionId uint
	Role         Role       `gorm:"foreignKey:RoleId"`
	Permission   Permission `gorm:"foreignKey:PermissionId"`
}
