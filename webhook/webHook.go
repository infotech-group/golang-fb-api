package webhook

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// Hook struct for handling facebook webhooks required information
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
// which will come via facebook web-hooks
func (h *Hook) StartHook(address, hookPath string) {
	router := mux.NewRouter()
	router.HandleFunc(hookPath, h.handleCheckRequest).Methods("GET")
	router.HandleFunc(hookPath, h.securityWrapper(h.handleEvents)).Methods("POST")
	http.ListenAndServe(address, router)
}

// handleCheckRequest handles check request from facebook with
// verify token and given challenge
// more information https://developers.facebook.com/docs/graph-api/webhooks
func (h *Hook) handleCheckRequest(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	challenge := r.URL.Query().Get("hub.challenge")
	verify := r.URL.Query().Get("hub.verify_token")
	if verify == h.verifyToken {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(challenge))
	}
}

// securityWrapper wraps every request for checking sha1 sum of each request
// to prevent any incorrect API calls and calls from another users
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

/* handleEvents handles incoming events from fb, unmarshaling
// incoming json to struct and sending it as a "instruction"
// how to build object, example:

	for change := range h.changes{
		event := make(change.Type)
		err := json.Unmarshal(change.Data,event)
		if err != nil{
			handle err
		}
	}

*/
func (h *Hook) handleEvents(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	event := new(Event)

	json.NewDecoder(r.Body).Decode(event)
	fmt.Printf("%+v\n", *event)
	for _, v := range event.Entry {
		for _, change := range v.GetChanges() {
			h.Changes <- change
		}
	}
	w.WriteHeader(http.StatusOK)
}
