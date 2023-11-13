package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load(".env")
	conn := os.Getenv("DB_URL")
	fmt.Println("Conn", conn)
	fmt.Println("text")
	data, err := http.Get("https://eodhd.com/api/real-time/AAPL.US?fmt=json&s=TSLA&api_token=6550af61516825.99143714")

	if err != nil {
		log.Fatal("Fatal: ", err)
	}
	fmt.Println(data.Body)
}
