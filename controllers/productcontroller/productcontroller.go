package productcontroller

import (
	"Go-crud/entities"
	"fmt"
	"net/http"
	"text/template"
)

func Indext(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("view/product/index.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(w, nil)
}
func Editx(w http.ResponseWriter, r *http.Request) {

}
func Addx(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		temp, err := template.ParseFiles("view/product/add.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(w, nil)
	} else if r.Method == http.MethodPost {
		r.ParseForm()
		var product entities.Product
		product.ProductName = r.Form.Get("ProductName")
		product.ProductNumber = r.Form.Get("ProductNumber")
		product.ProductPrice = r.Form.Get("ProductPrice")
		fmt.Println(product)
	}

}
func Deletex(w http.ResponseWriter, r *http.Request) {

}
