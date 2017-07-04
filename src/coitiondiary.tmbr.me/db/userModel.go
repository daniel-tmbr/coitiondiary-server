package db

import (
	"fmt"
	"time"
)

// User model
type User struct {
	Id          int       `json:"id"`
	Username    int       `json:"username"`
    Email       string    `json:"email"`
	Name        string    `json:"name"`
	Avatar      string    `json:"avatar"`
    Gender      string    `json:"gender"`
    Birthday    time.Time `json:"birthday"`
    Country     string    `json:"country"`
    City        string    `json:"city"`
    Profession  string    `json:"profession"`
	Created     time.Time `json:"created"`
	Updated     time.Time `json:"updated"`
}

// Users slice, containing User models
// type Users []User

// func (u *User) GetUser() error {
// 	return DB.QueryRow("SELECT * FROM user WHERE id=$1",
// 		u.Id).Scan(&u.Id, &u.Username, &u.Email, &u.Name, &u.Avatar, &u.Gender)
// }

// func (u *User) UpdateUser() error {
// 	date := time.Now().UTC().Format(time.RFC3339)
// 	_, err := DB.Exec("UPDATE user SET name=$1, date=$2, note=$3, updated=$4 WHERE id=$5",
// 		u.Name, u.Date, u.Note, date, u.Id)

// 	return err
// }

// func (u *User) DeleteUser() error {
// 	_, err := DB.Exec("DELETE FROM user WHERE id=$1", u.Id)
// 	return err
// }

// func (c *User) CreateUser() error {
// 	date := time.Now().UTC().Format(time.RFC3339)
// 	err := DB.QueryRow(
// 		"INSERT INTO user(name, date, note, created, updated) VALUES($1, $2, $3, $4, $5) RETURNING id",
// 		u.Name, u.Date, u.Note, date, date).Scan(&u.Id)

// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func GetUsers(since string) (Users, error) {

// 	fmt.Println("List since: " + since)
// 	rows, err := DB.Query("SELECT id, name, date, note, created, updated FROM user")
// 	// DB.Query("SELECT id, name, date, note, created, updated FROM User WHERE updated > $1",
// 	// since)

// 	if err != nil {
// 		return nil, err
// 	}

// 	defer rows.Close()

// 	users := Users{}

// 	for rows.Next() {
// 		var u User
// 		if err := rows.Scan(&u.Id, &u.Name, &u.Date, &u.Note, &u.Created, &u.Updated); err != nil {
// 			return nil, err
// 		}
// 		users = append(users, u)
// 	}

// 	return users, nil
}
