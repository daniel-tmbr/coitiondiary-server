package handlers

import (
	"fmt"
	"net/http"
)

func GetNearbyLocations(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func FindLocations(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	// w.WriteHeader(http.StatusOK)

	// if err := json.NewEncoder(w).Encode(db.Todos); err != nil {
	// 	panic(err)
	// }
}
