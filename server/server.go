package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/slausonio/go-webserver/environment"
)

type Server struct {
	Environment environment.Environment
}

func NewServer() *Server {
	return &Server{
		Environment: environment.NewEnvironment(),
	}
}

func (s *Server) Start(handler http.Handler) {
	startTS := time.Now().UnixMicro()
	serverAddr := fmt.Sprintf(":%s", s.Environment.Value(environment.Port))

	server := &http.Server{
		Addr:           serverAddr,
		Handler:        handler,
		ReadTimeout:    10 * time.Second,  // Time to read the entire request, including the body
		WriteTimeout:   10 * time.Second,  // Time to write the response
		IdleTimeout:    120 * time.Second, // Time a Keep-Alive connection will be kept idle before being reused
		MaxHeaderBytes: 1 << 20,           // Maximum header size
	}

	go func() {
		logrus.Fatal(server.ListenAndServe())

		s.printInfo(startTS)
	}()
}

func (s *Server) printInfo(start int64) {
	s.printSio()
	// e.printGopher()

	logrus.Infof("Server running on port: %v ", s.Environment.Value(environment.Port))
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
