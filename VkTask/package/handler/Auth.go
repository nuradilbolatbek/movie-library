package handler

import (
	"VkTask"
	"encoding/json"
	"net/http"
)

type UserSignInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signUp(w http.ResponseWriter, r *http.Request) {
	// Handle user sign-up
	var user VkTask.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		NewErrorResponse(w, http.StatusBadRequest, "invalid input body") // Using response package
		return
	}

	id, err := h.services.Authorization.CreateUser(user)
	if err != nil {
		NewErrorResponse(w, http.StatusInternalServerError, err.Error()) // Using response package
		return
	}

	respondJSON(w, http.StatusOK, map[string]interface{}{"id": id})

}

func (h *Handler) signIn(w http.ResponseWriter, r *http.Request) {
	var input UserSignInInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		NewErrorResponse(w, http.StatusBadRequest, "invalid input body") // Using response package
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		NewErrorResponse(w, http.StatusInternalServerError, err.Error()) // Using response package
		return
	}

	respondJSON(w, http.StatusOK, map[string]interface{}{"token": token})
}
