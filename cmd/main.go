package main

import (
	"imgui_try/components"
	"imgui_try/database"
	"imgui_try/theme"
	"imgui_try/types"
	"imgui_try/utils"
	"runtime"

	"github.com/AllenDang/cimgui-go/backend"
	"github.com/AllenDang/cimgui-go/backend/glfwbackend"
	"github.com/AllenDang/cimgui-go/imgui"
)

// Global backend variable to handle GUI rendering
var currentBackend backend.Backend[glfwbackend.GLFWWindowFlags]

func init() {
	runtime.LockOSThread()
	database.New().Migrate()
}

func main() {
	currentBackend, _ = backend.CreateBackend(glfwbackend.NewGLFWBackend())
	currentBackend.SetBgColor(theme.Background)
	currentBackend.CreateWindow("Simple Todo App", types.WindowWidth, types.WindowHeight)

	fonts := utils.LoadedFonts()

	io := imgui.CurrentIO()
	io.SetConfigFlags(io.ConfigFlags() | imgui.ConfigFlagsDockingEnable)
	io.SetFontDefault(fonts[0].Font)

	currentBackend.Run(renderLoop)
}

func renderLoop() {
	components.RenderMenubar()
	components.CreateDockspace()
	components.RenderTodoList()
}
