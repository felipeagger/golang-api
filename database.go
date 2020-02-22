package main

import (
	"fmt"
	"time"

	"database/sql"

	_ "github.com/lib/pq"
)

func ConectDB() (*sql.DB, error) {
	dbHost := "localhost"  //os.Getenv("POSTGRES_HOST")
	dbName := "go_api"     //os.Getenv("POSTGRES_DB")
	dbPasswd := "postgres" //os.Getenv("POSTGRES_PASSWD")

	urlConection := fmt.Sprintf("host=%s port=5432 user=postgres dbname=%s password=%s sslmode=disable",
		dbHost, dbName, dbPasswd)

	db, err := sql.Open("postgres", urlConection)

	if err != nil {
		panic("failed to connect database")
		return nil, err
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(time.Hour)

	return db, nil
}

func Create(user User) User {

	db, _ := ConectDB()

	var userID string
	err := db.QueryRow("INSERT INTO users (username, age, bio, link, avatar, score) VALUES($1, $2, $3, $4, $5, $6) returning id;",
		user.Username, user.Age, user.Bio, user.Link, user.Avatar, user.Score).Scan(&userID)
	checkErr(err)

	user.ID = userID
	return user
}

func SelectAllUsers() (usr []User) {

	db, _ := ConectDB()
	var users []User

	rows, err := db.Query("SELECT * FROM users")
	checkErr(err)

	for rows.Next() {
		var user User

		err = rows.Scan(&user.ID, &user.Username, &user.Age, &user.Bio, &user.Link,
			&user.Avatar, &user.CreatedAt, &user.Score)
		checkErr(err)

		users = append(users, user)
	}

	RedisSet("all_users", users)
	return users
}

func SelectUser(username string) (usr User) {

	db, _ := ConectDB()
	var user User

	rows, err := db.Query("SELECT * FROM users where username = $1",
		username)
	checkErr(err)

	for rows.Next() {

		err = rows.Scan(&user.ID, &user.Username, &user.Age, &user.Bio, &user.Link,
			&user.Avatar, &user.CreatedAt, &user.Score)
		checkErr(err)

	}

	return user
}

func Update(user User) int64 {

	db, _ := ConectDB()

	stmt, err := db.Prepare("update users set age=$1, bio=$2, link=$3, avatar=$4, score=$5 where username=$6")
	checkErr(err)

	res, err := stmt.Exec(user.Age, user.Bio, user.Link, user.Avatar, user.Score, user.Username)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	if affect > 0 {
		RedisSet("all_users", nil)
	}

	return affect
}

func Delete(username string) int64 {

	db, _ := ConectDB()

	stmt, err := db.Prepare("delete from users where username=$1")
	checkErr(err)

	res, err := stmt.Exec(username)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	if affect > 0 {
		RedisSet("all_users", nil)
	}

	return affect
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
