package storage

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func AddCommentaire(loginBoutique string, nomProduit string, idClient string, note string, commentaire string) error {
	var count int
	err := DB.QueryRow("SELECT COUNT(*) FROM commentaires WHERE login_boutique = $1 AND nom_produit = $2 AND idClient = $3", loginBoutique, nomProduit, idClient).Scan(&count)

	if err != nil {
		return fmt.Errorf("erreur lors de la vérification des commentaires: %v", err)
	}

	if count > 0 {
		return errors.New("un commentaire sur cet article existe déjà pour ce client")
	}

	_, errInsert := DB.Exec(
		"INSERT INTO commentaires (login_boutique, nom_produit, idClient, note, commentaire) VALUES ($1, $2, $3, $4, $5)",
		loginBoutique, nomProduit, idClient, note, commentaire)

	return errInsert
}

func GetAllComsProduit(loginBoutique string, nomProduit string) ([]map[string]interface{}, error) {
	rows, err := DB.Query("SELECT idClient, note, commentaire FROM commentaires WHERE login_boutique = $1 AND nom_produit = $2", loginBoutique, nomProduit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var commentaires []map[string]interface{}
	for rows.Next() {
		var idClient string
		var note string
		var commentaire string
		if err := rows.Scan(&idClient, &note, &commentaire); err != nil {
			return nil, err
		}
		commentaires = append(commentaires, map[string]interface{}{
			"login_boutique": loginBoutique,
			"nom_produit":    nomProduit,
			"idClient":       idClient,
			"note":           note,
			"commentaire":    commentaire,
		})
	}
	return commentaires, nil
}

func GetAllComsClient(idClient string) ([]map[string]interface{}, error) {
	rows, err := DB.Query("SELECT login_boutique, nom_produit, note, commentaire FROM commentaires WHERE idClient = $1", idClient)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var commentaires []map[string]interface{}
	for rows.Next() {
		var loginBoutique string
		var nomProduit string
		var note string
		var commentaire string
		if err := rows.Scan(&loginBoutique, &nomProduit, &note, &commentaire); err != nil {
			return nil, err
		}
		commentaires = append(commentaires, map[string]interface{}{
			"login_boutique": loginBoutique,
			"nom_produit":    nomProduit,
			"idClient":       idClient,
			"note":           note,
			"commentaire":    commentaire,
		})
	}
	return commentaires, nil
}

func DeleteCommentaire(loginBoutique string, nomProduit string, idClient string) error {
	result, err := DB.Exec("DELETE FROM commentaires WHERE login_boutique = $1 AND nom_produit = $2 AND idClient = $3", loginBoutique, nomProduit, idClient)

	if err != nil {
		return fmt.Errorf("failed to delete commentaires: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to check rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("commentaire not found")
	}

	return nil
}

func DeleteCommentaireByProduit(loginBoutique string, nomProduit string) error {
	result, err := DB.Exec("DELETE FROM commentaires WHERE login_boutique = $1 AND nom_produit = $2", loginBoutique, nomProduit)

	if err != nil {
		return fmt.Errorf("failed to delete commentaires: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to check rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("commentaire not found")
	}

	return nil
}

func UpdateCommentaire(loginBoutique string, nomProduit string, idClient string, updates map[string]interface{}) error {
	if len(updates) == 0 {
		return nil
	}

	query := "UPDATE commentaires SET "
	params := []interface{}{}
	i := 1

	allowedFields := map[string]bool{
		"note":        true,
		"commentaire": true,
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
	query += " WHERE login_boutique = $" + strconv.Itoa(i) + " AND nom_produit = $" + strconv.Itoa(i+1) + " AND idClient = $" + strconv.Itoa(i+2)
	params = append(params, loginBoutique, nomProduit, idClient)

	fmt.Printf("Generated SQL: %s\n", query)
	for i, param := range params {
		fmt.Printf("$%d = %v (type: %T)\n", i+1, param, param)
	}

	result, err := DB.Exec(query, params...)
	if err != nil {
		return fmt.Errorf("failed to update commentaire: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to check rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("commentaire not found")
	}

	return nil
}
