package main

import (
	"database/sql"
)

func HashPassword(password string) (string, error) {

	return password, nil
}

func registrarUsuario(db *sql.DB, nombre string, correo string, contrasena string) error {
	// Hashear la contraseña
	contrasenaHash, err := HashPassword(contrasena)
	if err != nil {
		return err
	}

	// Preparar la consulta SQL
	stmt, err := db.Prepare("INSERT INTO usuarios (nombre, correo, contrasena) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Ejecutar la consulta SQL
	_, err = stmt.Exec(nombre, correo, contrasenaHash)
	if err != nil {
		return err
	}

	return nil
}
