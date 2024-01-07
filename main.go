package siows

import (
	"net/http"
)

type Server interface {
	Env() Env
	Server() *http.Server
	Start(handler http.Handler)
}

type SioWS struct {
	env       Env
	sioServer Server
}

func NewSioWS(handler http.Handler) *SioWS {
	env := NewEnvironment()

	return &SioWS{
		env:       env,
		sioServer: NewServer(env),
	}
}
