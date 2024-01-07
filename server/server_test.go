package server

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"

	"github.com/slausonio/siotest"
	"github.com/slausonio/siows/environment"
	"github.com/slausonio/siows/server/mocks"
)

type handler struct{}

var HappyEnvMap = environment.SioGoEnv{
	environment.CurrentEnvKey: "test",
	environment.AppNameKey:    "go-webserver",
	environment.PortKey:       "8080"}

var CurrentEnvMap = environment.SioGoEnv{"test1": "test", "test2": "test2"}

func EnvSetup(t *testing.T) {
	t.Helper()

	siotest.CreateFile(t, environment.DefaultFilePath)
	siotest.CreateFile(t, fmt.Sprintf(environment.CurrentEnvFilePath, "test"))

	siotest.WriteEnvToFile(t, environment.DefaultFilePath, HappyEnvMap)
	siotest.WriteEnvToFile(t, fmt.Sprintf(environment.CurrentEnvFilePath, "test"), CurrentEnvMap)

}

func EnvCleanup(t *testing.T) {
	t.Helper()

	t.Cleanup(func() {
		siotest.RemoveFileAndDirs(t, environment.DefaultFilePath)
		siotest.RemoveFileAndDirs(t, fmt.Sprintf(environment.CurrentEnvFilePath, "test"))
	})
}

func createTestServerStruct(t *testing.T) (*Server, *mocks.Env) {
	mockEnv := mocks.NewEnv(t)

	return &Server{env: mockEnv, server: &http.Server{}}, mockEnv

}

func TestServer_Start(t *testing.T) {
	t.Run("Happy", func(t *testing.T) {
		testServer, _ := createTestServerStruct(t)

		//mockEnv.On("Value", environment.PortKey).Return("test")

		h := http.NewServeMux()

		go testServer.Start(h)

		//t.Cleanup(testServer.Kill)
	})

	tests := []struct {
		name        string
		environment environment.SioGoEnv
	}{
		{
			name: "basic",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				env: environment.NewEnvironment(),
			}
			h := http.NewServeMux()

			// This will start a Server which is not ideal but for now we are just ensuring it doesn't crash.
			go s.Start(h)
		})
	}
}
func TestServer_Kill(t *testing.T) {
	t.Run("happy", func(t *testing.T) {
		testServer, mockEnv := createTestServerStruct(t)
		mockEnv.On("Value", environment.PortKey).Return("8080")

		h := http.NewServeMux()

		testServer.Start(h)
		assert.NotPanics(t, testServer.Kill, "expected server kill to not panic")
	})

	t.Run("panics", func(t *testing.T) {
		t.Run("not started", func(t *testing.T) {
			testServer, _ := createTestServerStruct(t)

			testServer.Kill()
			assert.Panics(t, testServer.Kill, "expected server kill to panic")
		})

		t.Run("already closed", func(t *testing.T) {
			testServer, mockEnv := createTestServerStruct(t)
			mockEnv.On("Value", environment.PortKey).Return("8080")

			h := http.NewServeMux()

			testServer.Start(h)
			testServer.Kill()

			assert.Panics(t, testServer.Kill, "expected server kill to panic")
		})
	})
}

//func TestServer_printInfo(t *testing.T) {
//	tests := []struct {
//		name string
//	}{
//		{"basic"},
//		// TODO Add additional test cases when possible
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			s := &Server{
//				e: environment.NewEnvironment(),
//			}
//			s.printInfo(0)
//		})
//	}
//}
//
//func TestServer_printSio(t *testing.T) {
//	tests := []struct {
//		name string
//	}{
//		{"basic"},
//		// TODO Add additional test cases when possible
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			s := &Server{
//				e: environment.NewEnvironment(),
//			}
//			s.printSio()
//		})
//	}
//}
//
//func TestNewServer(t *testing.T) {
//	tests := []struct {
//		name string
//	}{
//		{"basic"},
//		// TODO Add additional test cases when possible
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := NewServer(); got == nil {
//				t.Errorf("NewServer() = %v, expected non-nil", got)
//			}
//		})
//	}
//}
