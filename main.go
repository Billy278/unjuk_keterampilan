package main

import (
	"os"
	"unjuk_keterampilan/config"
	router "unjuk_keterampilan/routes/v1"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	config.InitDatabase()
	config.InitValidator()
	e := echo.New()
	router.InitRouter(e)
	e.Start(getPort())

}

func init() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
}

func getPort() string {
	if envPort := os.Getenv("PORT"); envPort != "" {
		return ":" + envPort
	}
	return ":8080"
}
