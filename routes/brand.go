package brandRoutes

import (
	brand "go/controllers"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/brands", brand.GetAllBrands).Methods("GET")
	router.HandleFunc("/api/brands/{id}", brand.GetSingleBrand).Methods("GET")
	router.HandleFunc("/api/brands", brand.CreateBrand).Methods("POST")
	router.HandleFunc("/api/brands/{id}", brand.EditBrand).Methods("PUT")
	router.HandleFunc("/api/brands/{id}", brand.DeleteBrand).Methods("DELETE")

	return router
}