package main

import (
	"fmt"
	"os"

	"github.com/immanuel-supanova/EMS/server/settings"
)

func init() {
	settings.Config()
}
func main() {
	fmt.Println(os.Getenv("DB"))
}

// npx tailwindcss -i ./src/main.css -o ./src/style.css --watch
