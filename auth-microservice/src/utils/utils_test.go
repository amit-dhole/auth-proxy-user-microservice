package utils

import (
	"net/http"
	"testing"
)

func TestToString(t *testing.T) {

	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "defalt",
			args: args{
				v: "Abc",
			},
			want: "Abc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToString(tt.args.v); got != tt.want {
				t.Errorf("ToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerateResponse(t *testing.T) {
	type args struct {
		data []byte
		code int
		w    http.ResponseWriter
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GenerateResponse(tt.args.data, tt.args.code, tt.args.w)
		})
	}
}
