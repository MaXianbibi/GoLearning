package main

import (
	"log"
	"net/http"
)

// w = methode pour assembler la réponse
func home(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Hello from Snippetbox"))
}

// Add a snippetView handler function.
func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific snippet..."))
}

// Add a snippetCreate handler function.
func snippetCreate(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {

		w.Header().Set("Allow", "POST")

		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("Create a new snippet..."))
}

func main() {
	// mux = router
	// routeur = fait le lien entre les requêtes et les fonctions qui vont les traiter
	mux := http.NewServeMux()

	// on associe la route "/" avec la fonction home
	// le handler choisie toujours le plus precies
	//

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Print("starting server on :4000")

	// listen and serve = web server
	// arg1 = port
	// arg2 = router
	// boucle infinie
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
