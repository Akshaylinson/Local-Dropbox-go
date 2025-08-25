package main

import (
	"log"
	"net/http"
)

func main() {
	// Initialize DB
	err := InitDB("files.db")
	if err != nil {
		log.Fatal("Failed to initialize DB:", err)
	}

	// File server for static assets (UI)
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	// Routes
	http.HandleFunc("/upload", uploadHandler)      // POST: upload file
	http.HandleFunc("/files", listFilesHandler)    // GET: list uploaded files
	http.HandleFunc("/download/", downloadHandler) // GET: download file

	// Start server
	log.Println("🚀 Server running at http://localhost:8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Server failed:", err)
	}
}
