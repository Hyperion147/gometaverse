package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Hyperion147/gometaverse/internal/database"
	"github.com/Hyperion147/gometaverse/internal/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthHandler struct {
	db *gorm.DB
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{db: database.DB}
}

func (h *AuthHandler) SignUp (w http.ResponseWriter, r *http.Request) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Type string `json:"type"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error hasing password", http.StatusInternalServerError)
		return
	}

	user := models.User{
		Username: req.Username,
		Password: string(hashedPassword),
		Role: req.Type,
		AvatarId: 1,
	}

	result := h.db.Create(&user)
	if result.Error != nil {
		if result.Error.Error() == "duplicate key violates unique constraints \"idx_users_username\"" {
			http.Error(w, "Username already exists", http.StatusBadRequest)
			return
		}
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

}

