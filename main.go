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

var fragmentShaderSourceCodeOrange string = `#version 330 core
    	out vec4 FragColor;
    	void main() { FragColor = vec4(1.0, 0.5, 0.2, 1.0); }` + "\x00"

var fragmentShaderSourceCodeYellow string = `#version 330 core
		out vec4 FragColor;
		void main(){ FragColor = vec4(1.0, 1.0, 0.2, 1.0); }` + "\x00"

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
	source = gl.Str(fragmentShaderSourceCodeOrange)
	gl.ShaderSource(fragmentShader, 1, &source, nil)
	gl.CompileShader(fragmentShader)

	fragmentShaderYellow := gl.CreateShader(gl.FRAGMENT_SHADER)
	source = gl.Str(fragmentShaderSourceCodeYellow)
	gl.ShaderSource(fragmentShaderYellow, 1, &source, nil)
	gl.CompileShader(fragmentShaderYellow)

	shaderProgram := gl.CreateProgram()
	gl.AttachShader(shaderProgram, vertexShader)
	gl.AttachShader(shaderProgram, fragmentShader)
	gl.LinkProgram(shaderProgram)
	gl.UseProgram(shaderProgram)

	shaderProgramYellow := gl.CreateProgram()
	gl.AttachShader(shaderProgramYellow, vertexShader)
	gl.AttachShader(shaderProgramYellow, fragmentShaderYellow)
	gl.LinkProgram(shaderProgramYellow)

	triangle1 := []float32{
		//Triangle 1
		-0.8, -0.8, 0.0,
		0.0, -0.8, 0.0,
		-0.4, 0.0, 0.0,
	}

	triangle2 := []float32{
		//Triangle 2
		0.8, -0.8, 0.0,
		0.0, -0.8, 0.0,
		0.4, 0.0, 0.0,
	}

	triangle3 := []float32{
		//Triangle 3
		-0.4, 0.0, 0.0,
		0.4, 0.0, 0.0,
		0.0, 0.8, 0.0,
	}

	VBOTriangle1 := uint32(0)
	VBOTriangle2 := uint32(0)
	VBOTriangle3 := uint32(0)
	vaoTriangle1 := uint32(0)
	vaoTriangle2 := uint32(0)
	vaoTriangle3 := uint32(0)
	// Triangle 1 setup
	gl.GenBuffers(1, &VBOTriangle1)
	gl.BindBuffer(gl.ARRAY_BUFFER, VBOTriangle1)
	gl.BufferData(gl.ARRAY_BUFFER, len(triangle1)*4, gl.Ptr(triangle1), gl.STATIC_DRAW)

	gl.GenVertexArrays(1, &vaoTriangle1)
	gl.BindVertexArray(vaoTriangle1)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 3*4, nil)
	gl.EnableVertexAttribArray(0)

	// Triangle 2 setup
	gl.GenBuffers(1, &VBOTriangle2)
	gl.BindBuffer(gl.ARRAY_BUFFER, VBOTriangle2)
	gl.BufferData(gl.ARRAY_BUFFER, len(triangle2)*4, gl.Ptr(triangle2), gl.STATIC_DRAW)

	gl.GenVertexArrays(1, &vaoTriangle2)
	gl.BindVertexArray(vaoTriangle2)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 3*4, nil)
	gl.EnableVertexAttribArray(0)

	// Triangle 3 setup
	gl.GenBuffers(1, &VBOTriangle3)
	gl.BindBuffer(gl.ARRAY_BUFFER, VBOTriangle3)
	gl.BufferData(gl.ARRAY_BUFFER, len(triangle3)*4, gl.Ptr(triangle3), gl.STATIC_DRAW)

	gl.GenVertexArrays(1, &vaoTriangle3)
	gl.BindVertexArray(vaoTriangle3)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 3*4, nil)
	gl.EnableVertexAttribArray(0)

	for !window.ShouldClose() {
		processInput(window)

		gl.ClearColor(0.3, 0.0, 0.0, 1.0)
		gl.Clear(gl.COLOR_BUFFER_BIT)

		// Draw triangle 1 with orange shader
		gl.UseProgram(shaderProgram)
		gl.BindVertexArray(vaoTriangle1)
		gl.DrawArrays(gl.TRIANGLES, 0, 3)

		// Draw triangle 2 with yellow shader
		gl.UseProgram(shaderProgramYellow)
		gl.BindVertexArray(vaoTriangle2)
		gl.DrawArrays(gl.TRIANGLES, 0, 3)

		// Draw triangle 3 with whichever color you want
		gl.UseProgram(shaderProgram) // or shaderProgramYellow
		gl.BindVertexArray(vaoTriangle3)
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
