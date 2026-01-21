# Particle Simulator - Learning Roadmap & TODO

**Goal:** Build OpenGL foundation before implementing particle system  
**Timeline:** ~1 week of focused study  
**Status:** 🟡 In Progress

---

## Progress Tracker

- [ ] Environment setup complete (GLFW, OpenGL bindings)
- [ ] Basic window creation working
- [ ] Completed LearnOpenGL fundamentals
- [ ] Built camera system
- [ ] Understanding of SPH physics
- [ ] Ready to start particle implementation

---

## Phase 1: OpenGL Fundamentals (Days 1-3)

### Day 1: Triangles & Shaders

#### ✅ Task 1.1: Hello Window
**Status:** ✅ COMPLETE  
**What:** Open a window with OpenGL context  
**Resources:**
- Tutorial: https://learnopengl.com/Getting-started/Hello-Window
- Already implemented in `main.go`

---

#### 📝 Task 1.2: Hello Triangle
**Status:** ⬜ TODO  
**What:** Draw your first triangle using VBOs and shaders  
**Goal:** Understand the complete render pipeline

**Key Concepts to Learn:**
- Vertex Buffer Object (VBO) - stores vertex data on GPU
- Vertex Array Object (VAO) - describes how to read VBO data
- Vertex shader - transforms vertex positions
- Fragment shader - colors pixels
- Shader compilation and linking

**Resources:**
- Tutorial: https://learnopengl.com/Getting-started/Hello-Triangle
- Go OpenGL docs: https://pkg.go.dev/github.com/go-gl/gl/v4.1-core/gl

**C++ to Go Translation Guide:**

**Creating a VBO (C++):**
```cpp
unsigned int VBO;
glGenBuffers(1, &VBO);
glBindBuffer(GL_ARRAY_BUFFER, VBO);
glBufferData(GL_ARRAY_BUFFER, sizeof(vertices), vertices, GL_STATIC_DRAW);
```

**Creating a VBO (Go):**
```go
var vbo uint32
gl.GenBuffers(1, &vbo)
gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.STATIC_DRAW)
// Note: multiply by 4 because float32 is 4 bytes
```

**Creating a VAO (C++):**
```cpp
unsigned int VAO;
glGenVertexArrays(1, &VAO);
glBindVertexArray(VAO);
glVertexAttribPointer(0, 3, GL_FLOAT, GL_FALSE, 3*sizeof(float), (void*)0);
glEnableVertexAttribArray(0);
```

**Creating a VAO (Go):**
```go
var vao uint32
gl.GenVertexArrays(1, &vao)
gl.BindVertexArray(vao)
gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 3*4, nil)
gl.EnableVertexAttribArray(0)
// Note: stride is 3*4 (3 floats * 4 bytes each)
```

**Example Vertex Data (Works in both):**
```go
vertices := []float32{
    -0.5, -0.5, 0.0,  // bottom left
     0.5, -0.5, 0.0,  // bottom right
     0.0,  0.5, 0.0,  // top
}
```

**Your Tasks:**
- [ ] Create a VBO with triangle vertices
- [ ] Create a VAO to describe vertex layout
- [ ] Write a simple vertex shader (pass-through)
- [ ] Write a simple fragment shader (solid color)
- [ ] Compile and link shaders
- [ ] Draw the triangle with `gl.DrawArrays(gl.TRIANGLES, 0, 3)`
- [ ] Complete the exercises at the end of the tutorial

**Success Criteria:** Orange triangle on screen

**Notes/Questions:**
```
[Write your questions or observations here as you work]
```

---

#### 📝 Task 1.3: Shaders Deep Dive
**Status:** ⬜ TODO  
**What:** Learn GLSL language and uniform variables  
**Goal:** Understand how to pass data from CPU to shaders

**Key Concepts:**
- GLSL syntax (similar to C)
- Vertex shader inputs/outputs
- Fragment shader inputs/outputs
- Uniform variables (global constants)
- Data types: `vec2`, `vec3`, `vec4`, `mat4`

**Resources:**
- Tutorial: https://learnopengl.com/Getting-started/Shaders
- GLSL Quick Reference: https://www.khronos.org/files/opengl45-quick-reference-card.pdf

**C++ to Go Translation:**

**Setting a uniform (C++):**
```cpp
int vertexColorLocation = glGetUniformLocation(shaderProgram, "ourColor");
glUseProgram(shaderProgram);
glUniform4f(vertexColorLocation, 0.0f, 1.0f, 0.0f, 1.0f);
```

