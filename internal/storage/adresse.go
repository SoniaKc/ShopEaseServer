package storage

import (
	"database/sql"
	"errors"
	"fmt"
)

func AddAdresse(login string, nomAdresse string, numero string, nomRue string, codePostal int, ville string, pays string) error {
	var count int
	err := DB.QueryRow("SELECT * FROM adresses WHERE login = $1 AND nom_adresse = $2", login, nomAdresse).Scan(&count)

	if err == sql.ErrNoRows {
		_, errInsert := DB.Exec(
			"INSERT INTO adresses (login, nom_adresse, numero, nom_rue, code_postal, ville, pays) VALUES ($1, $2, $3, $4, $5, $6, $7)",
			login, nomAdresse, numero, nomRue, codePostal, ville, pays)
		return errInsert
	}

	if err != nil {
		return fmt.Errorf("erreur lors de la vérification des adresses: %v", err)
	}

	if count > 0 {
		return errors.New("une adresse avec ce nom existe déjà pour cet utilisateur")
	}

	return nil
}
