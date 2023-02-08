package main

import (
	"fmt"
	"log"
	"net/http"

	productRoutes "go/routes"
)

// Connection URI
const uri = "mongodb+srv://adminj:admin@cluster0.bsv3iaj.mongodb.net/test"

func main() {
	fmt.Println("Successfully connected and pinged.")
	r := productRoutes.Router()
	fmt.Printf("Starting server at port 8080...\n")
	 
    if err := http.ListenAndServe(":8080", r); err != nil {
        log.Fatal(err)
    }
}
