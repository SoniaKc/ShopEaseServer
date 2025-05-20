package storage

import (
	"database/sql"
	"errors"
	"fmt"
)

func AddParametre(login string, typeParam string, langue string, cookies string, notifications string) error {
	var count int
	err := DB.QueryRow("SELECT * FROM parameters WHERE login = $1 AND type = $2", login, typeParam).Scan(&count)

	if err == sql.ErrNoRows {
		_, errInsert := DB.Exec(
			"INSERT INTO parameters (login, type, langue, cookies, notifications) VALUES ($1, $2, $3, $4, $5)",
			login, typeParam, langue, cookies, notifications)
		return errInsert
	}

	if err != nil {
		return fmt.Errorf("error while checking parameters: %v", err)
	}

	if count > 0 {
		return errors.New("parameters already exist for this login and type")
	}

	_, errInsert := DB.Exec(
		"INSERT INTO parameters (login, type, langue, cookies, notifications) VALUES ($1, $2, $3, $4, $5)",
		login, typeParam, langue, cookies, notifications)

	return errInsert
}
