// Test version, not of significant importance
package main

import (
	"fmt"
	"log"

	"github.com/freedom-sketch/sub2go/internal/api/happ"
)

func main() {
	url := "https://example.com"

	encrypted, err := happ.Encrypt(url)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(encrypted)
}
