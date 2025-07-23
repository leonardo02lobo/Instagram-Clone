package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "¡Bienvenido al servidor nativo de Go!")
	})

	http.HandleFunc("/saludo", func(w http.ResponseWriter, r *http.Request) {
		nombre := r.URL.Query().Get("nombre")
		if nombre == "" {
			nombre = "Visitante"
		}
		fmt.Fprintf(w, "¡Hola, %s!", nombre)
	})

	fmt.Println("Servidor iniciado en http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
