package handlers

import (
	"RestApi/models"
	database "RestApi/utility"
	"encoding/json"
	"fmt"

	"net/http"
	"os"
)

func PostData(w http.ResponseWriter, r *http.Request) {
	connection := new(database.DBConnection).Database("users").Collection("infos")
	defer connection.Disconnect()
	connection.Connect(os.Getenv("MONGODB"))

	var infos models.Infos

	if err := json.NewDecoder(r.Body).Decode(&infos); err != nil {
		fmt.Println("Error decoding data ", err.Error())
		return
	}

	connection.Add(infos)
}
