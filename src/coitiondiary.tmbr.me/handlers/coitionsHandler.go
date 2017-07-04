package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"coitiondiary.tmbr.me/db"
	"github.com/gorilla/mux"
)

func ListCoitions(w http.ResponseWriter, r *http.Request) {
	since := r.FormValue("since")

	coitions, err := db.GetCoitions(since)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, coitions)
}

func CreateCoition(w http.ResponseWriter, r *http.Request) {

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

func UpdateCoition(w http.ResponseWriter, r *http.Request) {

	var coition db.Coition
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&coition); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
		return
	}
	defer r.Body.Close()

	if err := coition.UpdateCoition(); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, coition)

}

func GetCoition(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Coition ID")
		return
	}

	coition := db.Coition{Id: id}
	if err := coition.GetCoition(); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Coition not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, coition)
}

func DeleteCoition(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Coition ID")
		return
	}

	coition := db.Coition{Id: id}
	if err := coition.DeleteCoition(); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}
