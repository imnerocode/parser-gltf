package parser_gltf

import (
	"encoding/binary"
	"encoding/json"
	"errors"

	"github.com/imnerocode/parser-gltf/vo"
)

const (
	GLBHeaderSize   = 12
	ChunkHeaderSize = 8
)

func ReadGLBHeaders(gltfData *vo.GLTFData) error {
	// Read the GLB header
	header := make([]byte, GLBHeaderSize)
	_, err := gltfData.File.Read(header)
	if err != nil {
		return err
	}

	// Validate GLB header ('glTF' in ASCII)
	if string(header[0:4]) != "glTF" && string(header[0:4]) != "glTF\x00" {
		return errors.New("invalid file: not a GLB")
	}

	// Read version and file size
	version := binary.LittleEndian.Uint32(header[4:8])
	if version != 2 {
		return errors.New("unsupported version: only GLB version 2 is supported")
	}

	// Traverse the chunks
	currentOffset := int64(GLBHeaderSize)
	for {
		// Read the chunk header
		chunkHeader := make([]byte, ChunkHeaderSize)
		_, err := gltfData.File.ReadAt(chunkHeader, currentOffset)
		if err != nil {
			break
		}

		// Get chunk size and type
		chunkSize := binary.LittleEndian.Uint32(chunkHeader[0:4])
		chunkType := string(chunkHeader[4:8])

		// JSON Chunk
		if chunkType == "JSON" {
			gltfData.JSONOffset = currentOffset + ChunkHeaderSize
			gltfData.JSONSize = int64(chunkSize)
		}

		// BIN Chunk
		if chunkType == "BIN\x00" {
			gltfData.BINOffset = currentOffset + ChunkHeaderSize
			gltfData.BINSize = int64(chunkSize)
		}

		// Move to the next chunk
		currentOffset += ChunkHeaderSize + int64(chunkSize)
	}

	return nil
}

func ExtractJSONData(gltfData *vo.GLTFData) ([]byte, error) {
	// Check if JSON is found
	if gltfData.JSONSize == 0 {
		return nil, errors.New("no JSON found in the GLB file")
	}

	// Read the JSON
	jsonData := make([]byte, gltfData.JSONSize)
	_, err := gltfData.File.ReadAt(jsonData, gltfData.JSONOffset)
	if err != nil {
		return nil, err
	}

	return jsonData, nil
}

func ExtractBINData(gltfData *vo.GLTFData) ([]byte, error) {
	// Check if BIN is found
	if gltfData.BINSize == 0 {
		return nil, errors.New("no BIN found in the GLB file")
	}

	// Read the BIN
	binData := make([]byte, gltfData.BINSize)
	_, err := gltfData.File.ReadAt(binData, gltfData.BINOffset)
	if err != nil {
		return nil, err
	}

	return binData, nil
}

func ParseGLTFJSON(jsonData []byte) (*vo.GLTF, error) {
	var gltf vo.GLTF
	err := json.Unmarshal(jsonData, &gltf)
	if err != nil {
		return nil, errors.New("failed to parse GLTF JSON")
	}
	return &gltf, nil
}
