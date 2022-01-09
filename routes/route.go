package routes

import (
	"github.com/mert-erol/golang_wms/controller"
	"net/http"
)

func Application() {
	http.HandleFunc("/", controller.Dashboard)
	http.HandleFunc("/login", controller.Login)
	http.HandleFunc("/logout", controller.Logout)
	http.HandleFunc("/signup", controller.Signup)
	http.HandleFunc("/shelfs", controller.Shelfs)
	http.HandleFunc("/receipt", controller.AddReceipt)
	http.HandleFunc("/products", controller.AddProduct)

	http.ListenAndServe(":5567", nil)
}
