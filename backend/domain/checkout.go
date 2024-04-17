package domain

type (
	Checkout struct {
		ID     uint  `gorm:"primarykey"`
		BookID uint  `json:"-"`
		Book   *Book `json:"book,omitempty"`
		UserID uint  `json:"-"`
		User   *User `json:"user,omitempty"`
	} // @name Checkout
)
