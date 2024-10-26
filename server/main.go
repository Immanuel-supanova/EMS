package main

import (
	"log"
	"net/http"
)

func main() {

	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("../src"))))

}

// npx tailwindcss -i ./src/main.css -o ./src/style.css --watch
