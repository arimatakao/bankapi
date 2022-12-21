package api

import (
	"fmt"
	"net/http"
)

func (s *APIServer) handleTransfer(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "GET":
		s.handleGetTransferHistory(w, r)
	case "PUT":
		s.handleTransferFromTo(w, r)
	}
	return fmt.Errorf("method not allowed %s", r.Method)
}

func (s *APIServer) handleGetTransferHistory(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleTransferFromTo(w http.ResponseWriter, r *http.Request) error {
	return nil
}
