package handlers

import (
	"net/http"

	"github.com/EmmanoelDan/agro-pop/models"
	"github.com/EmmanoelDan/agro-pop/usecases"
	"github.com/gin-gonic/gin"
)

type RegisterUserHandler struct {
	RegisterUserUsecase *usecases.RegisterUserUseCase
}

func NewRegisterUserHandler(registerUserUsecase *usecases.RegisterUserUseCase) *RegisterUserHandler {
	return &RegisterUserHandler{RegisterUserUsecase: registerUserUsecase}
}

func (h *RegisterUserHandler) Register(ctx *gin.Context) {
	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
        ctx.JSON(400, gin.H{"error": err.Error()})
        return
    }

	newUser, err := h.RegisterUserUsecase.Register(user.Username, user.Password)

	if err != nil {
        ctx.JSON(400, gin.H{"error": err.Error()})
        return
    }

	ctx.JSON(http.StatusCreated, gin.H{"message": "Created user successfully!", "user": newUser.Username})
}