**Setting a uniform (Go):**
```go
vertexColorLocation := gl.GetUniformLocation(shaderProgram, gl.Str("ourColor\x00"))
gl.UseProgram(shaderProgram)
gl.Uniform4f(vertexColorLocation, 0.0, 1.0, 0.0, 1.0)
// Note: strings must be null-terminated with \x00
```

**Example Shaders:**

**Vertex Shader (GLSL):**
```glsl
#version 410 core
layout (location = 0) in vec3 aPos;
out vec4 vertexColor;

void main() {
    gl_Position = vec4(aPos, 1.0);
    vertexColor = vec4(0.5, 0.0, 0.0, 1.0);
}
```

**Fragment Shader (GLSL):**
```glsl
#version 410 core
out vec4 FragColor;
in vec4 vertexColor;
uniform vec4 ourColor;

void main() {
    FragColor = vertexColor * ourColor;
}
```

**Your Tasks:**
- [ ] Pass vertex colors from vertex shader to fragment shader
- [ ] Use uniforms to change triangle color from CPU
- [ ] Make triangle color change over time (use `glfw.GetTime()`)
- [ ] Complete the exercises

**Success Criteria:** Triangle that changes color over time

**Notes/Questions:**
```
[Your notes here]
```

---

### Day 2: Transformations & Math

#### 📝 Task 2.1: Textures (Optional/Skim)
**Status:** ⬜ TODO (Can skip or skim)  
**What:** Loading and applying textures  
**Resources:**
- Tutorial: https://learnopengl.com/Getting-started/Textures

**Note:** We won't use textures for particles, but understanding texture coordinates helps with understanding fragment shaders. You can skim this section.

---

#### 📝 Task 2.2: Transformations ⭐ CRITICAL
**Status:** ⬜ TODO  
**What:** Learn matrix mathematics for 3D graphics  
**Goal:** Understand how to move, rotate, and scale objects

**Key Concepts:**
- Translation (moving objects)
- Rotation (spinning objects)
- Scaling (resizing objects)
- Matrix multiplication (combining transformations)
- Homogeneous coordinates (4th dimension for translation)

**Resources:**
- Tutorial: https://learnopengl.com/Getting-started/Transformations
- go-gl/mathgl docs: https://pkg.go.dev/github.com/go-gl/mathgl/mgl32
- Matrix visualization: http://www.opengl-tutorial.org/beginners-tutorials/tutorial-3-matrices/

**C++ to Go Translation:**

**GLM (C++):**
```cpp
#include <glm/glm.hpp>
#include <glm/gtc/matrix_transform.hpp>

glm::mat4 trans = glm::mat4(1.0f);  // identity matrix
trans = glm::translate(trans, glm::vec3(1.0f, 1.0f, 0.0f));
trans = glm::rotate(trans, glm::radians(90.0f), glm::vec3(0.0, 0.0, 1.0));
trans = glm::scale(trans, glm::vec3(0.5, 0.5, 0.5));
```

**mathgl (Go):**
```go
import "github.com/go-gl/mathgl/mgl32"

trans := mgl32.Ident4()  // identity matrix
trans = trans.Mul4(mgl32.Translate3D(1.0, 1.0, 0.0))
trans = trans.Mul4(mgl32.HomogRotate3D(mgl32.DegToRad(90.0), mgl32.Vec3{0, 0, 1}))
trans = trans.Mul4(mgl32.Scale3D(0.5, 0.5, 0.5))
```

**Sending matrix to shader (C++):**
```cpp
unsigned int transformLoc = glGetUniformLocation(shaderProgram, "transform");
glUniformMatrix4fv(transformLoc, 1, GL_FALSE, glm::value_ptr(trans));
```

**Sending matrix to shader (Go):**
```go
transformLoc := gl.GetUniformLocation(shaderProgram, gl.Str("transform\x00"))
gl.UniformMatrix4fv(transformLoc, 1, false, &trans[0])
```

**Vertex shader with transformation:**
```glsl
#version 410 core
layout (location = 0) in vec3 aPos;
uniform mat4 transform;

void main() {
    gl_Position = transform * vec4(aPos, 1.0);
}
```

**Your Tasks:**
- [ ] Make triangle rotate continuously
- [ ] Make triangle move in a circle
- [ ] Make triangle scale up and down
- [ ] Combine all three transformations
- [ ] Complete ALL exercises (very important!)

**Success Criteria:** Spinning, moving, pulsating triangle

**Notes/Questions:**
```
[Your notes here]
```

---

### Day 3: 3D Space & Camera

#### 📝 Task 3.1: Coordinate Systems ⭐ CRITICAL
**Status:** ⬜ TODO  
**What:** Understand the journey from 3D world to 2D screen  
**Goal:** Master the transformation pipeline

