package storage

import (
	"database/sql"
	"errors"
	"fmt"
)

func AddPaiement(login string, nom_carte string, nom_personne_carte string, cvc int, date_expiration string) error {
	var count int
	err := DB.QueryRow("SELECT * FROM paiements WHERE login = $1 AND nom_carte = $2", login, nom_carte).Scan(&count)

	if err == sql.ErrNoRows {
		_, errInsert := DB.Exec(
			"INSERT INTO paiements (login, nom_carte, nom_personne_carte, cvc, date_expiration) VALUES ($1, $2, $3, $4, $5)",
			login, nom_carte, nom_personne_carte, cvc, date_expiration)
		return errInsert
	}

	if err != nil {
		return fmt.Errorf("error while checking paiements: %v", err)
	}

	if count > 0 {
		return errors.New("payment method already exists for this login and card name")
	}

	_, errInsert := DB.Exec(
		"INSERT INTO paiements (login, nom_carte, nom_personne_carte, cvc, date_expiration) VALUES ($1, $2, $3, $4, $5)",
		login, nom_carte, nom_personne_carte, cvc, date_expiration)
	return errInsert
}
