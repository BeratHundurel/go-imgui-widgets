package components

import (
	"imgui_try/theme"
	"imgui_try/types"

	"github.com/AllenDang/cimgui-go/imgui"
)

// Create the main dockspace to house windows and widgets
func CreateDockspace() {
	viewport := imgui.MainViewport()

	imgui.PushStyleColorVec4(imgui.ColWindowBg, theme.Background)
	imgui.PushStyleColorVec4(imgui.ColDockingEmptyBg, theme.Background)
	imgui.PushStyleColorVec4(imgui.ColDragDropTarget, theme.Accent)
	imgui.PushStyleColorVec4(imgui.ColTitleBg, theme.Background)
	imgui.PushStyleColorVec4(imgui.ColTitleBgActive, theme.Background)
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
	types.DockspaceID = imgui.IDStr("MyDockSpace")

	imgui.DockSpace(types.DockspaceID)
	imgui.PopStyleVar()
	imgui.PopStyleColorV(5)
	imgui.End() // End of dockspace window
}