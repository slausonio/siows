package siows

import (
	"net/http"
)

type Env interface {
	Value(key string) string
	Update(key, value string)
}

type SioServer interface {
	Env() SioWSEnv
	Server() *http.Server
	Start(handler http.Handler)
}

type SioWS struct {
	env       Env
	sioServer SioServer
}

func NewSioWS(handler http.Handler) *SioWS {
	env := NewEnvironment()

	return &SioWS{
		env:       env,
		sioServer: NewServer(env),
	}
}
