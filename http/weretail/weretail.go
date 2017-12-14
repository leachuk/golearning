package weretail

import (
	"net/http"
	"fmt"
	"log"
	"strings"
)

func HandleWeretail(w http.ResponseWriter, r *http.Request) {
	//time.Sleep(7 * time.Second) // Need a way of handling long running methods gracefully. Done with http.TimeoutHandler
	//parse URL
	u, err := r.URL.Parse(r.URL.Path)
	if err != nil {
		log.Fatal(err)
	}
	urlSplit := strings.Split(u.RequestURI(), "/")

	controller := urlSplit[2]
	switch controller {

	case "homepage":
		homepageController(w)

	case "aboutus":
		aboutUsController(w)

	default:
		fmt.Fprintf(w, "weretail, controller not recognised: %v", controller)
	}
	//http.Error(w, "Some Error", 500)
}

func homepageController(w http.ResponseWriter,) {
	fmt.Fprintf(w, "weretail homepage controller: %v", "foo")
}

func aboutUsController(w http.ResponseWriter,) {
	fmt.Fprintf(w, "weretail aboutus controller: %v", "foo")
}