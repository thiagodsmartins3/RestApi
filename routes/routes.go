package routes

import (
	handlers "RestApi/handlers"
	"fmt"
	"net/http"
)

func Routes() {
	http.HandleFunc("/infos", handlers.PostData)

	fmt.Println("Server running on port :8080")
	http.ListenAndServe(":8080", nil)
}
