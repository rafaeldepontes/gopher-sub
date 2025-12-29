package main

import (
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/rafaeldepontes/gopher-sub/internal/middleware"
	"github.com/rafaeldepontes/gopher-sub/internal/migration"
	"github.com/rafaeldepontes/gopher-sub/internal/tool"
	"github.com/rafaeldepontes/gopher-sub/pkg/database/postgres"
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
	defer postgres.Close()

	if err := migration.Init(); err != nil {
		handle.Log.Errorln("Migration error: ", err)
		panic(err)
	}

	handle.Log.Infoln("Application running on ", "http://localhost:"+serverPort)

	if err := http.ListenAndServe(":"+serverPort, r); err != nil {
		handle.Log.Errorln("Error initializing the server, ", err)
	}
}
