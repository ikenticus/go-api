package main

import (
	"encoding/json"
//	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"github.com/gorilla/mux"
)

/*
	TEST CURL:
	curl -H "Content-Type: application/json" -d '{"category": "generic", "test": true, "serial": 1}' http://localhost:8080/load/generic_test_document_key
*/
func Loader(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var docKey string
	docKey = strings.Replace(vars["docKey"], "/", "_", -1)

	var docJson interface{}
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &docJson); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	//fmt.Printf("JSON document is %s\n", string(body))
	DocStoreUpsert(docKey, string(body))
	/*
	d := DocStoreUpsert(docKey, docJson)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(d); err != nil {
		panic(err)
	}
	*/
}
