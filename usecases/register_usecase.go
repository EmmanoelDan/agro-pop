package usecases

import (
	"errors"

	"github.com/EmmanoelDan/agro-pop/models"
	"github.com/EmmanoelDan/agro-pop/repositories"
	"github.com/EmmanoelDan/agro-pop/utils"
)

type RegisterUserUseCase struct {
	RegisterUserRepo *repositories.UserRepository
}

func NewRegisterUserUseCase(registerUserRepo *repositories.UserRepository) *RegisterUserUseCase {
	return &RegisterUserUseCase{RegisterUserRepo: registerUserRepo}
}

func (u *RegisterUserUseCase) Register(username string, password string) (*models.User, error) {
	_, err := u.RegisterUserRepo.FindByUsername(username)
	if err == nil {
		return nil, errors.New("User already registered")
	}

	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
        return nil, errors.New("Invalid password")
    }
	
	user := &models.User{
		Username: username,
        Password: hashedPassword,
	}
	err = u.RegisterUserRepo.Create(user)

	if  err != nil {
		return nil, errors.New("Could not create user")
	}
	return user, nil
}