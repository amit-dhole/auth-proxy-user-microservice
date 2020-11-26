package services

import (
	"encoding/base64"
	"net/http"
	"net/http/httptest"
	"testing"

	apploader "github.com/auth-user-proxy-microservice/auth-microservice/src/app-loader"
	"github.com/auth-user-proxy-microservice/auth-microservice/src/config"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestAuthHandler(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	type payload struct {
		data interface{}
		err  error
	}

	type expected struct {
		body string
		code int
	}

	tests := []struct {
		name string
		expected
		payload
	}{
		{
			name: "successful response",
			payload: payload{
				data: nil,
			},
			expected: expected{
				code: 200,
			},
		},
	}
	for _, test := range tests {
		req, err := http.NewRequest(http.MethodGet, "/auth", nil)
		auth := "root"
		basicAuth := "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
		req.Header.Add("Proxy-Authorization", basicAuth)

		assert.Nil(t, err)
		config.SetFilePath("../config.json")

		// initializing global Application Service
		apploader.LoadApplicationServices()

		rw := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc(
			"/auth",
			AuthHandler,
		).Methods(http.MethodGet)
		router.ServeHTTP(rw, req)
		assert.Equal(t, rw.Code, test.code, test.name)
	}
}

func TestAuthHandler_Unauthorize(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	type payload struct {
		data interface{}
		err  error
	}

	type expected struct {
		body string
		code int
	}

	tests := []struct {
		name string
		expected
		payload
	}{
		{
			name: "Unauthorize response",
			payload: payload{
				data: nil,
			},
			expected: expected{
				code: 401,
			},
		},
	}
	for _, test := range tests {
		req, err := http.NewRequest(http.MethodGet, "/auth", nil)

		assert.Nil(t, err)
		config.SetFilePath("../config.json")

		// initializing global Application Service
		apploader.LoadApplicationServices()

		rw := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc(
			"/auth",
			AuthHandler,
		).Methods(http.MethodGet)
		router.ServeHTTP(rw, req)
		assert.Equal(t, rw.Code, test.code, test.name)
	}
}
func TestNewRouter(t *testing.T) {
	router := NewRouter()
	if router == nil {
		t.Error("expected initialized router but got nil")
	}
}
