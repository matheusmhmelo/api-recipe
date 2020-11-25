package internal

import (
	"github.com/gorilla/mux"
	"github.com/matheusmhmelo/api-recipe/internal/handler"
	"net/http"
)

type Server struct {
	Router *mux.Router
}

func NewServer() *Server {
	r := mux.NewRouter()
	s := Server{r}

	s.serve()

	return &s
}

func (s *Server) serve() {
	s.Router.Path("/heartbeat").HandlerFunc(handler.Heartbeat).Methods(http.MethodGet)
}