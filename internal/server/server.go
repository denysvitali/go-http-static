package server

import (
	"fmt"
	"net/http"
)

type Server struct {
	Path          string
	Port          string
	ListenAddress string
	TLS           bool
	CertFile      string
	CertKey       string
}

func New(path, port, listenAddress string, tls bool, certFile, certKey string) *Server {
	return &Server{
		Path:          path,
		Port:          port,
		ListenAddress: listenAddress,
		TLS:           tls,
		CertFile:      certFile,
		CertKey:       certKey,
	}
}

func (s *Server) Address() string {
	return fmt.Sprintf("%s:%s", s.ListenAddress, s.Port)
}

func (s *Server) Start() error {
	addr := s.Address()

	fileServer := http.FileServer(http.Dir(s.Path))
	http.Handle("/", fileServer)

	if s.TLS {
		if s.CertFile == "" {
			return fmt.Errorf("certificate file is required for TLS")
		}
		if s.CertKey == "" {
			return fmt.Errorf("certificate key is required for TLS")
		}
		return http.ListenAndServeTLS(addr, s.CertFile, s.CertKey, nil)
	}

	return http.ListenAndServe(addr, nil)
}