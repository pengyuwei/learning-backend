package main

import (
  "fmt"
  "log"
)

import (
  "context"
  "go.mongodb.org/mongo-driver/bson"
  "go.mongodb.org/mongo-driver/bson/primitive"
  "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Book struct {
  Id primitive.ObjectID `bson:"_id"`
  Name string
  Key string
}

func do_mongo() (int) {
  // mongodb://foo:bar@localhost:27017
  clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
  }
  collection := client.Database("db").Collection("test")
  
  res, err := collection.InsertOne(context.Background(), bson.M{"Name": "world"})
  if err != nil { return -1 }
  id := res.InsertedID
  fmt.Println("Inserted a single document: ", id)

  var one Book
  err = collection.FindOne(context.Background(), bson.M{"Name": "world"}).Decode(&one)
  if err != nil {
      log.Fatal(err)
  }
  log.Println("collection.FindOne: ", one)
  fmt.Printf("find %v:%v \n", one.Name, one.Key)

  cur, err := collection.Find(context.Background(), bson.D{})
  if err != nil { log.Fatal(err) }
  defer cur.Close(context.Background())
  for cur.Next(context.Background()) {
    // To decode into a struct, use cursor.Decode()
    result := struct {
      Name string
      Speak string
    } {}
    err := cur.Decode(&result)
    if err != nil { log.Fatal(err) }
    // do something with result...
    fmt.Printf("find %v:%v \n", result.Name, result.Speak)
  
    // To get the raw bson bytes use cursor.Current
    // raw := cur.Current
    // do something with raw...
  }

  result := struct{
    Foo string
    Bar int32
  }{}
  filter := bson.D{{"Name", "world"}}
  err = collection.FindOne(context.Background(), filter).Decode(&result)
  if err != nil { return -1 }

  err = client.Disconnect(context.Background())
  if err != nil {
    log.Fatal(err)
  }

  return 0
}

func main() {
  do_mongo()
  return
}