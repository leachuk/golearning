package weretail

import (
	"net/http"
	"fmt"
	"log"
	"strings"
	"io/ioutil"
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
		homepageController(w, r)

	case "aboutus":
		aboutUsController(w, r)

	default:
		fmt.Fprintf(w, "weretail, controller not recognised: %v", controller)
	}
	//http.Error(w, "Some Error", 500)
}

func homepageController(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "weretail homepage controller: %v", r.URL.Path)
	resp, err := http.Get("http://admin:admin@localhost:4502/content/we-retail/us/en/jcr:content.-1.json")
	if err != nil {
		log.Fatal(err)
	} else {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		//json.NewEncoder(w).Encode(string(body))
		fmt.Fprintf(w, "%v", string(body))
	}
}

func aboutUsController(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "weretail aboutus controller: %v", r.URL.Path)
}