package main

import (
	"fmt"
	"net/http"
	"github.com/leachuk/golearning/http/weretail"
	"time"
)


type String string

func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, s)
}


type Struct struct {
	Greeting string
	Punct string
	Who string
}

func (s *Struct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, s.Greeting, s.Punct, s.Who)
}

func handleQuery(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Got query: %v", r.URL.Query())
}

func handleSite1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Site 1: %v", r.URL.Path)
}

func handleSite2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Site 2: %v", r.URL.Path)
}


func main() {
	mux := http.NewServeMux()
	//http.Handle("/string", String("I'm a frayed knot."))
	//http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})
	//http.HandleFunc("/querytest", handleQuery)
	mux.HandleFunc("/site1/", handleSite1)
	//http.HandleFunc("/site2/", handleSite2)

	mux.HandleFunc("/weretail/", weretail.HandleWeretail)

	handler := http.TimeoutHandler(mux, 5*time.Second, "Request Timeout")

	server := &http.Server{
		Addr: "localhost:4000",
		Handler: handler,
		ReadTimeout: 30*time.Second,
		WriteTimeout: 35*time.Second,
	}
	server.ListenAndServe()
}