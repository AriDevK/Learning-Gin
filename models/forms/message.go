package forms

type MessageForm struct {
	Message string `form:"message" binding:"required"`
	Nick    string `form:"nick" binding:"required"`
}
