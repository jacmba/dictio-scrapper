package main

import (
	"dictio-scrapper/config"
	"fmt"
)

func main() {
	config.LoadConfig()
	fmt.Println("Hello scrapper!")
}
