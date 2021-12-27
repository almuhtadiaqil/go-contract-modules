package entity

import (
	"github.com/jinzhu/gorm"
)

type RolePermission struct {
	gorm.Model
	RoleId       uint
	PermissionId uint
}
