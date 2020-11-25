package config

import "testing"

func TestLoad(t *testing.T) {
	SetFilePath("../config.json")
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "default",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Load(); (err != nil) != tt.wantErr {
				t.Errorf("Load() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
