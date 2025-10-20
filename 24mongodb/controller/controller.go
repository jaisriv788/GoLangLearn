package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jaisriv788/mongodb/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const connectionString = "mongodb+srv://jai9450:jai9450@cluster1.ocfuth6.mongodb.net/?retryWrites=true&w=majority&appName=Cluster1"
const dbName = "netflix"
const collectionName = "watchlist"

var collection *mongo.Collection

// connect with mongodb
func init() {
	//Creating a context with timeout for the connection

	//This to use when we know that you are on the root of you program
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//Use this when do not yet have proper context
	//ctx := context.TODO()

	//client options
	clientOptions := options.Client().ApplyURI(connectionString)

	//connect to mongodb
	client, err := mongo.Connect(clientOptions)
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Could not ping to MongoDB:", err)
	}

	fmt.Println("✅ Successfully connected to MongoDB!")

	collection = client.Database(dbName).Collection(collectionName)

	//collection reference is ready
	fmt.Println("📁 Collection reference is ready:", collection.Name())
}

//Mongodb helpers

//insert one record

func insertOneMovie(movie model.Netflix) {
	fmt.Println(movie)
	inserted, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted 1 movie in db with id: ", inserted)
}

func updateOneMovie(movieId string) {
	id, err := bson.ObjectIDFromHex(movieId)

	if err != nil {
		log.Fatal("Invalid movie ID:", err)
	}

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()

	updation, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Fatal(err)
	}

	if updation.MatchedCount == 0 {
		fmt.Println("No match with ObjectID. Trying with string ID...")

		// TRY 2: Search with string ID (if stored as string)
		filterString := bson.M{"_id": movieId}
		updation, err = collection.UpdateOne(context.Background(), filterString, update)
		if err != nil {
			log.Println("Update error with string:", err)
			return
		}

		if updation.MatchedCount == 0 {
			fmt.Println("Still no movie found. Check your data structure.")

			// DEBUG: Find any document to see the actual structure
			var result bson.M
			err := collection.FindOne(context.Background(), bson.M{}).Decode(&result)
			if err != nil {
				log.Println("Debug query error:", err)
			} else {
				fmt.Printf("Sample document structure: %+v\n", result)
			}
			return
		}
	}

	fmt.Printf("Successfully updated! Matched: %d, Modified: %d\n",
		updation.MatchedCount, updation.ModifiedCount)

	fmt.Println("Updated 1 movie in db with id: ", updation.MatchedCount)
}

func deleteOneMovie(movieId string) {
	id, err := bson.ObjectIDFromHex(movieId)

	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": id}

	deletion, err := collection.DeleteOne(context.Background(), filter, nil)

	fmt.Println("Deleted 1 movie in db with id: ", deletion)
}

func deleteAllMovie() {
	filter := bson.D{{}}
	deleted, err := collection.DeleteMany(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Number of movies deleted: ", deleted, nil)
}

func allmovies() []bson.M {
	cursor, err := collection.Find(context.Background(), bson.D{{}})

	if err != nil {
		log.Fatal(err)
	}

	var movies []bson.M

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

//Actual controllets

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	allMovies := allmovies()

	json.NewEncoder(w).Encode(allMovies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie model.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)

	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)

	updateOneMovie(params["id"])

	json.NewEncoder(w).Encode(params["id"])
}

func DeleteOneMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)

	deleteOneMovie(params["id"])

	json.NewEncoder(w).Encode("Delted the movie.")
}

func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	deleteAllMovie()

	json.NewEncoder(w).Encode("Deleted all movie.")
}
