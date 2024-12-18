package main

import (
	"fmt"
	"net/http"
	"ru/zakat/server/router"
	"time"
)

type UserId int

type Event struct {
	UserId   UserId
	Title    string
	DateTime time.Time
}

func main() {
	router.InitRouter()

	fmt.Println("Listening on port 8080")
	http.ListenAndServe("127.0.0.1:8080", nil)
}
