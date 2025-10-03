package auth_test

import (
	"errors"
	"myapp/internal/auth"
	"myapp/internal/auth/moks"
	"testing"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

func Test_GetToken(t *testing.T) {
	tests := []struct {
		nameTest   string
		login      string
		passwordDb string
		password   string
		errorDB    error
	}{
		{
			nameTest:   "valid",
			login:      "pampam6752",
			passwordDb: "popakaki1",
			password:   "popakaki1",
			errorDB:    nil,
		},
		{
			nameTest:   "not found user",
			login:      "pampuk",
			password:   "pampamparam",
			passwordDb: "popakaki1",
			errorDB:    errors.New("the user was not found"),
		},
		{
			nameTest:   "invalid password",
			login:      "pammaram",
			passwordDb: "Pampuk",
			password:   "pampuk",
			errorDB:    errors.New("an empty password"),
		},
	}
	var testJwtKey = "lfpsdlfpsdflpffds[]aqwr32223sdfsflsp[fpsfsjhgbv]"
	for _, tt := range tests {
		t.Run(tt.nameTest, func(t *testing.T) {
			mockRapo := new(moks.MockAuthRepository)
			servise := auth.NewAuthService(mockRapo, testJwtKey)
			if tt.nameTest == "valid" || tt.nameTest == "invalid password" {
				bytes, _ := bcrypt.GenerateFromPassword([]byte(tt.passwordDb), bcrypt.DefaultCost)
				mockRapo.On("GetHashPassworld", tt.login).Return(string(bytes), nil)
			} else {
				mockRapo.On("GetHashPassworld", tt.login).Return("", errors.New("the user was not found"))
			}
			token, err := servise.GetToken(tt.login, tt.password)
			if tt.nameTest == "valid" {
				assert.NoError(t, err)
				parsedToken, _ := jwt.Parse(token, func(t *jwt.Token) (any, error) {
					return []byte(testJwtKey), nil
				})
				claims := parsedToken.Claims.(jwt.MapClaims)
				assert.Equal(t, tt.login, claims["sub"])
			}
			if tt.nameTest == "not found user" {
				assert.Error(t, err)
				assert.Equal(t, err.Error(), tt.errorDB.Error())
			}
			if tt.nameTest == "invalid password" {
				assert.Error(t, err)
				assert.Equal(t, err.Error(), tt.errorDB.Error())
				assert.Equal(t, token, "")
			}
		})
	}
}
func Test_Register(t *testing.T) {
	tests := []struct {
		nameTest   string
		login      string
		pasword    string
		shauldCall bool
		errors     error
	}{
		{
			nameTest:   "valid",
			login:      "patato",
			pasword:    "pam128",
			shauldCall: true,
			errors:     nil,
		},
		{
			nameTest:   "empty registration fields",
			login:      "",
			pasword:    "",
			shauldCall: false,
			errors:     errors.New("an empty username or password value"),
		},
		{
			nameTest:   "an empty password",
			login:      "makson",
			pasword:    "",
			shauldCall: false,
			errors:     errors.New("an empty username or password value"),
		},
		{
			nameTest:   "an emty password",
			login:      "",
			pasword:    "pampum",
			shauldCall: false,
			errors:     errors.New("an empty username or password value"),
		},
	}
	var testJwtKey = "lfpsdlfpsdflpffds[]aqwr32223sdfsflsp[fpsfsjhgbv]"
	for _, tt := range tests {
		t.Run(tt.nameTest, func(t *testing.T) {
			mockRapo := new(moks.MockAuthRepository)
			servise := auth.NewAuthService(mockRapo, testJwtKey)
			if tt.shauldCall {
				mockRapo.On("CreateUser", tt.login, mock.Anything).Return(nil)
			} else {
				mockRapo.On("CreateUser", tt.login, mock.Anything).Return(tt.errors)
			}

			err := servise.Register(tt.login, tt.pasword)
			if tt.errors != nil {
				assert.Equal(t, err.Error(), tt.errors.Error())
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			if tt.shauldCall {
				mockRapo.AssertCalled(t, "CreateUser", tt.login, mock.Anything)
			} else {
				mockRapo.AssertNotCalled(t, "CreateUser")
			}
		})
	}
}
