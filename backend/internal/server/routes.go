package server

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/golang-jwt/jwt/v5"
	"brie/internal/auth"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	r.Get("/", s.HelloWorldHandler)

	r.Get("/health", s.healthHandler)

    r.Post("/auth/sign-in",s.CreateUser)
    r.Post("/auth/log-in",s.LoginUser)

    // routes in which u need to be authenticated pr-protected router
    r.Group(func(pr chi.Router) {
        pr.Use(auth.JWTMiddleware)
        pr.Get("/protected",s.ProtectedHello)
		pr.Post("/auth/log-out",s.LogoutUser)
    })

	return r
}


func (s *Server) ProtectedHello(w http.ResponseWriter, r *http.Request) {
    claims, ok := r.Context().Value(auth.ClaimsContextKey).(jwt.MapClaims)
    if !ok {
        http.Error(w, "Failed to get claims", http.StatusUnauthorized)
        return
    }

    userID := claims["user_id"].(string)
    username := claims["username"].(string)

    w.Write([]byte("Hello " + username + "! Your ID is " + userID))
}


func (s *Server) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, _ := json.Marshal(s.db.Health())
	_, _ = w.Write(jsonResp)
}
