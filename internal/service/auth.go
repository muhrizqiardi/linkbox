package service

import (
	"fmt"
	"regexp"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/muhrizqiardi/linkbox/internal/constant"
	"github.com/muhrizqiardi/linkbox/internal/entities"
	"github.com/muhrizqiardi/linkbox/internal/entities/request"

	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	LogIn(payload request.LogInRequest) (string, error)
	CheckIsValid(token string) (entities.TokenClaims, error)
}

type authService struct {
	us     UserService
	secret string
}

func NewAuthService(us UserService, secret string) *authService {
	return &authService{us, secret}
}

func (as *authService) LogIn(payload request.LogInRequest) (string, error) {
	usernameRegex := regexp.MustCompile("^[a-z0-9_]{3,21}$")
	if !usernameRegex.MatchString(payload.Username) {
		return "", ErrInvalidUsername
	}
	user, err := as.us.GetOneByUsername(payload.Username)
	if err != nil {
		return "", err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password)); err != nil {
		return "", err
	}
	claim := entities.TokenClaims{
		UserID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	ss, err := token.SignedString([]byte(as.secret))
	if err != nil {
		return "", err
	}

	return ss, nil
}

func (as *authService) CheckIsValid(token string) (entities.TokenClaims, error) {
	parsedToken, err := jwt.ParseWithClaims(token, &entities.TokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(as.secret), nil
	})
	if err != nil {
		return entities.TokenClaims{}, err
	}
	claims, ok := parsedToken.Claims.(*entities.TokenClaims)
	if !ok || !parsedToken.Valid {
		return entities.TokenClaims{}, constant.ErrInvalidToken
	}

	return *claims, nil
}
