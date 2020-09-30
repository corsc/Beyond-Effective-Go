package server

type Config interface {
	// implementation removed
}

func New(cfg Config) *Server {
	return &Server{
		// implementation removed
	}
}

type Server struct{}

func (s *Server) Start() {
	// implementation removed
}
