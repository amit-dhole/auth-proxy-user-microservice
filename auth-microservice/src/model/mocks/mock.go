// Code generated by MockGen. DO NOT EDIT.
// Source: gitlab.connectwisedev.com/platform/platform-common-lib/src/web (interfaces: Server,ServerFactory,Resource,RequestContext,HTTPServer,Router)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	http "net/http"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	mux "github.com/gorilla/mux"
)

// MockServer is a mock of Server interface
type MockServer struct {
	ctrl     *gomock.Controller
	recorder *MockServerMockRecorder
}

// MockServerMockRecorder is the mock recorder for MockServer
type MockServerMockRecorder struct {
	mock *MockServer
}

// NewMockServer creates a new mock instance
func NewMockServer(ctrl *gomock.Controller) *MockServer {
	mock := &MockServer{ctrl: ctrl}
	mock.recorder = &MockServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockServer) EXPECT() *MockServerMockRecorder {
	return m.recorder
}

// CreateServer mocks base method
func (m *MockServer) CreateServer() {
	m.ctrl.Call(m, "CreateServer")
}

// CreateServer indicates an expected call of CreateServer
func (mr *MockServerMockRecorder) CreateServer() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateServer", reflect.TypeOf((*MockServer)(nil).CreateServer))
}

// GetRouter mocks base method
func (m *MockServer) GetRouter() *mux.Router {
	ret := m.ctrl.Call(m, "GetRouter")
	ret0, _ := ret[0].(*mux.Router)
	return ret0
}

// GetRouter indicates an expected call of GetRouter
func (mr *MockServerMockRecorder) GetRouter() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRouter", reflect.TypeOf((*MockServer)(nil).GetRouter))
}

// HTTP2ListenAndServeTLS mocks base method
func (m *MockServer) HTTP2ListenAndServeTLS() error {
	ret := m.ctrl.Call(m, "HTTP2ListenAndServeTLS")
	ret0, _ := ret[0].(error)
	return ret0
}

// HTTP2ListenAndServeTLS indicates an expected call of HTTP2ListenAndServeTLS
func (mr *MockServerMockRecorder) HTTP2ListenAndServeTLS() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HTTP2ListenAndServeTLS", reflect.TypeOf((*MockServer)(nil).HTTP2ListenAndServeTLS))
}

// ListenAndServe mocks base method
func (m *MockServer) ListenAndServe() error {
	ret := m.ctrl.Call(m, "ListenAndServe")
	ret0, _ := ret[0].(error)
	return ret0
}

// ListenAndServe indicates an expected call of ListenAndServe
func (mr *MockServerMockRecorder) ListenAndServe() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListenAndServe", reflect.TypeOf((*MockServer)(nil).ListenAndServe))
}

// RegisterOnShutdown mocks base method
func (m *MockServer) RegisterOnShutdown(arg0 func()) {
	m.ctrl.Call(m, "RegisterOnShutdown", arg0)
}

// RegisterOnShutdown indicates an expected call of RegisterOnShutdown
func (mr *MockServerMockRecorder) RegisterOnShutdown(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterOnShutdown", reflect.TypeOf((*MockServer)(nil).RegisterOnShutdown), arg0)
}

// SetRouter mocks base method
func (m *MockServer) SetRouter(arg0 *mux.Router) {
	m.ctrl.Call(m, "SetRouter", arg0)
}

// SetRouter indicates an expected call of SetRouter
func (mr *MockServerMockRecorder) SetRouter(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetRouter", reflect.TypeOf((*MockServer)(nil).SetRouter), arg0)
}

