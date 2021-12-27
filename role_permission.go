package go_contract_modules

import (
	"github.com/jinzhu/gorm"
)

type RolePermission struct {
	gorm.Model
	RoleId       uint
	PermissionId uint
}
