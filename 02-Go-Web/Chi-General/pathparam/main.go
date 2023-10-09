package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

/*
	Notas Contenido:
	- Esta función es la que debería ir en el genial.ly, la sección "Haciendo uso de la Ruta de nuestro endpoint" -> path param
*/
func main() {
	router := chi.NewRouter()
	router.Get("/movies", func(w http.ResponseWriter, r *http.Request) {
		// request
		name := chi.URLParam(r, "name")

		// response
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("Movie " + name))
	})

	http.ListenAndServe(":8080", router)
}