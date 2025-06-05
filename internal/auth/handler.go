package auth

import (
	"app/breafURL/configs"
	"app/breafURL/pkg/response"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
)

type AuthHandlerDeps struct {
	*configs.Config
}

type AuthHandler struct {
	*configs.Config
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config: deps.Config,
	}
	router.HandleFunc("POST /auth/login", handler.login())
	router.HandleFunc("POST /auth/register", handler.register())
}

func (handler *AuthHandler) login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var payload LoginRequest
		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			response.ResponseWithJson(w, err.Error(), http.StatusBadRequest)
			return
		}

		reg, _ := regexp.Compile(`[A-Za-z0-9\._%+\-]+@[A-Za-z0-9\.\-]+\.[A-Za-z]{2,}`)
		if !reg.MatchString(payload.Email) {
			response.ResponseWithJson(w, "Wrong email", http.StatusBadRequest)
			return
		}

		if payload.Password == "" {
			response.ResponseWithJson(w, "Wrong password", http.StatusBadRequest)
			return
		}

		data := LoginResponse{
			Token: "testing",
		}

		response.ResponseWithJson(w, data, http.StatusOK)
	}
}

func (handler *AuthHandler) register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Register")
	}
}
