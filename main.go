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
	fmt.Println("create")
	response.Header().Add("content-type", "application/json") // serve in
	var person Person
	json.NewDecoder(request.Body).Decode(&person)
	collection := client.Database("test").Collection("people")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, err := collection.InsertOne(ctx, person)
	if err != nil {
		fmt.Println("err insert", err.Error())
	}
	json.NewEncoder(response).Encode(result)

}

func main() {
	fmt.Println("Starting")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	var errdb interface{}                                                                                                                                      //create context with timeout
	client, errdb = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://test:test@api_mongo_db:27017/?authSource=admin&readPreference=primary&ssl=false")) //connect to the db TODO get params from conf
	if errdb != nil {
		fmt.Println("cannot connect to DB")
	}
	router := mux.NewRouter()                            // define the router
	router.HandleFunc("/person", Create).Methods("POST") // create the route
	http.ListenAndServe(":8086", router)                 // define the port TODO get from conf
}
