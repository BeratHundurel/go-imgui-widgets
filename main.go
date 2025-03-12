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

func init() {
	runtime.LockOSThread()
}

func main() {
	todoList = append(todoList, TodoItem{Text: "Example task 1", Completed: false})
	todoList = append(todoList, TodoItem{Text: "Example task 2", Completed: true})

	// Create the GLFW backend and set the background color
	currentBackend, _ = backend.CreateBackend(glfwbackend.NewGLFWBackend())
	currentBackend.SetBgColor(HexToVec4("#181616"))

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

	imgui.PushStyleColorVec4(imgui.ColWindowBg, HexToVec4("#181616"))
	imgui.PushStyleColorVec4(imgui.ColDockingEmptyBg, HexToVec4("#181616"))
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
	imgui.End() // End of dockspace window

	imgui.PopStyleVar()
	imgui.PopStyleColorV(2)
}

// Render the to-do list application window and logic
func renderTodoApp() {
	imgui.SetNextWindowSizeV(imgui.Vec2{X: 350, Y: 500}, imgui.CondFirstUseEver)
	imgui.PushStyleColorVec4(imgui.ColWindowBg, HexToVec4("#12120f"))
	if imgui.BeginV("Todo List", nil, imgui.WindowFlagsNone) {
		imgui.PushItemWidth(imgui.ContentRegionAvail().X - 60)
		enterPressed := imgui.InputTextWithHint("##newtodo", "Add new task...", &newTodoText, imgui.InputTextFlagsEnterReturnsTrue, nil)
		imgui.PopItemWidth()
		imgui.SameLine()

		// Add button for creating new to-do items
		imgui.PushStyleColorVec4(imgui.ColButton, imgui.Vec4{X: 0.2, Y: 0.5, Z: 0.3, W: 1.0})        // Button color
		imgui.PushStyleColorVec4(imgui.ColButtonHovered, imgui.Vec4{X: 0.3, Y: 0.6, Z: 0.4, W: 1.0}) // Hover color
		imgui.PushStyleColorVec4(imgui.ColButtonActive, imgui.Vec4{X: 0.1, Y: 0.4, Z: 0.2, W: 1.0})  // Active color
		imgui.PushStyleColorVec4(imgui.ColText, imgui.Vec4{X: 1.0, Y: 1.0, Z: 1.0, W: 1.0})
		addButtonPressed := imgui.ButtonV("Add", imgui.Vec2{X: 50, Y: 0})
		imgui.PopStyleColorV(4)

		// Add the new item when Enter is pressed or the Add button is clicked
		if (enterPressed || addButtonPressed) && newTodoText != "" {
			todoList = append(todoList, TodoItem{
				Text:      newTodoText,
				Completed: false,
			})
			newTodoText = ""
		}

		imgui.Separator()

		// Create a scrollable area for the to-do list items
		availHeight := imgui.ContentRegionAvail().Y - 30 
		imgui.BeginChildStrV("TodoListScroll", imgui.Vec2{X: 0, Y: availHeight}, 0, 0)

		// Variable to track which item to delete
		toDelete := -1

		// Iterate through the to-do list and render each item
		for i, item := range todoList {
			checkboxId := fmt.Sprintf("##check%d", i)
			if imgui.Checkbox(checkboxId, &todoList[i].Completed) {
				// Checkbox toggled: update completion status
			}

			imgui.SameLine()

			// Display the to-do item text, grayed out if completed
			if item.Completed {
				imgui.PushStyleColorVec4(imgui.ColText, imgui.Vec4{X: 0.5, Y: 0.5, Z: 0.5, W: 1.0})
				imgui.Text(todoList[i].Text)
				imgui.PopStyleColor() 
			} else {
				imgui.Text(todoList[i].Text)
			}

			imgui.SameLine()

			// Right-aligned delete button for each item
			currentX := imgui.CursorPosX()
			itemWidth := imgui.ContentRegionAvail().X
			imgui.SetCursorPosX(currentX + itemWidth - 30)
			deleteButtonId := fmt.Sprintf("X##%d", i)
			if imgui.Button(deleteButtonId) {
				toDelete = i 
			}
		}

		// Remove the item marked for deletion
		if toDelete >= 0 && toDelete < len(todoList) {
			todoList = slices.Delete(todoList, toDelete, toDelete+1)
		}

		imgui.EndChild()

		imgui.Separator()
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
	imgui.PopStyleColor() 
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
