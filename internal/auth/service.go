package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/hoangtu1372k2/common-go/reposity"
	"github.com/hoangtu1372k2/vms/internal/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Service struct {
	DB         *gorm.DB
	JWTSecret  []byte
	ExpireTime time.Duration
}

func (s *Service) Login(req *LoginRequest) (*TokenResponse, error) {
	dto, err := reposity.ReadItemWithFilterIntoDTO[model.CreateUser, model.User]("full_name= ?", req.Username)
	if err != nil {
		return nil, errors.New("invalid credentials username")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(dto.Password), []byte(req.Password)); err != nil {
		return nil, errors.New("invalid credentials password")
	}

	// Generate JWT token
	claims := jwt.MapClaims{
		"user_id":  dto.UserID,
		"username": dto.FullName,
		"email":    dto.Email,
		"role":     dto.Role,
		"exp":      time.Now().Add(s.ExpireTime).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(s.JWTSecret)
	if err != nil {
		return nil, err
	}

	return &TokenResponse{
		Token:     signedToken,
		ExpiresIn: int64(s.ExpireTime.Seconds()),
		TokenType: "Bearer",
	}, nil
}

func (s *Service) ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return s.JWTSecret, nil
	})
}
