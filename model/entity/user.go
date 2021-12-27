package entity

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	RoleId   uint
	Name     string
	Email    string
	Username string
	Password string
	Mobile   string
}
