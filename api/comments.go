package api

import (
	"NoteShareEFREI/backend"
	"NoteShareEFREI/database"
	"encoding/json"
	"net/http"

	"github.com/lestrrat-go/jwx/v4/jwt"
)

func GetCommentsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	hash := r.URL.Query().Get("sheet")
	if hash == "" {
		http.Error(w, "Missing sheet parameter", http.StatusBadRequest)
		return
	}

	sheetId, err := database.GetSheetIdByHash(hash)
	if err != nil {
		http.Error(w, "Sheet not found", http.StatusNotFound)
		return
	}

	comments, err := database.GetCommentsBySheetId(sheetId)
	if err != nil {
		http.Error(w, "Failed to get comments", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(comments)
}

func PostCommentHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse multipart form
	err := r.ParseMultipartForm(32 << 20) // 32MB
	if err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	// Check authentication
	token, err := jwt.ParseRequest(r, jwt.WithCookieKey("Http-Jwt"), jwt.WithVerify(false))
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	accountId, err := backend.ValidateJWT(token)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	hash := r.FormValue("sheet")
	content := r.FormValue("content")
	if hash == "" || content == "" {
		http.Error(w, "Missing parameters", http.StatusBadRequest)
		return
	}

	sheetId, err := database.GetSheetIdByHash(hash)
	if err != nil {
		http.Error(w, "Sheet not found", http.StatusNotFound)
		return
	}

	err = database.InsertComment(content, sheetId, accountId)
	if err != nil {
		http.Error(w, "Failed to add comment", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"status": "success"})
}
