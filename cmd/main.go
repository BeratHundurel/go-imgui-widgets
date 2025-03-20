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
	currentBackend.CreateWindow("Simple Todo App", types.WindowWidth, types.WindowHeight)
	currentBackend.SetBgColor(theme.Danger)
	
	fonts := utils.LoadedFonts()
	applyTheme()

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

func applyTheme() {
	ctx := imgui.CurrentContext()
	style := ctx.Style()
	colors := style.Colors()

	colors[imgui.ColTab] = theme.Background
	colors[imgui.ColTabHovered] = theme.Background
	colors[imgui.ColTabSelected] = theme.Background
	colors[imgui.ColTabSelectedOverline] = theme.Accent
	colors[imgui.ColTabDimmed] = theme.Background
	colors[imgui.ColTabDimmedSelected] = theme.Background
	colors[imgui.ColTabDimmedSelectedOverline] = theme.Accent
	
	colors[imgui.ColText] = theme.Text
	colors[imgui.ColTextDisabled] = theme.Disabled
	colors[imgui.ColTextSelectedBg] = theme.Accent
	
	colors[imgui.ColTitleBg] = theme.Background
	colors[imgui.ColTitleBgActive] = theme.Background
	colors[imgui.ColTitleBgCollapsed] = theme.Background
	
	colors[imgui.ColWindowBg] = theme.ElementBg
	colors[imgui.ColFrameBg] = theme.Background
	colors[imgui.ColPopupBg] = theme.ElementBg
	colors[imgui.ColBorder] = theme.Border
	
	colors[imgui.ColMenuBarBg] = theme.Background
	colors[imgui.ColScrollbarBg] = theme.ElementBg
	colors[imgui.ColScrollbarGrab] = theme.Accent
	colors[imgui.ColScrollbarGrabHovered] = theme.AccentHovered
	colors[imgui.ColScrollbarGrabActive] = theme.AccentHovered
	
	colors[imgui.ColSeparator] = theme.Border
	colors[imgui.ColSeparatorHovered] = theme.Accent
	colors[imgui.ColSeparatorActive] = theme.Accent
	
	colors[imgui.ColResizeGrip] = theme.Background
	colors[imgui.ColResizeGripHovered] = theme.Accent
	colors[imgui.ColResizeGripActive] = theme.Accent
	
	colors[imgui.ColSeparator] = theme.Border
	colors[imgui.ColSeparatorHovered] = theme.Accent
	colors[imgui.ColSeparatorActive] = theme.Accent
	
	colors[imgui.ColButton] = theme.Accent
	colors[imgui.ColButtonHovered] = theme.AccentHovered
	colors[imgui.ColButtonActive] = theme.AccentHovered
	
	colors[imgui.ColHeader] = theme.Background
	colors[imgui.ColHeaderHovered] = theme.Accent
	colors[imgui.ColHeaderActive] = theme.Accent
	
	colors[imgui.ColCheckMark] = theme.Accent
	
	colors[imgui.ColDockingEmptyBg] = theme.Background
	colors[imgui.ColDockingPreview] = theme.Accent
	
	style.SetColors(&colors)
	ctx.SetStyle(style)
}
