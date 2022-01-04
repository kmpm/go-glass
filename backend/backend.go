package backend

import (
    "fmt"
    "net/http"
)

func SetupRoutes() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Simple Server")
    })
}
