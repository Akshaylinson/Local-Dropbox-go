Local Dropbox
A lightweight, self-hosted file sharing application built with Go and SQLite that allows you to upload, manage, and download files through a web interface.

Features
📁 File upload with drag-and-drop support

📊 File listing with metadata (name, size, upload date)

⬇️ One-click downloads

🗃️ SQLite database for efficient storage

📱 Responsive design for mobile and desktop

🔒 Local hosting - your files stay on your machine

Technology Stack
Backend: Go (Golang)

Database: SQLite with modernc.org/sqlite driver

Frontend: Vanilla JavaScript, HTML5, CSS3

File Storage: Local filesystem with unique naming

Installation
Prerequisites
Go 1.21 or later

Modern web browser with JavaScript support

Setup
Clone the repository:

bash
git clone <your-repository-url>
cd local-dropbox
Initialize and download dependencies:

bash
go mod init local-dropbox
go mod tidy
Create necessary directories:

bash
mkdir -p static uploads
Build the application:

bash
go build -o local-dropbox
Usage
Start the server:

bash
./local-dropbox
Or run directly with:

bash
go run main.go db.go handlers.go
Open your browser and navigate to http://localhost:8080

Use the interface to:

Upload files using the file picker

View all uploaded files in the table

Download files by clicking the download link

Project Structure
text
local-dropbox/
├── main.go          # Server entry point and routing
├── db.go            # Database initialization
├── handlers.go      # HTTP request handlers
├── go.mod           # Go module definition
├── go.sum           # Dependency checksums
├── uploads/         # Directory for uploaded files
└── static/          # Frontend assets
    ├── index.html   # Main interface
    ├── style.css    # Styling
    └── app.js       # Frontend functionality
API Endpoints
GET / - Serves the web interface

POST /upload - Handles file uploads

GET /files - Returns JSON list of all files

GET /download/{id} - Downloads a specific file

Database Schema
The application uses a SQLite database with the following schema:

sql
CREATE TABLE files (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    original_name TEXT,
    stored_name TEXT,
    size INTEGER,
    uploaded_at DATETIME DEFAULT CURRENT_TIMESTAMP
);
Configuration
The application uses the following defaults:

Port: 8080

Database file: files.db

Upload directory: uploads/

Maximum file size: 20MB

Development
Adding New Features
Backend changes: Modify the appropriate Go files

Frontend changes: Update files in the static directory

Database changes: Update the schema in db.go

Testing
To test the application:

Start the server

Upload various file types

Verify they appear in the file list

Test download functionality

Check that the database is updated correctly

Troubleshooting
Common Issues
Port already in use: Change the port in main.go

Permission errors: Ensure write access to uploads directory

Database errors: Delete files.db to reset the database

File upload fails: Check available disk space

Logs
The application outputs logs to the console, including:

Server startup information

Database initialization status

File upload events

Error messages

Security Considerations
This application is designed for local use only

Do not expose to public networks without additional security

Files are stored with unique names but original names are preserved in database

No authentication is implemented

Contributing
Fork the repository

Create a feature branch

Make your changes

Test thoroughly

Submit a pull request

License
This project is open source and available under the MIT License.

Future Enhancements
Potential improvements for future versions:

User authentication

File sharing links

File categorization with folders

Search functionality

File previews for images and documents

API rate limiting

Administrative controls
