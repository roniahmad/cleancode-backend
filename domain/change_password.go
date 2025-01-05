package domain

type ChangePassword struct {
	OldPassword     string `form:"old_password" validate:"required"`
	NewPassword     string `form:"new_password" validate:"required"`
	ConfirmPassword string `form:"confirm_password" validate:"required"`
}
