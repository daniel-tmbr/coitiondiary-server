package db

import (
	"fmt"
	"time"
)

// Coition model
type Coition struct {
	Id      int       `json:"id"`
	Name    string    `json:"name"`
	Date    time.Time `json:"date"`
	Note    string    `json:"note"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}

// Coitions slice, containing Coition models
type Coitions []Coition

func (c *Coition) GetCoition() error {
	return DB.QueryRow("SELECT id, name, date, note, created, updated FROM coition WHERE id=$1",
		c.Id).Scan(&c.Id, &c.Name, &c.Date, &c.Note, &c.Created, &c.Updated)
}

func (c *Coition) UpdateCoition() error {
	date := time.Now().UTC().Format(time.RFC3339)
	_, err := DB.Exec("UPDATE coition SET name=$1, date=$2, note=$3, updated=$4 WHERE id=$5",
		c.Name, c.Date, c.Note, date, c.Id)

	return err
}

func (c *Coition) DeleteCoition() error {
	_, err := DB.Exec("DELETE FROM coition WHERE id=$1", c.Id)

	return err
}

func (c *Coition) CreateCoition() error {
	date := time.Now().UTC().Format(time.RFC3339)
	err := DB.QueryRow(
		"INSERT INTO coition(name, date, note, created, updated) VALUES($1, $2, $3, $4, $5) RETURNING id",
		c.Name, c.Date, c.Note, date, date).Scan(&c.Id)

	if err != nil {
		return err
	}

	return nil
}

func GetCoitions(since string) (Coitions, error) {

	fmt.Println("List since: " + since)
	rows, err := DB.Query("SELECT id, name, date, note, created, updated FROM coition")
	// DB.Query("SELECT id, name, date, note, created, updated FROM coition WHERE updated > $1",
	// since)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	coitions := Coitions{}

	for rows.Next() {
		var c Coition
		if err := rows.Scan(&c.Id, &c.Name, &c.Date, &c.Note, &c.Created, &c.Updated); err != nil {
			return nil, err
		}
		coitions = append(coitions, c)
	}

	return coitions, nil
}
