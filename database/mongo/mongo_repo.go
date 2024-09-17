package database

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/przant/zipcodes-api/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type MongoRepo struct {
	clt  *mongo.Client
	coll *mongo.Collection
}

func NewMongoRepo() (*MongoRepo, error) {
	if err := godotenv.Load(".env"); err != nil {
		return nil, err
	}

	usrName := os.Getenv("MONGODB_USERNAME")
	passwd := os.Getenv("MONGODB_PASSWORD")
	uri := fmt.Sprintf("mongodb://%s:%s@mongo-db:27017/", usrName, passwd)

	client, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	return &MongoRepo{
		clt: client,
	}, nil
}

func (mdb *MongoRepo) Close() {
	if err := mdb.clt.Disconnect(context.TODO()); err != nil {
		log.Fatalf("while disconnecting the MongoDB database: %s", err)
	}
}

func (mdb *MongoRepo) FetchByZipcode(zipcode string) (*models.Zipcode, error) {
	result := &bson.M{}
	ctx := context.TODO()

	coll := mdb.clt.Database(os.Getenv("MONGODB_DATABASE")).Collection(os.Getenv("MONGODB_COLLECTION"))
	log.Printf("%v\n", result)
	log.Printf("%v\n", coll)
	if err := coll.FindOne(ctx, bson.D{{Key: "zipcode", Value: zipcode}}).Decode(&result); err != nil {
		return nil, err
	}

	zip := &models.Zipcode{}
	err := json.NewDecoder(strings.NewReader(result.String())).Decode(zip)

	if err != nil {
		return nil, err
	}

	return zip, err
}

func (mdb *MongoRepo) FetchByCounty(county string) ([]models.Zipcode, error) {
	return nil, nil
}

func (mdb *MongoRepo) FetchByStateCounty(state, county string) ([]models.Zipcode, error) {
	return nil, nil
}

func (mdb *MongoRepo) FetchByStateCity(state, city string) ([]models.Zipcode, error) {
	return nil, nil
}

func (mdb *MongoRepo) FetchByCountyCity(county, city string) ([]models.Zipcode, error) {
	return nil, nil
}
