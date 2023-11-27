package domain

import (
	"testing"
)

func TestUser(t *testing.T) {
	tests := []struct {
		name     string
		username string
		password string
		role     string
	}{
		{
			name:     "Create valid user",
			username: "testUser",
			password: "testPass",
			role:     "testRole",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user := User{
				Username: tt.username,
				Password: tt.password,
				Role:     tt.role,
			}

			if user.Username != tt.username {
				t.Errorf("Expected username to be '%v', got '%v'", tt.username, user.Username)
			}
			if user.Password != tt.password {
				t.Errorf("Expected password to be '%v', got '%v'", tt.password, user.Password)
			}
			if user.Role != tt.role {
				t.Errorf("Expected role to be '%v', got '%v'", tt.role, user.Role)
			}
		})
	}
}
