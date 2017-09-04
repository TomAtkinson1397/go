package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"strings"
	"bytes"

	"github.com/gorilla/mux"
)

var keywords = []byte("keywords")

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/search={searchTerm}", search)
	fmt.Println("server listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func search(w http.ResponseWriter, r *http.Request) {
	var buf bytes.Buffer
	vars := mux.Vars(r)
	searchTerm := vars["searchTerm"]
	k, t := splitSearchTerm(searchTerm)
	tSize := strings.Count(t, "")
	url, err := getUrl(k)
	if err != nil {
		panic(err)
	}
	d, sd := splitUrl(url)
	buf.WriteString("https://")
	buf.WriteString(d)
	buf.WriteString("/")
	if tSize != 1 {
		buf.WriteString(sd)
		buf.WriteString(t)
	}
	url = buf.String()
	http.Redirect(w, r, url, 302)
	log.Println("db returned: ", url)
	log.Println("received request for ", html.EscapeString(r.URL.Path[8:]))
}

func splitUrl(url string) (string, string) {
	var domain string;
	var subdomain string;

	splits := strings.Split(url, "/")
	domain = splits[0]
	subdomain = splits[1]
	return domain, subdomain
}

func splitSearchTerm(searchTerm string) (string, string) {
	var keyword string
	var term string

	if strings.Contains(searchTerm, "-") {
		splits := strings.Split(searchTerm, "-")
		keyword = splits[0]
		term = splits[1]
		return keyword, term
	} else {
		return searchTerm, ""
	}
}
