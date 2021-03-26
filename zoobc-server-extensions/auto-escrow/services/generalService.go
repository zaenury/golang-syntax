package services

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	mongo "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func UpsertGeneral(client *mongo.Client, key string, height uint32) {
	collection := client.Database("zoosrv_lcl").Collection("General")

	opts := options.Update().SetUpsert(true)
	filter := bson.D{{"key", key}}
	update := bson.D{{"$set",
		bson.M{"LastHeight": height}},
	}

	_, err := collection.UpdateOne(context.TODO(), filter, update, opts)
	if err != nil {
		log.Fatal("blocks error = ", err)
	}
}
