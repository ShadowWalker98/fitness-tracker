package data

import (
	"FitnessTracker/internal/helpers"
	"FitnessTracker/internal/validator"
	"context"
	"database/sql"
	"fmt"
	"regexp"
)

const getUserByIdQuery = `SELECT first_name, last_name, email FROM users WHERE id=$1`
const addUserQuery = `
INSERT INTO users (first_name, last_name, email, password) VALUES ($1, $2, $3, $4);`

type User struct {
	ID        int64  `json:"-"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last-name"`
	Email     string `json:"email"`
	Password  string `json:"-"`
	Salt      string `json:"-"`
}

type UserModel struct {
	conn *sql.DB
}

func (u UserModel) Insert(user *User) error {
	return addUser(u.conn, user)
}

func (u UserModel) Update(user *User) error {
	return nil
}

func (u UserModel) Delete(id int64) error {
	return nil
}

func (u UserModel) Get(id int64) (*User, error) {
	return nil, nil
}

func ValidateUser(v *validator.Validator, u User) {
	v.Check(u.FirstName != "", "first-name", "must be provided")
	v.Check(u.LastName != "", "last-name", "must be provided")
	rx, err := regexp.Compile(validator.EmailRX)
	if err != nil {
		fmt.Println("Error while parsing email")
	}
	v.Check(validator.Matches(u.Email, rx), "email", "is not in the correct format")
	v.Check(len(u.Password) > 12, "password", "should be greater than 12 characters")
}

func addUser(conn *sql.DB, user *User) error {
	// adds a user to the database
	passwordHash, salt := helpers.GetHash(user.Password)
	ctx := context.TODO()
	_, err := conn.ExecContext(ctx, addUserQuery, &user.FirstName, &user.LastName, &user.Email, passwordHash)
	if err != nil {
		fmt.Println("Error occurred: ", err)
		return err
	}
	fmt.Println("Hashed Password: " + passwordHash)
	fmt.Println("Salt: " + salt)
	return nil
}
