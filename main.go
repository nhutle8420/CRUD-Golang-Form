package main

import (
	"Go-crud/controllers/productcontroller"
	"net/http"
)

func main() {

	http.HandleFunc("/", productcontroller.Indext)
	http.HandleFunc("/Product", productcontroller.Indext)
	http.HandleFunc("/Product/Index", productcontroller.Indext)
	http.HandleFunc("/Product/Add", productcontroller.Addx)
	http.HandleFunc("/Product/Edit", productcontroller.Editx)
	http.HandleFunc("/Product/Delete", productcontroller.Deletex)

	http.ListenAndServe(":8088", nil)

}
