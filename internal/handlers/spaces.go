package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Hyperion147/gometaverse/internal/models"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type SpaceHandler struct {
	db *gorm.DB
}

func (h *SpaceHandler) CreateSpace(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name      string `json:"name"`
		Width     int    `json:"width"`
		Height    int    `json:"height"`
		Thumbnail string `json:"thumbnail,omitempty"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid space body", http.StatusBadRequest)
		return
	}

	userID := r.Context().Value("userID").(int)

	space := models.Space{
		Name:      req.Name,
		Width:     req.Width,
		Height:    req.Height,
		OwnerID:   uint(userID),
		Thumbnail: req.Thumbnail,
	}

	result := h.db.Create(&space)
	if result.Error != nil {
		http.Error(w, "Error creating space", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(space)

}

func (h *SpaceHandler) GetAllSpaces(w http.ResponseWriter, r *http.Request) {
	var spaces []models.Space
	
	result := h.db.Preload("Owner").Preload("Elements.Element").Find(&spaces)
	if result.Error != nil {
		http.Error(w, "Error fetching spaces", http.StatusInternalServerError)
		return 
	}

	json.NewEncoder(w).Encode(spaces)

}

func (h *SpaceHandler) GetSpace(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	spaceID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid space ID", http.StatusBadRequest)
		return 
	}

	var space models.Space
	result := h.db.Preload("Owner").Preload("Elements.Element").First(&space, spaceID)
	if result.Error != nil {
		http.Error(w, "Error getting spaces", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(space)

}


