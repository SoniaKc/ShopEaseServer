package storage

import (
	"database/sql"
	"errors"
	"fmt"
)

func AddPanier(idProduit string, idClient string, quantite int) error {
	var count int
	err := DB.QueryRow("SELECT * FROM panier WHERE idProduit = $1 AND idClient = $2", idProduit, idClient).Scan(&count)

	if err != nil {
		return fmt.Errorf("erreur lors de la vérification du panier: %v", err)
	}

	if count > 0 {
		return errors.New("un panier avec ce produit existe déjà pour ce client")
	}

	_, errInsert := DB.Exec("INSERT INTO panier (idProduit, idClient, quantite) VALUES ($1, $2, $3)", idProduit, idClient, quantite)

	return errInsert
}

func GetQteInPanier(idProduit string, idClient string) (int, error) {
	var quantite int

	err := DB.QueryRow("SELECT quantite FROM panier WHERE idProduit = $1 AND idClient = $2", idProduit, idClient).Scan(&quantite)

	if err != nil {
		if err == sql.ErrNoRows {
			return 0, fmt.Errorf("cet article n'est pas dans ce panier")
		}
		return -1, fmt.Errorf("erreur lors de la récupération de la quantite dans le panier: %v", err)
	}

	return quantite, nil
}

func GetFullPanier(idClient string) ([]map[string]interface{}, error) {
	rows, err := DB.Query("SELECT idProduit, quantite FROM panier WHERE idClient = $1", idClient)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var panier []map[string]interface{}
	for rows.Next() {
		var idProduit string
		var quantite int
		if err := rows.Scan(&idProduit, &quantite); err != nil {
			return nil, err
		}
		panier = append(panier, map[string]interface{}{
			"idProduit": idProduit,
			"quantite":  quantite,
		})
	}
	return panier, nil
}

func DeletePanier(idProduit string, idClient string) error {
	result, err := DB.Exec("DELETE FROM panier WHERE idProduit = $1 AND idClient = $2", idProduit, idClient)

	if err != nil {
		return fmt.Errorf("failed to delete panier row: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to check rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("panier row not found")
	}

	return nil
}

func UpdateQteInPanier(idProduit string, idClient string, quantite int) error {
	if quantite == 0 {
		return DeletePanier(idProduit, idClient)
	}

	_, err := DB.Exec(`
        UPDATE panier SET quantite = $1 WHERE idProduit = $2 AND idClient = $3`,
		quantite, idProduit, idClient)
	return err
}
