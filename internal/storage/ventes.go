package storage

import (
	"errors"
	"fmt"
)

func AddVente(idTransaction string, LoginBoutique string, NomProduit string, idClient string, quantite string, total string, date_vente string, statut string) error {
	var count int

	err := DB.QueryRow("SELECT COUNT(*) FROM ventes WHERE idTransaction = $1 AND login_boutique = $2 AND nom_produit = $3", idTransaction, LoginBoutique, NomProduit).Scan(&count)

	if err != nil {
		return fmt.Errorf("erreur lors de la vérification de la vente: %v", err)
	}

	if count > 0 {
		return errors.New("cette vente sur ce produit existe déjà")
	}

	_, errInsert := DB.Exec("INSERT INTO ventes (idTransaction, login_boutique, nom_produit, idClient, quantite, total, date_vente, statut) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)", idTransaction, LoginBoutique, NomProduit, idClient, quantite, total, date_vente, statut)

	return errInsert
}

func GetAllTransaction(idTransaction string) ([]map[string]interface{}, error) {
	rows, err := DB.Query("SELECT login_boutique, nom_produit, idClient, quantite, total, date_vente, statut FROM ventes WHERE idTransaction& = $1", idTransaction)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ventes []map[string]interface{}
	for rows.Next() {
		var loginBoutique string
		var nomProduit string
		var idClient string
		var quantite string
		var total string
		var date_vente string
		var statut string
		if err := rows.Scan(&loginBoutique, &nomProduit, &idClient, &quantite, &total, &date_vente, &statut); err != nil {
			return nil, err
		}
		ventes = append(ventes, map[string]interface{}{
			"idTransaction":  idTransaction,
			"login_boutique": loginBoutique,
			"nom_produit":    nomProduit,
			"idClient":       idClient,
			"quantite":       quantite,
			"total":          total,
			"date_vente":     date_vente,
			"statut":         statut,
		})
	}
	return ventes, nil
}

func GetAllVentesClient(idClient string) (map[string][]map[string]interface{}, error) {
	rows, err := DB.Query(
		"SELECT login_boutique, nom_produit, idProduit, quantite, total, date_vente, statut FROM ventes WHERE idClient = $1 ORDER BY idTransaction, date_vente DESC",
		idClient)
	if err != nil {
		return nil, fmt.Errorf("erreur lors de la requête: %v", err)
	}
	defer rows.Close()

	result := make(map[string][]map[string]interface{})

	for rows.Next() {
		var (
			idTransaction string
			loginBoutique string
			nomProduit    string
			quantite      string
			total         string
			dateVente     string
			statut        string
		)

		if err := rows.Scan(&idTransaction, &loginBoutique, &nomProduit, &quantite, &total, &dateVente, &statut); err != nil {
			return nil, fmt.Errorf("erreur lors du scan: %v", err)
		}

		result[idTransaction] = append(result[idTransaction], map[string]interface{}{
			"login_boutique": loginBoutique,
			"nom_produit":    nomProduit,
			"quantite":       quantite,
			"total":          total,
			"date_vente":     dateVente,
			"statut":         statut,
		})
	}

	if len(result) == 0 {
		return nil, fmt.Errorf("aucune transaction trouvée pour ce client")
	}

	return result, nil
}

func GetAllVentesBoutique(loginBoutique string) (map[string]interface{}, error) {
	rows, err := DB.Query(
		"SELECT idTransaction, nom_produit, idClient, quantite, total, date_vente, statut FROM ventes WHERE login_boutique = $1 ORDER BY idTransaction, date_vente DESC",
		loginBoutique)

	if err != nil {
		return nil, fmt.Errorf("erreur lors de la requête: %v", err)
	}
	defer rows.Close()

	result := make(map[string]interface{})
	transactions := make(map[string][]map[string]interface{})

	for rows.Next() {
		var (
			idTransaction string
			nomProduit    string
			idClient      string
			quantite      string
			total         string
			dateVente     string
			statut        string
		)

		if err := rows.Scan(&idTransaction, &nomProduit, &idClient, &quantite, &total, &dateVente, &statut); err != nil {
			return nil, fmt.Errorf("erreur lors du scan: %v", err)
		}

		transactionItem := map[string]interface{}{
			"nom_produit": nomProduit,
			"quantite":    quantite,
			"total":       total,
			"date_vente":  dateVente,
			"statut":      statut,
			"client":      idClient,
		}

		transactions[idTransaction] = append(transactions[idTransaction], transactionItem)
	}

	if len(transactions) == 0 {
		return nil, fmt.Errorf("aucune transaction trouvée pour cette boutique")
	}

	result["boutique"] = loginBoutique
	result["total_transactions"] = len(transactions)
	result["transactions"] = transactions

	return result, nil
}

func DeleteAllTransaction(idTransaction string) error {
	result, err := DB.Exec("DELETE FROM transaction WHERE idTransaction = $1", idTransaction)

	if err != nil {
		return fmt.Errorf("failed to delete all transactions row: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to check rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("transaction row not found")
	}

	return nil
}

func UpdateTransactionStatut(idTransaction string, statut string) error {
	result, err := DB.Exec(`
        UPDATE ventes SET statut = $1 WHERE idTransaction = $2`, statut, idTransaction)

	if err != nil {
		return fmt.Errorf("failed to update all transaction row: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to check rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("transaction row not found")
	}

	return nil
}
