package storage

import (
	"database/sql"
	"errors"
	"fmt"
)

func AddClient(login string, password string, nom string, prenom string, email string, date_naissance string, telephone string) error {
	var count int
	err := DB.QueryRow("SELECT * FROM clients WHERE login = $1 AND email = $2", login, email).Scan(&count)

	if err == sql.ErrNoRows {
		_, errInsert := DB.Exec("INSERT INTO clients (login, password, nom, prenom, email, date_naissance, telephone) VALUES ($1, $2, $3, $4, $5, $6, $7)", login, password, nom, prenom, email, date_naissance, telephone)
		return errInsert
	}

	if err != nil {
		return fmt.Errorf("error while checking users: %v", err)
	}
	if count > 0 {
		return errors.New("user already exist, or email already used")
	}
	_, errInsert := DB.Exec("INSERT INTO clients (login, password, nom, prenom, email, date_naissance, telephone) VALUES ($1, $2, $3, $4, $5, $6, $7)", login, password, nom, prenom, email, date_naissance, telephone)
	return errInsert
}