**Key Concepts:**
- Local space (object coordinates)
- World space (scene coordinates)
- View space (camera coordinates)
- Clip space (projection applied)
- Screen space (final pixels)
- Model, View, Projection matrices (MVP)
- Perspective projection

**Resources:**
- Tutorial: https://learnopengl.com/Getting-started/Coordinate-Systems
- Interactive demo: https://jsantell.com/model-view-projection/

**The Complete Pipeline:**
```
Local Space (your model)
    ↓ Model Matrix (position in world)
World Space (scene)
    ↓ View Matrix (relative to camera)
View Space (from camera perspective)
    ↓ Projection Matrix (3D → 2D with perspective)
Clip Space
    ↓ (automatic by GPU)
Screen Space (pixels)
```

**C++ to Go Translation:**

**Creating perspective projection (C++):**
```cpp
glm::mat4 projection = glm::perspective(
    glm::radians(45.0f),  // FOV
    800.0f / 600.0f,      // aspect ratio
    0.1f,                 // near plane
    100.0f                // far plane
);
```

**Creating perspective projection (Go):**
```go
projection := mgl32.Perspective(
    mgl32.DegToRad(45.0), // FOV
    800.0 / 600.0,        // aspect ratio
    0.1,                  // near plane
    100.0,                // far plane
)
```

**Complete vertex shader with MVP:**
```glsl
#version 410 core
layout (location = 0) in vec3 aPos;

uniform mat4 model;
uniform mat4 view;
uniform mat4 projection;

void main() {
    gl_Position = projection * view * model * vec4(aPos, 1.0);
}
```

**Your Tasks:**
- [ ] Draw a 3D cube (not triangle)
- [ ] Apply model matrix (rotate the cube)
- [ ] Apply view matrix (move camera back)
- [ ] Apply projection matrix (perspective)
- [ ] Enable depth testing: `gl.Enable(gl.DEPTH_TEST)`
- [ ] Clear depth buffer: `gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)`
- [ ] Complete the exercises

**Success Criteria:** Rotating 3D cube with correct depth sorting

**Notes/Questions:**
```
[Your notes here]
```

---

#### 📝 Task 3.2: Camera ⭐ CRITICAL
**Status:** ⬜ TODO  
**What:** Implement a movable camera system  
**Goal:** Understand view matrix construction (needed for orbit camera)

**Key Concepts:**
- Camera position, target, up vector
- LookAt matrix
- Camera movement (WASD)
- Mouse look (FPS style)
- Camera coordinate system (right, up, forward vectors)

**Resources:**
- Tutorial: https://learnopengl.com/Getting-started/Camera
- Camera math explained: http://www.opengl-tutorial.org/beginners-tutorials/tutorial-6-keyboard-and-mouse/

**C++ to Go Translation:**

**Creating view matrix (C++):**
```cpp
glm::mat4 view = glm::lookAt(
    glm::vec3(0.0f, 0.0f, 3.0f),  // camera position
    glm::vec3(0.0f, 0.0f, 0.0f),  // look at point
    glm::vec3(0.0f, 1.0f, 0.0f)   // up vector
);
```

**Creating view matrix (Go):**
```go
view := mgl32.LookAtV(
    mgl32.Vec3{0, 0, 3},  // camera position
    mgl32.Vec3{0, 0, 0},  // look at point  
    mgl32.Vec3{0, 1, 0},  // up vector
)
```

**Keyboard input (Go):**
```go
// In your main loop
if window.GetKey(glfw.KeyW) == glfw.Press {
    cameraPos = cameraPos.Add(cameraFront.Mul(cameraSpeed * deltaTime))
}
if window.GetKey(glfw.KeyS) == glfw.Press {
    cameraPos = cameraPos.Sub(cameraFront.Mul(cameraSpeed * deltaTime))
}
// etc...
```

**Mouse input callback (Go):**
```go
var lastX, lastY float64 = 400, 300
var firstMouse bool = true

window.SetCursorPosCallback(func(w *glfw.Window, xpos, ypos float64) {
    if firstMouse {
        lastX = xpos
        lastY = ypos
        firstMouse = false
    }
    
    xoffset := xpos - lastX
    yoffset := lastY - ypos  // reversed: y ranges bottom to top
    lastX = xpos
    lastY = ypos
    
    sensitivity := 0.1
    xoffset *= sensitivity
    yoffset *= sensitivity
    
    yaw += float32(xoffset)
    pitch += float32(yoffset)
    
    // Clamp pitch
    if pitch > 89.0 {
        pitch = 89.0
    }
    if pitch < -89.0 {
        pitch = -89.0
    }
    
    // Calculate new front vector
    // ... (from tutorial)
})
```

