package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/arl/statsviz"
	example "github.com/arl/statsviz/_example"
	"github.com/gorilla/mux"
)

const (
	PORT = ":8080"
)

func pageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pageID := vars["id"]
	fileName := "files/" + pageID + ".html"
	_, err := os.Stat(fileName)
	if err != nil {
		fileName = "files/404.html"
	}
	http.ServeFile(w, r, fileName)
}

func work() {
	m := map[string][]byte{}

	for {
		b := make([]byte, 512+rand.Intn(16*1024))
		m[strconv.Itoa(len(m)%(10*100))] = b

		if len(m)%(10*100) == 0 {
			m = make(map[string][]byte)
		}
		time.Sleep(10 * time.Millisecond)
	}
}

func ArticlesCategoryHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category: %v\n", vars["category"])
}

func main() {
	go example.Work()

	rtr := mux.NewRouter()
	rtr.Methods("GET").Path("/debug/statsviz/ws").Name("GET /debug/statsviz/ws").HandlerFunc(statsviz.Ws)
	rtr.UseEncodedPath().Methods("GET").PathPrefix("/debug/statsviz").Name("GET /debug/statsviz/").Handler(statsviz.Index)
	rtr.HandleFunc("/pages/{id:[0-9]+}", pageHandler)
	http.Handle("/", rtr)
	rtr.HandleFunc("/article", ArticlesCategoryHandler).
		Methods("GET").
		Schemes("http")
	http.ListenAndServe(PORT, nil)
}
