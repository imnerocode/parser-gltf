package parser_gltf

import (
	"fmt"
	"io"

	"github.com/imnerocode/vo-structures"
	"github.com/qmuntal/gltf"
)

func ParseGltf(data io.Reader) (*vo.Model, error) {
	var doc gltf.Document

	if err := gltf.NewDecoder(data).Decode(&doc); err != nil {
		return nil, err
	}

	for _, mesh := range doc.Meshes {
		for _, primitive := range mesh.Primitives {
			idIndex := primitive.Indices
			fmt.Printf("Id indices: %v\n", idIndex)
		}
	}

	return nil, nil
}
