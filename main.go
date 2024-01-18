package siows

import (
	"net/http"

	"github.com/slausonio/siocore"
)

type Env interface {
	Value(key string) string
	Update(key, value string)
}

type SioServer interface {
	Env() siocore.Env
	Server() *http.Server
	Start()
}

type SioWS struct {
	env       Env
	sioServer SioServer
}

func (s SioWS) Env() Env {
	return s.env
}

func (s SioWS) SioServer() SioServer {
	return s.sioServer
}

func NewSioWS(handler http.Handler) *SioWS {

	appEnv := siocore.NewAppEnv()
	env := appEnv.Env()

	return &SioWS{
		env:       env,
		sioServer: NewServer(env, handler, log),
	}
}
