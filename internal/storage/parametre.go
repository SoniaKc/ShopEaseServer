package storage

import (
	"database/sql"
	"errors"
	"fmt"
	"shop-ease-server/internal/models"
	"strconv"
	"strings"
)

func AddParametre(login string, typeLogin string, langue string, cookies string, notifications string) error {
	var count int
	err := DB.QueryRow("SELECT * FROM parametres WHERE login = $1 AND type = $2", login, typeLogin).Scan(&count)

	if err == sql.ErrNoRows {
		_, errInsert := DB.Exec(
			"INSERT INTO parameters (login, type, langue, cookies, notifications) VALUES ($1, $2, $3, $4, $5)",
			login, typeLogin, langue, cookies, notifications)
		return errInsert
	}

	if err != nil {
		return fmt.Errorf("error while checking parameters: %v", err)
	}

	if count > 0 {
		return errors.New("parameters already exist for this login and type")
	}

	_, errInsert := DB.Exec(
		"INSERT INTO parametres (login, type, langue, cookies, notifications) VALUES ($1, $2, $3, $4, $5)",
		login, typeLogin, langue, cookies, notifications)

	return errInsert
}

func GetParametre(login string, typeLogin string) (*models.Parametre, error) {
	var parametre models.Parametre

	err := DB.QueryRow("SELECT * FROM parametres WHERE login = $1 AND type = $2", login, typeLogin).Scan(
		&parametre.Login,
		&parametre.Type,
		&parametre.Langue,
		&parametre.Cookies,
		&parametre.Notifications,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("parametres non trouvés")
		}
		return nil, fmt.Errorf("erreur lors de la récupération des parametres: %v", err)
	}

	return &parametre, nil
}

func DeleteParametre(login string, typeLogin string) error {
	result, err := DB.Exec("DELETE FROM parametres WHERE login = $1 AND type = $2", login, typeLogin)

	if err != nil {
		return fmt.Errorf("failed to delete parametres: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to check rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("parametres not found")
	}

	return nil
}

func UpdateParametre(login string, typeLogin string, updates map[string]interface{}) error {
	if len(updates) == 0 {
		return nil
	}

	query := "UPDATE parametres SET "
	params := []interface{}{}
	i := 1

	allowedFields := map[string]bool{
		"langue":        true,
		"cookies":       true,
		"notifications": true,
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
	query += " WHERE login = $" + strconv.Itoa(i) + " AND type = $" + strconv.Itoa(i+1)
	params = append(params, login, typeLogin)

	fmt.Printf("Generated SQL: %s\n", query)
	for i, param := range params {
		fmt.Printf("$%d = %v (type: %T)\n", i+1, param, param)
	}

	result, err := DB.Exec(query, params...)
	if err != nil {
		return fmt.Errorf("failed to update parametres: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to check rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("parametrs not found")
	}

	return nil
}
