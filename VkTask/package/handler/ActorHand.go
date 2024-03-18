package handler

import (
	"VkTask"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"
)

func (h *Handler) createActor(w http.ResponseWriter, r *http.Request) {
	userRole, err := GetUserRole(r)
	if err != nil {
		respondJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	if userRole != "admin" {
		respondJSON(w, http.StatusForbidden, "access denied")
		return
	}
	var actor VkTask.Actor
	err = json.NewDecoder(r.Body).Decode(&actor)
	if err != nil {
		NewErrorResponse(w, http.StatusBadRequest, "invalid request body")
		return
	}

	id, err := h.ActorService.CreateActor(actor)
	if err != nil {
		NewErrorResponse(w, http.StatusInternalServerError, "error creating actor")
		return
	}

	respondJSON(w, http.StatusCreated, map[string]int{"id": id})
}

func (h *Handler) getAllActors(w http.ResponseWriter, r *http.Request) {
	actors, err := h.ActorService.GetAllActors()
	if err != nil {
		NewErrorResponse(w, http.StatusInternalServerError, "error fetching actors")
		return
	}

	respondJSON(w, http.StatusOK, actors)
}

func (h *Handler) getActorById(w http.ResponseWriter, r *http.Request) {
	id, err := parseIDFromURL(r.URL.Path)
	if err != nil {
		NewErrorResponse(w, http.StatusBadRequest, "invalid actor id")
		return
	}

	actor, err := h.ActorService.GetActorByID(id)
	if err != nil {
		NewErrorResponse(w, http.StatusInternalServerError, "error fetching actor")
		return
	}

	respondJSON(w, http.StatusOK, actor)
}

func (h *Handler) updateActor(w http.ResponseWriter, r *http.Request) {
	id, err := parseIDFromURL(r.URL.Path)
	if err != nil {
		NewErrorResponse(w, http.StatusBadRequest, "invalid actor id")
		return
	}

	var actor VkTask.Actor
	err = json.NewDecoder(r.Body).Decode(&actor)
	if err != nil {
		NewErrorResponse(w, http.StatusBadRequest, "invalid request body")
		return
	}

	err = h.ActorService.UpdateActor(id, actor)
	if err != nil {
		NewErrorResponse(w, http.StatusInternalServerError, "error updating actor")
		return
	}

	NewStatusResponse(w, http.StatusOK, "actor updated successfully")
}

func (h *Handler) deleteActor(w http.ResponseWriter, r *http.Request) {
	id, err := parseIDFromURL(r.URL.Path)
	if err != nil {
		NewErrorResponse(w, http.StatusBadRequest, "invalid actor id")
		return
	}

	err = h.ActorService.DeleteActor(id)
	if err != nil {
		NewErrorResponse(w, http.StatusInternalServerError, "error deleting actor")
		return
	}

	NewStatusResponse(w, http.StatusOK, "actor deleted successfully")
}

func parseIDFromURL(urlPath string) (int, error) {
	parts := strings.Split(urlPath, "/")
	if len(parts) < 3 {
		return 0, errors.New("invalid URL path")
	}
	return strconv.Atoi(parts[2])
}
