package main

import (
	"net/http"

	"github.com/kmpm/go-glass/backend"
)

func main() {
	backend.SetupRoutes()
	http.ListenAndServe(":8080", nil)
}
