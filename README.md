# Distributed File Storage Server

## 1. Requirements

- Implemented in Rust or Golang
- Uses a Relational Database (RDB)
- Utilizes multi-threading
- Includes API documentation
- Dockerized
  - All programs should run with a single "docker compose up (--build)" command<a href="11"/>

## 2. Features

### API 1 - Upload File

- Separates file into multiple files<a href="15"/>
- Uploads separated files to Database in parallel using threads<a href="16"/>
- Returns file ID<a href="17"/>

### API 2 - Get Uploaded Files Data

[Image placeholder]<a href="18"/>

### API 3 - Download File by ID

- Retrieves files in parallel using threads and merges back into one file<a href="21"/>
- Returns the merged original file<a href="22"/>

## Additional Information

It is sufficient to save separated files on the database for this assignment.<a href="26"/>
