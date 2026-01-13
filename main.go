package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/getprojectinfo", handlegetprojectinfo)
	http.ListenAndServe(":8280", nil)
}

func handlegetprojectinfo(w http.ResponseWriter, r *http.Request) {
	var info = "a project info"
	fmt.Fprintf(w, info)
}
