package main

import (
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/rafaeldepontes/gopher-sub/internal/middleware"
	"github.com/rafaeldepontes/gopher-sub/internal/tool"
)

func main() {
	env := ".env"
	tool.ChecksEnv(&env)
	godotenv.Load(env)

	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		serverPort = "8000"
	}

	r := chi.NewRouter()
	handle := middleware.NewHandle()
	handle.ConfigHandler(r)

 	handle.Log.Infoln("Application running on ", "http://localhost:"+serverPort)

	http.ListenAndServe(":"+serverPort, r)
}
