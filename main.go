package main

import (
	"fmt"
	"runtime"
	"slices"
	"strconv"

	"github.com/AllenDang/cimgui-go/backend"
	"github.com/AllenDang/cimgui-go/backend/glfwbackend"
	"github.com/AllenDang/cimgui-go/imgui"
)

// Global backend variable to handle GUI rendering
var currentBackend backend.Backend[glfwbackend.GLFWWindowFlags]

// TodoItem represents a single to-do item with text and completion status
type TodoItem struct {
	Text      string
	Completed bool
}

type fonts struct {
	fontName string
	fontSize float32
	*imgui.Font
}

var (
	todoList     = []TodoItem{}
	newTodoText  = ""
	windowWidth  = 800
	windowHeight = 600
	dockspaceID  imgui.ID
)

var (
	Accent        = HexToVec4("#6f7a8c")
	AccentHovered = HexToVec4("#5a6574")
	Danger        = HexToVec4("#c4746e")
	DangerHovered = HexToVec4("#b76355")
	Muted         = HexToVec4("#7a8382")
	Text          = HexToVec4("#c5c9c5")
	Background    = HexToVec4("#181616")
	ElementBg     = HexToVec4("#1d1c19")
	Border        = HexToVec4("#393836")
	Disabled      = HexToVec4("#625e5a")
)

func init() {
	runtime.LockOSThread()
}

func main() {
	todoList = append(todoList, TodoItem{Text: "Example task 1", Completed: false})
	todoList = append(todoList, TodoItem{Text: "Example task 2", Completed: true})

	// Create the GLFW backend and set the background color
	currentBackend, _ = backend.CreateBackend(glfwbackend.NewGLFWBackend())
	currentBackend.SetBgColor(Background)

	// Create the window with a specific size and title
	currentBackend.CreateWindow("Simple Todo App", windowWidth, windowHeight)

	fonts := loadedFonts()

	// Enable docking feature in ImGui
	io := imgui.CurrentIO()
	io.SetConfigFlags(io.ConfigFlags() | imgui.ConfigFlagsDockingEnable)
	io.SetFontDefault(fonts[0].Font)

	// Register drop event callback (for file drops or similar)
	currentBackend.SetDropCallback(func(p []string) {
		fmt.Printf("drop triggered: %v\n", p)
	})

	// Register window close event callback
	currentBackend.SetCloseCallback(func() {
		fmt.Println("window is closing")
	})

	currentBackend.Run(renderLoop)
}

// Main render loop that continuously updates the UI
func renderLoop() {
	createDockspace()
	renderTodoApp()
}

// Create the main dockspace to house windows and widgets
func createDockspace() {
	viewport := imgui.MainViewport()

	imgui.PushStyleColorVec4(imgui.ColWindowBg, Background)
	imgui.PushStyleColorVec4(imgui.ColDockingEmptyBg, Background)
	imgui.PushStyleColorVec4(imgui.ColDragDropTarget, Accent)
	imgui.PushStyleColorVec4(imgui.ColTitleBg, Background)
	imgui.PushStyleColorVec4(imgui.ColTitleBgActive, Background)
	imgui.PushStyleVarVec2(imgui.StyleVarWindowPadding, imgui.Vec2{X: 0, Y: 0})

	imgui.SetNextWindowPos(viewport.Pos())
	imgui.SetNextWindowSize(viewport.Size())
	imgui.SetNextWindowViewport(viewport.ID())

	windowFlags := imgui.WindowFlagsNoTitleBar | imgui.WindowFlagsNoCollapse |
		imgui.WindowFlagsNoResize | imgui.WindowFlagsNoMove |
		imgui.WindowFlagsNoBringToFrontOnFocus | imgui.WindowFlagsNoNavFocus |
		imgui.WindowFlagsNoDocking

	// Begin the dockspace window with no title or WindowFlagsNonements
	imgui.BeginV("DockSpace", nil, windowFlags)

	// Create a dockspace for holding dockable windows
	dockspaceID = imgui.IDStr("MyDockSpace")

	imgui.DockSpace(dockspaceID)
	imgui.PopStyleVar()
	imgui.PopStyleColorV(5)
	imgui.End() // End of dockspace window
}

