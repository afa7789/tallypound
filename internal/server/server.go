package server

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/afa7789/tallypound/internal/compound"
	"github.com/gorilla/mux"
)

// Server is the definition of a REST server based on Gin
type Server struct {
	router *mux.Router
}

// New returns a new server that takes advantage of zerolog for logging
// and holds a reference to the app configuration
func NewServer(
	cc *compound.CacheCompound,
) *Server {
	r := mux.NewRouter()

	cCtrl := NewCompoundController(cc)

	// gorilla mux routing group
	r.HandleFunc("/stats", cCtrl.Stats)
	r.HandleFunc("/proposals", cCtrl.Proposals)

	return &Server{
		router: r,
	}
}

// Start starts the REST server
// simple mocking pattern to switch port if one is already in use.
func (s *Server) Start(port int) {
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), s.router)
	if err != nil {
		// Using this error treatment to try again on next port
		if strings.Contains(err.Error(), "address already in use") {
			fmt.Println("")
			log.Printf("PORT ALREADY IN USE::%d", port)
			port++
			log.Printf("TRYING NEXT PORT:%d\n", port)
			s.Start(port)
		} else {
			panic(err)
		}
	}
}
