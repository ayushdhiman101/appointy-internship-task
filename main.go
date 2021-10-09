package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var client *mongo.Client

func main() {

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017/?readPreference=primary&appname=MongoDB%20Compass&directConnection=true&ssl=false"))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	// quickstartDatabase := client.Database("apointy")
	// usersCollection := quickstartDatabase.Collection("users")

	// cursor, err := usersCollection.Find(ctx, bson.M{})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// var users []bson.M
	// if err = cursor.All(ctx, &users); err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(users)

	router := mux.NewRouter()

	router.HandleFunc("/users", newuser).Methods("POST")
	router.HandleFunc("/users/{id}", getuser).Methods("GET")
	router.HandleFunc("/posts", newpost).Methods("POST")
	router.HandleFunc("/posts/{id}", getpost).Methods("GET")
	router.HandleFunc("/posts/users/{id}/{pages}", getpostwithusername).Methods("GET")
	http.ListenAndServe(":9090", router)

}
