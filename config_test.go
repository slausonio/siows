package siows

import (
	"fmt"
	"github.com/google/uuid"
	"testing"
)

func TestConfigMethods(t *testing.T) {
	tests := []struct {
		name         string
		inputConfig  Config
		inputID      string
		inputName    string
		inputPort    string
		expectedID   string
		expectedName string
		expectedPort string
	}{
		{
			name:         "Test 1",
			inputConfig:  NewConfig(map[string]string{}),
			inputID:      "12345",
			inputName:    "Test",
			inputPort:    "8080",
			expectedID:   "12345",
			expectedName: "Test",
			expectedPort: "8080",
		},
		{
			name:         "Test 2",
			inputConfig:  NewConfig(map[string]string{}).WithID(uuid.NewString()),
			inputID:      "",
			inputName:    "Test",
			inputPort:    "80",
			expectedID:   "",
			expectedName: "Test",
			expectedPort: "80",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.inputConfig.WithID(tt.inputID).WithName(tt.inputName).WithPort(tt.inputPort)

			if c.ID() != tt.expectedID {
				t.Errorf("Expected ID to be '%s', got '%s'", tt.expectedID, c.ID())
			}

			if c.Name() != tt.expectedName {
				t.Errorf("Expected Name to be '%s', got '%s'", tt.expectedName, c.Name())
			}

			if c.Port() != tt.expectedPort {
				t.Errorf("Expected Port to be '%s', got '%s'", tt.expectedPort, c.Port())
			}

			expectedFmtString := "ID: " + tt.expectedID + ", Name: " + tt.expectedName + ", Port: " + fmt.Sprint(tt.expectedPort)
			if c.FmtString() != expectedFmtString {
				t.Errorf("Expected Config string to be '%s', got '%s'", expectedFmtString, c.FmtString())
			}
		})
	}
}
