package components

import (
	"imgui_try/database"
	"imgui_try/theme"
	"imgui_try/types"
	"slices"

	"github.com/AllenDang/cimgui-go/imgui"
)

func RenderMenubar() {
	if imgui.BeginMainMenuBar() {

		if imgui.BeginMenu("Add New") {
			types.State.IsModalOpen = true
			imgui.EndMenu()
		}

		if types.State.IsModalOpen {
			imgui.OpenPopupStr("New To-Do List")
			io := imgui.CurrentIO()
			displaySizeX := io.DisplaySize().X
			displaySizeY := io.DisplaySize().Y

			centerPos := imgui.Vec2{
				X: displaySizeX / 2,
				Y: displaySizeY / 2,
			}

			// Set the modal position at the center
			imgui.SetNextWindowPosV(
				centerPos,
				imgui.CondAlways,
				imgui.Vec2{X: 0.5, Y: 0.5}, // Centering pivot
			)

			imgui.SetNextWindowSize(imgui.Vec2{X: 400, Y: 400})
		}

		imgui.PushStyleVarVec2(imgui.StyleVarFramePadding, imgui.Vec2{X: 12, Y: 12})
		imgui.PushStyleVarVec2(imgui.StyleVarItemSpacing, imgui.Vec2{X: 12, Y: 16})

		if imgui.BeginPopupModalV("New To-Do List", nil, imgui.WindowFlagsAlwaysAutoResize) {

			imgui.InputTextWithHint("##NewToDoTitle", "Enter a title", &types.State.NewListTitle, 0, nil)

			imgui.PushStyleVarVec2(imgui.StyleVarFramePadding, imgui.Vec2{X: 8, Y: 8})
			imgui.PushStyleVarFloat(imgui.StyleVarFrameRounding, 4.0)
			if imgui.ButtonV("Add", imgui.Vec2{X: 100, Y: 0}) {
				if types.State.NewListTitle != "" {
					newList := types.TodoLists{Title: types.State.NewListTitle}
					types.State.Todos = append(types.State.Todos, newList)

					database.New().CreateTodoList(types.State.NewListTitle)

					types.State.IsModalOpen = false
					types.State.NewListTitle = ""
					imgui.CloseCurrentPopup()
				}
			}

			imgui.SameLine()

			// Cancel button
			imgui.PushStyleColorVec4(imgui.ColButton, theme.Danger)
			imgui.PushStyleColorVec4(imgui.ColButtonHovered, theme.DangerHovered)
			imgui.PushStyleColorVec4(imgui.ColButtonActive, theme.DangerHovered)
			if imgui.ButtonV("Cancel", imgui.Vec2{X: 100, Y: 0}) {
				types.State.IsModalOpen = false
				types.State.NewListTitle = "" // Clear the title input
				imgui.CloseCurrentPopup()
			}

			imgui.PopStyleVarV(2)
			imgui.PopStyleColorV(3)

			// End the modal window
			imgui.EndPopup()
		}

		imgui.PopStyleVarV(2)

		// "Lists" menu for displaying all available lists
		if imgui.BeginMenu("Lists") {
			// Display all available lists as selectable items
			for _, list := range types.State.Todos {
				isSelected := slices.Contains(types.State.CurrentListIds, list.Id)
				if imgui.SelectableBoolV(list.Title, isSelected, imgui.SelectableFlagsNone, imgui.Vec2{}) {
					// If the list is not already selected, add it to the selected list IDs
					if !isSelected {
						types.State.CurrentListIds = append(types.State.CurrentListIds, list.Id)
					} else {
						// Otherwise, remove it from the selected list IDs
						types.State.CurrentListIds = slices.DeleteFunc(types.State.CurrentListIds, func(E int) bool {
							return E == list.Id
						})
					}
				}
			}
			imgui.EndMenu()
		}
		imgui.EndMainMenuBar()
	}
}