// Render the to-do list application window and logic
func renderTodoApp() {
	imgui.SetNextWindowSizeV(imgui.Vec2{X: 350, Y: 500}, imgui.CondFirstUseEver)
	imgui.SetNextWindowSizeConstraints(imgui.Vec2{X: 300, Y: 400}, imgui.Vec2{X: 800, Y: 600})

	imgui.PushStyleColorVec4(imgui.ColWindowBg, ElementBg)
	imgui.PushStyleColorVec4(imgui.ColText, Text)
	imgui.PushStyleColorVec4(imgui.ColTitleBg, Background)
	imgui.PushStyleColorVec4(imgui.ColTitleBgActive, Background)
	imgui.PushStyleColorVec4(imgui.ColBorder, Accent)
	imgui.PushStyleColorVec4(imgui.ColResizeGrip, Accent)
	imgui.PushStyleColorVec4(imgui.ColResizeGripHovered, AccentHovered)
	imgui.PushStyleColorVec4(imgui.ColResizeGripActive, AccentHovered)
	imgui.PushStyleColorVec4(imgui.ColTitleBgCollapsed, Disabled)
	imgui.PushStyleColorVec4(imgui.ColFrameBg, ElementBg)
	imgui.PushStyleColorVec4(imgui.ColFrameBgActive, ElementBg)
	imgui.PushStyleColorVec4(imgui.ColScrollbarBg, ElementBg)
	imgui.PushStyleColorVec4(imgui.ColScrollbarGrab, Accent)

	imgui.PushStyleVarFloat(imgui.StyleVarFrameRounding, 4.0)
	imgui.PushStyleVarFloat(imgui.StyleVarWindowRounding, 6.0)
	imgui.PushStyleVarVec2(imgui.StyleVarFramePadding, imgui.Vec2{X: 8, Y: 4})

	if imgui.BeginV("Todo List", nil, imgui.WindowFlagsNone) {
		imgui.PushItemWidth(imgui.ContentRegionAvail().X - 60)
		imgui.PushStyleColorVec4(imgui.ColText, Muted)
		imgui.PushStyleColorVec4(imgui.ColFrameBg, Background)
		imgui.PushStyleVarVec2(imgui.StyleVarFramePadding, imgui.Vec2{X: 8, Y: 5})
		enterPressed := imgui.InputTextWithHint("##newtodo", "Add new task...", &newTodoText, imgui.InputTextFlagsEnterReturnsTrue, nil)
		imgui.PopItemWidth()
		imgui.PopStyleVar()
		imgui.PopStyleColorV(2)
		imgui.SameLine() // Add button for creating new to-do items

		imgui.PushStyleColorVec4(imgui.ColButton, Accent)
		imgui.PushStyleColorVec4(imgui.ColButtonHovered, AccentHovered)
		imgui.PushStyleColorVec4(imgui.ColText, Text)
		imgui.PushStyleVarVec2(imgui.StyleVarFramePadding, imgui.Vec2{X: 12.0, Y: 6.0})
		imgui.PushStyleVarVec2(imgui.StyleVarButtonTextAlign, imgui.Vec2{X: 0.5, Y: 0.5})
		imgui.PushStyleVarVec2(imgui.StyleVarItemSpacing, imgui.Vec2{X: 0.0, Y: 10.0})
		addButtonPressed := imgui.ButtonV("Add", imgui.Vec2{X: 50, Y: 0})
		imgui.PopStyleColorV(4)
		imgui.PopStyleVarV(3)

		// Add the new item when Enter is pressed or the Add button is clicked
		if (enterPressed || addButtonPressed) && newTodoText != "" {
			todoList = append(todoList, TodoItem{
				Text:      newTodoText,
				Completed: false,
			})
			newTodoText = ""
		}

		imgui.PushStyleVarVec2(imgui.StyleVarItemSpacing, imgui.Vec2{X: 0.0, Y: 10.0})
		imgui.Separator()
		imgui.PopStyleVar()

		// Create a scrollable area for the to-do list items
		availHeight := imgui.ContentRegionAvail().Y - 30
		imgui.BeginChildStrV("TodoListScroll", imgui.Vec2{X: 0, Y: availHeight}, 0, 0)

		// Variable to track which item to delete
		toDelete := -1

		// Iterate through the to-do list and render each item
		for i, item := range todoList {
			imgui.PushStyleVarVec2(imgui.StyleVarFramePadding, imgui.Vec2{X: 4, Y: 3})
			imgui.PushStyleVarVec2(imgui.StyleVarItemSpacing, imgui.Vec2{X: 0.0, Y: 10.0})

			imgui.PushStyleColorVec4(imgui.ColCheckMark, Accent)
			imgui.PushStyleColorVec4(imgui.ColFrameBg, Background)
			imgui.PushStyleColorVec4(imgui.ColFrameBgHovered, AccentHovered)
			imgui.PushStyleColorVec4(imgui.ColFrameBgActive, AccentHovered)

			checkboxId := fmt.Sprintf("##check%d", i)
			if imgui.Checkbox(checkboxId, &todoList[i].Completed) {
				// Checkbox toggled: update completion status
			}

			imgui.PopStyleColorV(4)
			imgui.SameLine()

			// Display the to-do item text, grayed out if completed
			if item.Completed {
				imgui.PushStyleColorVec4(imgui.ColText, Muted)
				imgui.Text(todoList[i].Text)
			} else {
				imgui.PushStyleColorVec4(imgui.ColText, Text)
				imgui.Text(todoList[i].Text)
			}
			imgui.PopStyleColor()

			imgui.SameLine()

			// Right-aligned delete button for each item
			currentX := imgui.CursorPosX()
			itemWidth := imgui.ContentRegionAvail().X
			imgui.SetCursorPosX(currentX + itemWidth - 40)

			imgui.PushStyleColorVec4(imgui.ColButton, Danger)
			imgui.PushStyleColorVec4(imgui.ColButtonHovered, DangerHovered)
			imgui.PushStyleColorVec4(imgui.ColText, Text)

			imgui.PushStyleVarFloat(imgui.StyleVarFrameRounding, 3.0)
			imgui.PushStyleVarVec2(imgui.StyleVarFramePadding, imgui.Vec2{X: 12.0, Y: 6.0})
			imgui.PushStyleVarVec2(imgui.StyleVarButtonTextAlign, imgui.Vec2{X: 0.5, Y: 0.5})
			deleteButtonId := fmt.Sprintf("X##%d", i)
			if imgui.Button(deleteButtonId) {
				toDelete = i
			}

			imgui.PopStyleVarV(5)
			imgui.PopStyleColorV(3)
		}

		// Remove the item marked for deletion
		if toDelete >= 0 && toDelete < len(todoList) {
			todoList = slices.Delete(todoList, toDelete, toDelete+1)
		}

		imgui.EndChild()
		imgui.Separator()

		imgui.PushStyleColorVec4(imgui.ColText, Accent)
		imgui.Text(fmt.Sprintf("Total: %d items", len(todoList)))

		completed := 0
		for _, item := range todoList {
			if item.Completed {
				completed++
			}
		}

		if len(todoList) > 0 {
			completionRate := float32(completed) / float32(len(todoList)) * 100
			imgui.SameLine()
			imgui.Text(fmt.Sprintf("Completed: %.1f%%", completionRate))
		}
	}
	imgui.PopStyleVarV(3)
	imgui.PopStyleColorV(13)
	imgui.End()
}

func HexToVec4(hex string) imgui.Vec4 {
	if hex[0] == '#' {
		hex = hex[1:]
	}

	r, _ := strconv.ParseUint(hex[0:2], 16, 8)
	g, _ := strconv.ParseUint(hex[2:4], 16, 8)
	b, _ := strconv.ParseUint(hex[4:6], 16, 8)

	return imgui.Vec4{
		X: float32(r) / 255.0,
		Y: float32(g) / 255.0,
		Z: float32(b) / 255.0,
		W: 1.0,
	}
}

func loadedFonts() []fonts {
	fonts := []fonts{
		{"Normal", 16, getImguiFont("fonts/FiraMonoNerdFont-Regular.otf", 18)},
		{"Bold", 16, getImguiFont("fonts/FiraMonoNerdFont-Bold.otf", 16)},
		{"Medium", 16, getImguiFont("fonts/FiraMonoNerdFont-Medium.otf", 16)},
	}

	return fonts
}

func getImguiFont(path string, size float32) *imgui.Font {
	io := imgui.CurrentIO()
	imgui_fonts := io.Fonts()

	return imgui_fonts.AddFontFromFileTTF(path, size)
}
