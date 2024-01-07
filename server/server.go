package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/slausonio/go-webserver/environment"
)

type Env interface {
	Value(key string) string
	Update(key, value string)
}

type Logger interface {
	Log(message string)
}

type SioGoServer interface {
	Env() Env
	Server() *http.Server
	Kill()
	Start(handler http.Handler)
	printInfo(start int64)
	printSio()
}

type Server struct {
	env    Env
	server *http.Server
}

func (s *Server) Env() Env {
	return s.env
}

func (s *Server) Server() *http.Server {
	return s.server
}

func (s *Server) Kill() {
	err := s.server.Close()
	if err != nil {
		panic(err)
	}
}

func NewServer() *Server {

	return &Server{
		env:    environment.NewEnvironment(),
		server: &http.Server{},
	}
}

func (s *Server) Start(handler http.Handler) {
	startTS := time.Now().UnixMicro()
	serverAddr := fmt.Sprintf(":%s", s.env.Value(environment.Port))

	s.server.Addr = serverAddr
	s.server.Handler = handler
	s.server.ReadTimeout = 10 * time.Second
	s.server.WriteTimeout = 10 * time.Second
	s.server.IdleTimeout = 120 * time.Second
	s.server.MaxHeaderBytes = 1 << 20

	go func() {
		logrus.Fatal(s.server.ListenAndServe())

		s.printInfo(startTS)
	}()
}

func (s *Server) printInfo(start int64) {
	s.printSio()
	// e.printGopher()

	logrus.Infof("Server running on port: %v ", s.env.Value(environment.Port))
	logrus.Infof("Server Started in %v", time.Now().UnixMicro()-start)
}

func (s *Server) printSio() {
	siogoASCII := `




 ▄▄▄▄▄▄▄▄▄▄▄  ▄▄▄▄▄▄▄▄▄▄▄  ▄▄▄▄▄▄▄▄▄▄▄  ▄▄▄▄▄▄▄▄▄▄▄  ▄▄▄▄▄▄▄▄▄▄▄ 
▐░░░░░░░░░░░▌▐░░░░░░░░░░░▌▐░░░░░░░░░░░▌▐░░░░░░░░░░░▌▐░░░░░░░░░░░▌
▐░█▀▀▀▀▀▀▀▀▀  ▀▀▀▀█░█▀▀▀▀ ▐░█▀▀▀▀▀▀▀█░▌▐░█▀▀▀▀▀▀▀▀▀ ▐░█▀▀▀▀▀▀▀█░▌
▐░▌               ▐░▌     ▐░▌       ▐░▌▐░▌          ▐░▌       ▐░▌
▐░█▄▄▄▄▄▄▄▄▄      ▐░▌     ▐░▌       ▐░▌▐░▌ ▄▄▄▄▄▄▄▄ ▐░▌       ▐░▌
▐░░░░░░░░░░░▌     ▐░▌     ▐░▌       ▐░▌▐░▌▐░░░░░░░░▌▐░▌       ▐░▌
 ▀▀▀▀▀▀▀▀▀█░▌     ▐░▌     ▐░▌       ▐░▌▐░▌ ▀▀▀▀▀▀█░▌▐░▌       ▐░▌
          ▐░▌     ▐░▌     ▐░▌       ▐░▌▐░▌       ▐░▌▐░▌       ▐░▌
 ▄▄▄▄▄▄▄▄▄█░▌ ▄▄▄▄█░█▄▄▄▄ ▐░█▄▄▄▄▄▄▄█░▌▐░█▄▄▄▄▄▄▄█░▌▐░█▄▄▄▄▄▄▄█░▌
▐░░░░░░░░░░░▌▐░░░░░░░░░░░▌▐░░░░░░░░░░░▌▐░░░░░░░░░░░▌▐░░░░░░░░░░░▌
 ▀▀▀▀▀▀▀▀▀▀▀  ▀▀▀▀▀▀▀▀▀▀▀  ▀▀▀▀▀▀▀▀▀▀▀  ▀▀▀▀▀▀▀▀▀▀▀  ▀▀▀▀▀▀▀▀▀▀▀ 
                                                                 



  `
	fmt.Println(siogoASCII)
}

// func (s *Server) printGopher() {
//   gopher := `
//
//          ,_---~~~~~----._
//   _,,_,*^____      _____``*g*\"*,
//  \/ __/ /'     ^.  /      \ ^@q   f
// [  @f | @))    |  | @))   l  0 _/
//  \` + '`' + `/   \~____ / __ \_____/    \
//   |           _l__l_           I
//   }          [______]           I
//   ]            | | |            |
//   ]             ~ ~             |
//   |                            |
//    |                           |
//
//
//   `
//
//   fmt = fmt.Println(gopher)
//
//   }
