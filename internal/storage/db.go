package storage

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitPostgres() error {
	connStr := os.Getenv("DATABASE_URL")
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	/*_, err = DB.Exec(`CREATE EXTENSION IF NOT EXISTS unaccent;`)
	if err != nil {
		return err
	}*/

	_, err = DB.Exec(`
    CREATE TABLE IF NOT EXISTS clients (
        login TEXT PRIMARY KEY,
        password TEXT NOT NULL,
        nom TEXT NOT NULL,
        prenom TEXT NOT NULL,
        email TEXT NOT NULL,
        date_naissance TEXT NOT NULL,
        telephone TEXT,
        image BYTEA
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
        iban TEXT NOT NULL,
        image BYTEA
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

	/*_, err = DB.Exec(`
	DROP TABLE IF EXISTS produits`)
	if err != nil {
		return err
	}*/

	_, err = DB.Exec(`
    CREATE TABLE IF NOT EXISTS produits (
        login_boutique TEXT NOT NULL,
        nom TEXT NOT NULL,
        categories TEXT,
        reduction TEXT,
        prix TEXT NOT NULL,
        description TEXT NOT NULL,
        image BYTEA,
        PRIMARY KEY(login_boutique, nom)
    )`)
	if err != nil {
		return err
	}

	/*_, err = DB.Exec(`
	  DROP TABLE IF EXISTS panier`)
	if err != nil {
		return err
	}*/
	_, err = DB.Exec(`
    CREATE TABLE IF NOT EXISTS panier (
        login_boutique TEXT NOT NULL,
        nom_produit TEXT NOT NULL,
        idClient TEXT NOT NULL,
        quantite TEXT NOT NULL,
        PRIMARY KEY(login_boutique, nom_produit, idClient)
    )`)
	if err != nil {
		return err
	}

	/*_, err = DB.Exec(`
	  DROP TABLE IF EXISTS favoris`)
	  if err != nil {
	      return err
	  }*/
	_, err = DB.Exec(`
    CREATE TABLE IF NOT EXISTS favoris (
        login_boutique TEXT NOT NULL,
        nom_produit TEXT NOT NULL,
        idClient TEXT NOT NULL,
		PRIMARY KEY(login_boutique, nom_produit, idClient)
    )`)
	if err != nil {
		return err
	}

	/*_, err = DB.Exec(`
	DROP TABLE IF EXISTS ventes`)
	if err != nil {
		return err
	}*/
	_, err = DB.Exec(`
    CREATE TABLE IF NOT EXISTS ventes (
        idTransaction TEXT NOT NULL,
        login_boutique TEXT NOT NULL,
        nom_produit TEXT NOT NULL,
        idClient TEXT NOT NULL,
        nom_adresse TEXT NOT NULL,
        nom_paiement TEXT NOT NULL,
        quantite TEXT NOT NULL,
        total TEXT NOT NULL,
        date_vente TEXT NOT NULL,
        statut TEXT NOT NULL,
		PRIMARY KEY(idTransaction, login_boutique, nom_produit)
    )`)
	if err != nil {
		return err
	}

	_, err = DB.Exec(`
    CREATE TABLE IF NOT EXISTS commentaires (
        login_boutique TEXT NOT NULL,
        nom_produit TEXT NOT NULL,
        idClient TEXT NOT NULL,
        note TEXT NOT NULL,
        commentaire TEXT,
		PRIMARY KEY(login_boutique, nom_produit, idClient)
    )`)
	if err != nil {
		return err
	}

	return err
}
