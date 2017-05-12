package fb

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// Hook struct for handling facebook webhooks
type Hook struct {
	verifyToken string
	secret      []byte
	Changes     chan TypedChange
}

// NewHook returns new instance of WebHook
func NewHook(VerifyToken string, AppSecret []byte) *Hook {
	return &Hook{
		verifyToken: VerifyToken,
		secret:      AppSecret,
		Changes:     make(chan TypedChange, 0),
	}
}

// StartHook starts handlers for incoming requests
func (h *Hook) StartHook(address, hookPath string) {
	router := mux.NewRouter()
	router.HandleFunc(hookPath, h.handleCheckRequest).Methods("GET")
	router.HandleFunc(hookPath, h.securityWrapper(h.handleEvents)).Methods("POST")
	http.ListenAndServe(address, router)
}

func (h *Hook) handleCheckRequest(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	challenge := r.URL.Query().Get("hub.challenge")
	verify := r.URL.Query().Get("hub.verify_token")
	if verify == h.verifyToken {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(challenge))
	}
}

func (h *Hook) securityWrapper(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestHash := strings.Split(r.Header.Get("X-Hub-Signature"), "=")[1]
		hash, err := calculateHash(r, h.secret)
		if requestHash == fmt.Sprintf("%x", hash) && err == nil {
			fmt.Println("Done! Let`s serve it")
			handlerFunc.ServeHTTP(w, r)
		} else {
			w.WriteHeader(http.StatusBadRequest)
			r.Body.Close()
		}
	})
}

func (h *Hook) handleEvents(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	event := new(Event)
	json.NewDecoder(r.Body).Decode(event)
	fmt.Printf("%+v\n", *event)
	for _, v := range event.Entry {
		for _, change := range v.getChanges() {
			h.Changes <- change
		}
	}
	w.WriteHeader(http.StatusOK)
}
