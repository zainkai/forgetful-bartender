package persistence

import (
	"context"
	"fmt"
	"time"

	"github.com/mongodb/mongo-go-driver/bson/primitive"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"

	config "github.com/zainkai/forgetful-bartender/configs"
	"github.com/zainkai/forgetful-bartender/pkg/logger"
)

const pkgName string = "persistence"

// Move to pkg/models
type SpiritType string

const (
	Vodka   SpiritType = "VODKA"
	Gin     SpiritType = "GIN"
	Whiskey SpiritType = "WHISKEY"
	Rum     SpiritType = "RUM"
	Beer    SpiritType = "BEER"
	Wine    SpiritType = "WINE"
	Other   SpiritType = "OTHER"
)

type DB struct {
	collectionName string
	databaseName   string
	parent         (*mongo.Client)
	conn           (*mongo.Collection)
}

type Ingredient struct {
	Spirit          SpiritType `json:"Spirit" binding:"required"`
	Amount          string     `json:"Amount" binding:"required"`
	SuggestedSpirit []string   `json:"SuggestedSpirit" binding:"required"`
}

type Drink struct {
	Name           string       `json:"Name" binding:"required"`           // Secondary index
	MainSpiritType SpiritType   `json:"MainSpiritType" binding:"required"` // Primary index
	Description    string       `json:"Description"`
	Instructions   []string     `json:"Instructions" binding:"required"`
	Ingredients    []Ingredient `json:"Ingredients" binding:"required"`
}

func Init() *DB {
	config.LoadConfig()
	configuration := *config.Config
	dbConfig := configuration.Database
	client, err := mongo.Connect(context.Background(), dbConfig.URL)

	if err != nil {
		logger.Log(pkgName, "Could not connect to MongoDb", err)
	}

	db := client.Database(dbConfig.Name).Collection(dbConfig.Collection)
	fmt.Println("Connected to database")
	return &DB{
		collectionName: dbConfig.Collection,
		databaseName:   dbConfig.Name,
		parent:         client,
		conn:           db}
}

func (dbPtr *DB) Disconnect() {
	fmt.Println("Disconnecting from database...")
	err := dbPtr.parent.Disconnect(context.Background())

	if err != nil {
		logger.Log(pkgName, "Could not disconnect to MongoDb", err)
	}
	fmt.Println("Disconnected from database")
}

func (dbPtr *DB) CreateDrink(data Drink) (*mongo.InsertOneResult, error) {
	res, err := dbPtr.conn.InsertOne(context.Background(), data)

	if err != nil {
		logger.Log(pkgName, "Could not create Drink", err)
	}
	return res, err
}

func (dbPtr *DB) GetOneDrink(drinkId string) (*Drink, error) {
	var drink Drink

	drinkObjId, err1 := primitive.ObjectIDFromHex(drinkId)
	if err1 != nil {
		logger.Log(pkgName, "drinkId not in hex format", err1)
		return nil, err1
	}

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	err := dbPtr.conn.FindOne(ctx, bson.M{"_id": drinkObjId}).Decode(&drink)

	if err != nil {
		logger.Log(pkgName, "drinkId could not be found", err)
		return nil, err

	}

	return &drink, err
}
