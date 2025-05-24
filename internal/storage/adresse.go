package storage

import (
	"database/sql"
	"errors"
	"fmt"
	"shop-ease-server/internal/models"
	"strconv"
	"strings"
)

func AddAdresse(login string, nomAdresse string, numero string, nomRue string, codePostal string, ville string, pays string) error {
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

func GetAdresse(login string, nomAdresse string) (*models.Adresse, error) {
	var adresse models.Adresse

	err := DB.QueryRow("SELECT * FROM adresses WHERE login = $1 AND nom_adresse = $2", login, nomAdresse).Scan(
		&adresse.Login,
		&adresse.NomAdresse,
		&adresse.Numero,
		&adresse.NomRue,
		&adresse.CodePostal,
		&adresse.Ville,
		&adresse.Pays,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("adresse non trouvée")
		}
		return nil, fmt.Errorf("erreur lors de la récupération de l'adresse : %v", err)
	}

	return &adresse, nil
}

func DeleteAdresse(login string, nomAdresse string) error {
	result, err := DB.Exec("DELETE FROM adresses WHERE login = $1 AND nom_adresse = $2", login, nomAdresse)

	if err != nil {
		return fmt.Errorf("failed to delete adresse: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to check rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("adresse not found")
	}

	return nil
}

func UpdateAdresse(login string, nomAdresse string, updates map[string]interface{}) error {
	if len(updates) == 0 {
		return nil
	}

	query := "UPDATE adresses SET "
	params := []interface{}{}
	i := 1

	allowedFields := map[string]bool{
		"numero":      true,
		"nom_rue":     true,
		"code_postal": true,
		"ville":       true,
		"pays":        true,
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
	query += " WHERE login = $" + strconv.Itoa(i) + " AND nom_adresse = $" + strconv.Itoa(i+1)
	params = append(params, login, nomAdresse)

	fmt.Printf("Generated SQL: %s\n", query)
	for i, param := range params {
		fmt.Printf("$%d = %v (type: %T)\n", i+1, param, param)
	}

	result, err := DB.Exec(query, params...)
	if err != nil {
		return fmt.Errorf("failed to update adresse: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to check rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("adresse not found")
	}

	return nil
}
