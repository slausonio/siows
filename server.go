package siows

import (
	"fmt"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

// Server represents a server instance that handles HTTP requests.
type Server struct {
	env    Env
	server *http.Server
}

// Env returns the env variable of the Server.
func (s *Server) Env() Env {
	return s.env
}

// Server is the method of type `Server` that returns the underlying `http.Server` instance.
func (s *Server) Server() *http.Server {
	return s.server
}

// NewServer initializes and returns a new instance of the Server struct.
func NewServer(env Env) *Server {
	return &Server{
		env:    env,
		server: &http.Server{},
	}
}

// Start starts the server with the provided handler.
func (s *Server) Start(handler http.Handler) {
	startTS := time.Now().UnixMicro()
	serverAddr := fmt.Sprintf(":%s", s.env.Value(PortKey))

	s.server.Addr = serverAddr
	s.server.Handler = handler
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
	// e.printGopher()

	logrus.Infof("Server running on port: %v ", s.env.Value(PortKey))
	logrus.Infof("Server Started in %v μs", time.Now().UnixMicro()-start)
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
