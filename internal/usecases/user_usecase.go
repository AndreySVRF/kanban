package usecases

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"kanban/internal/entities"
	"kanban/internal/repositories"
)

type UserUseСase interface {
	Register(username, password string) error
	Login(username, password string) (string, error) // Возвращает JWT
}

type userUseСase struct {
	userRepo  repositories.UserRepository
	jwtSecret string
}

func NewUserUseСase(userRepo repositories.UserRepository, jwtSecret string) UserUseСase {
	return &userUseСase{
		userRepo:  userRepo,
		jwtSecret: jwtSecret,
	}
}

func (u *userUseСase) Register(username, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user := entities.User{
		Username: username,
		Password: string(hashedPassword),
	}
	return u.userRepo.CreateUser(user)
}

func (u *userUseСase) Login(username, password string) (string, error) {
	user, err := u.userRepo.GetUserByUsername(username)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(u.jwtSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
