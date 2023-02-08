package brandRoutes

import (
	products "go/controllers"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/products", products.GetProducts).Methods("GET")
	router.HandleFunc("/api/products/{id}", products.GetSingleProduct).Methods("GET")
	router.HandleFunc("/api/products", products.CreateProduct).Methods("POST")
	router.HandleFunc("/api/products/{id}", products.EditProduct).Methods("PUT")
	router.HandleFunc("/api/products/{id}", products.DeleteProduct).Methods("DELETE")

	return router
}