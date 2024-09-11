# Distributed File Storage Server

## 1. Requirements

- Implemented in Rust or Golang
- Uses a Relational Database (RDB)
- Utilizes multi-threading
- Includes API documentation
- Dockerized
  - All programs should run with a single "docker compose up (--build)" command

## 2. Features

### API 1 - Upload File

- Separates file into multiple files
- Uploads separated files to Database in parallel using threads
- Returns file ID

### API 2 - Get Uploaded Files Data

### API 3 - Download File by ID

- Retrieves files in parallel using threads and merges back into one file
- Returns the merged original file

## 3. API Documentation

### API 1 - Upload File

- **Endpoint:** `/upload`
- **Method:** POST
- **Request Body:** Multipart form-data with file
- **Response:**
  - Status: 200 OK
  - Body: `{ "file": "unique_file_id" }`

### API 2 - Get Uploaded Files Data

- **Endpoint:** `/files`
- **Method:** GET
- **Response:**
  - Status: 200 OK
  - Body:
    ```json
    [
      {
        "file_id": "unique_file_id",
        "file_name": "example.txt",
        "upload_date": "2023-04-15T10:30:00Z",
        "size": 1024
      }
      // ... more file entries ...
    ]
    ```

### API 3 - Download File by ID

- **Endpoint:** `/download/{file_id}`
- **Method:** GET
- **Parameters:**
  - `file_id`: Unique identifier of the file
- **Response:**
  - Status: 200 OK
  - Body: File content (original merged file)
  - Headers:
    - `Content-Type`: Appropriate MIME type for the file
    - `Content-Disposition`: `attachment; filename="original_filename.ext"`

## Additional Information

It is sufficient to save separated files on the database for this assignment.<a href="26"/>
