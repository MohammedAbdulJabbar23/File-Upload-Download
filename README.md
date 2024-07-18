# File Upload and Download with Gin and PostgreSQL

This is a simple web application built with the Gin framework in Go, allowing users to upload files and download them later through links. The application uses PostgreSQL for storing file metadata.

## Features

- **File Upload**: Users can upload files through a web interface.
- **File Download**: Uploaded files can be downloaded through unique links.
- **PostgreSQL Integration**: File metadata is stored in a PostgreSQL database.

## Prerequisites

- [Go](https://golang.org/doc/install) (1.16 or higher)
- [PostgreSQL](https://www.postgresql.org/download/) (any recent version)

## Setup

1. **Clone the repository**:
    ```sh
    git https://github.com/MohammedAbdulJabbar23/File-Upload-Download.git
    cd File-Upload-Download
    ```

2. **Install dependencies**:
    ```sh
    go get -u github.com/gin-gonic/gin
    go get -u github.com/lib/pq
    ```

3. **Configure PostgreSQL**:
    - Create a database and user in PostgreSQL.
    - Update the PostgreSQL connection string in `main.go`:
      ```go
      db, err = sql.Open("postgres", "postgres://username:password@localhost:5432/database_name?sslmode=disable")
      ```

4. **Run the application**:
    ```sh
    go run main.go
    ```

5. **Access the application**:
    - Open your browser and navigate to `http://localhost:8080` for the upload page.

## Project Structure


- `main.go`: The main Go application file.
- `static/index.html`: The HTML file for uploading files.
- `static/styles.css`: The CSS file for styling.
- `uploads`: Directory where uploaded files are stored.

## API Endpoints

- `POST /upload`: Endpoint for uploading files.
- `GET /download/:fileID`: Endpoint for downloading a specific file.

## Example Usage

1. **Upload a File**:
    - Select a file using the file input and click "Upload".
    - After successful upload, an alert will confirm the upload.

2. **Download Files**:
    - Navigate to the download page to view available files.
    - Click on a file link to download it.

