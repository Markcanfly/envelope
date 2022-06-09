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
	tokenData, err := TokenAuth(w, r)
	if err != nil {
		return
	}
	var message models.Message
	_ = json.NewDecoder(r.Body).Decode(&message)
	message.User = tokenData.User.ID
	createMessage(message)
	json.NewEncoder(w).Encode(message)
}

func createMessage(message models.Message) {
	if message.User == primitive.NilObjectID {
		log.Fatal("didn't inject creator user into message, dev error")
	}
	_, err := messageCollection.InsertOne(context.Background(), message)

	if err != nil {
		log.Fatal(err)
	}
}

func DeleteMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	params := mux.Vars(r)
	tokenData, err := TokenAuth(w, r)
	if err != nil {
		return
	}
	deleteOneMessage(params["id"], tokenData.User)
	json.NewEncoder(w).Encode(params[""])
}

func deleteOneMessage(hexid string, user *models.User) {
	id, _ := primitive.ObjectIDFromHex(hexid)
	filter := bson.M{"_id": id}
	var message *models.Message
	err := messageCollection.FindOne(context.Background(), filter).Decode(&message)
	if err != nil {
		log.Fatal(err)
	}
	if message.User != user.ID {
		log.Fatal("user doesn't own message")
	}
	_, err = messageCollection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
}
