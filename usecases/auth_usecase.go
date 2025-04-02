package usecases

import (
	"github.com/EmmanoelDan/agro-pop/repositories"
	"github.com/EmmanoelDan/agro-pop/utils"
)

type AuthUseCase struct {
	UserRepo repositories.UserRepository
}

func NewAuthUseCase(userRepo repositories.UserRepository) *AuthUseCase {
    return &AuthUseCase{UserRepo: userRepo}
}

func (a *AuthUseCase) Login(username string, password string) (string, error) {
	user, err := a.UserRepo.FindByUsername(username)
    if err != nil {
        return "", err
    }

    if !utils.ComparePassword(user.Password, password) {
        return "", err
    }

    token, err := utils.GenerateJWT(username)
    if err != nil {
        return "", err
    }

    return token, nil
}