package usecase

import (
	"go-rest-api/models"
	"go-rest-api/repository"
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
	userRepository repository.IUserRepository
}

func NewUserUsecase(ur repository.IUserRepository) IUserUsecase {
	return &userUsecase{ur}
}

func (uu *userUsecase) SignUp(user models.User) (models.UserResponse, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return models.UserResponse{}, err
	}

	newUser := models.User{Email: user.Email, Password: string(hash)}
	if err := uu.userRepository.CreateUser(&newUser); err != nil {
		return models.UserResponse{}, err
	}

	resUser := models.UserResponse{
		Id:    newUser.Id,
		Email: newUser.Email,
	}

	return resUser, nil
}

func (uu *userUsecase) Login(user models.User) (string, error) {
	storedUser := models.User{}

	if err := uu.userRepository.GetUserByEmail(&storedUser, user.Email); err != nil {
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
