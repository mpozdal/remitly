package api

import (
	"log"
	"net/http"

	"mpozdal/remitly/db"
	"mpozdal/remitly/services/swift"

	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	dbm  *db.DBManager
}

func NewAPIServer(addr string, dbm *db.DBManager) *APIServer {
	return &APIServer{
		addr: addr,
		dbm:  dbm,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/v1").Subrouter()
	swiftService := swift.NewSwiftService(s.dbm)
	swiftHandler := swift.NewHandler(swiftService)
	swiftHandler.RegisterRoutes(subrouter)

	log.Println("Starting server at", s.addr)
	return http.ListenAndServe(s.addr, router)

}
