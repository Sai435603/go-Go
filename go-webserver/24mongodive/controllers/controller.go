package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"mongo-go/model"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://saivaraprasad073_db_user:tBYGRb3TkuJ10xaQ@cluster0.rqj1fra.mongodb.net/?appName=Cluster0"
const dbName = "netflix"
const collectionName = "watchlist"

//**important rev to connection

var collection *mongo.Collection

// connect with mongoDB
// init() is a special method that runs only once at start

func init() {
	//client option
	clientOption := options.Client().ApplyURI(connectionString)

	//connect to mongo
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MongoDB connection success :)")
	collection = client.Database(dbName).Collection(collectionName)

	fmt.Println("Collection instance is ready")
}

//MONODB helpers - file

//insert 1 movie

func insertOneMovie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.TODO(), movie)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted 1 movie in db with id : ", inserted.InsertedID)
}

//update 1 record

func updateOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Modified recourds count : ", result.ModifiedCount)
}

func deleteOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}

	deleteCount, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Movie got delete with delete count : ", deleteCount)
}

//delete all records from mongodb

func deleteAllMovie() int64 {
	filter := bson.D{{}}
	deleteResult, err := collection.DeleteMany(context.Background(), filter, nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Number of rows deleted : ", deleteResult.DeletedCount)

	return deleteResult.DeletedCount
}

func getAllMovies() []primitive.M {
	cursor, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var movies []primitive.M

	for cursor.Next(context.Background()) {
		var movie bson.M
		err := cursor.Decode(&movie)

		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}
	defer cursor.Close(context.Background())
	return movies
}

//Actual Controllers - file

func GetMyAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie model.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])

}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")
	params := mux.Vars(r)
	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")
	count := deleteAllMovie()
	json.NewEncoder(w).Encode(count)
}
