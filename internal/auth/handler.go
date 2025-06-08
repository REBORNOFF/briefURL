package auth

import (
	"app/breafURL/configs"
	"app/breafURL/pkg/request"
	"app/breafURL/pkg/response"
	"fmt"
	"net/http"
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
		body, err := request.HandleBody[LoginRequest](w, r)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(body)

		data := LoginResponse{
			Token: "testing",
		}

		response.ResponseWithJson(w, data, http.StatusOK)
	}
}

func (handler *AuthHandler) register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := request.HandleBody[RegisterRequest](w, r)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(body)
	}
}
