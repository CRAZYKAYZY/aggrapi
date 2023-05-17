package api

import (
	"net/http"

	db "github.com/CRAZYKAYZY/aggrapi/db/sqlc"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

type Server struct {
	store  *db.Store
	router *chi.Mux
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()

	v1Router.Get("/healthz", server.HandlerReadiness)
	v1Router.Get("/err", server.HandlerErr)
	v1Router.Post("/users", server.HandlerUsersCreate)

	router.Mount("/v1", v1Router)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	srv := &http.Server{
		Addr:    address,
		Handler: server.router,
	}

	return srv.ListenAndServe()
}
