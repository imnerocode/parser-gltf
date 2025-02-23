# GLTF Parser in Go

This project is a GLTF parser implemented in Go. It reads `.glb` files, extracts JSON and binary data, and deserializes the JSON into structured Go objects.

## Features
- Reads GLB headers and validates the file version.
- Extracts JSON and BIN chunks from GLB files.
- Deserializes JSON data into structured Go structs.

---

## Project Structure
.
├── main.go                # Entry point of the application
├── parser_gltf            # Parser package
│   ├── reader.go          # Reads GLB headers and extracts data
│   └── parser.go          # Deserializes JSON data
└── vo
    └── models.go          # Data structures for GLTF

---

## Requirements
- Go 1.19 or later

---

## Installation

Clone the repository:
    git clone https://github.com/imnerocode/parser-gltf
    cd your-repo-name

---

## Usage

1. Place your `.glb` file in the project directory.
2. Update the `filePath` in `main.go`:
    filePath := "path/to/your/file.glb"
3. Run the application:
    go run main.go

---

## License
This project is licensed under the MIT License.
