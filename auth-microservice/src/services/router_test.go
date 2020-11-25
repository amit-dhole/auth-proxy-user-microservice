package services

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestAuthHandler(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	reqContMock := func() *cwmock.MockRequestContext {
		rcmock := cwmock.NewMockRequestContext(ctrl)
		rcmock.EXPECT().GetResponse().Return(httptest.NewRecorder()).AnyTimes()
		req, _ := http.NewRequest("", "", bytes.NewReader([]byte("")))
		rcmock.EXPECT().GetRequest().Return(req).AnyTimes()
		return rcmock
	}()

	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AuthHandler(tt.args.w, tt.args.r)
		})
	}
}

func TestNewRouter(t *testing.T) {
	router := NewRouter()
	if router == nil {
		t.Error("expected initialized router but got nil")
	}
}
