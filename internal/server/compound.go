package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/afa7789/tallypound/internal/compound"
)

// Proposals return an array with all proposals.
func Proposals(w http.ResponseWriter, r *http.Request) {
	p, err := compound.Proposals()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error: %s", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	a, err := json.Marshal(p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error: %s", err)
		return
	}

	_, _ = w.Write(a)
}

// Stats return an object with the stats.
func Stats(w http.ResponseWriter, r *http.Request) {
	s, err := compound.Stats()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error: %s", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	a, err := json.Marshal(s)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error: %s", err)
		return
	}

	_, _ = w.Write(a)
}
