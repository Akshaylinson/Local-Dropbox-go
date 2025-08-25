package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// File model for response
type FileResponse struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Size       int64  `json:"size"`
	UploadedAt string `json:"uploaded_at"`
	Download   string `json:"download"`
}

// 📌 Upload Handler
func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Parse multipart form (max 20 MB file size)
	err := r.ParseMultipartForm(20 << 20)
	if err != nil {
		http.Error(w, "File too big or bad request", http.StatusBadRequest)
		return
	}

	// Get uploaded file
	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Failed to read file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Ensure uploads folder exists
	os.MkdirAll("uploads", os.ModePerm)

	// Generate unique filename with timestamp
	timestamp := time.Now().Unix()
	storedName := fmt.Sprintf("%d_%s", timestamp, handler.Filename)
	filepath := filepath.Join("uploads", storedName)

	dst, err := os.Create(filepath)
	if err != nil {
		http.Error(w, "Unable to save file", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	size, _ := io.Copy(dst, file)

	// Insert into DB
	_, err = db.Exec(`INSERT INTO files (original_name, stored_name, size) VALUES (?, ?, ?)`,
		handler.Filename, storedName, size)
	if err != nil {
		http.Error(w, "DB insert failed", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("✅ File uploaded successfully"))
}

// 📌 List Files Handler
func listFilesHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query(`SELECT id, original_name, size, uploaded_at FROM files ORDER BY uploaded_at DESC`)
	if err != nil {
		http.Error(w, "DB query failed", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var files []FileResponse
	for rows.Next() {
		var id int
		var name string
		var size int64
		var uploadedAt string
		err = rows.Scan(&id, &name, &size, &uploadedAt)
		if err != nil {
			http.Error(w, "Failed to read row", http.StatusInternalServerError)
			return
		}

		files = append(files, FileResponse{
			ID:         id,
			Name:       name,
			Size:       size,
			UploadedAt: uploadedAt,
			Download:   fmt.Sprintf("/download/%d", id),
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(files)
}

// 📌 Download Handler
func downloadHandler(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/download/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid file ID", http.StatusBadRequest)
		return
	}

	var originalName, storedName string
	err = db.QueryRow(`SELECT original_name, stored_name FROM files WHERE id = ?`, id).Scan(&originalName, &storedName)
	if err == sql.ErrNoRows {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "DB query failed", http.StatusInternalServerError)
		return
	}

	filepath := filepath.Join("uploads", storedName)

	// Set appropriate headers for download
	w.Header().Set("Content-Disposition", "attachment; filename="+originalName)
	w.Header().Set("Content-Type", "application/octet-stream")

	http.ServeFile(w, r, filepath)
}
