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
	dto, err := reposity.ReadItemWithFilterIntoDTO[model.CreateUser, model.User]("username= ?", req.Username)
	if err != nil {
		return nil, errors.New("invalid credentials username")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(dto.Password), []byte(req.Password)); err != nil {
		return nil, errors.New("invalid credentials password")
	}

	// Generate JWT token
	claims := jwt.MapClaims{
		"user_id":  dto.UserID,
		"username": dto.Username,
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

func (s *Service) Register(req *RegisterRequest) error {
	// Check if username already exists
	var existingUser model.CreateUser
	if err := s.DB.Where("username = ?", req.Username).First(&existingUser).Error; err == nil {
		return errors.New("username already exists")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Create new user
	user := model.CreateUser{
		Username: req.Username,
		Password: string(hashedPassword),
		Email:    req.Email,
		FullName: req.FullName,
		Role:     req.Role,
	}

	if err := s.DB.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

func (s *Service) ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return s.JWTSecret, nil
	})
}