// ShutDown mocks base method
func (m *MockServer) ShutDown(arg0 context.Context) error {
	ret := m.ctrl.Call(m, "ShutDown", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ShutDown indicates an expected call of ShutDown
func (mr *MockServerMockRecorder) ShutDown(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShutDown", reflect.TypeOf((*MockServer)(nil).ShutDown), arg0)
}

// MockServerFactory is a mock of ServerFactory interface
type MockServerFactory struct {
	ctrl     *gomock.Controller
	recorder *MockServerFactoryMockRecorder
}

// MockServerFactoryMockRecorder is the mock recorder for MockServerFactory
type MockServerFactoryMockRecorder struct {
	mock *MockServerFactory
}

// NewMockServerFactory creates a new mock instance
func NewMockServerFactory(ctrl *gomock.Controller) *MockServerFactory {
	mock := &MockServerFactory{ctrl: ctrl}
	mock.recorder = &MockServerFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockServerFactory) EXPECT() *MockServerFactoryMockRecorder {
	return m.recorder
}

// MockResource is a mock of Resource interface
type MockResource struct {
	ctrl     *gomock.Controller
	recorder *MockResourceMockRecorder
}

// MockResourceMockRecorder is the mock recorder for MockResource
type MockResourceMockRecorder struct {
	mock *MockResource
}

// NewMockResource creates a new mock instance
func NewMockResource(ctrl *gomock.Controller) *MockResource {
	mock := &MockResource{ctrl: ctrl}
	mock.recorder = &MockResourceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockResource) EXPECT() *MockResourceMockRecorder {
	return m.recorder
}

// MockRequestContext is a mock of RequestContext interface
type MockRequestContext struct {
	ctrl     *gomock.Controller
	recorder *MockRequestContextMockRecorder
}

// MockRequestContextMockRecorder is the mock recorder for MockRequestContext
type MockRequestContextMockRecorder struct {
	mock *MockRequestContext
}

// NewMockRequestContext creates a new mock instance
func NewMockRequestContext(ctrl *gomock.Controller) *MockRequestContext {
	mock := &MockRequestContext{ctrl: ctrl}
	mock.recorder = &MockRequestContextMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRequestContext) EXPECT() *MockRequestContextMockRecorder {
	return m.recorder
}

// GetData mocks base method
func (m *MockRequestContext) GetData() ([]byte, error) {
	ret := m.ctrl.Call(m, "GetData")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetData indicates an expected call of GetData
func (mr *MockRequestContextMockRecorder) GetData() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetData", reflect.TypeOf((*MockRequestContext)(nil).GetData))
}

// GetRemoteAddr mocks base method
func (m *MockRequestContext) GetRemoteAddr() (string, error) {
	ret := m.ctrl.Call(m, "GetRemoteAddr")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRemoteAddr indicates an expected call of GetRemoteAddr
func (mr *MockRequestContextMockRecorder) GetRemoteAddr() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRemoteAddr", reflect.TypeOf((*MockRequestContext)(nil).GetRemoteAddr))
}

// GetRequest mocks base method
func (m *MockRequestContext) GetRequest() *http.Request {
	ret := m.ctrl.Call(m, "GetRequest")
	ret0, _ := ret[0].(*http.Request)
	return ret0
}

// GetRequest indicates an expected call of GetRequest
func (mr *MockRequestContextMockRecorder) GetRequest() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRequest", reflect.TypeOf((*MockRequestContext)(nil).GetRequest))
}

// GetRequestDcDateTimeUTC mocks base method
func (m *MockRequestContext) GetRequestDcDateTimeUTC() time.Time {
	ret := m.ctrl.Call(m, "GetRequestDcDateTimeUTC")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// GetRequestDcDateTimeUTC indicates an expected call of GetRequestDcDateTimeUTC
func (mr *MockRequestContextMockRecorder) GetRequestDcDateTimeUTC() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRequestDcDateTimeUTC", reflect.TypeOf((*MockRequestContext)(nil).GetRequestDcDateTimeUTC))
}

// GetResponse mocks base method
func (m *MockRequestContext) GetResponse() http.ResponseWriter {
	ret := m.ctrl.Call(m, "GetResponse")
	ret0, _ := ret[0].(http.ResponseWriter)
	return ret0
}

// GetResponse indicates an expected call of GetResponse
func (mr *MockRequestContextMockRecorder) GetResponse() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResponse", reflect.TypeOf((*MockRequestContext)(nil).GetResponse))
}

