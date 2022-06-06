package middleware

import (
	"context"
	"encoding/json"
	"envelope/models"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllOpenedMessages(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	payload := getAllOpenedMessages()
	json.NewEncoder(w).Encode(payload)
}

func getAllOpenedMessages() []models.Message {
	cur, err := messageCollection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var results []models.Message
	for cur.Next(context.Background()) {
		var result models.Message
		e := cur.Decode(&result)
		if e != nil {
			log.Fatal(e)
		}
		if result.IsOpened() {
			results = append(results, result)			
		}
	}

	cur.Close(context.Background())
	return results
}

func CreateMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type",  "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var message models.Message
	_ = json.NewDecoder(r.Body).Decode(&message)
	createMessage(message)
	json.NewEncoder(w).Encode(message)
}

func createMessage(message models.Message) {
	_, err := messageCollection.InsertOne(context.Background(), message)

	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println("Inserted a Single Record ", insertResult.InsertID)
}

func DeleteMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	params := mux.Vars(r)
	deleteOneMessage(params["id"])
	json.NewEncoder(w).Encode(params[""])
}

func deleteOneMessage(hexid string) {
	id, _ := primitive.ObjectIDFromHex(hexid)
	filter := bson.M{"_id": id}
	_, err := messageCollection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
}
