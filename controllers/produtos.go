package controllers

import (
	"html/template"
	"loja/check"
	"loja/models"
	"net/http"
	"strconv"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {

	tdProducts := models.AllProdutos()

	temp.ExecuteTemplate(w, "Index", tdProducts)
}

func New(w http.ResponseWriter, r *http.Request) {

	temp.ExecuteTemplate(w, "New", nil)

}

func Edit(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	tdProduct := models.BuscaProduto(id)

	temp.ExecuteTemplate(w, "Edit", tdProduct)
}

func Update(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

		id := r.FormValue("id")
		nome := r.FormValue("nome")
		preco := r.FormValue("preco")
		descricao := r.FormValue("descricao")
		quantidade := r.FormValue("quantidade")

		idConvertido, err := strconv.Atoi(id)
		check.Check(err)

		precoConvertido, err := strconv.ParseFloat(preco, 64)
		check.Check(err)

		quantConvertida, err := strconv.Atoi(quantidade)
		check.Check(err)

		models.AtualizaProduto(idConvertido, nome, precoConvertido, descricao, quantConvertida)

	}

	http.Redirect(w, r, "/", 301)
}

func Insert(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

		nome := r.FormValue("nome")
		preco := r.FormValue("preco")
		descricao := r.FormValue("descricao")
		quantidade := r.FormValue("quantidade")

		precoConvertido, err := strconv.ParseFloat(preco, 64)
		check.Check(err)

		quantConvertida, err := strconv.Atoi(quantidade)
		check.Check(err)

		models.InsereProduto(nome, precoConvertido, descricao, quantConvertida)

	}

	http.Redirect(w, r, "/", 301)

}

func Delete(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	models.RemoveProduto(id)

	http.Redirect(w, r, "/", 301)

}
