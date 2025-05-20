package storage

import (
    "database/sql"
    "fmt"
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
            username TEXT NOT NULL,
            role TEXT NOT NULL,
            firstName TEXT NOT NULL,
            lastName TEXT NOT NULL,
            phone TEXT,
            email TEXT UNIQUE NOT NULL,
            pswrd TEXT NOT NULL
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