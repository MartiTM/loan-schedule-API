package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/MartiTM/loan-schedule-API/handler"
)

func main() {
	
	http.HandleFunc("/", handler.CalcScheduler)
	
	log.Print("The is Server Running on http://localhost:3000/")
	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}