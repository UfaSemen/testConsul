package main

import (
	"testing"
)

func Test_serverStart(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "no consul running",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := serverStart(); (err != nil) != tt.wantErr {
				t.Errorf("serverStart() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
