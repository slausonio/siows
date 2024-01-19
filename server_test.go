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

	t.Run("Happy", func(t *testing.T) {

		h := http.NewServeMux()
		testServer := NewServer(happyEnvMap, h)

		testServer.Start()
	})

	t.Run("panics", func(t *testing.T) {
		t.Skip()

		t.Run("duplicate address", func(t *testing.T) {
			h := http.NewServeMux()
			testServer := NewServer(happyEnvMap, h)

			assert.Panics(t,
				func() {
					testServer.Start()
				}, "expected server kill to panic")
		})
	})
}

func TestServer_getters(t *testing.T) {
	t.Parallel()

	h := http.NewServeMux()
	testServer := NewServer(happyEnvMap, h)
	testServer.Start()

	t.Run("server", func(t *testing.T) {

		assert.NotNil(t, testServer.Env(), "expected test server env to not be nil")

	})

	t.Run("env", func(t *testing.T) {

		assert.NotNil(t, testServer.Env(), "expected test server env to not be nil")
	})

}
