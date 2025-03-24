package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Função de manipulador para a rota "/"
func homeHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Teste Inicial com a GoLang!")
}

// Função de manipulador para a rota "/about"
func aboutHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Primeiras implementações com GoLang 🐀🪽")
}

func main() {
    // Criar um novo roteador
    r := mux.NewRouter()

    // Registrar rotas e suas funções de manipulador
    r.HandleFunc("/", homeHandler).Methods("GET")
    r.HandleFunc("/about", aboutHandler).Methods("GET")

    // Definir o servidor para escutar na porta 8080 usando o roteador mux
    if err := http.ListenAndServe(":8080", r); err != nil {
        fmt.Println("Erro ao iniciar o servidor:", err)
    }
}