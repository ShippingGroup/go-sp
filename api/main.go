package main

import (
	"fmt"
	"log"
	"net/http"

	"encoding/json"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/contacto", Contact)
	router.HandleFunc("/peliculas", MovieList)
	router.HandleFunc("/pelicula/{id}", MovieShow)

	server := http.ListenAndServe(":8080", router)

	log.Fatal(server)

}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola mundo desde mi servidor web con GO")
}

func Contact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Esta es la página de contacto!")
}

func MovieList(w http.ResponseWriter, r *http.Request) {
	movies := Movies{
		Movie{"Sin límetes", 2013, "Desconocido"},
		Movie{"El Demoledor", 1999, "Russell Crox"},
		Movie{"Corazón de León", 2020, "Martín Sifrón"},
	}
	//fmt.Fprintf(w, "Listado de películas")
	json.NewEncoder(w).Encode(movies)
}

func MovieShow(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movie_id := params["id"]
	fmt.Fprintf(w, "Seleccionaste la película número %s", movie_id)
}
