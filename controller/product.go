package controller

import (
	"database/sql"
	"github.com/mert-erol/golang_wms/database"
	"net/http"
)

func AddProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.ServeFile(w, r, "views/products.gohtml")
		return
	}

	db = database.ConnectDB()

	productName := r.FormValue("product_name")
	stockCode := r.FormValue("stock_code")

	var product string

	err := db.QueryRow("SELECT product_name FROM products WHERE product_name=?", productName).Scan(&product)

	switch {
	case err == sql.ErrNoRows:

		_, err = db.Exec("INSERT INTO products(product_name, stock_code) VALUES(?, ?)", productName, stockCode)
		if err != nil {
			http.Error(w, "Server hatası, ürün oluşturulamadı.", 500)
			return
		}

		w.Write([]byte("Ürün Oluşturuldu!"))
		return
	case err != nil:
		http.Error(w, "Server hatası, ürün oluşturulamadı3.", 500)
		return
	default:
		http.Redirect(w, r, "/products", 301)
	}
}
