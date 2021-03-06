// Copyright 2014 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package utils

import (
	"bufio"
	"fmt"
	"os"

	"github.com/luxengine/lux/glm"
)

// MeshObject is a VUN mesh
type MeshObject struct {
	Vertices []glm.Vec3
	UVs      []glm.Vec2
	Normals  []glm.Vec3
}

// LoadObject loads the wavefront object in the given file.
func LoadObject(fname string, invertV bool) *MeshObject {
	file, err := os.Open(fname)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	vertices, uvs, normals := make([]glm.Vec3, 0), make([]glm.Vec2, 0), make([]glm.Vec3, 0)
	vIndices, uvIndices, nIndices := make([]uint, 0), make([]uint, 0), make([]uint, 0)

	for line, err := reader.ReadString('\n'); err == nil; line, err = reader.ReadString('\n') {
		lineReader := bufio.NewReader(stringReader(line))
		header, _ := lineReader.ReadString(' ')
		restOfLine, _ := lineReader.ReadString('\n')

		switch header[:len(header)-1] {
		case "v":
			vert := glm.Vec3{}
			count, _ := fmt.Sscanf(restOfLine, "%f %f %f\n", &vert.X, &vert.Y, &vert.Z)
			if count != 3 {
				panic("Wrong vert count")
			}
			vertices = append(vertices, vert)

		case "vt":
			uv := glm.Vec2{}
			count, _ := fmt.Sscanf(restOfLine, "%f %f\n", &uv.X, &uv.Y)
			if count != 2 {
				panic("Wrong uv count")
			}
			if invertV {
				// For DDS textures
				uv = glm.Vec2{X: uv.X, Y: 1 - uv.Y}
			}
			uvs = append(uvs, uv)
		case "vn":
			norm := glm.Vec3{}
			count, _ := fmt.Sscanf(restOfLine, "%f %f %f\n", &norm.X, &norm.Y, &norm.Z)
			if count != 3 {
				panic("Wrong norm count")
			}
			normals = append(normals, norm)
		case "f":
			//vert1, vert2, vert3 string
			vIndex, uvIndex, nIndex := [3]uint{}, [3]uint{}, [3]uint{}
			matches, _ := fmt.Sscanf(restOfLine, "%d/%d/%d %d/%d/%d %d/%d/%d\n", &vIndex[0], &uvIndex[0], &nIndex[0], &vIndex[1], &uvIndex[1], &nIndex[1], &vIndex[2], &uvIndex[2], &nIndex[2])
			if matches != 9 {
				panic("Can't read file")
			}
			vIndices = append(vIndices, vIndex[:]...)
			uvIndices = append(uvIndices, uvIndex[:]...)
			nIndices = append(nIndices, nIndex[:]...)
		default:
			// eat line
		}
	}

	//fmt.Println(vertices)

	obj := &MeshObject{make([]glm.Vec3, 0, len(vIndices)), make([]glm.Vec2, 0, len(uvIndices)), make([]glm.Vec3, 0, len(nIndices))}
	for i := range vIndices {
		vIndex, uvIndex, nIndex := vIndices[i], uvIndices[i], nIndices[i]

		vert, uv, norm := vertices[vIndex-1], uvs[uvIndex-1], normals[nIndex-1]

		obj.Vertices = append(obj.Vertices, vert)
		obj.UVs = append(obj.UVs, uv)
		obj.Normals = append(obj.Normals, norm)
	}

	return obj
}

type stringReader string

func (s stringReader) Read(byt []byte) (n int, err error) {
	copy(byt, string(s))
	return len(byt), nil
}
