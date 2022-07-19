package main

import (
	"database/sql"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
	"github.com/gofrs/uuid"
	_ "github.com/lib/pq"
)

func connectDB() *sql.DB {
	db, err := sql.Open(
		"postgres",
		"host=localhost port=5432 user=najmiddin password=1234 dbname=test sslmode=disable",
	)
	if err != nil {
		panic(err)
	}
	return db
}

type User struct {
	Id        string
	FirstName string
	LastName  string
	Email     string
	Password  string
}

func CreateUser(user User) (bool, error) {
	db := connectDB()
	defer db.Close()

	id, err := uuid.NewV4()
	if err != nil {
		log.Panic(err)
		return false, err
	}
	user.Id = id.String()

	pass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(pass)
	query := `INSERT INTO users(
		id, first_name, last_name, email, password
		) VALUES($1,$2,$3,$4,$5)`
	_, err = db.Exec(query, user.Id, user.FirstName, user.LastName,
		user.Email, user.Password)
	if err != nil {
		log.Panic(err)
		return false, err
	}
	return true, nil
}

func UpdateUser(user User) (bool, error) {
	db := connectDB()
	defer db.Close()

	query := `UPDATE users SET first_name = $1, last_name = $2, email = $3, password = $4
	WHERE id = $5`
	_, err := db.Exec(query, user.FirstName, user.LastName,
		user.Email, user.Password, user.Id)
	if err != nil {
		log.Panic(err)
		return false, err
	}

	return true, nil
}

func UserList() ([]User, error) {
	db := connectDB()
	defer db.Close()

	query := `SELECT * FROM users`
	rows, err := db.Query(query)
	if err != nil {
		log.Panic(err)
		return nil, err
	}
	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(
			&user.Id,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.Password,
		)
		if err != nil {
			log.Panic(err)
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func GetUser(id string) (User, error) {
	db := connectDB()
	defer db.Close()
	var user User
	query := `SELECT * FROM users WHERE id = $1`
	rows, err := db.Query(query, id)
	if err != nil {
		log.Panic(err)
		return User{}, err
	}
	for rows.Next() {
		err = rows.Scan(
			&user.Id,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.Password,
		)
		if err != nil {
			log.Panic(err)
			return User{}, err
		}
	}
		
		return user, nil
}

func DeleteUser(id string) (bool, error) {
	db := connectDB()
	defer db.Close()

	query := `DELETE FROM users WHERE id = $1`
	_, err := db.Exec(query, id)

	if err != nil {
		log.Panic(err)
		return false, err
	}

	return true, nil
}

func main() {
	// res, err := CreateUser(User{
	// 	FirstName: "Carroll",
	// 	LastName:  "John",
	// 	Email:     "john@gmail.com",
	// 	Password:  "password",
	// })

	// res, err := UpdateUser(User{
	// 	Id: 2358a327-5809-463a-8d2d-e3dfdafe893b,
	// 	FirstName: "Carrick",
	// 	LastName:  "Keln",
	// 	Email:     "Keln@gmail.com",
	// 	Password:  "pass",
	// })

	// res, err := UserList()

	res, err := GetUser("cfb4b7d8-a644-4bd3-985d-2ed7dac484a6")

	// res, err := DeleteUser("2358a327-5809-463a-8d2d-e3dfdafe893b")
	
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}
