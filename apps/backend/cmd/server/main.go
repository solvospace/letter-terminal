package main

import (
	"log"
	"backend/internal/app"
)

func main() {
	application:= app.New();
	appError := application.Run();

	if  appError != nil {
		log.Fatal(appError)
	}
}
