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

	for !window.ShouldClose() {
		draw(window, glslProgram)
	}
	log.Println("Application end")
}

// initGlfw initializes glfw and returns a window to use
func initGlfw() *glfw.Window {
	if err := glfw.Init(); err != nil {
		panic(err)
	}

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4) // or 2
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(width, height, "GOL", nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()

	return window
}

func initOpenGL() {
	if err := gl.Init(); err != nil {
		panic(err)
	}
	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("OpenGL Version", version)
}

func draw(window *glfw.Window, program GLSLProgram) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	program.Use()
	glfw.PollEvents()
	window.SwapBuffers()
}
