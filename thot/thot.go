package thot

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"gopkg.in/tylerb/graceful.v1"
)

var (
	subscribers = make(map[string]*Subscriber)
	port        = flag.Int("port", 8080, "Server port")
)

func subscribeHandler(router *mux.Router, subscribers Subscribers) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var request Subscriber
		body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
		fatal(err)

		if err := json.Unmarshal(body, &request); err != nil {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")

			w.WriteHeader(422) // unprocessable entity
			if err := json.NewEncoder(w).Encode(err); err != nil {
				panic(err)
			}
			return
		}
		fmt.Println(request)
		subscribers.HandleFunc(
			router,
			&request,
			ForwardRequest(request),
		)
		w.WriteHeader(200)
		log.Printf("New endpoint subscribed.\n\n")
	}
}

// Run ...
func Run() {
	flag.Parse()
	app := negroni.New()
	app.Use(negroni.NewRecovery())
	app.Use(negroni.NewLogger())

	router := mux.NewRouter()

	router.
		HandleFunc("/subscribe", subscribeHandler(router, subscribers)).
		Methods("POST").
		Name("Subscribe")

	handler := cors.Default().Handler(router)
	app.UseHandler(handler)

	addr := fmt.Sprintf(":%d", *port)
	s := &http.Server{
		Addr:    addr,
		Handler: app,
	}

	log.Printf("Running proxy server on %s\n", addr)
	err := graceful.ListenAndServe(s, 3*time.Minute)
	fatal(err)
}
