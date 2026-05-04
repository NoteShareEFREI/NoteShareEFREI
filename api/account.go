package api

import (
    "NoteShareEFREI/backend"
    "NoteShareEFREI/database"
    "NoteShareEFREI/routers"
    "encoding/json"
    "net/http"

    "github.com/lestrrat-go/jwx/v4/jwt"
)

func DeleteAccountHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    // Check if user is authenticated
    token, err := jwt.ParseRequest(r, jwt.WithCookieKey("Http-Jwt"), jwt.WithVerify(false))
    if err != nil {
        w.WriteHeader(http.StatusUnauthorized)
        json.NewEncoder(w).Encode(map[string]string{"error": "Unauthorized"})
        return
    }

    // Validate JWT token
    accountId, err := backend.ValidateJWT(token)
    if err != nil {
        w.WriteHeader(http.StatusUnauthorized)
        json.NewEncoder(w).Encode(map[string]string{"error": "Invalid token"})
        return
    }

    // Verify it's a POST request
    if r.Method != "POST" {
        w.WriteHeader(http.StatusMethodNotAllowed)
        json.NewEncoder(w).Encode(map[string]string{"error": "Method not allowed"})
        return
    }

    // Delete the account from the database
    err = database.DeleteAccount(accountId)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(map[string]string{"error": "Failed to delete account"})
        return
    }

    routers.LogoutHandler(w, r)
}