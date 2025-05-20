package storage

import (
	"database/sql"
	"errors"
	"fmt"
)

func AddBoutique(login string, password string, nom string, email string, telephone string,
	siret string, forme_juridique string, siege_social string, pays_enregistrement string, iban string) error {

	var count int
	err := DB.QueryRow("SELECT * FROM boutiques WHERE login = $1 OR email = $2", login, email).Scan(&count)

	if err == sql.ErrNoRows {
		_, errInsert := DB.Exec("INSERT INTO boutiques (login, password, nom, email, telephone, siret, forme_juridique, siege_social, pays_enregistrement, iban) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)",
			login, password, nom, email, telephone, siret, forme_juridique, siege_social, pays_enregistrement, iban)
		return errInsert
	}

	if err != nil {
		return fmt.Errorf("error while checking boutiques: %v", err)
	}

	if count > 0 {
		return errors.New("boutique already exists with this login or email")
	}

	_, errInsert := DB.Exec("INSERT INTO boutiques (login, password, nom, email, telephone, siret, forme_juridique, siege_social, pays_enregistrement, iban) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)",
		login, password, nom, email, telephone, siret, forme_juridique, siege_social, pays_enregistrement, iban)

	return errInsert
}
