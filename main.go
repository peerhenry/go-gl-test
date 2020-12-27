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
	glProgram := initOpenGL()
	for !window.ShouldClose() {
		draw(window, glProgram)
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

func initOpenGL() uint32 {
	if err := gl.Init(); err != nil {
		panic(err)
	}
	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("OpenGL Version", version)

	glProgram := gl.CreateProgram()
	gl.LinkProgram(glProgram)
	return glProgram
}

func draw(window *glfw.Window, glProgram uint32) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.UseProgram(glProgram)
	glfw.PollEvents()
	window.SwapBuffers()
}
