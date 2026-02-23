package valid

import (
	"fmt"
	"log"
)

func main() {
	log.Println("good") // want "logging found"
	fmt.Println("good") // skip
}
