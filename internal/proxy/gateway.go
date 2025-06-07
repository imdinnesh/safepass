package proxy

import (
	"net/http"
	"github.com/imdinnesh/safepass/internal/config"
)

func NewGateway(cfg *config.Config) (http.Handler, error) {
	// TODO: setup reverse proxy with middleware stack
	mux := http.NewServeMux()

	// Temporary placeholder
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Safepass Gateway Running"))
	})

	return mux, nil
}
