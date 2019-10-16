package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func init() {
	initDB()
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/json", jsonTest)
	http.HandleFunc("/paramtest", paramTest)
	// http.HandleFunc("/list", getrecords)
	// http.HandleFunc("/getbyid", getbyid)

	listhandler := http.HandlerFunc(getrecords)
	byidhandler := http.HandlerFunc(getbyid)

	http.Handle("/list", loggerMiddleware(listhandler))
	http.Handle("/getbyid", loggerMiddleware(byidhandler))

	fmt.Println("Server listening on port 3001...")
	http.ListenAndServe(":3001", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bare minimum API server in go with JSON and Database calls")
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<html><body><h1>Hello from golang!</h1></body></html>")
}

func jsonTest(w http.ResponseWriter, r *http.Request) {
	// set header first before writing to response
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode("JSON from golang!")
}

func paramTest(w http.ResponseWriter, r *http.Request) {
	tmp, ok := r.URL.Query()["key"]
	if !ok || len(tmp[0]) < 1 {
		http.Error(w, "No parameter passed", 500)
		return
	}

	fmt.Fprintf(w, "key = %s", tmp[0])
}

func getrecords(w http.ResponseWriter, r *http.Request) {
	list, err := getRecordsFromDB()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// set header first before writing to response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(list)
}

func getbyid(w http.ResponseWriter, r *http.Request) {
	// get querystring from url
	tmp, ok := r.URL.Query()["id"]
	if !ok || len(tmp[0]) < 1 {
		http.Error(w, "Roll number missing", 403)
		return
	}

	id, err := strconv.Atoi(tmp[0])
	if err != nil {
		http.Error(w, "Id number must be number", 500)
		return
	}

	itm, err := getByIDFromDB(id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// set header first before writing to response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(itm)
}
