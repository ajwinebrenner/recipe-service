package server

import (
	"embed"
	"log"

	"github.com/ajwinebrenner/recipe-service/internal/resolvers"

	"github.com/graph-gophers/graphql-go"
)

//go:embed graphql/schema.graphql
var f embed.FS

func (s *Server) makeSchema() {
	bstr, err := f.ReadFile("graphql/schema.graphql")
	if err != nil {
		log.Fatalf("error reading schema %v", err)
	}
	schemaStr := string(bstr)
	rootResolver := resolvers.NewRootResolver(s.DB)
	opts := []graphql.SchemaOpt{graphql.UseFieldResolvers()}

	s.Schema, err = graphql.ParseSchema(schemaStr, rootResolver, opts...)
	if err != nil {
		log.Fatalf("error creating schema %v", err)
	}
}
