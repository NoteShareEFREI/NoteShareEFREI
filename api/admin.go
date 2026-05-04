package api

import (
	"NoteShareEFREI/backend"
	"NoteShareEFREI/database"
	"encoding/json"
	"net/http"

	"github.com/lestrrat-go/jwx/v4/jwt"
)

func CheckAdminHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Check if user is authenticated
	token, err := jwt.ParseRequest(r, jwt.WithCookieKey("Http-Jwt"), jwt.WithVerify(false))
	if err != nil {
		json.NewEncoder(w).Encode(map[string]bool{"isAdmin": false})
		return
	}

	// Validate JWT token
	accountId, err := backend.ValidateJWT(token)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]bool{"isAdmin": false})
		return
	}

	// Check if user is admin
	isAdmin, err := database.IsAdmin(accountId)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]bool{"isAdmin": false})
		return
	}

	json.NewEncoder(w).Encode(map[string]bool{"isAdmin": isAdmin})
}
