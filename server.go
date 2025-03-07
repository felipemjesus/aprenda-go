package main

import (
    "fmt"
    "log"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/hello-word" {
        http.Error(w, "404 não encontrado.", http.StatusNotFound)
        return
    }

    if r.Method != "GET" {
        http.Error(w, "Método não suportado.", http.StatusNotFound)
        return
    }

    fmt.Fprintf(w, "Hello Word!")
}

func main() {
    http.HandleFunc("/hello-word", handler)

    fmt.Printf("Servidor iniciado na porta 8080\n")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}
