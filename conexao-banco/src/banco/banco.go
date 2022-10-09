package banco

import (
	"database/sql"
	"fmt"
	"webapp/src/config"

	_ "github.com/go-sql-driver/mysql"
)

func Conectar() (*sql.DB, error) {
	db, erro := sql.Open("mysql", config.StringConexaoBanco)
	if erro != nil {
		fmt.Println(db)
		return nil, erro

	}
	if erro = db.Ping(); erro != nil {

		db.Close()
		return nil, erro
	}
	fmt.Println("conectado")
	return db, nil
}
