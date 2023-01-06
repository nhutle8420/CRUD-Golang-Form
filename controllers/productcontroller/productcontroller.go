package productcontroller

import (
	"Go-crud/entities"
	"Go-crud/libraries"
	"Go-crud/models"
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)

var products = models.NewProductModel()
var validation = libraries.NewValidation()

func Indext(w http.ResponseWriter, r *http.Request) {
	listproduct, _ := products.FinAll()
	var data = map[string]interface{}{
		"product": listproduct,
	}
	temp, err := template.ParseFiles("view/product/index.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(w, data)
}
func Editx(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		queryString := r.URL.Query()
		id, _ := strconv.ParseInt(queryString.Get("id"), 10, 64)
		var pro entities.Product
		products.Find(id, &pro)

		data := map[string]interface{}{
			"t": pro,
		}
		temp, err := template.ParseFiles("view/product/edit.html")
		if err != nil {
			fmt.Println("Controller 1")
			panic(err)
		}
		temp.Execute(w, data)
	} else if r.Method == http.MethodPost {
		r.ParseForm()
		var product entities.Product
		product.ProductID, _ = strconv.ParseInt(r.Form.Get("id"), 10, 64)
		product.ProductName = r.Form.Get("ProductName")
		product.ProductNumber = r.Form.Get("ProductNumber")
		product.ProductPrice = r.Form.Get("ProductPrice")
		var data = make(map[string]interface{})
		vErrors := validation.Struct(product)
		fmt.Print(product)
		if vErrors != nil {
			data["product"] = product
			data["validation"] = vErrors
		} else {
			data["message"] = "Data pasien berhasil disimpan"
			products.Update(product)
		}
		temp, _ := template.ParseFiles("view/product/add.html")
		temp.Execute(w, data)

	}

}
func Addx(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		temp, err := template.ParseFiles("view/product/add.html")
		if err != nil {
			fmt.Println("Controller 1")
			panic(err)
		}
		temp.Execute(w, nil)
	} else if r.Method == http.MethodPost {
		r.ParseForm()
		var product entities.Product
		product.ProductName = r.Form.Get("ProductName")
		product.ProductNumber = r.Form.Get("ProductNumber")
		product.ProductPrice = r.Form.Get("ProductPrice")

		var data = make(map[string]interface{})

		vErrors := validation.Struct(product)

		if vErrors != nil {
			data["message"] = product
			data["validation"] = vErrors
		} else {
			data["pesan"] = "Data pasien berhasil disimpan"
			products.Create(product)
		}
		temp, _ := template.ParseFiles("view/product/add.html")
		temp.Execute(w, data)

	}

}
func Deletex(w http.ResponseWriter, r *http.Request) {
	queryString := r.URL.Query()
	id, _ := strconv.ParseInt(queryString.Get("id"), 10, 64)
	products.Delete(id)
	fmt.Print(id)
	http.Redirect(w, r, "/Product", http.StatusSeeOther)
}
