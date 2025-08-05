package auth_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/bootdotdev/learn-cicd-starter/internal/auth"
	"github.com/stretchr/testify/assert"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name           string
		inputHeaders   string
		setHeader      bool
		expectedString string
		expectedError  error
	}{
		{
			name:           "successfully get api key",
			inputHeaders:   "ApiKey testKey",
			setHeader:      true,
			expectedString: "testKey",
			expectedError:  nil,
		},
		{
			name:           "fail with no auth header",
			inputHeaders:   "",
			setHeader:      false,
			expectedString: "",
			expectedError:  auth.ErrNoAuthHeaderIncluded,
		},
		{
			name:           "fail with malformed auth header",
			inputHeaders:   "ApiKey",
			setHeader:      true,
			expectedString: "",
			expectedError:  errors.New("malformed authorization header"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			headers := http.Header{}
			if tt.setHeader {
				headers.Set("Authorization", tt.inputHeaders)
			}

			output, err := auth.GetAPIKey(headers)
			assert.Equal(t, tt.expectedString, output)
			assert.Equal(t, tt.expectedError, err)
		})
	}
}
