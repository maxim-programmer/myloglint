package invalid

import (
	"fmt"
	"log"
)

func main() {
	log.Println("Bad") // want "log messages must begin with a lowercase letter"
	fmt.Println("Bad") // skip
}
