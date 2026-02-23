package invalid

import (
	"fmt"
	"log"
)

func main() {
	log.Println("Bad") // want "logging found"
	fmt.Println("Bad") // skip
}
