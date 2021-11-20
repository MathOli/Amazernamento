package models

import (
	"loja/check"
	"loja/db"
)

type Produto struct {
	Id              int
	Nome, Descricao string
	Preco           float64
	Quantidade      int
}

func AllProdutos() []Produto {

	db := db.InitBd()

	selectProdutos, err := db.Query("select * from produtos order by id asc")

	check.Check(err)

	p := Produto{}
	produtos := []Produto{}

	for selectProdutos.Next() {
		var (
			id, quantidade  int
			nome, descricao string
			preco           float64
		)

		err = selectProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		check.Check(err)

		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}

	defer db.Close()

	return produtos

}

func InsereProduto(nome string, preco float64, descricao string, quantidade int) {

	db := db.InitBd()

	insertProduto, err := db.Prepare("insert into produtos(nome, descricao, preco, quantidade) values($1, $2, $3, $4)")
	check.Check(err)

	insertProduto.Exec(nome, descricao, preco, quantidade)

	defer db.Close()
}

func RemoveProduto(id string) {

	db := db.InitBd()

	deleteProduto, err := db.Prepare("delete from produtos where id = $1")
	check.Check(err)
	deleteProduto.Exec(id)

	defer db.Close()
}

func BuscaProduto(id string) Produto {

	db := db.InitBd()

	selectProduto, err := db.Query("select * from produtos where id = $1", id)
	check.Check(err)

	p := Produto{}
	for selectProduto.Next() {
		var (
			id, quantidade  int
			nome, descricao string
			preco           float64
		)

		err = selectProduto.Scan(&id, &nome, &descricao, &preco, &quantidade)
		check.Check(err)

		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

	}

	defer db.Close()
	return p
}

func AtualizaProduto(id int, nome string, preco float64, descricao string, quantidade int) {

	db := db.InitBd()

	updateProduto, err := db.Prepare("update produtos set nome = $1, descricao = $2, preco = $3, quantidade = $4 where id = $5")
	check.Check(err)

	updateProduto.Exec(nome, descricao, preco, quantidade, id)

	defer db.Close()
}
