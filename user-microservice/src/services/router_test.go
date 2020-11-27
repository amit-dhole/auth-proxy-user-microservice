package services

import (
	"net"
	"net/http"
	"net/http/httptest"
	"testing"

	apploader "github.com/auth-user-proxy-microservice/user-microservice/src/app-loader"
	"github.com/auth-user-proxy-microservice/user-microservice/src/config"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestHomeHandler_Success(t *testing.T) {

	WebServer := "127.0.0.1:8083"
	var jsonBlob = []byte(``)
	ts1 := StartMockWebServer(WebServer, http.StatusOK, jsonBlob)
	defer ts1.Close()

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
			name: "default",
			payload: payload{
				data: nil,
			},
			expected: expected{
				code: 200,
			},
		},
	}
	for _, test := range tests {
		req, err := http.NewRequest(http.MethodGet, "/user/profile", nil)

		assert.Nil(t, err)
		config.SetFilePath("../config.json")

		// initializing global Application Service
		apploader.LoadApplicationServices()

		rw := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc(
			"/user/profile",
			UserProfileHandler,
		).Methods(http.MethodGet)
		router.ServeHTTP(rw, req)
		assert.Equal(t, rw.Code, test.code, test.name)
	}
}

func TestHomeHandler_Success_withUser(t *testing.T) {

	WebServer := "127.0.0.1:8083"
	var jsonBlob = []byte(``)
	ts1 := StartMockWebServer(WebServer, http.StatusOK, jsonBlob)
	defer ts1.Close()

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
			name: "default",
			payload: payload{
				data: nil,
			},
			expected: expected{
				code: 200,
			},
		},
	}
	for _, test := range tests {
		req, err := http.NewRequest(http.MethodGet, "/user/profile", nil)
		req.Header.Add("user", "root")
		assert.Nil(t, err)
		config.SetFilePath("../config.json")

		// initializing global Application Service
		apploader.LoadApplicationServices()

		rw := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc(
			"/user/profile",
			UserProfileHandler,
		).Methods(http.MethodGet)
		router.ServeHTTP(rw, req)
		assert.Equal(t, rw.Code, test.code, test.name)
	}
}

func TestServiceNameHandler_ServiceName(t *testing.T) {

	WebServer := "127.0.0.1:8083"
	var jsonBlob = []byte(``)
	ts1 := StartMockWebServer(WebServer, http.StatusOK, jsonBlob)
	defer ts1.Close()

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
			name: "default",
			payload: payload{
				data: nil,
			},
			expected: expected{
				code: 200,
			},
		},
	}
	for _, test := range tests {
		req, err := http.NewRequest(http.MethodGet, "/microservice/name", nil)
		assert.Nil(t, err)
		config.SetFilePath("../config.json")

		// initializing global Application Service
		apploader.LoadApplicationServices()

		rw := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc(
			"/microservice/name",
			ServiceNameHandler,
		).Methods(http.MethodGet)
		router.ServeHTTP(rw, req)
		assert.Equal(t, rw.Code, test.code, test.name)
	}
}

//StartMockWebServer starts mocking for web server
func StartMockWebServer(WebServer string, StatusCode int, jsonBlob []byte) *httptest.Server {
	ts := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//var jsonBlob = []byte(`[{"vendor": "lmi","nodeId": "eb828cb7-3a87-11e9-a4de-02af0800762e","parentNodeId": "","level": 0,"attributes": "{\"companyidDesktop\": 11111,\"pskDesktop\": \"00_desktop\",\"companyidServer\": 22222,\"pskServer\": \"00_server\"}"}]`)
		w.WriteHeader(StatusCode)
		w.Write(jsonBlob)
	}))
	l, _ := net.Listen("tcp", WebServer)
	ts.Listener = l
	ts.Start()
	return ts
}

func TestNewRouter(t *testing.T) {
	router := NewRouter()
	if router == nil {
		t.Error("expected initialized router but got nil")
	}
}
