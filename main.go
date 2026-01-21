package main

import (
	"log"
	"runtime"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

const windowWidth = 800
const windowHeight = 600
const windowTitle = "Particle Simulator"

var shaderSourceCode string = `#version 330 core
layout (location = 0) in vec3 aPos;
void main()
{
   gl_Position = vec4(aPos.x, aPos.y, aPos.z, 1.0);
}` + "\x00"

var fragmentShaderSourceCode string = `#version 330 core
    	out vec4 FragColor;
    	void main() { FragColor = vec4(1.0, 0.5, 0.2, 1.0); }` + "\x00"

func init() {
	runtime.LockOSThread()
}

func main() {
	//Initialize GLFW
	err := glfw.Init()
	if err != nil {
		log.Fatal("failed to initialize GLFW:", err)
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(windowWidth, windowHeight, windowTitle, nil, nil)
	if err != nil {
		log.Fatal("failed to create window:", err)
	}
	window.MakeContextCurrent()
	err = gl.Init()
	if err != nil {
		log.Fatal("failed to initialize OpenGL:", err)
	}

	vertexShader := gl.CreateShader(gl.VERTEX_SHADER)
	source := gl.Str(shaderSourceCode)
	gl.ShaderSource(vertexShader, 1, &source, nil)
	gl.CompileShader(vertexShader)

	fragmentShader := gl.CreateShader(gl.FRAGMENT_SHADER)
	source = gl.Str(fragmentShaderSourceCode)
	gl.ShaderSource(fragmentShader, 1, &source, nil)
	gl.CompileShader(fragmentShader)

	shaderProgram := gl.CreateProgram()
	gl.AttachShader(shaderProgram, vertexShader)
	gl.AttachShader(shaderProgram, fragmentShader)
	gl.LinkProgram(shaderProgram)
	gl.UseProgram(shaderProgram)

	vertices := []float32{
		-0.5, -0.5, 0.0,
		0.5, -0.5, 0.0,
		0.0, 0.5, 0.0,
	}

	VBO := uint32(0)
	gl.GenBuffers(1, &VBO)
	gl.BindBuffer(gl.ARRAY_BUFFER, VBO)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.STATIC_DRAW)

	vao := uint32(0)
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)

	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)
	gl.EnableVertexAttribArray(0)

	for !window.ShouldClose() {
		processInput(window)

		gl.ClearColor(0.3, 0.0, 0.0, 1.0)
		gl.Clear(gl.COLOR_BUFFER_BIT)

		gl.BindVertexArray(vao)
		gl.DrawArrays(gl.TRIANGLES, 0, 3)

		window.SwapBuffers()
		glfw.PollEvents()
	}

	window.MakeContextCurrent()

}

func processInput(window *glfw.Window) {
	if window.GetKey(glfw.KeyEscape) == glfw.Press {
		window.SetShouldClose(true)
	}
}
