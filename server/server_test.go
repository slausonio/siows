package server

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/slausonio/siotest"
	"github.com/slausonio/siows/environment"
)

type handler struct{}

var HappyEnvMap = environment.SioWSEnv{
	environment.CurrentEnvKey: "test",
	environment.AppNameKey:    "go-webserver",
	environment.PortKey:       "8080"}

var CurrentEnvMap = environment.SioWSEnv{"test1": "test", "test2": "test2"}

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

func createTestServerStruct(t *testing.T, env environment.SioWSEnv) *Server {
	return &Server{env: env, server: &http.Server{}}
}

func TestServer_Start(t *testing.T) {
	t.Run("Happy", func(t *testing.T) {
		testServer := createTestServerStruct(t, HappyEnvMap)

		//mockEnv.On("Value", environment.PortKey).Return("test")

		h := http.NewServeMux()

		testServer.Start(h)

		//t.Cleanup(testServer.Kill)
	})

	t.Run("panics", func(t *testing.T) {
		t.Skip()

		t.Run("duplicate address", func(t *testing.T) {
			testServer := createTestServerStruct(t, HappyEnvMap)

			h := http.NewServeMux()

			testServer.Start(h)
			testServer.Start(h)

			assert.Panics(t, testServer.Kill, "expected server kill to panic")
		})
	})
}
