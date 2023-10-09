package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

/*
	Notas Contenido:
	- Esta función es la que debería ir en el genial.ly, la sección "Haciendo uso de la Ruta de nuestro endpoint" -> query param
*/
func main() {
	router := chi.NewRouter()
	router.Get("/movies", func(w http.ResponseWriter, r *http.Request) {
		// request
		titleLike := r.URL.Query().Get("name")
		awards := r.URL.Query().Get("awards")
		
		// response
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte(
			"Movie: " + titleLike + " Awards: " + awards,
		))
	})

	http.ListenAndServe(":8080", router)
}