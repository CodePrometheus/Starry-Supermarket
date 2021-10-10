package forms

type EmailForm struct {
	Email string `form:"email" json:"email" binding:"required,email"`
	Type  uint   `form:"type" json:"type" binding:"required"`
}
