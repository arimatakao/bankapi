package api

import (
	"encoding/json"
	"net/http"

	"github.com/arimatakao/bankapi/storage"
	"github.com/gorilla/mux"
)

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

type apifunc func(w http.ResponseWriter, r *http.Request) error

func makeHTTPHandleFunc(f apifunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, APIError{Error: err.Error()})
		}
	}
}

type APIError struct {
	Error string
}

type APIServer struct {
	listenAddr string
	db         storage.Storager
}

func NewAPIServer(addr string, db storage.Storager) *APIServer {
	return &APIServer{
		listenAddr: addr,
		db:         db,
	}
}

func (s *APIServer) Run() {
	r := mux.NewRouter()

	r.HandleFunc("/signup", makeHTTPHandleFunc(s.handleSignUp)).Methods("POST")
	r.HandleFunc("/signin", makeHTTPHandleFunc(s.handleSignIn)).Methods("POST")

	r.HandleFunc("/account", makeHTTPHandleFunc(s.handleAccount)).Methods("GET", "POST", "DELETE")
	r.HandleFunc("/account/login", makeHTTPHandleFunc(s.handleChangeAccountLogin)).Methods("PUT")
	r.HandleFunc("/account/password", makeHTTPHandleFunc(s.handleChangeAccountPassword)).Methods("PUT")
	r.HandleFunc("/account/balance", makeHTTPHandleFunc(s.handleAccountBalance)).Methods("GET")

	r.HandleFunc("/transfer", makeHTTPHandleFunc(s.handleTransfer)).Methods("GET", "PUT")

	http.ListenAndServe(s.listenAddr, r)
}
