package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/getprojectinfo", handlegetProjectInfo)
	http.HandleFunc("/getprojectname", handlegetProjectName)
	http.ListenAndServe(":8280", nil)
}

func handlegetProjectInfo(w http.ResponseWriter, r *http.Request) {
	var info = "a project info"
	fmt.Fprintf(w, info)
}

func handlegetProjectName(w http.ResponseWriter, r *http.Request) {
	var name = "a project name"
	fmt.Fprintf(w, name)
}
