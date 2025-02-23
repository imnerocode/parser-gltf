package vo

import "os"

type GLTFData struct {
	File       *os.File
	JSONOffset int64
	JSONSize   int64
	BINOffset  int64
	BINSize    int64
}

type Accessor struct {
	BufferView    int       `json:"bufferView"`
	ComponentType int       `json:"componentType"`
	Count         int       `json:"count"`
	Type          string    `json:"type"`
	Max           []float32 `json:"max,omitempty"`
	Min           []float32 `json:"min,omitempty"`
}

type GLTF struct {
	Asset       Asset        `json:"asset"`
	Buffers     []Buffer     `json:"buffers"`
	BufferViews []BufferView `json:"bufferViews"`
	Meshes      []Mesh       `json:"meshes"`
	Nodes       []Node       `json:"nodes"`
	Scenes      []Scene      `json:"scenes"`
	Accessor    []Accessor   `json:"accessors"`
}

type Primitive struct {
	Attributes map[string]int `json:"attributes"`
	Indices    int            `json:"indices,omitempty"`
	Material   int            `json:"material,omitempty"`
}

type Asset struct {
	Generator string `json:"generator"`
	Version   string `json:"version"`
}

type Buffer struct {
	ByteLength int `json:"byteLength"`
}

type BufferView struct {
	Buffer     int `json:"buffer"`
	ByteLength int `json:"byteLength"`
	ByteOffset int `json:"byteOffset"`
}

type Mesh struct {
	Name       string      `json:"name"`
	Primitives []Primitive `json:"primitives"`
}

type Node struct {
	Mesh int    `json:"mesh,omitempty"`
	Name string `json:"name"`
}

type Scene struct {
	Name  string `json:"name"`
	Nodes []int  `json:"nodes"`
}
