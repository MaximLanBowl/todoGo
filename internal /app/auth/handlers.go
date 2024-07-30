package auth

import (
	"encoding/json"
	"net/http"

	"github.com/MaximLanBowl/todoGo.git/pkg/models"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User 
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest) 
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var creds *Credentials
}