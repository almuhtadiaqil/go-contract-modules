package go_contract_modules

import (
	"github.com/jinzhu/gorm"
)

type Permission struct {
	gorm.Model
	Title           string
	Description     string
	RolesPermission []RolePermission
}
