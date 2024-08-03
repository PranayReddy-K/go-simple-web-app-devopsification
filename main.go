package main

import (
	"log"
	"net/http"
)

// Handler for serving Issues_Faced.html
func issuesPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/Issues_Faced.html")
}

// Handler for serving Project_Info.html
func projectInfoPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/Project_Info.html")
}

// Handler for serving Workflow_and_Tools.html
func workflowAndToolsPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/Workflow_and_Tools.html")
}

func main() {
	// Map URLs to their respective handler functions
	http.HandleFunc("/issues", issuesPage)
	http.HandleFunc("/project-info", projectInfoPage)
	http.HandleFunc("/workflow-tools", workflowAndToolsPage)

	// Start the server
	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
