package controllers

import (
	"encoding/json"
	"fmt"
	"go/models"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var products []models.Product
var product []models.Product

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Endpoint hit: GetAllProducts")
	w.Header().Set("Content-Type", "application/json")

	resp := make(map[string]interface{})
	resp["message"] = "Status OK products are here"
	resp["data"] = products

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonResp)
	return
}

func GetSingleProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Println(("Endpoint hit: GetSingleProduct"))
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	resp := make(map[string]interface{})
	fmt.Println(params["id"])
	fmt.Println(products)
	for _, product := range products{
		if product.ProductId == params["id"]{
			resp["message"] = "Status Ok product is here"
			resp["data"] = product
			jsonResp, err := json.Marshal(resp)
			if err != nil {
				log.Fatalf("Error happened in JSON marshal. Err: %s", err)
			}

			w.WriteHeader(http.StatusOK)
			w.Write(jsonResp)
			return
		} 
	}

	resp["message"] = "Not found"
	resp["data"] = nil

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}

	w.WriteHeader(http.StatusNotFound)
	w.Write(jsonResp)
	return
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Println(("Endpoint hit: CreateProduct"))
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]interface{})
	resp["message"] = "Status OK product was created"
	resp["data"] = product
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
	return
}

func EditProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Println(("Endpoint hit: EditProduct"))
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]interface{})
	resp["message"] = "Status OK product was edited"
	resp["data"] = product
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
	return
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Println(("Endpoint hit: DeleteProduct"))
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
	resp["message"] = "Status OK product was deleted"
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
	return
}