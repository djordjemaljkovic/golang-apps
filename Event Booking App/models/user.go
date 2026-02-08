package models

import (
	"errors"

	"example.com/events/db"
	"example.com/events/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {
	query := `INSERT INTO users(email,password) VALUES (?,?)`
	st, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer st.Close()

	hashed, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

	result, err := st.Exec(u.Email, hashed)

	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()
	u.ID = userId
	return err
}

func (u *User) Validate() error {
	query := `SELECT id,password FROM users WHERE email = ?`
	row := db.DB.QueryRow(query, u.Email)

	var retrievedPass string
	err := row.Scan(&u.ID, &retrievedPass)

	if err != nil {
		return err
	}

	passValid := utils.CheckHash(u.Password, retrievedPass)

	if !passValid {
		return errors.New("Credentials invalid")
	}

	return nil
}
