package server

import (
	"github.com/joho/godotenv"
	"github.com/slausonio/go-webserver/environment"
	"net/http"
	"os"
	"testing"
)

type handler struct{}

var happyEnv = environment.Environment{environment.AppName: "go-webserver",
	environment.Port: "8080"}

func setCurrentEnvForTest(t *testing.T) {
	t.Helper()

	err := os.Setenv(environment.CurrentEnv, "test")
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(os.Clearenv)
}

func writeEnvToFile(t *testing.T, env environment.Environment) {
	t.Helper()

	err := godotenv.Write(env, "env/.env")
	if err != nil {
		t.Fatal(err)
	}
}

func TestServer_Start(t *testing.T) {
	t.Run("Happy", func(t *testing.T) {
		setCurrentEnvForTest(t)
		writeEnvToFile(t, happyEnv)
		s := &Server{
			Environment: environment.NewEnvironment(),
		}
		h := http.NewServeMux()

		// This will start a Server which is not ideal but for now we are just ensuring it doesn't crash.
		go s.Start(h)
	})

	tests := []struct {
		name        string
		environment environment.Environment
	}{
		{
			name: "basic",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				Environment: environment.NewEnvironment(),
			}
			h := http.NewServeMux()

			// This will start a Server which is not ideal but for now we are just ensuring it doesn't crash.
			go s.Start(h)
		})
	}
}

func TestServer_printInfo(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"basic"},
		// TODO Add additional test cases when possible
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				Environment: environment.NewEnvironment(),
			}
			s.printInfo(0)
		})
	}
}

func TestServer_printSio(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"basic"},
		// TODO Add additional test cases when possible
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				Environment: environment.NewEnvironment(),
			}
			s.printSio()
		})
	}
}

func TestNewServer(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"basic"},
		// TODO Add additional test cases when possible
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewServer(); got == nil {
				t.Errorf("NewServer() = %v, expected non-nil", got)
			}
		})
	}
}
