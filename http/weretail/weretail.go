package weretail

import (
	"net/http"
	"fmt"
	"time"
)

func HandleWeretail(w http.ResponseWriter, r *http.Request) {
	time.Sleep(7 * time.Second) // Need a way of handling long running methods gracefully
	fmt.Fprintf(w, "weretail: %v", r.URL.Path)
	//http.Error(w, "Some Error", 500)
}