// GetVars mocks base method
func (m *MockRequestContext) GetVars() map[string]string {
	ret := m.ctrl.Call(m, "GetVars")
	ret0, _ := ret[0].(map[string]string)
	return ret0
}

// GetVars indicates an expected call of GetVars
func (mr *MockRequestContextMockRecorder) GetVars() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVars", reflect.TypeOf((*MockRequestContext)(nil).GetVars))
}

// MockHTTPServer is a mock of HTTPServer interface
type MockHTTPServer struct {
	ctrl     *gomock.Controller
	recorder *MockHTTPServerMockRecorder
}

// MockHTTPServerMockRecorder is the mock recorder for MockHTTPServer
type MockHTTPServerMockRecorder struct {
	mock *MockHTTPServer
}

// NewMockHTTPServer creates a new mock instance
func NewMockHTTPServer(ctrl *gomock.Controller) *MockHTTPServer {
	mock := &MockHTTPServer{ctrl: ctrl}
	mock.recorder = &MockHTTPServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHTTPServer) EXPECT() *MockHTTPServerMockRecorder {
	return m.recorder
}

// RegisterOnShutdown mocks base method
func (m *MockHTTPServer) RegisterOnShutdown(arg0 func()) {
	m.ctrl.Call(m, "RegisterOnShutdown", arg0)
}

// RegisterOnShutdown indicates an expected call of RegisterOnShutdown
func (mr *MockHTTPServerMockRecorder) RegisterOnShutdown(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterOnShutdown", reflect.TypeOf((*MockHTTPServer)(nil).RegisterOnShutdown), arg0)
}

// ShutDown mocks base method
func (m *MockHTTPServer) ShutDown(arg0 context.Context) error {
	ret := m.ctrl.Call(m, "ShutDown", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ShutDown indicates an expected call of ShutDown
func (mr *MockHTTPServerMockRecorder) ShutDown(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShutDown", reflect.TypeOf((*MockHTTPServer)(nil).ShutDown), arg0)
}

// Start mocks base method
func (m *MockHTTPServer) Start(arg0 context.Context) error {
	ret := m.ctrl.Call(m, "Start", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start
func (mr *MockHTTPServerMockRecorder) Start(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockHTTPServer)(nil).Start), arg0)
}

// MockRouter is a mock of Router interface
type MockRouter struct {
	ctrl     *gomock.Controller
	recorder *MockRouterMockRecorder
}

// MockRouterMockRecorder is the mock recorder for MockRouter
type MockRouterMockRecorder struct {
	mock *MockRouter
}

// NewMockRouter creates a new mock instance
func NewMockRouter(ctrl *gomock.Controller) *MockRouter {
	mock := &MockRouter{ctrl: ctrl}
	mock.recorder = &MockRouterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRouter) EXPECT() *MockRouterMockRecorder {
	return m.recorder
}

// AddHandle mocks base method
func (m *MockRouter) AddHandle(arg0 string, arg1 http.Handler, arg2 ...string) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "AddHandle", varargs...)
}

// AddHandle indicates an expected call of AddHandle
func (mr *MockRouterMockRecorder) AddHandle(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddHandle", reflect.TypeOf((*MockRouter)(nil).AddHandle), varargs...)
}

// ServeHTTP mocks base method
func (m *MockRouter) ServeHTTP(arg0 http.ResponseWriter, arg1 *http.Request) {
	m.ctrl.Call(m, "ServeHTTP", arg0, arg1)
}

// ServeHTTP indicates an expected call of ServeHTTP
func (mr *MockRouterMockRecorder) ServeHTTP(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServeHTTP", reflect.TypeOf((*MockRouter)(nil).ServeHTTP), arg0, arg1)
}

// Use mocks base method
func (m *MockRouter) Use(arg0 ...func(http.Handler) http.Handler) {
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Use", varargs...)
}

// Use indicates an expected call of Use
func (mr *MockRouterMockRecorder) Use(arg0 ...interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Use", reflect.TypeOf((*MockRouter)(nil).Use), arg0...)
}