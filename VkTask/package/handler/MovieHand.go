package handler

import (
	"VkTask"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func (h *Handler) createMovie(w http.ResponseWriter, r *http.Request) {
	// Handle movie creation

	userRole, err := GetUserRole(r)
	if err != nil {
		respondJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	if userRole != "admin" {
		respondJSON(w, http.StatusForbidden, "access denied")
		return
	}

	var movie VkTask.Movie
	err = json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		respondJSON(w, http.StatusBadRequest, "invalid request body")
		return
	}

	id, err := h.MovieService.CreateMovie(movie)
	if err != nil {
		respondJSON(w, http.StatusInternalServerError, "error creating movie")
		return
	}

	respondJSON(w, http.StatusCreated, map[string]int{"id": id})

}

func (h *Handler) getAllMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := h.MovieService.GetAllMovies()
	if err != nil {
		NewErrorResponse(w, http.StatusInternalServerError, "error fetching movies")
		return
	}

	respondJSON(w, http.StatusOK, movies)
}

func (h *Handler) getMovieById(w http.ResponseWriter, r *http.Request) {
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 3 {
		NewErrorResponse(w, http.StatusBadRequest, "invalid URL path")
		return
	}

	id, err := strconv.Atoi(pathParts[2])
	if err != nil {
		NewErrorResponse(w, http.StatusBadRequest, "invalid movie id")
		return
	}

	movie, err := h.MovieService.GetMovieByID(id)
	if err != nil {
		NewErrorResponse(w, http.StatusInternalServerError, "error fetching movie")
		return
	}

	respondJSON(w, http.StatusOK, movie)
}
func (h *Handler) updateMovie(w http.ResponseWriter, r *http.Request) {
	userRole, err := GetUserRole(r)
	if err != nil {
		respondJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	if userRole != "admin" {
		respondJSON(w, http.StatusForbidden, "access denied")
		return
	}
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 3 {
		NewErrorResponse(w, http.StatusBadRequest, "invalid URL path")
		return
	}

	id, err := strconv.Atoi(pathParts[2])
	if err != nil {
		NewErrorResponse(w, http.StatusBadRequest, "invalid movie id")
		return
	}

	var movie VkTask.Movie
	err = json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		NewErrorResponse(w, http.StatusBadRequest, "invalid request body")
		return
	}

	err = h.MovieService.UpdateMovie(id, movie)
	if err != nil {
		NewErrorResponse(w, http.StatusInternalServerError, "error updating movie")
		return
	}

	NewStatusResponse(w, http.StatusOK, "movie updated successfully")
}

func (h *Handler) deleteMovie(w http.ResponseWriter, r *http.Request) {
	userRole, err := GetUserRole(r)
	if err != nil {
		respondJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	if userRole != "admin" {
		respondJSON(w, http.StatusForbidden, "access denied")
		return
	}
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 3 {
		NewErrorResponse(w, http.StatusBadRequest, "invalid URL path")
		return
	}

	id, err := strconv.Atoi(pathParts[2])
	if err != nil {
		NewErrorResponse(w, http.StatusBadRequest, "invalid movie id")
		return
	}

	err = h.MovieService.DeleteMovie(id)
	if err != nil {
		NewErrorResponse(w, http.StatusInternalServerError, "error deleting movie")
		return
	}

	NewStatusResponse(w, http.StatusOK, "movie deleted successfully")
}
func (h *Handler) getAllMoviesSorted(w http.ResponseWriter, r *http.Request) {
	sortParam := r.URL.Query().Get("sort")

	movies, err := h.MovieService.GetAllMoviesSorted(sortParam)
	if err != nil {
		NewErrorResponse(w, http.StatusInternalServerError, "error retrieving sorted movies")
		return
	}

	respondJSON(w, http.StatusOK, movies)
}
