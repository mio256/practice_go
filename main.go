package main

import "fmt"

func main() {
	response_code := map[int]string{
		200: "OK",
		201: "Created",
		400: "Bad Request",
		403: "Forbidden",
		404: "Not Found",
		500: "Internal Server Error",
	}

	fmt.Printf("%v\n", response_code)

	for k, v := range response_code {
		fmt.Printf("%d : %s\n", k, v)
	}
}
