package config

import (
	"fmt"

	"db/db"
)

func Init() {
	db.Init()

	fmt.Println("Servidor de autenticación iniciado")
}
