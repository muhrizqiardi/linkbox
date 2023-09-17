package auth

import (
	"errors"
	"regexp"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/muhrizqiardi/linkbox/linkbox/pkg/common"
	"golang.org/x/crypto/bcrypt"
)

var ErrInvalidUsername error = errors.New("Username can only contains alphanumeric character and underscore, and can only have at least 3 characters and 21 characters maximum")

type service struct {
	us     common.UserService
	secret string
}

func NewService(us common.UserService, secret string) *service {
	return &service{us, secret}
}

func (s *service) LogIn(payload common.LogInDTO) (string, error) {
	usernameRegex := regexp.MustCompile("^[a-z0-9_]{3,21}$")
	if !usernameRegex.MatchString(payload.Username) {
		return "", ErrInvalidUsername
	}
	user, err := s.us.GetOneByUsername(payload.Username)
	if err != nil {
		return "", err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password)); err != nil {
		return "", err
	}
	claim := common.TokenClaims{
		UserID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	ss, err := token.SignedString([]byte(s.secret))
	if err != nil {
		return "", err
	}

	return ss, nil
}

func (s *service) CheckIsValid(token string) (common.TokenClaims, string, error) {
	parsedToken, err := jwt.ParseWithClaims(token, &common.TokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(s.secret), nil
	})
	claims, ok := parsedToken.Claims.(*common.TokenClaims)
	if !ok || !parsedToken.Valid {
		return common.TokenClaims{}, "", nil
	}

	newClaims := common.TokenClaims{
		UserID: claims.UserID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, newClaims)
	ss, err := newToken.SignedString([]byte(s.secret))
	if err != nil {
		return common.TokenClaims{}, "", err
	}

	return *claims, ss, nil
}
