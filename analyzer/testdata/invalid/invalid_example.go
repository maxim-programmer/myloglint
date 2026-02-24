package invalid

import (
	"log"
)

var password = "pass"

func main() {
	log.Println("Starting server on port 8080")  // want "log messages must begin with a lowercase letter"
	log.Println("Failed to connect to database") // want "log messages must begin with a lowercase letter"

	log.Println("запуск сервера")                   // want "log messages must be in english only"
	log.Println("ошибка подключения к базе данных") // want "log messages must be in english only"

	log.Println("server started! 🚀")                // want "log messages must not contain special characters or emojis"
	log.Println("connection failed!!!")             // want "log messages must not contain special characters or emojis"
	log.Println("warning: something went wrong...") // want "log messages must not contain special characters or emojis"

	log.Println("user password: " + password) // want "log messages should not contain potentially sensitive data"
}
