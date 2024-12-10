package main

import (
	"fmt"
	"log"

	"github.com/clrajapaksha/to-do-list-app/cmd/api"
)

func main() {

	fmt.Println("Starting server...")
	server := api.NewAPIServer(":8080", nil)

	if err := server.Run(); err != nil {
		log.Fatal(err.Error())
	}

	// mux := http.NewServeMux()

	// mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Hello World")
	// })

	// if err := http.ListenAndServe("localhost:8080", mux); err != nil {
	// 	// fmt.Println(err.Error())
	// 	log.Println(err.Error())
	// }
}
