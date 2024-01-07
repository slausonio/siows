package siows

import (
	"fmt"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

// SioWSServer represents a server instance that handles HTTP requests.
type SioWSServer struct {
	env    Env
	server *http.Server
}

// Env returns the env variable of the SioWSServer.
func (s *SioWSServer) Env() Env {
	return s.env
}

// SioWSServer is the method of type `SioWSServer` that returns the underlying `http.Server` instance.
func (s *SioWSServer) Server() *http.Server {
	return s.server
}

// NewServer initializes and returns a new instance of the SioWSServer struct.
func NewServer(env Env) *SioWSServer {
	return &SioWSServer{
		env:    env,
		server: &http.Server{},
	}
}

// Start starts the server with the provided handler.
func (s *SioWSServer) Start(handler http.Handler) {
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
func (s *SioWSServer) printInfo(start int64) {
	s.printSio()
	// e.printGopher()

	logrus.Infof("SioWSServer running on port: %v ", s.env.Value(PortKey))
	logrus.Infof("SioWSServer Started in %v μs", time.Now().UnixMicro()-start)
}

// printSio prints the Siogo ASCII art to the console.
func (s *SioWSServer) printSio() {
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
