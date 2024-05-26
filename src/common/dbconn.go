package common

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DbConnection struct {
	Client   *mongo.Client
	ServerDb *mongo.Database

	LexicalCategoryCollection *mongo.Collection
	VocabularyCollection      *mongo.Collection
}

var DbConn DbConnection

func InitDbConnection() error {
	var err error

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	DbConn.Client, err = mongo.Connect(
		ctx,
		options.Client().ApplyURI(GlobalServerConfig.Database.Uri))
	if err != nil {
		return err
	}

	err = DbConn.Client.Ping(ctx, nil)
	if err != nil {
		return errors.New("Error pinging database server: " + err.Error())
	}

	DbConn.ServerDb = DbConn.Client.Database(GlobalServerConfig.Database.DatabaseName)

	DbConn.LexicalCategoryCollection = DbConn.ServerDb.Collection("lexical_category")
	if DbConn.LexicalCategoryCollection == nil {
		DeinitDbConnection()
		return errors.New("missing lexical category collection")
	}

	DbConn.VocabularyCollection = DbConn.ServerDb.Collection("vocabulary")
	if DbConn.VocabularyCollection == nil {
		DeinitDbConnection()
		return errors.New("missing vocabulary collection")
	}

	return nil
}

func DeinitDbConnection() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	DbConn.Client.Disconnect(ctx)
}
