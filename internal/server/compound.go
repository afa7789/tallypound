package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/afa7789/tallypound/internal/compound"
)

type CompoundController struct {
	cc *compound.CacheCompound
}

func NewCompoundController(cc *compound.CacheCompound) *CompoundController {
	return &CompoundController{
		cc: cc,
	}
}

// Proposals return an array with all proposals.
func (ctl *CompoundController) Proposals(w http.ResponseWriter, r *http.Request) {
	p, err := ctl.cc.Proposals()
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
func (ctl *CompoundController) Stats(w http.ResponseWriter, r *http.Request) {
	s, err := ctl.cc.Stats()
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

// CachedStats return an object with the stats.
func (ctl *CompoundController) CachedStats(w http.ResponseWriter, r *http.Request) {
	s, err := ctl.cc.CachedStats()
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
