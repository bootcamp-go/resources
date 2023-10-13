package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func HandlerHealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func HandlerGetItems(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("GetItems"))
}

func HandlerSearchItems(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("SearchItems"))
}

/*
	Notas Contenido:
	- Esta es la función "Ejemplo sin Agrupar" del archivo "main.go"
*/
// func main() {
// 	// create a new router
// 	router := chi.NewRouter()

// 	// register endpoints: method + path + handler
// 	router.Get("/health", HandlerHealthCheck)
// 	router.Get("/items", HandlerGetItems)
// 	router.Get("/items/search", HandlerSearchItems)

// 	// run server
// 	http.ListenAndServe(":8080", router)
// }

func HandlerProfileAbout(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ProfileAbout"))
}

func HandlerProfileData(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ProfileData"))
}

func HandlerProfilePictures(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ProfilePictures"))
}

func HandlerProfileFriends(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ProfileFriends"))
}

/*
	Notas Contenido:
	- Esta es la función "Ejemplo Agrupado" del archivo "main.go"
*/
func main() {
	// create a new router
	router := chi.NewRouter()

	// register endpoints: method + path + handler
	// - health: checks if the service is up and running
	router.Get("/health", HandlerHealthCheck)
	// - profile group: handles all profile related endpoints
	router.Route("/profile", func(r chi.Router) {
		r.Get("/about", HandlerProfileAbout)
		r.Get("/data", HandlerProfileData)
		r.Get("/pictures", HandlerProfilePictures)
		r.Get("/friends", HandlerProfileFriends)
	})
		
	// run server
	if err := http.ListenAndServe(":8080", router); err != nil {
		fmt.Println(err)
		return
	}
}