// main.go

package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SpiderInfo struct {
	// ... (other fields)
	Family       string    `json:"family" bson:"family"`
	Genus        string    `json:"genus" bson:"genus"`
	Species      string    `json:"species" bson:"species"`
	Author       string    `json:"author" bson:"author,omitempty"`
	PublishYear  string    `json:"publish_year" bson:"publish_year,omitempty"`
	Country      string    `json:"country" bson:"country,omitempty"`
	CountryOther string    `json:"country_other" bson:"country_other,omitempty"`
	Altitude     string    `json:"altitude" bson:"altitude,omitempty"`
	Method       string    `json:"method" bson:"method,omitempty"`
	Habital      string    `json:"habitat" bson:"habitat,omitempty"`
	Microhabital string    `json:"microhabitat" bson:"microhabitat,omitempty"`
	Designate    string    `json:"designate" bson:"designatel,omitempty"`
	Address      []Address `json:"address" bson:"address,omitempty"`
	Paper        []string  `json:"paper" bson:"paper,omitempty"`
	Status       string    `json:"status" bson:"status"`
	CreatedAt    time.Time `json:"created_at,omitempty" bson:"created_at"`
	UpdatedAt    time.Time `json:"updated_at,omitempty" bson:"updated_at"`
}

type Address struct {
	// ... (other fields)
	Province string     `json:"province" bson:"province"`
	District string     `json:"district" bson:"district"`
	Locality string     `json:"locality" bson:"locality"`
	Position []Position `json:"position" bson:"position"`
}

type Position struct {
	// ... (other fields)
	Name      string  `json:"name" bson:"name"`
	Latitude  float64 `json:"latitude" bson:"latitude"`
	Longitude float64 `json:"longitude" bson:"longitude"`
}

func main() {
	// Parse command-line arguments
	inputFileName := flag.String("i", "", "Input file name (JSON)")
	queryIndex := flag.Int("q", -1, "Query item index (0-based)")
	flag.Parse()

	if *inputFileName == "" && *queryIndex == -1 {
		log.Fatal("Please provide either an input file name using the -i flag or a query item index using the -q flag.")
	}

	// Set your MongoDB connection string
	uri := "mongodb://user:1234@localhost:27017"

	// Set client options
	clientOptions := options.Client().ApplyURI(uri)

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Ping the database to check if the connection is successful
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	// Disconnect from MongoDB
	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			log.Fatal(err)
		}
	}()

	// Use the client to create a new database
	dbName := "spiderTh"
	client.Database(dbName)

	// Create a new collection for SpiderInfo
	collectionName := "spiderInfo"
	collection := client.Database(dbName).Collection(collectionName)

	if *inputFileName != "" {
		// Read data from JSON file
		data, err := ioutil.ReadFile(*inputFileName)
		if err != nil {
			log.Fatal(err)
		}

		// Unmarshal JSON data
		var spiderData []SpiderInfo
		if err := json.Unmarshal(data, &spiderData); err != nil {
			log.Fatal(err)
		}

		// Insert data into the collection
		for _, data := range spiderData {
			if err := insertSpiderInfo(collection, data); err != nil {
				log.Fatal(err)
			}
		}

		fmt.Println("Data inserted successfully!")
	}

	if *queryIndex != -1 {
		// Query data from the collection by index
		result, err := queryByIndex(collection, *queryIndex)
		if err != nil {
			log.Fatal(err)
		}

		// Print the result
		fmt.Println("Query Result:")
		fmt.Println(result)
	}
}

func insertSpiderInfo(collection *mongo.Collection, data SpiderInfo) error {
	// Insert document
	_, err := collection.InsertOne(context.Background(), data)
	return err
}

func queryByIndex(collection *mongo.Collection, index int) (SpiderInfo, error) {
	// Query data from the collection
	cursor, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		return SpiderInfo{}, err
	}

	defer cursor.Close(context.Background())

	// Iterate through the cursor to find the document at the specified index
	var result SpiderInfo
	currentIndex := 0
	for cursor.Next(context.Background()) {
		if currentIndex == index {
			if err := cursor.Decode(&result); err != nil {
				return SpiderInfo{}, err
			}
			return result, nil
		}
		currentIndex++
	}

	if err := cursor.Err(); err != nil {
		return SpiderInfo{}, err
	}

	// If the specified index is out of bounds, return an error
	return SpiderInfo{}, fmt.Errorf("index %d is out of bounds", index)
}
