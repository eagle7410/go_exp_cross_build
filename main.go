package main

import (
	"fmt"
	"go_exp_cross_build/lib"
	"log"
	"net/http"
)

const port = "6060"

func init() {
	lib.OpenLogFile()
}

func index(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "HElloy")
}

func main() {

	http.HandleFunc("/", index)

	lib.LogAppRun(port)
	log.Fatal(http.ListenAndServe(":"+port, lib.LogRequest(http.DefaultServeMux)))
}
