package main

import (
    "database/sql"
    "log"
    "os"
    _ "github.com/mattn/go-sqlite3" 
)

//Models:
type UserModel struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

func setupDatabase(dbPath string) (*sql.DB, error) {
    log.Printf("Conectando a la base de datos en: %s", dbPath)
    db, err := sql.Open("sqlite3", dbPath)
    if err != nil {
        return nil, err
    }

    // Es buena idea hacer ping para verificar la conexión inmediatamente
    if err = db.Ping(); err != nil {
        db.Close() // Cerrar si el ping falla
        return nil, err
    }

    log.Println("Base de datos conectada exitosamente.")
    return db, nil
}
