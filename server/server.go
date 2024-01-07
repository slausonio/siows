package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/slausonio/siows/environment"
)

// Server represents a server instance that handles HTTP requests.
type Server struct {
	env    environment.SioWSEnv
	server *http.Server
}

// Env returns the env variable of the Server.
func (s *Server) Env() environment.SioWSEnv {
	return s.env
}

// Server is the method of type `Server` that returns the underlying `http.Server` instance.
func (s *Server) Server() *http.Server {
	return s.server
}

// Kill terminates the server by closing the underlying http.Server instance.
// It panics if an error is encountered while closing the server.
func (s *Server) Kill() {
	err := s.server.Close()
	if err != nil {
		panic(err)
	}
}

// NewServer initializes and returns a new instance of the Server struct.
func NewServer(env environment.SioWSEnv) *Server {
	return &Server{
		env:    env,
		server: &http.Server{},
	}
}

// Start starts the server with the provided handler.
func (s *Server) Start(handler http.Handler) {
	startTS := time.Now().UnixMicro()
	serverAddr := fmt.Sprintf(":%s", s.env.Value(environment.PortKey))

	s.server.Addr = serverAddr
	s.server.Handler = handler
	s.server.ReadTimeout = 10 * time.Second
	s.server.WriteTimeout = 10 * time.Second
	s.server.IdleTimeout = 120 * time.Second
	s.server.MaxHeaderBytes = 1 << 20

	err := s.server.ListenAndServe()
	if err != nil {
		panic(err)
	}

	s.printInfo(startTS)
}

// printInfo prints information about the server.
// calls the [s.printSio] method and logs the server's port number and start time.
func (s *Server) printInfo(start int64) {
	s.printSio()
	// e.printGopher()

	logrus.Infof("Server running on port: %v ", s.env.Value(environment.PortKey))
	logrus.Infof("Server Started in %v", time.Now().UnixMicro()-start)
}

// printSio prints the Siogo ASCII art to the console.
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
