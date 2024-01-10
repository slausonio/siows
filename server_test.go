package siows

import (
	"net/http"
	"testing"

	"github.com/slausonio/siocore"
	"github.com/stretchr/testify/assert"
)

var happyEnvMap = siocore.Env{
	siocore.EnvKeyCurrentEnv: "test",
	siocore.EnvKeyAppName:    "go-webserver",
	siocore.EnvKeyPort:       "8080",
}

func TestServer_Start(t *testing.T) {
	t.Parallel()

	t.Run("Happy", func(t *testing.T) {

		testServer := NewServer(happyEnvMap)
		h := http.NewServeMux()

		testServer.Start(h)
	})

	t.Run("panics", func(t *testing.T) {
		t.Skip()

		t.Run("duplicate address", func(t *testing.T) {
			testServer := NewServer(happyEnvMap)

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
	t.Parallel()

	testServer := NewServer(happyEnvMap)

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
