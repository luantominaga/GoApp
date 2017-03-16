package main

import (
	"database/sql"
	"net/http"
	"fmt"

	_ "github.com/go-sql-driver/mysql"


)

var db *sql.DB
var err error

// // Insertin to lista //

func PostData(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.ServeFile(res, req, "template/formulario.html")
		return
	}

	//Dados
	product_name := req.FormValue("product_name")
	product_price := req.FormValue("product_price")
	product_description := req.FormValue("product_description")
	product_amount := req.FormValue("product_amount")

	//Botões
	adicionar := req.FormValue("adicionar")
	//atualizar := req.FormValue("atualizar")
	//delete := req.FormValue("delete")


	var id string

	var databaseProduto string
	var databasePreco string
	var databaseDescricao string
	var databaseQuantidade string


	err := db.QueryRow("SELECT id, Produto, Preco, Descricao, Quantidade FROM users WHERE id=?", id).Scan(&databaseProduto, &databasePreco, &databaseDescricao , &databaseQuantidade)


	//Adicionar
	if adicionar == "adicionar"{
	_, err = db.Exec("INSERT INTO users(Produto, Preco, Descricao, Quantidade) VALUES(?, ?, ?, ?)", product_name, product_price, product_description, product_amount)
		if err != nil {
			http.Error(res, "Server não conectado", 500)
			return
		}
		return
		}

		/*
		//Atualizar
		if atualizar == "atualizar"{
			_, err = db.Exec("update users set Produto= ?, Preco= ?, Descricao= ?, Quantidade= ? where id= ? LIMIT 1", product_name, product_price, product_description, product_amount)
				if err != nil {
					http.Error(res, "Server não conectado", 500)
					return
				}

		}

		//Deletar
		if delete == "delete"{
			_, err = db.Exec("delete from users where id= 1")
				if err != nil {
					http.Error(res, "Server não conectado", 500)
					return
				}

		}
		*/

}
// --Lista-- //


//Criar tabela
func CreateTable(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.ServeFile(res, req, "template/criartabela.html")
		return
	}

	criar := req.FormValue("criar")
	apagar := req.FormValue("apagar")

			if criar == "criar"{
			stmt, err := db.Prepare("CREATE TABLE users (id int NOT NULL AUTO_INCREMENT, Produto varchar(40), Preco float(12,4), Descricao varchar(255), Quantidade int(255) , PRIMARY KEY (id));")
		if err != nil {
			fmt.Println(err.Error())
		}
		_, err = stmt.Exec()
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("Criado com sucesso")
			return
		}
		}

		if apagar == "apagar"{
			stmt, err := db.Prepare("DROP TABLE users")
			if err != nil {
				fmt.Println(err.Error())
				}
			_, err = stmt.Exec()
			if err != nil {
			fmt.Println(err.Error())
			} else {
				fmt.Println("Deletado com sucesso")
				return
			}
		}
		}

//Página principal
func homePage(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "index.html")
}


// Função principal //
func main() {
	db, err = sql.Open("mysql", "root: @/webapp")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	err = db.Ping()

	if err != nil {
		panic(err.Error())
	}


	http.HandleFunc("/template/formulario", PostData)
	http.HandleFunc("/template/criartabela", CreateTable)
	http.HandleFunc("/", homePage)
	http.ListenAndServe(":8080", nil)

}
// --Função principal-- //
