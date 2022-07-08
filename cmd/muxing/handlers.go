package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func nameParamHandlers(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	fmt.Fprintf(w, "Hello, %s!", v["PARAM"])
}

func badHandlers(w http.ResponseWriter, r *http.Request) {
	status := http.StatusText(http.StatusInternalServerError)
	http.Error(w, status, http.StatusInternalServerError)
}

func dataHandlers(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		// log.Printf("data body read: %s", err.Error())
		http.Error(w, "can not read request body", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "I got message:\n%s", body)
}

func headersHandlers(w http.ResponseWriter, r *http.Request) {
	result := 0
	for _, v := range []string{"a", "b"} {
		head := r.Header.Get(v)
		num, err := strconv.Atoi(head)
		if err != nil {
			// log.Printf("header read: %s", err.Error())
			status := fmt.Sprintf("header \"%s\" is not a number %s", v, head)
			http.Error(w, status, http.StatusBadRequest)
			return
		}
		result += num
	}
	w.Header().Add("a+b", strconv.Itoa(result))
}
