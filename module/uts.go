package package_uts

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/bson"
	"os"
)

var MongoString string = os.Getenv("MONGOSTRING")

func MongoConnect(dbname string) (db *mongo.Database) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoString))
	if err != nil {
		fmt.Printf("MongoConnect: %v\n", err)
	}
	return client.Database(dbname)
}

func InsertOneDoc(db string, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}

func InsertMuseumCollection(name string, description string, items []interface{}, location string, openingHours string, establishedYear int, director string, website string, visitorCount int) (insertedID interface{}) {
	var collection = map[string]interface{}{
		"_id":             primitive.NewObjectID(),
		"name":            name,
		"description":     description,
		"items":           items,
		"location":        location,
		"openingHours":    openingHours,
		"establishedYear": establishedYear,
		"director":        director,
		"website":         website,
		"visitorCount":    visitorCount,
	}
	return InsertOneDoc("ats", "museum", collection)
}

func InsertMuseumItem(name string, description string, year int, artist string, medium string, dimensions string, origin string, acquisition string, condition string) (insertedID interface{}) {
	var item = map[string]interface{}{
		"_id":         primitive.NewObjectID(),
		"name":        name,
		"description": description,
		"year":        year,
		"artist":      artist,
		"medium":      medium,
		"dimensions":  dimensions,
		"origin":      origin,
		"acquisition": acquisition,
		"condition":   condition,
	}
	return InsertOneDoc("ats", "museum", item)
}

func GetAllMuseumCollections() (collections []interface{}) {
	collection := MongoConnect("ats").Collection("museum")
	filter := bson.M{}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetAllMuseumCollections:", err)
		return
	}
	err = cursor.All(context.TODO(), &collections)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func GetAllMuseumItems() (items []interface{}) {
	collection := MongoConnect("ats").Collection("museum")
	filter := bson.M{}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetAllMuseumItems:", err)
		return
	}
	err = cursor.All(context.TODO(), &items)
	if err != nil {
		fmt.Println(err)
	}
	return
}
