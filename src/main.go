package main

import (
	"log"
	"runtime"

	"github.com/go-gl/gl/v4.1-core/gl" // OR: github.com/go-gl/gl/v2.1/gl
	"github.com/go-gl/glfw/v3.2/glfw"
)

const (
	width  = 500
	height = 500
)

func main() {
	log.Println("Application starting...")
	runtime.LockOSThread()
	window := initGlfw()
	defer glfw.Terminate()

	initOpenGL()
	glslProgram := NewGLSLProgram()
	vertex := readFile("shaders/vertex.glsl")
	fragment := readFile("shaders/fragment.glsl")
	glslProgram.CompileAndAttachShader(vertex, gl.VERTEX_SHADER)
	glslProgram.CompileAndAttachShader(fragment, gl.FRAGMENT_SHADER)
	glslProgram.Link()

	triangle := createTriangle()
	objects := []drawable{triangle}

	for !window.ShouldClose() {
		draw(window, glslProgram, objects)
	}
	log.Println("Application end")
}

func initOpenGL() {
	if err := gl.Init(); err != nil {
		panic(err)
	}
	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("OpenGL Version", version)
}

func draw(window *glfw.Window, program GLSLProgram, objects []drawable) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	program.Use()

	for _, object := range objects {
		object.draw()
	}

	glfw.PollEvents()
	window.SwapBuffers()
}
