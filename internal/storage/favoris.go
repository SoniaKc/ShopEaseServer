package storage

import (
	"errors"
	"fmt"
)

func AddFavori(idProduit string, idClient string) error {
	var count int
	err := DB.QueryRow("SELECT COUNT(*) FROM panier WHERE idProduit = $1 AND idClient = $2", idProduit, idClient).Scan(&count)

	if err != nil {
		return fmt.Errorf("erreur lors de la vérification du panier: %v", err)
	}

	if count > 0 {
		return errors.New("ce favori existe déjà pour ce client")
	}

	_, errInsert := DB.Exec("INSERT INTO favoris (idProduit, idClient) VALUES ($1, $2)", idProduit, idClient)

	return errInsert
}

func GetAllFavoris(idClient string) ([]map[string]interface{}, error) {
	rows, err := DB.Query("SELECT idProduit, idClient FROM favoris WHERE idClient = $1", idClient)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var favoris []map[string]interface{}
	for rows.Next() {
		var idProduit string
		var idClient string
		if err := rows.Scan(&idProduit, &idClient); err != nil {
			return nil, err
		}
		favoris = append(favoris, map[string]interface{}{
			"idProduit": idProduit,
			"idClient":  idClient,
		})
	}
	return favoris, nil
}

func DeleteFavoris(idProduit string, idClient string) error {
	result, err := DB.Exec("DELETE FROM favoris WHERE idProduit = $1 AND idClient = $2", idProduit, idClient)

	if err != nil {
		return fmt.Errorf("failed to delete favori row: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to check rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("favori row not found")
	}

	return nil
}
