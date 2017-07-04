package handlers

import (
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func Register(w http.ResponseWriter, r *http.Request) {

	var coition db.Coition
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&coition); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := coition.CreateCoition(); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, coition)
}

func Login(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	// w.WriteHeader(http.StatusOK)

	// if err := json.NewEncoder(w).Encode(db.Todos); err != nil {
	// 	panic(err)
	// }
}

func Logout(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	// todoId := vars["todoId"]
	// fmt.Fprintln(w, "Todo show:", todoId)
}

// Private methods

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
