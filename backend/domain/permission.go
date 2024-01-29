package domain

type (
	Permission struct {
		ID         uint    `gorm:"primarykey"`
		Permission string  `gorm:"permission"`
		Roles      []*Role `gorm:"many2many:role_permissions"`
	} // @name Permission
)
