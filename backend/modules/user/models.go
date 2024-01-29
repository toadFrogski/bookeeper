package user

import (
	"gg/domain"
	"gg/utils/dto"
)

type (
	RegisterUserForm struct {
		Email    string `json:"email" binding:"required"`
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	} // @name RegisterUserForm

	LoginUserForm struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	} // @name LoginUserForm

	Auth struct {
		Token string `json:"token"`
	} // @name Auth

	AuthResponse = dto.Response[Auth]        // @name AuthResponse
	UserResponse = dto.Response[domain.User] // @name UserResponse
)
