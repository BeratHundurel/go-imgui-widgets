package components

import (
	"imgui_try/database"
	"imgui_try/types"
	"slices"

	"github.com/AllenDang/cimgui-go/imgui"
)

func RenderMenubar() {
	if imgui.BeginMainMenuBar() {
		if imgui.BeginMenu("File") {
			// Add New List option
			if imgui.MenuItemBool("Add New List") {
				types.State.IsModalOpen = true
			}
			imgui.EndMenu()
		}

		// Show modal if types.ShowModal is true
		if types.State.IsModalOpen {
			imgui.OpenPopupStr("New To-Do List")
		}

		// Create the modal window for adding a new list
		if imgui.BeginPopupModalV("New To-Do List", nil, imgui.WindowFlagsAlwaysAutoResize) {
			// Title input field
			imgui.InputTextWithHint("Title", "Enter a title for the new list...", &types.State.NewListTitle, 0, nil)

			// Add button
			if imgui.Button("Add") {
				if types.State.NewListTitle != "" {
					// Create a new to-do list and append it
					newList := types.TodoLists{Title: types.State.NewListTitle}
					types.State.Todos = append(types.State.Todos, newList)

					// Save the new list to the database
					database.New().CreateTodoList(types.State.NewListTitle)

					types.State.IsModalOpen = false
					types.State.NewListTitle = ""
					imgui.CloseCurrentPopup()
				}
			}

			// Cancel button
			if imgui.Button("Cancel") {
				types.State.IsModalOpen = false
				types.State.NewListTitle = "" // Clear the title input
				imgui.CloseCurrentPopup()
			}
			
			// End the modal window
			imgui.EndPopup()
		}

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
