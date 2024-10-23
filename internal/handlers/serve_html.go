package handlers

import "net/http"

func ServeIndexHTML(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }

    http.ServeFile(w, r, "web/index.html")
}

func ServeAuthHTML(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "web/auth.html")
}