package main

import (
	"fmt"
	"net/http"
	"os"

	config "private-chat/config"

	handle "private-chat/routes"
	utils "private-chat/utils"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()

	fmt.Println(
		fmt.Sprintf("%s%s%s%s", "Server will start at http://",os.Getenv("HOST"), ":", os.Getenv("PORT")),
	)

	config.ConnectDatabase()

	route := mux.NewRouter()

	handle.AddApproutes(route)

	serverPath := ":" + os.Getenv("PORT")

	cors := utils.GetCorsConfig()

	http.ListenAndServe(serverPath, cors.Handler(route))
}