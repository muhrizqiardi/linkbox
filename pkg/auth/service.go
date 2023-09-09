package auth

import (
	"errors"
	"regexp"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/muhrizqiardi/linkbox/linkbox/pkg/user"
	"golang.org/x/crypto/bcrypt"
)

var ErrInvalidUsername error = errors.New("Username can only contains alphanumeric character and underscore, and can only have at least 3 characters and 21 characters maximum")

type Service interface {
	LogIn(payload LogInDTO) (string, error)
	CheckIsValid(token string) (TokenClaims, string, error)
}

type service struct {
	us     user.Service
	secret string
}

func NewService(us user.Service, secret string) *service {
	return &service{us, secret}
}

func (s *service) LogIn(payload LogInDTO) (string, error) {
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
	claim := TokenClaims{
		user.ID,
		jwt.RegisteredClaims{
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

func (s *service) CheckIsValid(token string) (TokenClaims, string, error) {
	parsedToken, err := jwt.ParseWithClaims(token, &TokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(s.secret), nil
	})
	claims, ok := parsedToken.Claims.(*TokenClaims)
	if !ok || !parsedToken.Valid {
		return TokenClaims{}, "", nil
	}

	newClaims := TokenClaims{
		claims.UserID,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, newClaims)
	ss, err := newToken.SignedString(s.secret)
	if err != nil {
		return TokenClaims{}, "", err
	}

	return *claims, ss, nil
}
