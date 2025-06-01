package storage

import (
	"database/sql"
	"errors"
	"fmt"
	"shop-ease-server/internal/models"
	"strconv"
	"strings"
)

func AddBoutique(login string, password string, nom string, email string, telephone string,
	siret string, forme_juridique string, siege_social string, pays_enregistrement string, iban string, image []byte) error {

	var count int
	err := DB.QueryRow("SELECT COUNT(*) FROM boutiques WHERE login = $1 OR email = $2", login, email).Scan(&count)

	if err != nil {
		return fmt.Errorf("error while checking boutiques: %v", err)
	}

	if count > 0 {
		return errors.New("boutique already exists with this login or email")
	}

	_, errInsert := DB.Exec("INSERT INTO boutiques (login, password, nom, email, telephone, siret, forme_juridique, siege_social, pays_enregistrement, iban, image) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)",
		login, password, nom, email, telephone, siret, forme_juridique, siege_social, pays_enregistrement, iban, image)

	return errInsert
}

func GetBoutique(login string) (*models.Boutique, error) {
	var boutique models.Boutique

	err := DB.QueryRow("SELECT * FROM boutiques WHERE login = $1", login).Scan(
		&boutique.Login,
		&boutique.Password,
		&boutique.Nom,
		&boutique.Email,
		&boutique.Telephone,
		&boutique.Siret,
		&boutique.Forme_juridique,
		&boutique.Siege_social,
		&boutique.Pays_enregistrement,
		&boutique.Iban,
		&boutique.Image,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("boutique non trouvée")
		}
		return nil, fmt.Errorf("erreur lors de la récupération de la boutique: %v", err)
	}

	return &boutique, nil
}

func DeleteBoutique(login string) error {
	result, err := DB.Exec("DELETE FROM boutiques WHERE login = $1", login)

	if err != nil {
		return fmt.Errorf("failed to delete boutique: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to check rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("boutique not found")
	}

	DeleteProduitsByBoutique(login)
	DeleteParametre(login, "boutique")

	return nil
}

func UpdateBoutique(login string, updates map[string]interface{}) error {
	if len(updates) == 0 {
		return nil
	}

	query := "UPDATE boutiques SET "
	params := []interface{}{}
	i := 1

	allowedFields := map[string]bool{
		"password":            true,
		"nom":                 true,
		"email":               true,
		"telephone":           true,
		"siret":               true,
		"forme_juridique":     true,
		"siege_social":        true,
		"pays_enregistrement": true,
		"iban":                true,
		"image":               true,
	}

	for field, value := range updates {
		if !allowedFields[field] {
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
	query += " WHERE login = $" + strconv.Itoa(i)
	params = append(params, login)
	/*
		fmt.Printf("Generated SQL: %s\n", query)
		for i, param := range params {
			fmt.Printf("$%d = %v (type: %T)\n", i+1, param, param)
		}
	*/
	result, err := DB.Exec(query, params...)
	if err != nil {
		return fmt.Errorf("failed to update boutique: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to check rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("boutique not found")
	}

	return nil
}
