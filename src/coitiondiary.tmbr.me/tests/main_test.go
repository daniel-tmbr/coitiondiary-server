package main_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"
	"time"

	"coitiondiary.tmbr.me/db"
	"coitiondiary.tmbr.me/main"
)

const (
	host     = "localhost"
	port     = 1234
	user     = "danieltmbr"
	password = ""
	dbname   = "coitiondiary_test"
)

var a main.App

func TestMain(m *testing.M) {

	a = main.App{}
	a.Initialize(host, user, password, dbname, port)

	code := m.Run()

	clearTables()

	os.Exit(code)
}

// Private methods

func clearTables() {

	tables := []string{
		"coitiontags",
		"coitionmedias",
		"participant",
		"tag",
		"coition",
		"location",
		"media",
		"user",
	}

	for _, table := range tables {
		clearTable(table)
	}
}

func clearTable(name string) {
	db.DB.Exec("DELETE FROM " + name)
	db.DB.Exec("ALTER SEQUENCE " + name + "_id_seq RESTART WITH 1")
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func addCoitions(count int) {
	if count < 1 {
		count = 1
	}

	date := time.Now().UTC().Format(time.RFC3339)
	for i := 0; i < count; i++ {
		db.DB.Exec("INSERT INTO coition(name, date, note, created, updated) VALUES($1, $2, $3, $4, $5)",
			"Test coition no.:"+strconv.Itoa(i), date, "Test note", date, date)
	}
}

// Public test methods

func TestEmptyTable(t *testing.T) {
	clearTables()

	req, _ := http.NewRequest("GET", "/coitions", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body != "[]" {
		t.Errorf("Expected an empty array. Got %s", body)
	}
}

func TestGetNonExistentCoition(t *testing.T) {
	clearTables()

	req, _ := http.NewRequest("GET", "/coitions/11", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusNotFound, response.Code)

	var m map[string]string
	json.Unmarshal(response.Body.Bytes(), &m)
	if m["error"] != "Coition not found" {
		t.Errorf("Expected the 'error' key of the response to be set to 'Coition not found'. Got '%s'", m["error"])
	}
}

func TestCreateCoition(t *testing.T) {
	clearTables()

	date := time.Now().UTC().Format(time.RFC3339) //.Format("yyyy-MM-ddThh:MM:ss-0700")
	payload := []byte(`{"name": "Test sex", "date": "` + date + `","note": "Test sex note"}`)

	req, _ := http.NewRequest("POST", "/coitions", bytes.NewBuffer(payload))
	response := executeRequest(req)

	checkResponseCode(t, http.StatusCreated, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["name"] != "Test sex" {
		t.Errorf("Expected coition name to be 'Test sex'. Got '%v'", m["name"])
	}

	if m["date"] != date {
		t.Errorf("Expected coition date to be '%v'. Got '%v'", date, m["date"])
	}

	if m["note"] != "Test sex note" {
		t.Errorf("Expected coition price to be '11.22'. Got '%v'", m["note"])
	}

	// the id is compared to 1.0 because JSON unmarshaling converts numbers to
	// floats, when the target is a map[string]interface{}
	if m["id"] != 1.0 {
		t.Errorf("Expected coition ID to be '1'. Got '%v'", m["id"])
	}
}

func TestGetCoition(t *testing.T) {
	clearTables()
	addCoitions(1)

	req, _ := http.NewRequest("GET", "/coitions/1", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)
}

func TestUpdateCoition(t *testing.T) {

	clearTables()
	addCoitions(1)

	req, _ := http.NewRequest("GET", "/coitions/1", nil)
	response := executeRequest(req)
	var originalCoition map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &originalCoition)

	originalCoition["name"] = "Updated test coition"
	payload, _ := json.Marshal(&originalCoition)
	//payload := []byte(`{"name":"test product - updated name","price":11.22}`)

	req, _ = http.NewRequest("PUT", "/coitions", bytes.NewBuffer(payload))
	response = executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	var updatedCoition map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &updatedCoition)

	if updatedCoition["id"] != originalCoition["id"] {
		t.Errorf("Expected the id to remain the same (%v). Got %v", originalCoition["id"], updatedCoition["id"])
	}

	if updatedCoition["name"] != originalCoition["name"] {
		t.Errorf("Expected the name to change from '%v' to 'Updated test coition'. Got '%v'", originalCoition["name"], updatedCoition["name"])
	}
}

func TestDeleteCoition(t *testing.T) {
	clearTables()
	addCoitions(1)

	req, _ := http.NewRequest("GET", "/coitions/1", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

	req, _ = http.NewRequest("DELETE", "/coitions/1", nil)
	response = executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	req, _ = http.NewRequest("GET", "/coitions/1", nil)
	response = executeRequest(req)
	checkResponseCode(t, http.StatusNotFound, response.Code)
}
