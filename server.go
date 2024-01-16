package siows

import (
	"fmt"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/slausonio/siocore"
)

// Server represents a server instance that handles HTTP requests.
type Server struct {
	env     siocore.Env
	config  Config
	server  *http.Server
	handler http.Handler
}

// Env returns the env variable of the SioWSServer.
func (s *Server) Env() siocore.Env {
	return s.env
}

// Server is the method of type `SioWSServer` that returns the underlying `http.Server` instance.
func (s *Server) Server() *http.Server {
	return s.server
}

// NewServer initializes and returns a new instance of the SioWSServer struct.
func NewServer(env siocore.Env) *Server {
	config := NewConfig(env)
	return &Server{
		env:    env,
		config: config,
		server: &http.Server{ReadHeaderTimeout: 5 * time.Second},
	}
}

// Start starts the server with the provided handler.
func (s *Server) Start() {
	startTS := time.Now().UnixMicro()
	serverAddr := fmt.Sprintf(":%s", s.config.port)

	s.server.Addr = serverAddr
	s.server.Handler = s.handler
	s.server.ReadTimeout = 10 * time.Second
	s.server.WriteTimeout = 10 * time.Second
	s.server.IdleTimeout = 120 * time.Second
	s.server.MaxHeaderBytes = 1 << 20

	go func() {
		err := s.server.ListenAndServe()
		if err != nil {
			logrus.Panic(err)
			panic(err)
		}

		defer func(server *http.Server) {
			err := server.Close()
			if err != nil {
				logrus.Panic(err)
				panic(err)
			}
		}(s.server)
	}()

	s.printInfo(startTS)
}

// printInfo prints information about the server.
// calls the [s.printSio] method and logs the server's port number and start time.
func (s *Server) printInfo(start int64) {
	s.printSio()

	logrus.Infof("SioWSServer running on port: %v ", s.config.Port())
	logrus.Infof("SioWSServer Started in %v μs", time.Now().UnixMicro()-start)
}

// printSio prints the Siogo ASCII art to the console.
func (s *Server) printSio() {
	siogoASCII := `
	


	 ▄▄▄▄▄▄▄▄▄▄▄  ▄▄▄▄▄▄▄▄▄▄▄  ▄▄▄▄▄▄▄▄▄▄▄  ▄▄▄▄▄▄▄▄▄▄▄  ▄▄▄▄▄▄▄▄▄▄▄       ▄         ▄  ▄▄▄▄▄▄▄▄▄▄▄ 
	▐░░░░░░░░░░░▌▐░░░░░░░░░░░▌▐░░░░░░░░░░░▌▐░░░░░░░░░░░▌▐░░░░░░░░░░░▌     ▐░▌       ▐░▌▐░░░░░░░░░░░▌
	▐░█▀▀▀▀▀▀▀▀▀  ▀▀▀▀█░█▀▀▀▀ ▐░█▀▀▀▀▀▀▀█░▌▐░█▀▀▀▀▀▀▀▀▀ ▐░█▀▀▀▀▀▀▀█░▌     ▐░▌       ▐░▌▐░█▀▀▀▀▀▀▀▀▀ 
	▐░▌               ▐░▌     ▐░▌       ▐░▌▐░▌          ▐░▌       ▐░▌     ▐░▌       ▐░▌▐░▌          
	▐░█▄▄▄▄▄▄▄▄▄      ▐░▌     ▐░▌       ▐░▌▐░▌ ▄▄▄▄▄▄▄▄ ▐░▌       ▐░▌     ▐░▌   ▄   ▐░▌▐░█▄▄▄▄▄▄▄▄▄ 
	▐░░░░░░░░░░░▌     ▐░▌     ▐░▌       ▐░▌▐░▌▐░░░░░░░░▌▐░▌       ▐░▌     ▐░▌  ▐░▌  ▐░▌▐░░░░░░░░░░░▌
	 ▀▀▀▀▀▀▀▀▀█░▌     ▐░▌     ▐░▌       ▐░▌▐░▌ ▀▀▀▀▀▀█░▌▐░▌       ▐░▌     ▐░▌ ▐░▌░▌ ▐░▌ ▀▀▀▀▀▀▀▀▀█░▌
			  ▐░▌     ▐░▌     ▐░▌       ▐░▌▐░▌       ▐░▌▐░▌       ▐░▌     ▐░▌▐░▌ ▐░▌▐░▌          ▐░▌
	 ▄▄▄▄▄▄▄▄▄█░▌ ▄▄▄▄█░█▄▄▄▄ ▐░█▄▄▄▄▄▄▄█░▌▐░█▄▄▄▄▄▄▄█░▌▐░█▄▄▄▄▄▄▄█░▌     ▐░▌░▌   ▐░▐░▌ ▄▄▄▄▄▄▄▄▄█░▌
	▐░░░░░░░░░░░▌▐░░░░░░░░░░░▌▐░░░░░░░░░░░▌▐░░░░░░░░░░░▌▐░░░░░░░░░░░▌     ▐░░▌     ▐░░▌▐░░░░░░░░░░░▌
	 ▀▀▀▀▀▀▀▀▀▀▀  ▀▀▀▀▀▀▀▀▀▀▀  ▀▀▀▀▀▀▀▀▀▀▀  ▀▀▀▀▀▀▀▀▀▀▀  ▀▀▀▀▀▀▀▀▀▀▀       ▀▀       ▀▀  ▀▀▀▀▀▀▀▀▀▀▀ 
                                                                                                

												  ,#@&%(/,                                  
									.&&(,,,,,,,,,,,,,,,,,,,,,,,,,,,,/&%.                    
							  .@/,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,*@      .*,       
				 %@*,,,/@* @,,,,,,,*%@@&#,,,,,,,,,,,,,,,,,,,&&        &&,,,,,,&,,,,,,,,,*%  
			  %,,,,,,,,,%*,,,,%*             @,,,,,,,,,,,&     /@@@@      %,,,,,,%*/,,,,,,,*
			 &,,,,%@@@&,,,,,@       @@@@@@     &,,,,,,,@      /@@@@@@      %,,,,,,*@@@,,,,,/
			(,,,,#@@@%,,,,,%        @@#/@@      &,,,,,@        #@,&@        @,,,,,,,@,,,,,,(
			 %,,,,,*/,,,,,*.          ..         ,,,,,@                     @,,,,,,,,@,,,,% 
			  *%,,,(,,,,,,./                    #,,,,,*                     /,,,,,,,,,#@*   
				  @,,,,,,,,,*                  #,,,,,,,,%                 @,,,,,,,,,,,%     
				 ,.,,,,,,,,,,(*              @,,,@@@@@@@&,#(           &/,,,,,,,,,,,,,,%    
				 @,,,,,,,,,,,,,,,(@@/*/%@#.,,,,/@@@@@@@@@@,,,,,,,,,,,,,,,,,,,,,,,,,,,,,%    
				 &,,,,,,,,,,,,,,,,,,,,,,,,,,(***************%,,,,,,,,,,,,,,,,,,,,,,,,,,,,   
				 #,,,,,,,,,,,,,,,,,,,,,,,,,,@****************#,,,,,,,,,,,,,,,,,,,,,,,,,,%   
				 #,,,,,,,,,,,,,,,,,,,,,,,,,,,*@@%   @   .@@@#,,,,,,,,,,,,,,,,,,,,,,,,,,,@   
				 %,,,,,,,,,,,,,,,,,,,,,,,,,,,,,&    @    %,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,@   
				 @,,,,,,,,,,,,,,,,,,,,,,,,,,,,,*,  ,@    #,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,&

  `
	fmt.Println(siogoASCII)
}
