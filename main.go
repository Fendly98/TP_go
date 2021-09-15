package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func tpIntermediaire(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		fmt.Fprintf(w, time.Now().Format("15h04"))
	}
}

func tpIntermediaire2(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:

		author := req.FormValue("author")
		entry := req.FormValue("entry")

		f, _ := os.OpenFile("save", os.O_RDWR|os.O_CREATE, 0755)
		fmt.Fprintln(f, author+" : "+entry+"\n")
		fmt.Fprintf(w, author+" : "+entry)

	}
}

func main() {
	http.HandleFunc("/", tpIntermediaire)
	http.HandleFunc("/add", tpIntermediaire2)
	http.ListenAndServe(":4567", nil)
}
