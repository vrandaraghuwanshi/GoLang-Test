package main

import (
	"fmt"
	"time"
)

func main() {

	// years := "26-01-2020".Sub("21-01-2020").Hours() / 24 / 365
	start, _ := time.Parse(time.RFC3339, "2020-01-21T15:04:05Z")

	fmt.Println(t)

	start, _ := time.Parse(time.RFC3339, "2020-01-02T15:04:05Z")
	fmt.Println(t)


}
