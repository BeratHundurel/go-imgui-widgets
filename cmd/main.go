package main

import (
	"imgui_try/components"
	"imgui_try/database"
	"imgui_try/theme"
	"imgui_try/types"
	"imgui_try/utils"
	"slices"
	"runtime"

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
	database.New().Migrate()

	currentBackend, _ = backend.CreateBackend(glfwbackend.NewGLFWBackend())
	currentBackend.SetBgColor(theme.Background)
	currentBackend.CreateWindow("Simple Todo App", types.WindowWidth, types.WindowHeight)

	fonts := utils.LoadedFonts()

	io := imgui.CurrentIO()
	io.SetConfigFlags(io.ConfigFlags() | imgui.ConfigFlagsDockingEnable)
	io.SetFontDefault(fonts[0].Font)

	types.State = types.AppState{
		IsModalOpen:    false,
		NewTodoText:    "",
		NewListTitle:   "",
		CurrentListIds: []int{},
		Todos:          database.GetAllTodos(), // Get all todos from the database
	}

	currentBackend.Run(renderLoop)
}

func renderLoop() {
	components.RenderMenubar()
	components.CreateDockspace()
	
	for _, todo := range types.State.Todos {
		if slices.Contains(types.State.CurrentListIds, todo.Id) {
			components.RenderTodoList(todo)
		}
	}
}
