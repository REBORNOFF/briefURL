package main

import (
	"app/breafURL/configs"
	"app/breafURL/internal/auth"
	"fmt"
	"net/http"
)

func main() {
	config := configs.LoadConfig()
	router := http.NewServeMux()

	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config: config,
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	fmt.Println("Server started on 8080 port")
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err.Error())
	}
}
