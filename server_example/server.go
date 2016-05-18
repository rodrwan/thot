package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type greeting struct {
	Name string `json:"name"`
}

// HTTPServer ...
func HTTPServer() {
	app := negroni.New()
	app.Use(negroni.NewRecovery())
	app.Use(negroni.NewLogger())
	router := mux.NewRouter()

	router.
		HandleFunc("/hello", handler).
		Methods("POST").
		Name("Hello")

	handler := cors.Default().Handler(router)
	app.UseHandler(handler)

	log.Println("Running Example HTTP server")
	log.Fatal(http.ListenAndServe(":8000", app))
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	body, _ := ioutil.ReadAll(r.Body)
	var cheer greeting

	if err := json.Unmarshal(body, &cheer); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")

		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	response := fmt.Sprintf("Hello %s\n", cheer.Name)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}

func main() {
	HTTPServer()
}