**Your Tasks:**
- [ ] Implement WASD camera movement
- [ ] Implement mouse look (FPS camera)
- [ ] Add camera speed control
- [ ] Keep camera from flipping (clamp pitch)
- [ ] Complete the exercises

**Success Criteria:** You can fly around your 3D scene

**Notes/Questions:**
```
[Your notes here]
```

---

## Phase 2: Video Reinforcement (Days 4-5)

### 📺 The Cherno's OpenGL Series

**Watch these episodes (take notes):**

#### Episode 3: Using Modern OpenGL
- [ ] Watch: https://www.youtube.com/watch?v=H2E3yO0J7TM&list=PLlrATfBNZ98foTJPJ_Ev03o2oq3-GGOS2&index=3
- [ ] Key takeaway: _____________________

#### Episode 5: Vertex Buffers and Drawing
- [ ] Watch: https://www.youtube.com/watch?v=0p9VxImr7Y0&list=PLlrATfBNZ98foTJPJ_Ev03o2oq3-GGOS2&index=5
- [ ] Key takeaway: _____________________

#### Episode 7: Writing a Shader
- [ ] Watch: https://www.youtube.com/watch?v=5W7JLgFCkwI&list=PLlrATfBNZ98foTJPJ_Ev03o2oq3-GGOS2&index=7
- [ ] Key takeaway: _____________________

#### Episode 13: Maths in OpenGL
- [ ] Watch: https://www.youtube.com/watch?v=adsDbg_6E8o&list=PLlrATfBNZ98foTJPJ_Ev03o2oq3-GGOS2&index=13
- [ ] Key takeaway: _____________________

---

### 🎯 Side Quest: Point Sprite Renderer

**Goal:** Adapt your triangle code to render multiple points  
**Why:** This is 80% of particle rendering!

**Tasks:**
- [ ] Change `gl.DrawArrays(gl.TRIANGLES, ...)` to `gl.DrawArrays(gl.POINTS, ...)`
- [ ] Create array of 100 random 3D positions
- [ ] Set `gl_PointSize` in vertex shader
- [ ] Make points round using `gl_PointCoord` in fragment shader
- [ ] Color points based on their position

**Example point rounding in fragment shader:**
```glsl
void main() {
    // Make point circular
    vec2 coord = gl_PointCoord - vec2(0.5);
    if (length(coord) > 0.5) {
        discard;  // Don't draw this pixel
    }
    fragColor = vec4(1.0, 0.5, 0.2, 1.0);
}
```

**Success Criteria:** 100 round, colored points floating in 3D space

---

## Phase 3: Physics Fundamentals (Days 6-7)

### 📚 Spatial Data Structures

#### Task 4.1: Understand Spatial Hashing
**Status:** ⬜ TODO  
**What:** Learn how to avoid O(n²) neighbor searches  
**Goal:** Understand grid-based optimization for particle interactions

**Resources:**
- Article: https://www.gamedev.net/tutorials/programming/general-and-gameplay-programming/spatial-hashing-r2697/
- Video explanation: https://www.youtube.com/watch?v=sx4IIQL0x7c

**Key Concepts:**
- Why checking every pair is slow (5000 particles = 12.5 million checks!)
- Grid cells for spatial partitioning
- Hash function to map position → cell
- Only check neighbors in nearby cells

**Your Understanding Check:**
```
Why is O(n²) bad for 5000 particles?
Answer: _____________________

How does a grid help?
Answer: _____________________

What's a typical cell size?
Answer: _____________________
```

---

### 🌊 Fluid Dynamics (SPH)

#### Task 4.2: SPH Conceptual Understanding
**Status:** ⬜ TODO  
**What:** Learn Smoothed Particle Hydrodynamics theory  
**Goal:** Understand the physics before coding

**Resources:**
- Paper (sections 1-3 only): https://matthias-research.github.io/pages/publications/sca03.pdf
- Video: https://www.youtube.com/watch?v=rSKMYc1CQHE (Sebastian Lague - Coding Adventure)
- Alternative: https://www.dgp.toronto.edu/public_user/stam/reality/Research/pdf/GDC03.pdf

**Key Concepts to Understand:**
- Particles represent fluid mass
- Kernel functions (smoothing functions)
- Density calculation
- Pressure forces
- Viscosity forces
- Surface tension

**Write in your own words:**
```
What is a kernel function?
Answer: _____________________

Why do we need density calculation?
Answer: _____________________

What creates pressure forces?
Answer: _____________________
```

