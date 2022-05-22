package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ajwinebrenner/recipe-service/internal/config"
	"github.com/graph-gophers/graphql-go"
	"gorm.io/gorm"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	Conf   *config.Config
	DB     *gorm.DB
	Router chi.Router
	Schema *graphql.Schema
}

func Start(c *config.Config) {

	db, err := ConnectToDatabase(c.DatabaseURL)
	if err != nil {
		log.Fatal("could not connect to database: ", err)
	}

	s := Server{
		Conf:   c,
		DB:     db,
		Router: chi.NewRouter(),
	}

	s.makeSchema() //

	s.routes() // uses all fields

	addr := fmt.Sprintf("%s:%s", s.Conf.Host, s.Conf.Port)
	log.Fatal(http.ListenAndServe(addr, s.Router))
}
