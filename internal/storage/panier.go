package storage

import (
	"database/sql"
	"errors"
	"fmt"
	"shop-ease-server/internal/models"
)

func AddPanier(loginBoutique string, nomProduit string, idClient string, quantite string) error {
	var count int
	err := DB.QueryRow("SELECT * FROM panier WHERE login_boutique = $1 AND nom_produit = $2 AND idClient = $3", loginBoutique, nomProduit, idClient).Scan(&count)

	if err != nil {
		return fmt.Errorf("erreur lors de la vérification du panier: %v", err)
	}

	if count > 0 {
		return errors.New("un panier avec ce produit existe déjà pour ce client")
	}

	_, errInsert := DB.Exec("INSERT INTO panier (login_boutique, nom_produit, idClient, quantite) VALUES ($1, $2, $3, $4)", loginBoutique, nomProduit, idClient, quantite)

	return errInsert
}

func GetQteInPanier(loginBoutique string, nomProduit string, idClient string) (*models.Panier, error) {
	var panier models.Panier

	err := DB.QueryRow("SELECT quantite FROM panier WHERE login_boutique = $1 AND nom_produit = $2 AND idClient = $3", loginBoutique, nomProduit, idClient).Scan(
		&panier.LoginBoutique,
		&panier.NomProduit,
		&panier.IdClient,
		&panier.Quantite,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("cet article n'est pas dans ce panier")
		}
		return nil, fmt.Errorf("erreur lors de la récupération de la quantite dans le panier: %v", err)
	}

	return &panier, nil
}

func GetFullPanier(idClient string) ([]map[string]interface{}, error) {
	rows, err := DB.Query("SELECT login_boutique, nom_produit, quantite FROM panier WHERE idClient = $1", idClient)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var panier []map[string]interface{}
	for rows.Next() {
		var loginBoutique string
		var nomProduit string
		var quantite string
		if err := rows.Scan(&loginBoutique, &nomProduit, &quantite); err != nil {
			return nil, err
		}
		panier = append(panier, map[string]interface{}{
			"login_boutique": loginBoutique,
			"nom_produit":    nomProduit,
			"quantite":       quantite,
		})
	}
	return panier, nil
}

func DeletePanier(loginBoutique string, nomProduit string, idClient string) error {
	result, err := DB.Exec("DELETE FROM panier WHERE login_boutique = $1 AND nom_produit = $2 AND idClient = $3", loginBoutique, nomProduit, idClient)

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

func UpdateQteInPanier(loginBoutique string, nomProduit string, idClient string, quantite string) error {
	if quantite == "0" {
		return DeletePanier(loginBoutique, nomProduit, idClient)
	}

	_, err := DB.Exec(`
        UPDATE panier SET quantite = $1 WHERE login_boutique = $2 AND nom_produit = $3 AND idClient = $4`,
		quantite, loginBoutique, nomProduit, idClient)
	return err
}
