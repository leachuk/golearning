package weretail

import (
	"net/http"
	"fmt"
	"log"
	"strings"
 	"io/ioutil"
	"time"
)

type Homepage struct {
	Title  	string	`json:"jcr:title"`
	Root 	map[string]string	`json:"root"`
}

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
	var client = &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get("http://admin:admin@localhost:4502/content/we-retail/us/en/jcr:content.-1.json")
	if err != nil {
		log.Fatal(err)
	} else {
		//return response as direct pass through
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, string(body))

		//return response as parsed json structure
		//defer resp.Body.Close()
		//
		//homepage := new(Homepage)
		//json.NewDecoder(resp.Body).Decode(homepage)
		//fmt.Fprint(w, homepage.Root["sling:resourceType"])
	}
}

func aboutUsController(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "weretail aboutus controller: %v", r.URL.Path)
}