package safe

import (
	"errors"
	"testing"
)

func TestTry(t *testing.T) {
	tests := []struct {
		name    string
		args    func()
		wantErr bool
	}{
		{name: "panic", args: func() {
			panic("panic")
		}, wantErr: true},
		{name: "not problem", args: func() {
		}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Try(tt.args); (err != nil) != tt.wantErr {
				t.Errorf("TryE() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTryE(t *testing.T) {
	tests := []struct {
		name    string
		args    func() error
		wantErr bool
	}{
		{name: "panic", args: func() error {
			panic("panic")
		}, wantErr: true},
		{name: "error", args: func() error {
			return errors.New("error")
		}, wantErr: true},
		{name: "not problem", args: func() error {
			return nil
		}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := TryE(tt.args); (err != nil) != tt.wantErr {
				t.Errorf("TryE() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
