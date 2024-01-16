package siows

import (
	"log/slog"
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
	log       *slog.Logger
}

func (s SioWS) Env() Env {
	return s.env
}

func (s SioWS) SioServer() SioServer {
	return s.sioServer
}

func (s SioWS) Log() *slog.Logger {
	return s.log
}

func NewSioWS(handler http.Handler, log *slog.Logger) *SioWS {

	appEnv := siocore.NewAppEnv(log)
	env := appEnv.Env()

	return &SioWS{
		env:       env,
		sioServer: NewServer(env),
		log:       log,
	}
}
