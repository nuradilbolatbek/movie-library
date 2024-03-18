package handler

import (
	"VkTask/package/service"
	"net/http"
	"strings"
)

type Handler struct {
	services     *service.Service
	authService  *service.AuthService
	MovieService *service.MovieManagementService
	ActorService *service.ActorManagementService
}

func NewHandler(services *service.Service, authService *service.AuthService) *Handler {
	return &Handler{services: services, authService: authService}
}

// ServeHTTP is the main router for incoming HTTP requests
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//log.Printf("Received request: %s %s", r.Method, r.URL.Path) // Log the incoming request
	userIdentityMiddleware := UserIdentity(h.authService)

	switch {
	// User Authentication Endpoints
	case r.URL.Path == "/auth/sign-up" && r.Method == "POST":
		h.signUp(w, r)
	case r.URL.Path == "/auth/sign-in" && r.Method == "GET":
		h.signIn(w, r)

		// Actor Management Endpoints
		userIdentityMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			switch {
			case r.URL.Path == "/api/actors" && r.Method == "POST":
				h.createActor(w, r)
			case r.URL.Path == "/api/actors" && r.Method == "GET":
				h.getAllActors(w, r)
			case strings.HasPrefix(r.URL.Path, "/api/actors/") && r.Method == "PUT":
				h.updateActor(w, r)
			case strings.HasPrefix(r.URL.Path, "/api/actors/") && r.Method == "DELETE":
				h.deleteActor(w, r)
			}
		})).ServeHTTP(w, r)
	// Movie Management Endpoints
	case strings.HasPrefix(r.URL.Path, "/api/movies"):
		userIdentityMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			switch r.Method {
			case "POST":
				h.createMovie(w, r)
			case "GET":
				h.getAllMovies(w, r)
			case "PUT":
				h.updateMovie(w, r)
			case "DELETE":
				h.deleteMovie(w, r)
			}
		})).ServeHTTP(w, r)

	// Search and Filtering Endpoints
	case r.URL.Path == "/api/search/movies" && r.Method == "GET":
		h.searchMovies(w, r)
	case r.URL.Path == "/api/search/actors" && r.Method == "GET":
		h.searchActors(w, r)

	// Default Case for Undefined Routes
	default:
		http.NotFound(w, r)
	}
}
