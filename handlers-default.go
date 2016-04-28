package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to GO API!\n")
}

func Presenter(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var docKey string
	docKey = strings.Replace(vars["docKey"], "/", "_", -1)
	//fmt.Fprint(w, "Doc Key: ", docKey, "\n")

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	var docJson interface{}
	docJson = DocStoreGet(docKey)
	if docJson != nil {
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(docJson); err != nil {
			panic(err)
		}
		return
	}

	// 404 when no document found
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
		panic(err)
	}
}