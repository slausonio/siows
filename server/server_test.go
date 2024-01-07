package server

import (
	"fmt"
	"github.com/slausonio/go-webserver/environment"
	"github.com/slausonio/go-webserver/server/mocks"
	"github.com/slausonio/siotest"
	"net/http"
	"testing"
)

type handler struct{}

var HappyEnvMap = environment.SioGoEnv{
	environment.CurrentEnv: "test",
	environment.AppName:    "go-webserver",
	environment.Port:       "8080"}

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

	return &Server{env: mockEnv}, mockEnv

}

func TestServer_Start(t *testing.T) {
	t.Run("Happy", func(t *testing.T) {
		testServer, _ := createTestServerStruct(t)

		//mockEnv.On("Value", environment.Port).Return("test")

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
