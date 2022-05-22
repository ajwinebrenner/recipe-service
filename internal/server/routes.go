package server

import (
	"net/http"

	"github.com/graph-gophers/graphql-go/relay"
)

func (s *Server) routes() {
	//with chi, method of server struct
	//s.router.?
	//how to route with graphql?
	s.Router.Method(http.MethodPost, "/query", &relay.Handler{Schema: s.Schema})
}
