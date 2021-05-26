package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {
	http.HandleFunc("/calc/faverage", fluorescentAverageHandler)
	http.HandleFunc("/calc/active_cells", activeCellHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func fluorescentAverageHandler(w http.ResponseWriter, req *http.Request) {
	if validateRequest(w, req, http.MethodPost) != nil {
		return
	}
}

func activeCellHandler(w http.ResponseWriter, req *http.Request) {
	if validateRequest(w, req, http.MethodPost) != nil {
		return
	}
}

func errorResponse(w http.ResponseWriter, err error, status int) {
	http.Error(w, fmt.Sprint(err), status)
}

func validateRequest(w http.ResponseWriter, req *http.Request, expectedMethod string) error {
	if req.Method != expectedMethod {
		err := errors.New("access denied")
		errorResponse(w, err, http.StatusBadRequest)
		return err
	}

	dump, err := httputil.DumpRequest(req, true)
	if err != nil {
		errorResponse(w, err, http.StatusInternalServerError)
		return err
	}
	fmt.Println(string(dump))
	return nil
}
