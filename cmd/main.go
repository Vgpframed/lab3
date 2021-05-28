package main

import (
	db "lab3/internal/DB"
	product "lab3/internal/products"
	"os"
)
var Filename = string(os.Args[1])
func main() {
	var Keyboard *product.Product
//	var Mouse *product.Product

	Keyboards := Keyboard.GetProduct(Filename, "1")
	//Mouses := Mouse.GetProduct(Filename, "0")
	db.InsertProducts(&Keyboards, "1")
	//db.InsertProducts(&Mouses, "0")
	db.SelectAllProducts()
	db.UserSelect()
	db.UpdDel()
}