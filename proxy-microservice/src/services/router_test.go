package services

import (
	"net/http"
	"testing"
)

func TestHomeHandler(t *testing.T) {
	type args struct {
		w   http.ResponseWriter
		in1 *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "default",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			HomeHandler(tt.args.w, tt.args.in1)
		})
	}
}
