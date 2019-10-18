package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/samtech09/train-api/models"
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

	http.HandleFunc("/create", createItem)

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

func createItem(w http.ResponseWriter, r *http.Request) {
	// create decoder
	decoder := json.NewDecoder(r.Body)
	itm := models.Item{}

	err := decoder.Decode(&itm)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// validate fields
	if len(itm.Title) < 1 || itm.Price < 1 {
		http.Error(w, "Price or title invalid", 500)
		return
	}

	err = saveItem(itm)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Fprint(w, 1)

	//curl -o - -d '{"ID":4, "Title": "Post-book"}' -H "Content-Type: application/json" -X POST http://localhost:3001/create
}
