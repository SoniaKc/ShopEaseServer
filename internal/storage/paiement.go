package storage

import (
	"database/sql"
	"errors"
	"fmt"
	"shop-ease-server/internal/models"
	"strconv"
	"strings"
)

func AddPaiement(login string, nom_carte string, nom_personne_carte string, numero string, cvc string, date_expiration string) error {
	var count int
	err := DB.QueryRow("SELECT * FROM paiements WHERE login = $1 AND nom_carte = $2", login, nom_carte).Scan(&count)

	if err == sql.ErrNoRows {
		_, errInsert := DB.Exec(
			"INSERT INTO paiements (login, nom_carte, nom_personne_carte, numero, cvc, date_expiration) VALUES ($1, $2, $3, $4, $5, $6)",
			login, nom_carte, nom_personne_carte, numero, cvc, date_expiration)
		return errInsert
	}

	if err != nil {
		return fmt.Errorf("error while checking paiements: %v", err)
	}

	if count > 0 {
		return errors.New("paiement method already exists for this login and card name")
	}

	_, errInsert := DB.Exec(
		"INSERT INTO paiements (login, nom_carte, nom_personne_carte, numero, cvc, date_expiration) VALUES ($1, $2, $3, $4, $5, $6)",
		login, nom_carte, nom_personne_carte, numero, cvc, date_expiration)
	return errInsert
}

func GetPaiement(login string, nomCarte string) (*models.Paiement, error) {
	var paiement models.Paiement

	err := DB.QueryRow("SELECT * FROM paiements WHERE login = $1 AND nom_carte = $2", login, nomCarte).Scan(
		&paiement.Login,
		&paiement.NomCarte,
		&paiement.NomPersonneCarte,
		&paiement.Numero,
		&paiement.CVC,
		&paiement.DateExpiration,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("carte paiement non trouvée")
		}
		return nil, fmt.Errorf("erreur lors de la récupération de carte paiement: %v", err)
	}

	return &paiement, nil
}

func GetAllPaiement(login string) ([]map[string]interface{}, error) {
	rows, err := DB.Query("SELECT * FROM paiement WHERE login = $1", login)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var paiements []map[string]interface{}
	for rows.Next() {
		var loginRecup string
		var nomCarte string
		var nomPersonneCarte string
		var numero string
		var cvc string
		var dateExpiration string

		if err := rows.Scan(&loginRecup, &nomCarte, &nomPersonneCarte, &numero, &cvc, &dateExpiration); err != nil {
			return nil, err
		}
		paiements = append(paiements, map[string]interface{}{
			"login":              loginRecup,
			"nom_carte":          nomCarte,
			"nom_personne_carte": nomPersonneCarte,
			"nunero":             numero,
			"cvc":                cvc,
			"date_expiration":    dateExpiration,
		})
	}
	return paiements, nil
}

func DeletePaiement(login string, nomCarte string) error {
	result, err := DB.Exec("DELETE FROM paiements WHERE login = $1 AND nom_carte = $2", login, nomCarte)

	if err != nil {
		return fmt.Errorf("failed to delete paiement: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to check rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("paiement not found")
	}

	return nil
}

func UpdatePaiement(login string, nomCarte string, updates map[string]interface{}) error {
	if len(updates) == 0 {
		return nil
	}

	query := "UPDATE paiements SET "
	params := []interface{}{}
	i := 1

	allowedFields := map[string]bool{
		"nom_personne_carte": true,
		"numero":             true,
		"cvc":                true,
		"date_expiration":    true,
	}

	for field, value := range updates {
		if !allowedFields[field] {
			continue
		}

		if strVal, ok := value.(string); ok && strVal == "" {
			continue
		}

		if value == nil {
			continue
		}
		query += fmt.Sprintf("%s = $%d, ", field, i)
		params = append(params, value)
		i++
	}

	if len(params) == 0 {
		return errors.New("aucun champ valide à mettre à jour")
	}

	query = strings.TrimSuffix(query, ", ")
	query += " WHERE login = $" + strconv.Itoa(i) + " AND nom_carte = $" + strconv.Itoa(i+1)
	params = append(params, login, nomCarte)

	fmt.Printf("Generated SQL: %s\n", query)
	for i, param := range params {
		fmt.Printf("$%d = %v (type: %T)\n", i+1, param, param)
	}

	result, err := DB.Exec(query, params...)
	if err != nil {
		return fmt.Errorf("failed to update paiement: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to check rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("paiement not found")
	}

	return nil
}
