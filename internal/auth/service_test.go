package auth_test

import (
	"errors"
	"myapp/internal/auth"
	"myapp/internal/auth/moks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_Register(t *testing.T) {
	tests := []struct {
		login      string
		pasword    string
		shauldCall bool
		errors     error
	}{
		{
			login:      "patato",
			pasword:    "pam128",
			shauldCall: true,
			errors:     nil,
		},
		{
			login:      "",
			pasword:    "",
			shauldCall: false,
			errors:     errors.New("an empty username or password value"),
		},
		{
			login:      "makson",
			pasword:    "",
			shauldCall: false,
			errors:     errors.New("an empty username or password value"),
		},
		{
			login:      "",
			pasword:    "pampum",
			shauldCall: false,
			errors:     errors.New("an empty username or password value"),
		},
	}
	var testJwtKey = "lfpsdlfpsdflpffds[]aqwr32223sdfsflsp[fpsfsjhgbv]"
	for _, tt := range tests {
		t.Run("pops", func(t *testing.T) {
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
