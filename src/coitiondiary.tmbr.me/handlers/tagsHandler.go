package handlers

import (
	"fmt"
	"net/http"
)

func CreateTag(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func FindTags(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}
