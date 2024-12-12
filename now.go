package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/beevik/ntp"
)

func Now() {
	timeResult, err := ntp.Time("time.google.com")
	if err != nil {
		log.Fatalf("Ошибка в получении времени: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Текущее время: %v\n", timeResult.Format(time.UnixDate))
}
