package mongo_connection

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/pratikdev/learning-golang/mongoApi/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var Collection *mongo.Collection

// initiating mongodb connection
func init() {
	// loading env vars
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	const db_name = "netflix-go"
	const collection_name = "watchList"
	connection_string := os.Getenv("MONGODB_CONNECTION_STRING")

	// client options
	clientOption := options.Client().ApplyURI(connection_string)

	// database options
	dbOption := options.Database()

	// connect to db
	client, err := mongo.Connect(clientOption)
	if err != nil {
		log.Fatal(err)
	}

	Collection = client.Database(db_name, dbOption).Collection(collection_name)

	fmt.Println("Collection instance is ready")
}

func InsertMovie(movie model.Netflix) {
	_, err := Collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}
}

func UpdateMovie(movieId string) {
	movieObjectId, err := bson.ObjectIDFromHex(movieId)
	if err != nil {
		log.Fatal(err)
	}

	_, resultErr := Collection.UpdateByID(context.Background(), movieObjectId, bson.M{"$set": bson.M{"watched": true}})
	if resultErr != nil {
		log.Fatal(resultErr)
	}
}

func DeleteMovie(movieId string) {
	movieObjectId, err := bson.ObjectIDFromHex(movieId)
	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": movieObjectId}
	_, resultErr := Collection.DeleteOne(context.Background(), filter)
	if resultErr != nil {
		log.Fatal(resultErr)
	}
}

func DeleteAllMovies() {
	_, err := Collection.DeleteMany(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
}

func GetMovie(movieId string) bson.M {
	movieObjectId, err := bson.ObjectIDFromHex(movieId)
	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": movieObjectId}
	result := Collection.FindOne(context.Background(), filter)

	var movie bson.M
	result.Decode(&movie)

	return movie
}

func GetAllMovies() []bson.M {
	ctx := context.Background()
	cursor, err := Collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	var movies []bson.M

	// loop thru `cursor.next`
	for cursor.Next(ctx) {
		var movie bson.M
		if err := cursor.Decode(&movie); err != nil {
			log.Fatal(err)
		}

		movies = append(movies, movie)
	}

	return movies
}
