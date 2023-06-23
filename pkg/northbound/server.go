// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

// Package northbound houses implementations of various application-oriented interfaces
// for the ONOS configuration subsystem.
package northbound

import (
	"crypto/tls"
	"fmt"
	"github.com/onosproject/onos-lib-go/pkg/grpc/auth"
	"google.golang.org/grpc/credentials"
	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"

	"github.com/onosproject/onos-lib-go/pkg/certs"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"google.golang.org/grpc"
)

var log = logging.GetLogger("northbound")

// Service provides service-specific registration for grpc services.
type Service interface {
	Register(s *grpc.Server)
}

// Server provides NB gNMI server for onos-config.
type Server struct {
	cfg      *ServerConfig
	services []Service
	server   *grpc.Server
}

// SecurityConfig security configuration
type SecurityConfig struct {
	AuthenticationEnabled bool
	AuthorizationEnabled  bool
}

// ServerConfig comprises a set of server configuration options.
type ServerConfig struct {
	CaPath      *string
	KeyPath     *string
	CertPath    *string
	Port        int16
	Insecure    bool
	SecurityCfg *SecurityConfig
}

// NewServer initializes gNMI server using the supplied configuration.
func NewServer(cfg *ServerConfig) *Server {
	return &Server{
		services: []Service{},
		cfg:      cfg,
	}
}

// NewServerConfig creates a server config created with the specified end-point security details.
// Deprecated: Use NewServerCfg instead
func NewServerConfig(caPath string, keyPath string, certPath string, port int16, insecure bool) *ServerConfig {
	return &ServerConfig{
		Port:        port,
		Insecure:    insecure,
		CaPath:      &caPath,
		KeyPath:     &keyPath,
		CertPath:    &certPath,
		SecurityCfg: &SecurityConfig{},
	}
}

// NewServerCfg creates a server config created with the specified end-point security details.
func NewServerCfg(caPath string, keyPath string, certPath string, port int16, insecure bool, secCfg SecurityConfig) *ServerConfig {
	return &ServerConfig{
		Port:        port,
		Insecure:    insecure,
		CaPath:      &caPath,
		KeyPath:     &keyPath,
		CertPath:    &certPath,
		SecurityCfg: &secCfg,
	}
}

// NewInsecureServerConfig creates an insecure server configuration for the specified port.
func NewInsecureServerConfig(port int16) *ServerConfig {
	return &ServerConfig{
		Port:     port,
		Insecure: true,
		SecurityCfg: &SecurityConfig{
			AuthenticationEnabled: false,
			AuthorizationEnabled:  false,
		},
	}
}

// AddService adds a Service to the server to be registered on Serve.
func (s *Server) AddService(r Service) {
	s.services = append(s.services, r)
}

// Serve starts the NB gNMI server.
func (s *Server) Serve(started func(string), grpcOpts ...grpc.ServerOption) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.cfg.Port))
	if err != nil {
		return err
	}
	tlsCfg := &tls.Config{}

	if s.cfg.Insecure && s.cfg.CertPath == nil && s.cfg.KeyPath == nil {
		// nothing
		log.Debug("Running in insecure mode")
	} else if *s.cfg.CertPath == "" && *s.cfg.KeyPath == "" {
		// Load default Certificates
		clientCerts, err := tls.X509KeyPair([]byte(certs.DefaultLocalhostCrt), []byte(certs.DefaultLocalhostKey))
		if err != nil {
			log.Error("Error loading default certs")
			return err
		}
		tlsCfg.Certificates = []tls.Certificate{clientCerts}
	} else {
		log.Infof("Loading certs: %s %s", *s.cfg.CertPath, *s.cfg.KeyPath)
		clientCerts, err := tls.LoadX509KeyPair(*s.cfg.CertPath, *s.cfg.KeyPath)
		if err != nil {
			log.Info("Error loading default certs")
		}
		tlsCfg.Certificates = []tls.Certificate{clientCerts}
	}

	if s.cfg.Insecure {
		// RequestClientCert will ask client for a certificate but won't
		// require it to proceed. If certificate is provided, it will be
		// verified.
		tlsCfg.ClientAuth = tls.RequestClientCert
	} else {
		tlsCfg.ClientAuth = tls.RequireAndVerifyClientCert
	}

	if s.cfg.CaPath == nil {
		log.Debug("Running with no CA certificates")
	} else if *s.cfg.CaPath == "" {
		log.Info("Loading default CA onfca")
		tlsCfg.ClientCAs, err = certs.GetCertPoolDefault()
	} else {
		tlsCfg.ClientCAs, err = certs.GetCertPool(*s.cfg.CaPath)
	}

	if err != nil {
		return err
	}

	opts := make([]grpc.ServerOption, 0, 5)
	if len(tlsCfg.Certificates) > 0 {
		opts = append(opts, grpc.Creds(credentials.NewTLS(tlsCfg)))
	}

	if s.cfg.SecurityCfg.AuthenticationEnabled {
		log.Info("Authentication Enabled")
		opts = append(opts, grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				grpc_auth.UnaryServerInterceptor(auth.AuthenticationInterceptor),
			)))
		opts = append(opts, grpc.StreamInterceptor(
			grpc_middleware.ChainStreamServer(
				grpc_auth.StreamServerInterceptor(auth.AuthenticationInterceptor))))

	}

	opts = append(opts, grpcOpts...)

	s.server = grpc.NewServer(opts...)
	for i := range s.services {
		s.services[i].Register(s.server)
	}
	started(lis.Addr().String())

	log.Infof("Starting RPC server on address: %s", lis.Addr().String())
	return s.server.Serve(lis)
}

// Stop stops the server.
func (s *Server) Stop() {
	s.server.Stop()
}

// GracefulStop stops the server gracefully.
func (s *Server) GracefulStop() {
	s.server.GracefulStop()
}

// StartInBackground starts serving in the background, returning an error if any issue is encountered
func (s *Server) StartInBackground() error {
	doneCh := make(chan error)
	go func() {
		err := s.Serve(func(started string) {
			log.Info("Started NBI on ", started)
			close(doneCh)
		})
		if err != nil {
			doneCh <- err
		}
	}()
	return <-doneCh
}
