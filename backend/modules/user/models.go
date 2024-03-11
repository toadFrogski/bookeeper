package user

import (
	"bookeeper/domain"
	"bookeeper/utils/dto"
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
		Token   string `json:"token"`
		Expires int64  `json:"expires"`
	} // @name Auth

	ValidationError struct {
		Type        string `json:"type"`
		Description string `json:"description"`
	} //@name ValidationError

	NamedValidationErrors struct {
		Name   string            `json:"name"`
		Errors []ValidationError `json:"errors"`
	} // @name NamedValidationErrors

	AuthResponse                  = dto.Response[Auth]                  // @name AuthResponse
	UserResponse                  = dto.Response[domain.User]           // @name UserResponse
	NamedValidationErrorsResponse = dto.Response[NamedValidationErrors] // @name NamedValidationErrorsResponse
)
