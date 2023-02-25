package database

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	envErr := godotenv.Load()
	if envErr != nil {
		fmt.Printf("Erro ao carregar credenciais %v\n", envErr)
	}
	var (
		user     = os.Getenv("MSSQL_DB_USER")
		password = os.Getenv("MSSQL_DB_PASSWORD")
		port     = os.Getenv("MSSQL_DB_PORT")
		database = os.Getenv("MSSQL_DB_DATABASE")
	)
	connectionString := fmt.Sprintf("user id=%s;password=%s;port=%s;database=%s", user, password, port, database)
	fmt.Println(connectionString)
	db, err := gorm.Open(sqlserver.Open(connectionString), &gorm.Config{})
	if err != nil {
		fmt.Printf("Falha ao conectar no banco %v\n", err)
	}

	return db
}
