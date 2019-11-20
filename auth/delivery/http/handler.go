package http

import (
	"app/auth/usecase"
	"app/server/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	authUseCase *usecase.UserUseCase
}

func New(uc *usecase.UserUseCase) *Handler {
	return &Handler{authUseCase: uc}
}

func (h *Handler) SignIn(ctx *gin.Context) {

	input := new(SignInForm)

	if err := ctx.BindJSON(&input); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, &response.Basic{
			Status:  response.StatusInvalidRequestParams,
			Message: err.Error(),
		})
		return
	}

	token, err := h.authUseCase.SignIn(ctx, input.Username, input.Password)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	ctx.JSON(http.StatusOK, SignInResponse{Token: token})
}

func (h *Handler) SignUp(ctx *gin.Context) {
	input := new(SignUpForm)

	if err := ctx.BindJSON(&input); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, &response.Basic{
			Status:  response.StatusInvalidRequestParams,
			Message: err.Error(),
		})
		return
	}

	if input.Password1 != input.Password2 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.Basic{
			Status:  response.StatusInvalidRequestParams,
			Message: "passwords are not identical",
		})
		return
	}

	if err := h.authUseCase.SignUp(ctx, input.Username, input.Password1); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.Basic{
			Status:  response.StatusInvalidRequestParams,
			Message: err.Error(),
		})
		return
	}
	ctx.Status(http.StatusOK)

}