**Kernel Functions (Reference):**

**Poly6 Kernel (for density):**
```
W(r, h) = (315 / (64π h⁹)) * (h² - r²)³   if 0 ≤ r ≤ h
        = 0                                if r > h

where:
- r = distance between particles
- h = smoothing radius (influence distance)
```

**Spiky Kernel (for pressure):**
```
W(r, h) = (15 / (π h⁶)) * (h - r)³   if 0 ≤ r ≤ h
        = 0                           if r > h
```

**Don't worry about implementing these yet!** Just understand conceptually what they do.

---

## Phase 4: Self-Assessment

### ✅ Before Moving to Implementation

Can you answer these questions confidently?

#### OpenGL Fundamentals:
- [ ] What is a VBO and what data does it store?
- [ ] What is a VAO and why do we need it?
- [ ] What does a vertex shader do?
- [ ] What does a fragment shader do?
- [ ] How do we pass data from CPU to shaders? (uniforms)

#### 3D Mathematics:
- [ ] What are the three main transformation matrices? (Model, View, Projection)
- [ ] What does the view matrix represent?
- [ ] What does the projection matrix do?
- [ ] How do you create a view matrix from camera position and target?
- [ ] What is the difference between orthographic and perspective projection?

#### Camera:
- [ ] Can you implement a basic orbit camera from scratch?
- [ ] Do you understand the LookAt function?
- [ ] Can you convert mouse input to camera rotation?

#### Physics:
- [ ] Why is spatial hashing important for particle systems?
- [ ] What is SPH and what problem does it solve?
- [ ] What are the main forces in fluid simulation? (pressure, viscosity, surface tension)

#### Go + OpenGL:
- [ ] How do you create and bind a VBO in Go?
- [ ] How do you send a matrix to a shader in Go?
- [ ] How do you handle null-terminated strings in Go OpenGL? (`\x00`)
- [ ] What does `gl.Ptr()` do?

---

## 📝 Your Implementation Checklist

Once you've completed the learning phase:

- [ ] I have working triangle code
- [ ] I have working 3D cube with camera
- [ ] I have working point sprite renderer
- [ ] I understand all key concepts above
- [ ] I've pushed my learning examples to a repo
- [ ] I'm ready to start particle system implementation

**My Learning Repo:** _____________________  
**Date Completed:** _____________________

---

## 🆘 Getting Stuck? Common Issues

### OpenGL Context Issues
```go
// Make sure you're calling from main thread
runtime.LockOSThread()

// Make sure context is current before gl.Init()
window.MakeContextCurrent()
if err := gl.Init(); err != nil { ... }
```

### Shader Compilation Errors
```go
// Always check compilation status
var status int32
gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
if status == gl.FALSE {
    // Get and print error log
}
```

### Matrix Math Not Working
```go
// Remember: matrices multiply right-to-left
projection * view * model  // Correct order!

// In Go with mathgl:
mvp := projection.Mul4(view).Mul4(model)
```

### Nothing Rendering
- Check `gl.GetError()` after each GL call
- Verify shader compilation succeeded
- Ensure VAO is bound before drawing
- Check viewport is set correctly: `gl.Viewport(0, 0, width, height)`
- Enable depth test if 3D: `gl.Enable(gl.DEPTH_TEST)`

---

## 📌 Quick Reference Links

### Documentation
- **go-gl OpenGL**: https://pkg.go.dev/github.com/go-gl/gl/v4.1-core/gl
- **go-gl GLFW**: https://pkg.go.dev/github.com/go-gl/glfw/v3.3/glfw
- **go-gl mathgl**: https://pkg.go.dev/github.com/go-gl/mathgl/mgl32
- **OpenGL Reference**: https://www.khronos.org/registry/OpenGL-Refpages/gl4/
- **GLSL Reference**: https://www.khronos.org/opengl/wiki/OpenGL_Shading_Language

### Communities
- **LearnOpenGL Discord**: https://discord.gg/learnopengl
- **r/opengl**: https://reddit.com/r/opengl
- **Go Gophers Slack**: #gamedev channel

---

## 💭 Personal Notes & Insights
```
[Use this space to write down "aha!" moments, 
confusions that got cleared up, or anything 
you want to remember]
```

---

## Next Steps

When you're ready to build the particle system, come back with:

1. ✅ Your learning repo showing completed exercises
2. ✅ Any questions or confusions
3. ✅ Your point sprite demo (if you did the side quest)
4. ✅ This TODO file with checkboxes marked

Then we'll start Phase 1 of the particle simulator together!

**Good luck! 🚀**
