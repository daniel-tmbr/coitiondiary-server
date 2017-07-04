package handlers

import (
	"fmt"
	"net/http"
)

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	// w.WriteHeader(http.StatusOK)

	// if err := json.NewEncoder(w).Encode(db.Todos); err != nil {
	// 	panic(err)
	// }
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	// todoId := vars["todoId"]
	// fmt.Fprintln(w, "Todo show:", todoId)
}

func UploadUserAvatar(w http.ResponseWriter, r *http.Request) {
	// var todo models.Todo
	// body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	// if err != nil {
	// 	panic(err)
	// }
	// if err := r.Body.Close(); err != nil {
	// 	panic(err)
	// }
	// if err := json.Unmarshal(body, &todo); err != nil {
	// 	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	// 	w.WriteHeader(422) // unprocessable entity
	// 	if err := json.NewEncoder(w).Encode(err); err != nil {
	// 		panic(err)
	// 	}
	// }

	// t := db.RepoCreateTodo(todo)
	// w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	// w.WriteHeader(http.StatusCreated)
	// if err := json.NewEncoder(w).Encode(t); err != nil {
	// 	panic(err)
	// }
}

func FindUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func CreateDummyPartner(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	// w.WriteHeader(http.StatusOK)

	// if err := json.NewEncoder(w).Encode(db.Todos); err != nil {
	// 	panic(err)
	// }
}

func UpdateDummyPartner(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	// todoId := vars["todoId"]
	// fmt.Fprintln(w, "Todo show:", todoId)
}

func MergeDummyPartner(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	// todoId := vars["todoId"]
	// fmt.Fprintln(w, "Todo show:", todoId)
}
