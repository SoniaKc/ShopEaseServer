package storage

import (
	"database/sql"
	"errors"
	"fmt"
	"shop-ease-server/internal/models"
	"strconv"
	"strings"
)

func AddProduit(login_boutique string, nom string, categories string, reduction string, prix string, description string) error {
	var count int
	err := DB.QueryRow("SELECT COUNT(*) FROM produits WHERE login_boutique = $1 AND nom = $2", login_boutique, nom).Scan(&count)

	if err != nil {
		return fmt.Errorf("erreur lors de la vérification des produits: %v", err)
	}

	if count > 0 {
		return errors.New("un produit avec ce nom existe déjà pour cette boutique")
	}

	_, errInsert := DB.Exec(
		"INSERT INTO produits (login_boutique, nom, categories, reduction, prix, description) VALUES ($1, $2, $3, $4, $5, $6)",
		login_boutique, nom, categories, reduction, prix, description)

	return errInsert
}

func GetProduit(login_boutique string, nom string) (*models.Produit, error) {
	var produit models.Produit

	err := DB.QueryRow("SELECT * FROM produits WHERE login_boutique = $1 AND nom = $2", login_boutique, nom).Scan(
		&produit.LoginBoutique,
		&produit.Nom,
		&produit.Categories,
		&produit.Reduction,
		&produit.Prix,
		&produit.Description,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("produit non trouvé")
		}
		return nil, fmt.Errorf("erreur lors de la récupération du produit : %v", err)
	}

	return &produit, nil
}

func GetAllProduit(loginBoutique string) ([]map[string]interface{}, error) {
	rows, err := DB.Query("SELECT nom, categories, reduction, prix, description FROM produits WHERE login_boutique = $1", loginBoutique)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var produits []map[string]interface{}
	for rows.Next() {
		var nom string
		var categories string
		var reduction string
		var prix string
		var description string
		if err := rows.Scan(&nom, &categories, &reduction, &prix, &description); err != nil {
			return nil, err
		}
		produits = append(produits, map[string]interface{}{
			"login_boutique": loginBoutique,
			"nom":            nom,
			"categories":     categories,
			"reduction":      reduction,
			"prix":           prix,
			"description":    description,
		})
	}
	return produits, nil
}

func DeleteProduit(login_boutique string, nom string) error {
	result, err := DB.Exec("DELETE FROM produits WHERE login_boutique = $1 AND nom = $2", login_boutique, nom)

	if err != nil {
		return fmt.Errorf("failed to delete produit: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to check rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("produit not found")
	}

	return nil
}

func UpdateProduit(login_boutique string, nom string, updates map[string]interface{}) error {
	if len(updates) == 0 {
		return nil
	}

	query := "UPDATE produits SET "
	params := []interface{}{}
	i := 1

	allowedFields := map[string]bool{
		//"nom":         true,
		"categories":  true,
		"reduction":   true,
		"prix":        true,
		"description": true,
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
	query += " WHERE login_boutique = $" + strconv.Itoa(i) + " AND nom = $" + strconv.Itoa(i+1)
	params = append(params, login_boutique, nom)

	fmt.Printf("Generated SQL: %s\n", query)
	for i, param := range params {
		fmt.Printf("$%d = %v (type: %T)\n", i+1, param, param)
	}

	result, err := DB.Exec(query, params...)
	if err != nil {
		return fmt.Errorf("failed to update produit: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to check rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("produit not found")
	}

	return nil
}
