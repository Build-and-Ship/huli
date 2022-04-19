package main

import (
	"fmt"
)

func displayVersion() {
	const VERSION string = "0.0.1"
	fmt.Printf("Huli CLI v%s by Jio Mamaoag(@mamaoag)\n", VERSION)
}

func scanEndpoint(url string) {
	// TODO
	fmt.Printf("Scanning on %s:\n", url)
	fmt.Printf("Feature work in progress.\n")
}

func main() {
	var base_url string

	displayVersion()

	fmt.Printf("To get started. Enter your base api endpoint for scanning >\n")
	fmt.Scan(&base_url)
	fmt.Println("===========================================")
	fmt.Printf("Base URL: %s\n", base_url)

	scanEndpoint(base_url)
}
