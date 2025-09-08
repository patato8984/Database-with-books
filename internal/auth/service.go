package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	authRapo *AuthRepository
	jwtKye   string
}

func NewAuthService(authRapo *AuthRepository, kye string) *AuthService {
	return &AuthService{authRapo: authRapo, jwtKye: kye}
}

func (s *AuthService) Register(login, password string) error {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	if er := s.authRapo.CreateUser(login, string(hashPassword)); er != nil {
		return er
	}
	return nil
}
func (s *AuthService) GetToken(login, password string) (string, error) {
	hashPassword, err := s.authRapo.GetHashPassworld(login)
	if err != nil {
		return "", err
	}
	if er := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password)); er != nil {
		return "", er
	}
	clearTocen := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{"sub": login, "exp": time.Now().Add(time.Hour * 72).Unix()})
	tocen, errr := clearTocen.SignedString([]byte(s.jwtKye))
	if errr != nil {
		return "", errr
	}
	return tocen, nil
}
