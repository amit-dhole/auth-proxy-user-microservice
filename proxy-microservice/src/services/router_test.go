package services

import (
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	apploader "github.com/auth-user-proxy-microservice/proxy-microservice/src/app-loader"
	"github.com/auth-user-proxy-microservice/proxy-microservice/src/config"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestHomeHandler_Success(t *testing.T) {

	WebServer := "127.0.0.1:8081"
	var jsonBlob = []byte(``)
	ts := StartMockWebServer(WebServer, http.StatusOK, jsonBlob)
	defer ts.Close()

	WebServer1 := "127.0.0.1:8083"
	ts1 := StartMockWebServer(WebServer1, http.StatusOK, jsonBlob)
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
				code: 500,
			},
		},
	}
	for _, test := range tests {
		req, err := http.NewRequest(http.MethodGet, "/", nil)

		assert.Nil(t, err)
		config.SetFilePath("../config.json")

		// initializing global Application Service
		apploader.LoadApplicationServices()

		rw := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc(
			"/",
			HomeHandler,
		).Methods(http.MethodGet)
		router.ServeHTTP(rw, req)
		assert.Equal(t, rw.Code, test.code, test.name)
	}
}

func TestHomeHandler_Error(t *testing.T) {

	WebServer := "127.0.0.1:8081"
	ts := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var jsonBlob []byte
		Query := r.URL.String()
		if strings.Contains(Query, "/auth") {
			jsonBlob = []byte(``)
		}
		w.WriteHeader(http.StatusInternalServerError)
		WriteJSON(t, w, jsonBlob)
	}))
	l, lErr := net.Listen("tcp", WebServer)
	if lErr != nil {
		t.Fatal(lErr)
	}
	ts.Listener = l
	ts.Start()
	defer ts.Close()

	WebServer1 := "127.0.0.1:8083"
	var jsonBlob = []byte(``)
	ts1 := StartMockWebServer(WebServer1, http.StatusInternalServerError, jsonBlob)
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
		req, err := http.NewRequest(http.MethodGet, "/", nil)

		assert.Nil(t, err)
		config.SetFilePath("../config.json")

		// initializing global Application Service
		apploader.LoadApplicationServices()

		rw := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc(
			"/",
			HomeHandler,
		).Methods(http.MethodGet)
		router.ServeHTTP(rw, req)
		assert.Equal(t, rw.Code, test.code, test.name)
	}
}

func TestHomeHandler_User_Success(t *testing.T) {

	WebServer := "127.0.0.1:8081"
	ts := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var jsonBlob []byte
		Query := r.URL.String()
		if strings.Contains(Query, "/auth") {
			jsonBlob = []byte(``)
		}
		w.WriteHeader(http.StatusOK)
		WriteJSON(t, w, jsonBlob)
	}))
	l, lErr := net.Listen("tcp", WebServer)
	if lErr != nil {
		t.Fatal(lErr)
	}
	ts.Listener = l
	ts.Start()
	defer ts.Close()

	WebServer1 := "127.0.0.1:8083"
	ts1 := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var jsonBlob []byte
		Query := r.URL.String()
		if strings.Contains(Query, "/user/profile") {
			jsonBlob = []byte(`{
				"name": "root",
				"dob": "30-nov-2015",
				"age": "5",
				"email": "dd@dd.com",
				"phone-number": "9912345678"
			   }`)
		}
		w.WriteHeader(http.StatusOK)
		WriteJSON(t, w, jsonBlob)
	}))
	l, lErr1 := net.Listen("tcp", WebServer1)
	if lErr1 != nil {
		t.Fatal(lErr)
	}
	ts1.Listener = l
	ts1.Start()
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
		req, err := http.NewRequest(http.MethodGet, "/", nil)

		assert.Nil(t, err)
		config.SetFilePath("../config.json")

		// initializing global Application Service
		apploader.LoadApplicationServices()

		rw := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc(
			"/",
			HomeHandler,
		).Methods(http.MethodGet)
		router.ServeHTTP(rw, req)
		assert.Equal(t, rw.Code, test.code, test.name)
	}
}

