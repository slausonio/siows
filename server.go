package siows

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/slausonio/siocore"
)

// Server represents a server instance that handles HTTP requests.
type Server struct {
	env    siocore.Env
	server *http.Server
	config Config
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
func NewServer(env siocore.Env, handler http.Handler) *Server {
	config := NewConfig(env)
	serverAddr := fmt.Sprintf(":%s", config.port)

	return &Server{
		env:    env,
		config: config,
		server: &http.Server{
			Addr:              serverAddr,
			Handler:           handler,
			ReadTimeout:       10 * time.Second,
			WriteTimeout:      10 * time.Second,
			IdleTimeout:       120 * time.Second,
			MaxHeaderBytes:    1 << 20,
			ReadHeaderTimeout: 5 * time.Second},
	}
}

// Start starts the server with the provided handler.
func (s *Server) Start() {
	startTS := time.Now().UnixMicro()
	done := make(chan bool)

	go func() {
		s.printSio()

		err := s.server.ListenAndServe()
		if err != nil {
			slog.Error("server start error: ", err)
			panic(err)
		}
		done <- true // Signal that server has stopped
	}()

	s.printInfo(startTS)

	<-done
}

// printInfo prints information about the server.
// calls the [s.printSio] method and logs the server's port number and start time.
func (s *Server) printInfo(start int64) {

	slog.Info(fmt.Sprintf("SioWSServer running on port: %v ", s.config.Port()))
	slog.Info(fmt.Sprintf("SioWSServer Started in %v μs", time.Now().UnixMicro()-start))
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
