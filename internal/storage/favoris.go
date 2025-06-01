package storage

import (
	"errors"
	"fmt"
)

func AddFavori(loginBoutique string, nomProduit string, idClient string) error {
	var count int
	err := DB.QueryRow("SELECT COUNT(*) FROM favoris WHERE login_boutique = $1 AND nom_produit = $2 AND idClient = $3", loginBoutique, nomProduit, idClient).Scan(&count)

	if err != nil {
		return fmt.Errorf("erreur lors de la vérification du panier: %v", err)
	}

	if count > 0 {
		return errors.New("ce favori existe déjà pour ce client")
	}

	_, errInsert := DB.Exec("INSERT INTO favoris (login_boutique, nom_produit, idClient) VALUES ($1, $2, $3)", loginBoutique, nomProduit, idClient)

	return errInsert
}

func GetAllFavoris(idClient string) ([]map[string]interface{}, error) {
	rows, err := DB.Query("SELECT login_boutique, nom_produit, idClient FROM favoris WHERE idClient = $1", idClient)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var favoris []map[string]interface{}
	for rows.Next() {
		var loginBoutique string
		var nomProduit string
		var idClient string
		if err := rows.Scan(&loginBoutique, &nomProduit, &idClient); err != nil {
			return nil, err
		}
		favoris = append(favoris, map[string]interface{}{
			"login_boutique": loginBoutique,
			"nom_produit":    nomProduit,
			"idClient":       idClient,
		})
	}
	return favoris, nil
}

func DeleteFavoris(loginBoutique string, nomProduit string, idClient string) error {
	result, err := DB.Exec("DELETE FROM favoris WHERE login_boutique = $1 AND nom_produit = $2 AND idClient = $3", loginBoutique, nomProduit, idClient)

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

func DeleteFavorisByProduit(loginBoutique string, nomProduit string) error {
	result, err := DB.Exec("DELETE FROM favoris WHERE login_boutique = $1 AND nom_produit = $2", loginBoutique, nomProduit)

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

func DeleteFavorisByClient(idClient string) error {
	result, err := DB.Exec("DELETE FROM favoris WHERE idClient = $1", idClient)

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
