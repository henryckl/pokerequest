package main

import (
	"github.com/henryckl/pokerequest/database"
)

func main() {
	db := database.Init()
	// Teste com procedure
	db.Exec("insertTeste", "testando db")
}
