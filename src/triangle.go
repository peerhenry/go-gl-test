package main

import "github.com/go-gl/gl/v4.1-core/gl"

type drawable interface {
	draw()
}

var (
	triangle = []float32{
		0, 0.5, 0, // top
		-0.5, -0.5, 0, // left
		0.5, -0.5, 0, // right
	}
)

type Triangle struct {
	vao uint32
}

func createTriangle() Triangle {
	vao := makeVao(triangle)
	return Triangle{vao}
}

func (t Triangle) draw() {
	gl.BindVertexArray(t.vao)
	gl.DrawArrays(gl.TRIANGLES, 0, int32(len(triangle)/3))
}
