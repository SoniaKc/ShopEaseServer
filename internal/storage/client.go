package storage

import (
	"database/sql"
	"errors"
	"fmt"
	"shop-ease-server/internal/models"
	"strconv"
	"strings"
)

func AddClient(login string, password string, nom string, prenom string, email string, date_naissance string, telephone string) error {
	var count int
	err := DB.QueryRow("SELECT * FROM clients WHERE login = $1 OR email = $2", login, email).Scan(&count)

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

func GetClient(login string) (*models.Client, error) {
	var client models.Client

	err := DB.QueryRow("SELECT * FROM clients WHERE login = $1", login).Scan(
		&client.Login,
		&client.Password,
		&client.Nom,
		&client.Prenom,
		&client.Email,
		&client.DateNaissance,
		&client.Telephone,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("client non trouvé")
		}
		return nil, fmt.Errorf("erreur lors de la récupération du client: %v", err)
	}

	return &client, nil
}

func DeleteClient(login string) error {
	result, err := DB.Exec("DELETE FROM clients WHERE login = $1", login)

	if err != nil {
		return fmt.Errorf("failed to delete client: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to check rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("client not found")
	}

	return nil
}

func UpdateClient(login string, updates map[string]interface{}) error {
	if len(updates) == 0 {
		return nil
	}

	query := "UPDATE clients SET "
	params := []interface{}{}
	i := 1

	allowedFields := map[string]bool{
		"password":       true,
		"nom":            true,
		"prenom":         true,
		"email":          true,
		"date_naissance": true,
		"telephone":      true,
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
	query += " WHERE login = $" + strconv.Itoa(i)
	params = append(params, login)

	fmt.Printf("Generated SQL: %s\n", query)
	for i, param := range params {
		fmt.Printf("$%d = %v (type: %T)\n", i+1, param, param)
	}

	result, err := DB.Exec(query, params...)
	if err != nil {
		return fmt.Errorf("failed to update client: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to check rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("client not found")
	}

	return nil
}
