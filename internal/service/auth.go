package service

import (
	"errors"
	"github.com/alihaqberdi/goga_go/internal/dtos"
	"github.com/alihaqberdi/goga_go/internal/models"
	"github.com/alihaqberdi/goga_go/internal/pkg/jwt_manager"
	"github.com/alihaqberdi/goga_go/internal/repo"
	"github.com/alihaqberdi/goga_go/internal/service/caching"
	"golang.org/x/crypto/bcrypt"
	"regexp"
)

type Auth struct {
	Repo       *repo.Repo
	Cache      *caching.Cache
	JwtManager *jwt_manager.JwtManager
}

func (s *Auth) Register(data *dtos.Register) (*dtos.AuthRes, error) {
	user, err := s.register(data)
	if err != nil {
		return nil, err
	}

	return s.authRes(user)
}

func (s *Auth) Login(data *dtos.Login) (*dtos.AuthRes, error) {
	user, err := s.Repo.Users.GetByUsername(data.Username)
	if err != nil {
		return nil, err
	}

	if !s.verifyPassword(user.PasswordHash, data.Password) {
		return nil, errors.New("invalid password")
	}

	return s.authRes(user)
}

func (s *Auth) register(data *dtos.Register) (*models.User, error) {
	if data.Email == "" || data.Username == "" {
		return nil, errors.New("username or email cannot be empty")
	}

	if data.Password == "" {
		return nil, errors.New("password cannot be empty")
	}

	if !s.validateEmail(data.Email) {
		return nil, errors.New("invalid email format")
	}

	if s.Repo.Users.ExistsByEmail(data.Email) {
		return nil, errors.New("email already exists")
	}

	if s.Repo.Users.ExistsByUsername(data.Username) {
		return nil, errors.New("username already exists")
	}

	passwordHash, err := s.hashPassword(data.Password)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Username:     data.Username,
		PasswordHash: passwordHash,
		Role:         data.Role,
		Email:        data.Email,
	}

	err = s.Repo.Users.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Auth) authRes(user *models.User) (*dtos.AuthRes, error) {

	token, err := s.JwtManager.Generate(dtos.JwtUser{
		Id:       user.ID,
		Username: user.Username,
		Role:     user.Role,
	})
	if err != nil {
		return nil, err
	}

	res := &dtos.AuthRes{
		User: dtos.User{
			ID:       user.ID,
			Username: user.Username,
			Role:     user.Role,
			Email:    user.Email,
		},
		Token: token,
	}

	return res, nil
}

func (s *Auth) validateEmail(email string) bool {
	var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	return emailRegex.MatchString(email)
}

func (s *Auth) hashPassword(password string) (string, error) {
	// Generate a hashed password using bcrypt with a cost factor of 12
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (s *Auth) verifyPassword(hashedPassword, password string) bool {
	// Compare the hashed password with the entered password
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil // returns true if password matches
}
