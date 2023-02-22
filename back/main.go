package main

import (
	"flag"
	"fmt"
	"os"
	"support-back/db"
	"support-back/routes"
	"support-back/ws"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var DB = make(map[string]string)
var DBPath string

func main() {
	// Define flags
	var env string

	// Get flag env
	flag.StringVar(&env, "env", "dev", "Default environment: dev")
	flag.Parse()

	fmt.Print(env)

	// Load .env file
	if err := godotenv.Load("./env/" + env); err != nil {
		fmt.Print("No .env file found")
	}

	// Get DB Path
	DBPath, _ = os.LookupEnv("DB_PATH")

	// Db Connect and Close
	db.InitDb(DBPath)
	defer db.CloseDb()

	// Init WS
	ws.InitHub()
	hub := ws.GetHub()
	go hub.RunWS()

	// Init Gin
	r := gin.Default()
	r.SetTrustedProxies(nil)
	routes.InitRouter(r)

	// Run Server
	r.Run(":8081")
}
