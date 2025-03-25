package components

import (
	"imgui_try/database"
	"imgui_try/theme"
	"imgui_try/types"
	"slices"

	"github.com/AllenDang/cimgui-go/imgui"
)

func RenderMenubar() {
	imgui.PushStyleVarVec2(imgui.StyleVarFramePadding, imgui.Vec2{X: 24, Y: 12})
	if imgui.BeginMainMenuBar() {
		imgui.PushStyleVarVec2(imgui.StyleVarItemSpacing, imgui.Vec2{X: 24, Y: 12})

		if imgui.BeginMenu("Add New") {
			types.State.IsModalOpen = true
			imgui.EndMenu()
		}

		if types.State.IsModalOpen {
			imgui.OpenPopupStr("New To-Do List")
			imgui.SetNextWindowSize(imgui.Vec2{X: 400, Y: 400})
		}

		imgui.PushStyleVarVec2(imgui.StyleVarFramePadding, imgui.Vec2{X: 12, Y: 12})
		imgui.PushStyleVarVec2(imgui.StyleVarItemSpacing, imgui.Vec2{X: 12, Y: 16})
		imgui.PushStyleVarFloat(imgui.StyleVarWindowRounding, 12)
		if imgui.BeginPopupModalV("New To-Do List", nil, imgui.WindowFlagsAlwaysAutoResize) {
			
			imgui.PushItemWidth(380)
			imgui.InputTextWithHint("##NewToDoTitle", "Enter a title", &types.State.NewListTitle, 0, nil)
			imgui.PopItemWidth()

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

			imgui.PushStyleColorVec4(imgui.ColButton, theme.Danger)
			imgui.PushStyleColorVec4(imgui.ColButtonHovered, theme.DangerHovered)
			imgui.PushStyleColorVec4(imgui.ColButtonActive, theme.DangerHovered)
			if imgui.ButtonV("Cancel", imgui.Vec2{X: 100, Y: 0}) {
				types.State.IsModalOpen = false
				types.State.NewListTitle = "" 
				imgui.CloseCurrentPopup()
			}

			imgui.PopStyleVarV(2)
			imgui.PopStyleColorV(3)

			imgui.EndPopup()
		}

		imgui.PopStyleVarV(3)

		if imgui.BeginMenu("Available Lists") {
			imgui.PushStyleVarVec2(imgui.StyleVarFramePadding, imgui.Vec2{X: 12, Y: 24})
			imgui.PushStyleVarVec2(imgui.StyleVarItemSpacing, imgui.Vec2{X: 12, Y: 24})
			itemSize := imgui.Vec2{X: 200, Y: 0} 
			for _, list := range types.State.Todos {
				isSelected := slices.Contains(types.State.CurrentListIds, list.Id)

				if imgui.SelectableBoolV(list.Title, isSelected, imgui.SelectableFlagsNone, itemSize) {
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
			imgui.PopStyleVarV(2)
			imgui.EndMenu()
		}

		imgui.PopStyleVarV(2)
		imgui.EndMainMenuBar()
	}
}
