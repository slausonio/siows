package siows

import (
	"github.com/slausonio/siows/environment"
	"github.com/slausonio/siows/server"
	"net/http"
)

type Server interface {
	Env() environment.Env
	Server() *http.Server
	Start(handler http.Handler)
	Kill()
}

type SioWS struct {
	env       environment.Env
	sioServer Server
}

func NewSioWS(handler http.Handler) *SioWS {
	env := environment.NewEnvironment()

	return &SioWS{
		env:       env,
		sioServer: server.NewServer(env),
	}
}
