package usecase

import (
	"app/auth"
	"app/auth/repository"
	"context"
	"crypto/sha1"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type AuthClaims struct {
	jwt.StandardClaims
	User *auth.User `json:"user"`
}

type UserUseCase struct {
	userRepository   *repository.UserRepository
	salt             string
	signingKey       []byte
	tokenDurationTTL int64
}

func New(r *repository.UserRepository) *UserUseCase {
	return &UserUseCase{
		userRepository: r,
	}
}

func (u *UserUseCase) SignIn(ctx context.Context, username, password string) (string, error) {
	password = u.encodePassword(password)

	user, err := u.userRepository.Get(ctx, username, password)
	if err != nil {
		return "", err
	}

	claims := AuthClaims{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(u.tokenDurationTTL)).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(u.signingKey)
}

func (u *UserUseCase) SignUp(ctx context.Context, username, password string) error {
	password = u.encodePassword(password)
	return u.userRepository.Create(ctx, username, password)
}

func (u *UserUseCase) encodePassword(rawPassword string) string {
	pwd := sha1.New()
	pwd.Write([]byte(rawPassword))
	pwd.Write([]byte(u.salt))
	return fmt.Sprintf("%x", pwd.Sum(nil))
}
