package services

import (
	"context"
	"errors"
	"time"

	"github.com/IvanGelium/main-service/internal/domain"
	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	repo domain.AuthRepository
}

func NewAuthService(r domain.AuthRepository) domain.AuthService {
	return &authService{repo: r}
}

func (serv *authService) Login(ctx context.Context, input domain.LoginInput) (domain.UserJWT, error) {
	user, err := serv.repo.GetUserByEmail(ctx, input.Email)
	if err != nil {
		return domain.UserJWT{}, err
	}
	err = serv.verifyPassword(user.PasswordHash, input.Password)
	if err != nil {
		return domain.UserJWT{}, err
	}

	token, err := serv.generateToken(user.ID.String(), user.AccountId.String())

	result := domain.UserJWT{
		ID:        user.ID,
		Email:     user.Email,
		Token:     token,
		AccountId: user.AccountId,
	}

	return result, nil
}

func (serv *authService) verifyPassword(hashedPassword string, loginPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(loginPassword))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return errors.New("Невероный логин или пароль")
		}
		return err
	}
	return err
}

func (serv *authService) generateToken(userId string, accountId string) (string, error) {
	jwtSecret := []byte("secret_key")
	claims := jwt.MapClaims{
		"user_id":    userId,
		"account_id": accountId,
		"exp":        jwt.NewNumericDate(time.Now().Add(time.Minute * 15)),
		"iat":        jwt.NewNumericDate(time.Now()),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func (serv *authService) SignUp(ctx context.Context, data domain.SignupInput) (domain.UserJWT, error) {
	passwordHash, err := serv.hashPassword(data.Password)
	if err != nil {
		return domain.UserJWT{}, err
	}

	payload := domain.User{
		Email:        data.Email,
		Nickname:     &data.Nickname,
		PasswordHash: passwordHash,
	}
	id := uuid.Must(uuid.NewV7())
	user, err := serv.repo.CreateUser(ctx, payload, id)
	if err != nil {
		return domain.UserJWT{}, err
	}
	token, err := serv.generateToken(user.ID.String(), user.AccountId.String())
	if err != nil {
		return domain.UserJWT{}, err
	}

	result := domain.UserJWT{
		ID:        user.ID,
		AccountId: user.AccountId,
		Email:     user.Email,
		Token:     token,
	}
	return result, nil
}

func (serv *authService) hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
