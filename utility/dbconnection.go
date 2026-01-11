package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

type DBConnection struct {
	Server       *options.ServerAPIOptions
	Options      *options.ClientOptions
	Client       *mongo.Client
	DBCollection string
	DBName       string
}

func (db *DBConnection) setup(serverURI string) {
	db.Server = options.ServerAPI(options.ServerAPIVersion1)
	db.Options = options.Client().ApplyURI(serverURI).SetServerAPIOptions(db.Server)
}

func (db *DBConnection) Connect(serverURI string) *DBConnection {
	var err error
	db.setup(serverURI)

	db.Client, err = mongo.Connect(db.Options)
	if err != nil {
		fmt.Print("Could not connect to database")
	}

	return db
}

func (db *DBConnection) IsRunning() bool {
	if err := db.Client.Ping(context.TODO(), readpref.Primary()); err != nil {
		fmt.Print("Could not ping database")
		return false
	} else {
		return true
	}
}

func (db *DBConnection) Disconnect() {
	if err := db.Client.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}

func (db *DBConnection) Collection(collection string) *DBConnection {
	db.DBCollection = collection
	return db
}

func (db *DBConnection) Database(data string) *DBConnection {
	db.DBName = data
	return db
}

func (db *DBConnection) Add(data any) {
	collection := db.Client.Database(db.DBName).Collection(db.DBCollection)

	inserResult, err := collection.InsertOne(context.Background(), data)

	if err != nil {
		fmt.Print("Error adding data")
		return
	}

	fmt.Print(inserResult.InsertedID)
}
