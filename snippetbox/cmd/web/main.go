package main

import (
    "log"
    "net/http"
    "flag"
)

func main() {
    mux := http.NewServeMux()



    addr := flag.String("addr", ":4000", "HTTP network address")
    flag.Parse()

    /// tous les fhichiers dans le dossier static sont servis
    fileServer := http.FileServer(http.Dir("./ui/static/"))
    mux.Handle("/static/", http.StripPrefix("/static", fileServer))






    mux.HandleFunc("/", home)
    mux.HandleFunc("/snippet/view", snippetView)
    mux.HandleFunc("/snippet/create", snippetCreate)

    log.Print("starting server on :", *addr)
    
    err := http.ListenAndServe(*addr, mux)
    log.Fatal(err)
}
