package main

import "fmt"

func main() {
	myMap := map[string]string{
		"language": "Golang",
		"creator":  "Google",
	}
	fmt.Println(myMap["language"])
}
