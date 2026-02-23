package invalid

import (
	"fmt"
	"log"
)

func main() {
	log.Println("Bad") // want "log messages must begin with a lowercase letter"
	fmt.Println("Bad") // skip
	log.Println("english русский") // want "log messages must be in english only"
	log.Println("dot...") // want "log messages must not contain special characters or emojis"
	log.Println("emoji 👍") // want "log messages must not contain special characters or emojis"
}
