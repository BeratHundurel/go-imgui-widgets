package main

import (
	"runtime"

	"imgui_try/components"
	"imgui_try/theme"
	"imgui_try/types"
	"imgui_try/utils"

	"github.com/AllenDang/cimgui-go/backend"
	"github.com/AllenDang/cimgui-go/backend/glfwbackend"
	"github.com/AllenDang/cimgui-go/imgui"
)

// Global backend variable to handle GUI rendering
var currentBackend backend.Backend[glfwbackend.GLFWWindowFlags]

func init() {
	runtime.LockOSThread()
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
	components.CreateDockspace()
	renderMenubar()
	components.RenderTodoApp()
}

func renderMenubar() {
    if imgui.BeginMainMenuBar() {
        if imgui.BeginMenu("File") {
            if imgui.BeginMenu("Add New List") {
            	if imgui.MenuItemBool("Add New List") {
					types.TodoList = append(types.TodoList, types.TodoItem{Text: "New List"})
				}
				imgui.EndMenu()
            }
            imgui.EndMenu()
        }
        imgui.EndMainMenuBar()
    }
}