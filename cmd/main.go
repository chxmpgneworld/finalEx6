package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "[Server]", log.Ldate|log.Ltime|log.Lshortfile)
	srv := server.NewRouter(logger)
	err := srv.Http.ListenAndServe()
	if err != nil {
		logger.Fatal(err)
	}
}
