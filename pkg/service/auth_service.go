package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"time"
	"webapp/model"
	"webapp/pkg/repository"

	"github.com/golang-jwt/jwt"
)

const (
	salt       = "hjqrhjqw124617ajfhajs"
	signingKey = "qrkjk#4#%35FSFJlja#4353KSFjH"
	tokenTTL   = 12 * time.Hour
)


type AuthService struct {
	repo repository.Authorization
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (r *AuthService) CreateUser(user model.User) (int, error){
	user.Password = r.generatePassword(user.Password)
	return r.repo.CreateUser(user)
}	

func (r *AuthService) generatePassword(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum(nil))
}

func (r *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := r.repo.GetUser(username, r.generatePassword(password))
	
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	return token.SignedString([]byte(signingKey))
}

func (r *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid singing method")
		}

		return []byte(signingKey), nil
	})
	
	if err != nil {
		return 0, err
	}
	
	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}

//in process...
func (r *AuthService) GetUser(username, password string) (model.User, error) {
	return model.User{}, nil
}