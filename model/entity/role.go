package entity

import (
	"github.com/jinzhu/gorm"
)

type Role struct {
	gorm.Model
	Name            string
	Label           string
	Description     string
	Users           []User
	RolesPermission []RolePermission
}
