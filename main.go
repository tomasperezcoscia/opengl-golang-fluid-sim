package main

import (
	"fmt"
	"log"
	"runtime"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

const (
	windowWidth  = 1920
	windowHeight = 1080
	windowTitle  = "Particle Simulator"
)

func init() {
	// OpenGL requires calls to be made from the main OS thread
	runtime.LockOSThread()
}

func main() {
	// Initialize GLFW
	if err := glfw.Init(); err != nil {
		log.Fatal("failed to initialize GLFW:", err)
	}
	defer glfw.Terminate()

	// Configure OpenGL version and profile
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	glfw.WindowHint(glfw.Resizable, glfw.True)

	// Create window
	window, err := glfw.CreateWindow(windowWidth, windowHeight, windowTitle, nil, nil)
	if err != nil {
		log.Fatal("failed to create window:", err)
	}
	window.MakeContextCurrent()

	// Initialize OpenGL bindings
	if err := gl.Init(); err != nil {
		log.Fatal("failed to initialize OpenGL:", err)
	}

	// Print OpenGL version for verification
	version := gl.GoStr(gl.GetString(gl.VERSION))
	fmt.Println("OpenGL version:", version)

	// Configure OpenGL
	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)
	gl.ClearColor(0.1, 0.1, 0.15, 1.0) // Dark blue background

	// Main render loop
	for !window.ShouldClose() {
		// Clear the screen
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		// Swap buffers and poll events
		window.SwapBuffers()
		glfw.PollEvents()
	}
}
