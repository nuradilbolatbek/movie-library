package handler

import "net/http"

func (h *Handler) searchMovies(w http.ResponseWriter, r *http.Request) {
	// Handle movie search
	queryValues := r.URL.Query()
	movieTitleFragment := queryValues.Get("title")
	actorNameFragment := queryValues.Get("actor")

	// Assuming MovieService has a SearchMovies method
	movies, err := h.MovieService.SearchMovies(movieTitleFragment, actorNameFragment)
	if err != nil {
		NewErrorResponse(w, http.StatusInternalServerError, "error searching movies")
		return
	}

	respondJSON(w, http.StatusOK, movies)

}

func (h *Handler) searchActors(w http.ResponseWriter, r *http.Request) {
	// Handle actor search
	queryValues := r.URL.Query()
	actorNameFragment := queryValues.Get("name")

	// Assuming ActorService has a SearchActors method
	actors, err := h.ActorService.SearchActors(actorNameFragment)
	if err != nil {
		NewErrorResponse(w, http.StatusInternalServerError, "error searching actors")
		return
	}

	respondJSON(w, http.StatusOK, actors)
}
