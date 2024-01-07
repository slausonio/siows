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

func TestServer_Start(t *testing.T) {
	t.Parallel()

	t.Run("Happy", func(t *testing.T) {
		testServer := NewServer(HappyEnvMap)
		h := http.NewServeMux()

		testServer.Start(h)
	})

	t.Run("panics", func(t *testing.T) {
		t.Skip()

		t.Run("duplicate address", func(t *testing.T) {
			testServer := NewServer(HappyEnvMap)

			h := http.NewServeMux()

			testServer.Start(h)

			assert.Panics(t,
				func() {
					testServer.Start(h)
				}, "expected server kill to panic")
		})
	})
}

func TestServer_getters(t *testing.T) {
	testServer := NewServer(HappyEnvMap)

	t.Run("server", func(t *testing.T) {
		h := http.NewServeMux()

		testServer.Start(h)
		assert.NotNil(t, testServer.Env(), "expected test server env to not be nil")

	})

	t.Run("env", func(t *testing.T) {
		h := http.NewServeMux()

		testServer.Start(h)
		assert.NotNil(t, testServer.Env(), "expected test server env to not be nil")
	})

}
