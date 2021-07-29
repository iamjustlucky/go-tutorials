package main

import (
	"hello-world/handler"
	"log"
	"net/http"
)

func main(){
    mux := http.NewServeMux()

    mux.HandleFunc("/", handler.HandleIndex)
    mux.HandleFunc("/home", handler.HandleHome)
    mux.HandleFunc("/about", handler.HandleAbout)
    mux.HandleFunc("/product", handler.HandleProduct)
    mux.HandleFunc("/post-get", handler.PostGet)
    mux.HandleFunc("/form", handler.Form)
    mux.HandleFunc("/process", handler.Process)
    

    fileServer := http.FileServer(http.Dir("assets"))
    mux.Handle("/static/", http.StripPrefix("/static", fileServer))

    log.Println("Starting web on port 9000")

    err := http.ListenAndServe(":9000", mux)
    log.Fatal(err)
}