func TestHomeHandler_InValid_User_Success(t *testing.T) {

	WebServer := "127.0.0.1:8081"
	ts := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var jsonBlob []byte
		Query := r.URL.String()
		if strings.Contains(Query, "/auth") {
			jsonBlob = []byte(``)
		}
		w.WriteHeader(http.StatusOK)
		WriteJSON(t, w, jsonBlob)
	}))
	l, lErr := net.Listen("tcp", WebServer)
	if lErr != nil {
		t.Fatal(lErr)
	}
	ts.Listener = l
	ts.Start()
	defer ts.Close()

	WebServer1 := "127.0.0.1:8083"
	ts1 := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var jsonBlob []byte
		Query := r.URL.String()
		if strings.Contains(Query, "/user/profile") {
			jsonBlob = []byte(``)
		}
		w.WriteHeader(http.StatusInternalServerError)
		WriteJSON(t, w, jsonBlob)
	}))
	l, lErr1 := net.Listen("tcp", WebServer1)
	if lErr1 != nil {
		t.Fatal(lErr)
	}
	ts1.Listener = l
	ts1.Start()
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
		req, err := http.NewRequest(http.MethodGet, "/", nil)

		assert.Nil(t, err)
		config.SetFilePath("../config.json")

		// initializing global Application Service
		apploader.LoadApplicationServices()

		rw := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc(
			"/",
			HomeHandler,
		).Methods(http.MethodGet)
		router.ServeHTTP(rw, req)
		assert.Equal(t, rw.Code, test.code, test.name)
	}
}

func TestHomeHandler_ServiceName(t *testing.T) {

	WebServer := "127.0.0.1:8081"
	ts := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var jsonBlob []byte
		Query := r.URL.String()
		if strings.Contains(Query, "/auth") {
			jsonBlob = []byte(``)
		}
		w.WriteHeader(http.StatusInternalServerError)
		WriteJSON(t, w, jsonBlob)
	}))
	l, lErr := net.Listen("tcp", WebServer)
	if lErr != nil {
		t.Fatal(lErr)
	}
	ts.Listener = l
	ts.Start()
	defer ts.Close()

	WebServer1 := "127.0.0.1:8083"
	ts1 := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var jsonBlob []byte
		Query := r.URL.String()
		if strings.Contains(Query, "/microservice/name") {
			jsonBlob = []byte(``)
		}
		w.WriteHeader(http.StatusOK)
		WriteJSON(t, w, jsonBlob)
	}))
	l, lErr1 := net.Listen("tcp", WebServer1)
	if lErr1 != nil {
		t.Fatal(lErr)
	}
	ts1.Listener = l
	ts1.Start()
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
				data: "user-microservice",
			},
			expected: expected{
				code: 500,
			},
		},
	}
	for _, test := range tests {
		req, err := http.NewRequest(http.MethodGet, "/", nil)

		assert.Nil(t, err)
		config.SetFilePath("../config.json")

		// initializing global Application Service
		apploader.LoadApplicationServices()

		rw := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc(
			"/",
			HomeHandler,
		).Methods(http.MethodGet)
		router.ServeHTTP(rw, req)
		assert.Equal(t, rw.Code, test.code, test.name)
	}
}

func TestHomeHandler_Service_Success(t *testing.T) {

	WebServer := "127.0.0.1:8081"
	ts := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var jsonBlob []byte
		Query := r.URL.String()
		if strings.Contains(Query, "/auth") {
			jsonBlob = []byte(``)
		}
		w.WriteHeader(http.StatusInternalServerError)
		WriteJSON(t, w, jsonBlob)
	}))
	l, lErr := net.Listen("tcp", WebServer)
	if lErr != nil {
		t.Fatal(lErr)
	}
	ts.Listener = l
	ts.Start()
	defer ts.Close()

	WebServer1 := "127.0.0.1:8083"
	ts1 := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var jsonBlob []byte
		Query := r.URL.String()
		if strings.Contains(Query, "/microservice/name") {
			jsonBlob = []byte(`{
				"name": "user-microservice"
			   }`)
		}
		w.WriteHeader(http.StatusOK)
		WriteJSON(t, w, jsonBlob)
	}))
	l, lErr1 := net.Listen("tcp", WebServer1)
	if lErr1 != nil {
		t.Fatal(lErr)
	}
	ts1.Listener = l
	ts1.Start()
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
				data: "user-microservice",
			},
			expected: expected{
				code: 200,
			},
		},
	}
	for _, test := range tests {
		req, err := http.NewRequest(http.MethodGet, "/", nil)

		assert.Nil(t, err)
		config.SetFilePath("../config.json")

		// initializing global Application Service
		apploader.LoadApplicationServices()

		rw := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc(
			"/",
			HomeHandler,
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

//WriteJSON : Write Json of unit test case
func WriteJSON(t *testing.T, w io.Writer, reqJSON []byte) {
	_, wErr := w.Write(reqJSON)
	if wErr != nil {
		t.Fatal(wErr)
	}
}

func TestNewRouter(t *testing.T) {
	router := NewRouter()
	if router == nil {
		t.Error("expected initialized router but got nil")
	}
}
