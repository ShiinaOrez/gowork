package main

import (
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/bson"

    "context"
    "time"
    "fmt"
)

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://47.107.48.100:27017"))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Connect successful.")

	collection := client.Database("weekdaydb2018_19_2").Collection("all_week2018_19_2")
	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
	res, err := collection.InsertOne(ctx, bson.M{
		"bno": "7",
		"weekNo": "week2",
		"mon": bson.M{
			"1": []string{"7105", "7107"},
			"2": []string{"7105", "7107"},
		},
	})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	id := res.InsertedID
	fmt.Println("Insert successful.", id)

    var result struct {
    	Mon  bson.M {
    		"1",
    	}
    }
	filter := bson.M{"bno": "7", "weekNo": "week2"}
	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
	err = collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Query successful.", result)
	return
}