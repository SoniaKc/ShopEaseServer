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
        CREATE TABLE IF NOT EXISTS users (
            id SERIAL PRIMARY KEY,
            firstname TEXT NOT NULL,
			lastname TEXT NOT NULL
        )`)
    if err != nil {
        return err
    }
    _, err = DB.Exec(`
    CREATE TABLE IF NOT EXISTS plantCollections (
        id SERIAL PRIMARY KEY,
        userID INTEGER REFERENCES users NOT NULL,
        collectionName TEXT NOT NULL
    )`)
    if err != nil {
        return err
    }

    _, err = DB.Exec(`
        CREATE TABLE IF NOT EXISTS plants (
            id SERIAL PRIMARY KEY,
            plantCollectionID INTEGER REFERENCES plantCollections NOT NULL,
            plantName TEXT NOT NULL,
            azoteFixing REAL NOT NULL,
            upgradeGround REAL NOT NULL,
            waterFixing REAL NOT NULL
        )`)
    if err != nil {
        return err
    }

    return err
}

func isValidTableName(name string) bool {
    // Only allow alphanumeric and underscores
    matched,_ := regexp.MatchString(`^[a-zA-Z][a-zA-Z0-9_]*$`, name)
    return matched
}

func sanitizeTableName(name string) string {
    // Quote the identifier to handle special characters
    return `"` + strings.ReplaceAll(name, `"`, `""`) + `"`
}