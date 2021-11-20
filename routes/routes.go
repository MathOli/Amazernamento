package routes

import (
	control "loja/controllers"
	"net/http"
)

func LoadRoutes() {

	http.HandleFunc("/", control.Index)
	http.HandleFunc("/new", control.New)
	http.HandleFunc("/insert", control.Insert)
	http.HandleFunc("/delete", control.Delete)
	http.HandleFunc("/edit", control.Edit)
	http.HandleFunc("/update", control.Update)

}
