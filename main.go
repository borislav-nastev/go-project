package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Connection URI
const uri = "mongodb+srv://adminj:admin@cluster0.bsv3iaj.mongodb.net/test"

func main() {
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	// Ping the primary
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected and pinged.")

	fmt.Printf("Starting server at port 8080\n")
	
	// TODO need to be moved to folder routes
	// http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request){
    //     fmt.Fprintf(w, "Hello!")
    // })
	 
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}
