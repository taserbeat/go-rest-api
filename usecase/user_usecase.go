package usecase

import (
	"go-rest-api/models"
	"go-rest-api/repository"
	"go-rest-api/validator"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

// ユーザーのユースケースのインターフェース
type IUserUsecase interface {
	// サインアップ(新規登録)を行う
	SignUp(user models.User) (models.UserResponse, error)

	// ログインする
	// 戻り値はJWTトークンとエラーとなる
	Login(user models.User) (string, error)
}

type userUsecase struct {
	ur repository.IUserRepository
	uv validator.IUserValidator
}

func NewUserUsecase(ur repository.IUserRepository, uv validator.IUserValidator) IUserUsecase {
	return &userUsecase{ur, uv}
}

func (uu *userUsecase) SignUp(user models.User) (models.UserResponse, error) {
	// サインアップ時の入力値をバリデーションする
	if err := uu.uv.UserValidate(user); err != nil {
		return models.UserResponse{}, err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return models.UserResponse{}, err
	}

	newUser := models.User{Email: user.Email, Password: string(hash)}
	if err := uu.ur.CreateUser(&newUser); err != nil {
		return models.UserResponse{}, err
	}

	resUser := models.UserResponse{
		Id:    newUser.Id,
		Email: newUser.Email,
	}

	return resUser, nil
}

func (uu *userUsecase) Login(user models.User) (string, error) {
	// ログイン時にバリデーションを行う
	if err := uu.uv.UserValidate(user); err != nil {
		return "", err
	}

	storedUser := models.User{}

	if err := uu.ur.GetUserByEmail(&storedUser, user.Email); err != nil {
		return "", err
	}

	// DBのパスワードとWebアプリからのパスワードを検証する
	err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": storedUser.Id,
		"exp":    time.Now().Add(time.Hour * 12).Unix(),
	})
	tokenStr, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}
