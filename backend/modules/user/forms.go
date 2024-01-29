package user

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

	AuthResponse struct {
		Token string `json:"token"`
	} // @name AuthResponse
)
