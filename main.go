package main

import (
	"log"
	"net/http"
	"strconv"
)

func jsonMessage(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	log.Printf("%v", r.URL)
	w.Write([]byte(`{"name":"XYZ"}`))
}

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	log.Printf("%v", r.URL)
	w.Write([]byte("Hello from Snippetbox"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 0 {
		http.Error(w, "No id provided", http.StatusBadRequest)
		log.Print(err)
		return
	}

	log.Printf("%v", r.URL)
	w.Write([]byte("Display a specific snippet...\n" + strconv.Itoa(id)))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	log.Printf("%v", r.URL)
	w.Write([]byte("Create a new snippet..."))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)
	mux.HandleFunc("/json", jsonMessage)

	log.Print("Starting server on :4000")
	err := http.ListenAndServe("0.0.0.0:4000", mux)
	log.Fatal(err)
}
