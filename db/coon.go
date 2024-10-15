package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "go_db"
	port     = 5432
	user     = "postgres"
	password = "1234" // O valor da senha deve ser uma string
	dbname   = "postgres"
)

func ConnectDB() (*sql.DB, error) {
	// Corrigindo o espaço faltante na string de conexão
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Abrindo a conexão
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, fmt.Errorf("erro ao abrir a conexão com o banco de dados: %v", err)
	}

	// Verificando a conexão
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("erro ao testar a conexão com o banco de dados: %v", err)
	}

	fmt.Println("Conectado ao banco de dados " + dbname)

	return db, nil
}
