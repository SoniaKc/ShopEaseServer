package storage

import (
	"database/sql"
	"os"
	"regexp"
	"strings"

	_ "github.com/lib/pq" // Driver SQLite
)

var DB *sql.DB

func InitPostgres() error {
	// Récupérez l'URL de connexion depuis les variables d'environnement
	connStr := os.Getenv("DATABASE_URL")
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	_, err = DB.Exec(`
        CREATE TABLE IF NOT EXISTS clients (
            login TEXT PRIMARY KEY,
            password TEXT NOT NULL,
            nom TEXT NOT NULL,
			prenom TEXT NOT NULL,
            email TEXT NOT NULL,
            date_naissance TEXT NOT NULL,
            telephone TEXT
        )`)
	if err != nil {
		return err
	}

	_, err = DB.Exec(`
    CREATE TABLE IF NOT EXISTS boutiques (
        login TEXT PRIMARY KEY,
        password TEXT NOT NULL,
        nom TEXT NOT NULL,
        email TEXT NOT NULL,
        telephone TEXT,
        siret TEXT NOT NULL,
        forme_juridique TEXT NOT NULL,
        siege_social TEXT NOT NULL,
        pays_enregistrement TEXT NOT NULL,
        iban TEXT NOT NULL
    )`)
	if err != nil {
		return err
	}

	_, err = DB.Exec(`
    CREATE TABLE IF NOT EXISTS parametres (
        login TEXT NOT NULL,
        type TEXT NOT NULL,
        langue TEXT NOT NULL,
        cookies TEXT NOT NULL,
        notifications TEXT NOT NULL,
        PRIMARY KEY(login, type)
    )`)
	if err != nil {
		return err
	}

	_, err = DB.Exec(`
    CREATE TABLE IF NOT EXISTS paiements (
            login TEXT NOT NULL,
            nom_carte TEXT NOT NULL,
            nom_personne_carte TEXT NOT NULL,
            numero TEXT NOT NULL,
            cvc TEXT NOT NULL,
            date_expiration TEXT NOT NULL,
            PRIMARY KEY(login, nom_carte)
        )`)
	if err != nil {
		return err
	}

	_, err = DB.Exec(`
    CREATE TABLE IF NOT EXISTS adresses (
        login TEXT NOT NULL,
        nom_adresse TEXT NOT NULL,
        numero TEXT NOT NULL,
        nom_rue TEXT NOT NULL,
        code_postal TEXT NOT NULL,
        ville TEXT NOT NULL,
        pays TEXT NOT NULL,
        PRIMARY KEY(login, nom_adresse)
    )`)
	if err != nil {
		return err
	}

	_, err = DB.Exec(`
    CREATE TABLE IF NOT EXISTS produits (
        login_boutique TEXT NOT NULL,
        nom TEXT NOT NULL,
        categories TEXT,
        reduction TEXT,
        prix TEXT NOT NULL,
        description TEXT NOT NULL,
        PRIMARY KEY(login_boutique, nom)
    )`)
	if err != nil {
		return err
	}

	_, err = DB.Exec(`
    CREATE TABLE IF NOT EXISTS panier (
        idProduit TEXT NOT NULL,
        idClient TEXT NOT NULL,
        quantite, INTEGER NOT NULL,
        PRIMARY KEY(idProduit, idClient),
    )`)
	if err != nil {
		return err
	}

	_, err = DB.Exec(`
    CREATE TABLE IF NOT EXISTS favoris (
        idProduit TEXT NOT NULL,
        idClient TEXT NOT NULL,
        PRIMARY KEY(idProduit, idClient),
    )`)
	if err != nil {
		return err
	}

	_, err = DB.Exec(`
    CREATE TABLE IF NOT EXISTS ventes (
        idTransaction TEXT NOT NULL PRIMARY KEY,
        idProduit TEXT NOT NULL,
        idClient TEXT NOT NULL,
        quantite TEXT NOT NULL,
        total TEXT NOT NULL,
        date_vente TEXT NOT NULL,
        statut TEXT NOT NULL,
    )`)
	if err != nil {
		return err
	}

	return err
}

func isValidTableName(name string) bool {
	// Only allow alphanumeric and underscores
	matched, _ := regexp.MatchString(`^[a-zA-Z][a-zA-Z0-9_]*$`, name)
	return matched
}

func sanitizeTableName(name string) string {
	// Quote the identifier to handle special characters
	return `"` + strings.ReplaceAll(name, `"`, `""`) + `"`
}
