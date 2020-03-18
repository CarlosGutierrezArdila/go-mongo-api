package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

type Person struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Firstname string             `json:"firstname,omitempty" bson:"firstname,omitempty"`
	Lastname  string             `json:"lastname,omitempty" bson:"lastname,omitempty"`
}

// func to create
func Create(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json") // serve in
	var person Person
	json.NewDecoder(request.Body).Decode(&person)
	collection := client.Database("test").Collection("people")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, _ := collection.InsertOne(ctx, person)
	json.NewEncoder(response).Encode(result)

}

func main() {
	fmt.Println("Starting")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)                    //create context with timeout
	client, _ = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017")) //connect to the db TODO get params from conf
	router := mux.NewRouter()                                                              // define the router
	router.HandleFunc("/person", Create).Methods("POST")                                   // create the route
	http.ListenAndServe(":8086", router)                                                   // define the port TODO get from conf
}
