package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/rs/zerolog/log"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello")
		log.Info().Msg("receive hello")
	})

	fmt.Println("start :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "")
		io.WriteString(os.Stderr, err.Error())
		os.Exit(1)
	}
}
