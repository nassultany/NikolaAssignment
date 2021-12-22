package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

func redirect(w http.ResponseWriter, r *http.Request) {
	// new url path
	newURL := url.URL{Path: r.URL.Query().Get("module")}
	if newURL.Path == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid request: No module specified."))
		return
	}

	// create new query, maintaining all values from old query (except module)
	newQuery := r.URL.Query()
	delete(newQuery, "module")
	newQuery.Add("redirectFrom", r.URL.Path+"?"+r.URL.RawQuery)
	newURL.RawQuery = newQuery.Encode()

	http.Redirect(w, r, newURL.String(), 301)
	// http.Redirect(w, r, newURL.String()+"&redirectFrom="+r.URL.Path+"?"+r.URL.RawQuery, 301)
}

func root(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("You made it!\n"))
	// return query values
	j := json.NewEncoder(w)
	j.SetEscapeHTML(false)
	j.Encode(r.URL.Query())
}

func main() {
	port := flag.Int("port", 8080, "port to start server on")
	flag.Parse()

	http.HandleFunc("/api/testing", redirect)
	http.HandleFunc("/", root)
	fmt.Println("Starting server on port:", *port)
	err := http.ListenAndServe(":"+strconv.Itoa(*port), nil)
	if err != nil {
		log.Fatal("Error: ", err)
	}
}
