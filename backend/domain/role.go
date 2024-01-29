package domain

import "gg/utils/constants"

type (
	Role struct {
		ID         uint           `gorm:"primarykey"`
		Name       constants.Role `gorm:"column:name"`
		Users      []*User        `gorm:"many2many:user_roles;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
		Permission []*Permission  `gorm:"many2many:role_permissions"`
	} // @name Role
)
