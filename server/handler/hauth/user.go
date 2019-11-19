package hauth

import "github.com/gin-gonic/gin"

type UserHandler struct {
}

type RegisterForm struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (h *UserHandler) SignUp(ctx gin.Context) {

}

func (h *UserHandler) SignIn() {

}

func (h *UserHandler) SignOut() {

}
