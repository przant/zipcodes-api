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
		clt:  client,
		coll: client.Database(os.Getenv("MONGODB_DATABASE")).Collection(os.Getenv("MONGODB_COLLECTION")),
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

	if err := mdb.coll.FindOne(ctx, bson.D{{Key: "zipcode", Value: zipcode}}).Decode(&result); err != nil {
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
	ctx := context.TODO()

	cursor, err := mdb.coll.Find(ctx, bson.D{{Key: "county", Value: county}})
	if err != nil {
		return nil, err
	}

	return createResp(ctx, cursor)
}

func (mdb *MongoRepo) FetchByStateCounty(state, county string) ([]models.Zipcode, error) {
	ctx := context.TODO()

	cursor, err := mdb.coll.Find(ctx, bson.D{{Key: "state", Value: state}, {Key: "county", Value: county}})
	if err != nil {
		return nil, err
	}

	return createResp(ctx, cursor)
}

func (mdb *MongoRepo) FetchByStateCity(state, city string) ([]models.Zipcode, error) {
	ctx := context.TODO()

	cursor, err := mdb.coll.Find(ctx, bson.D{{Key: "state", Value: state}, {Key: "city", Value: city}})
	if err != nil {
		return nil, err
	}

	return createResp(ctx, cursor)
}

func (mdb *MongoRepo) FetchByCountyCity(county, city string) ([]models.Zipcode, error) {
	ctx := context.TODO()

	cursor, err := mdb.coll.Find(ctx, bson.D{{Key: "county", Value: county}, {Key: "city", Value: city}})
	if err != nil {
		return nil, err
	}

	return createResp(ctx, cursor)
}

func createResp(ctx context.Context, cursor *mongo.Cursor) ([]models.Zipcode, error) {
	results := make([]bson.M, 0)
	err := cursor.All(ctx, &results)

	if err != nil {
		return nil, err
	}

	zips := make([]models.Zipcode, 0)
	for _, result := range results {
		zip := models.Zipcode{}
		err = json.NewDecoder(strings.NewReader(result.String())).Decode(&zip)
		if err != nil {
			return nil, err
		}
		zips = append(zips, zip)
	}

	return zips, nil
}
