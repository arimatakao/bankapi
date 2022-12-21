package api

import (
	"fmt"
	"net/http"

	"github.com/arimatakao/bankapi/types"
)

func (s *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "GET":
		s.handleGetAccount(w, r)
	case "POST":
		s.hanleCreateAccount(w, r)
	case "DELETE":
		s.handleDeleteAccount(w, r)

	}
	return fmt.Errorf("method not allowed %s", r.Method)
}

func (s *APIServer) hanleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	account := types.NewAccount("Abc", "Def")
	WriteJSON(w, http.StatusAccepted, account)
	return nil
}

func (s *APIServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleChangeAccountLogin(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleChangeAccountPassword(w http.ResponseWriter, t *http.Request) error {
	return nil
}

func (s *APIServer) handleAccountBalance(w http.ResponseWriter, r *http.Request) error {
	return nil
}
