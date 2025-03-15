package components

import (
	"fmt"
	"imgui_try/database"
	"imgui_try/theme"
	"imgui_try/types"
	"slices"

	"github.com/AllenDang/cimgui-go/imgui"
)

// Render the to-do list application window and logic
func RenderTodoList(todo types.TodoLists) {
	imgui.SetNextWindowSizeV(imgui.Vec2{X: 350, Y: 500}, imgui.CondFirstUseEver)
	imgui.SetNextWindowSizeConstraints(imgui.Vec2{X: 300, Y: 400}, imgui.Vec2{X: 800, Y: 600})

	imgui.PushStyleColorVec4(imgui.ColWindowBg, theme.ElementBg)
	imgui.PushStyleColorVec4(imgui.ColText, theme.Text)
	imgui.PushStyleColorVec4(imgui.ColTitleBg, theme.Background)
	imgui.PushStyleColorVec4(imgui.ColTitleBgActive, theme.Background)
	imgui.PushStyleColorVec4(imgui.ColBorder, theme.Accent)
	imgui.PushStyleColorVec4(imgui.ColResizeGrip, theme.Accent)
	imgui.PushStyleColorVec4(imgui.ColResizeGripHovered, theme.AccentHovered)
	imgui.PushStyleColorVec4(imgui.ColResizeGripActive, theme.AccentHovered)
	imgui.PushStyleColorVec4(imgui.ColTitleBgCollapsed, theme.Disabled)
	imgui.PushStyleColorVec4(imgui.ColFrameBg, theme.ElementBg)
	imgui.PushStyleColorVec4(imgui.ColFrameBgActive, theme.ElementBg)
	imgui.PushStyleColorVec4(imgui.ColScrollbarBg, theme.ElementBg)
	imgui.PushStyleColorVec4(imgui.ColScrollbarGrab, theme.Accent)

	imgui.PushStyleVarFloat(imgui.StyleVarFrameRounding, 4.0)
	imgui.PushStyleVarFloat(imgui.StyleVarWindowRounding, 6.0)
	imgui.PushStyleVarVec2(imgui.StyleVarFramePadding, imgui.Vec2{X: 8, Y: 4})

	if imgui.BeginV(todo.Title, nil, imgui.WindowFlagsNone) {
		imgui.PushItemWidth(imgui.ContentRegionAvail().X - 60)
		imgui.PushStyleColorVec4(imgui.ColText, theme.Muted)
		imgui.PushStyleColorVec4(imgui.ColFrameBg, theme.Background)
		imgui.PushStyleVarVec2(imgui.StyleVarFramePadding, imgui.Vec2{X: 8, Y: 5})
		enterPressed := imgui.InputTextWithHint("##newtodo", "Add new task...", &types.State.NewTodoText, imgui.InputTextFlagsEnterReturnsTrue, nil)
		imgui.PopItemWidth()
		imgui.PopStyleVar()
		imgui.PopStyleColorV(2)
		imgui.SameLine() // Add button for creating new to-do items

		imgui.PushStyleColorVec4(imgui.ColButton, theme.Accent)
		imgui.PushStyleColorVec4(imgui.ColButtonHovered, theme.AccentHovered)
		imgui.PushStyleColorVec4(imgui.ColText, theme.Text)
		imgui.PushStyleVarVec2(imgui.StyleVarFramePadding, imgui.Vec2{X: 12.0, Y: 6.0})
		imgui.PushStyleVarVec2(imgui.StyleVarButtonTextAlign, imgui.Vec2{X: 0.5, Y: 0.5})
		imgui.PushStyleVarVec2(imgui.StyleVarItemSpacing, imgui.Vec2{X: 0.0, Y: 10.0})
		addButtonPressed := imgui.ButtonV("Add", imgui.Vec2{X: 50, Y: 0})
		imgui.PopStyleColorV(4)
		imgui.PopStyleVarV(3)

		// Add the new item when Enter is pressed or the Add button is clicked
		if (enterPressed || addButtonPressed) && types.State.NewTodoText != "" {
			todo.Items = append(todo.Items, types.TodoItem{
				Text:      types.State.NewTodoText,
				Completed: false,
			})

			database.CreateTodoItem(todo.Id, types.State.NewTodoText)
			types.State.NewTodoText = ""
		}

		imgui.PushStyleVarVec2(imgui.StyleVarItemSpacing, imgui.Vec2{X: 0.0, Y: 10.0})
		imgui.Separator()
		imgui.PopStyleVar()

		// Create a scrollable area for the to-do list items
		availHeight := imgui.ContentRegionAvail().Y - 30
		imgui.BeginChildStrV("types.TodoListScroll", imgui.Vec2{X: 0, Y: availHeight}, 0, 0)

		// Variable to track which item to delete
		toDelete := -1

		// Iterate through the to-do list and render each item
		for i, item := range todo.Items {
			imgui.PushStyleVarVec2(imgui.StyleVarFramePadding, imgui.Vec2{X: 4, Y: 3})
			imgui.PushStyleVarVec2(imgui.StyleVarItemSpacing, imgui.Vec2{X: 0.0, Y: 10.0})

			imgui.PushStyleColorVec4(imgui.ColCheckMark, theme.Accent)
			imgui.PushStyleColorVec4(imgui.ColFrameBg, theme.Background)
			imgui.PushStyleColorVec4(imgui.ColFrameBgHovered, theme.AccentHovered)
			imgui.PushStyleColorVec4(imgui.ColFrameBgActive, theme.AccentHovered)

			checkboxId := fmt.Sprintf("##check%d", i)
			if imgui.Checkbox(checkboxId, &item.Completed) {
				// Checkbox toggled: update completion status
			}

			imgui.PopStyleColorV(4)
			imgui.SameLine()

			// Display the to-do item text, grayed out if completed
			if item.Completed {
				imgui.PushStyleColorVec4(imgui.ColText, theme.Muted)
				imgui.Text(item.Text)
			} else {
				imgui.PushStyleColorVec4(imgui.ColText, theme.Text)
				imgui.Text(item.Text)
			}
			imgui.PopStyleColor()

			imgui.SameLine()

			// Right-aligned delete button for each item
			currentX := imgui.CursorPosX()
			itemWidth := imgui.ContentRegionAvail().X
			imgui.SetCursorPosX(currentX + itemWidth - 40)

			imgui.PushStyleColorVec4(imgui.ColButton, theme.Danger)
			imgui.PushStyleColorVec4(imgui.ColButtonHovered, theme.DangerHovered)
			imgui.PushStyleColorVec4(imgui.ColText, theme.Text)

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
		if toDelete >= 0 && toDelete < len(todo.Items) {
			todo.Items = slices.Delete(todo.Items, toDelete, toDelete+1)
		}

		imgui.EndChild()
		imgui.Separator()

		imgui.PushStyleColorVec4(imgui.ColText, theme.Accent)
		imgui.Text(fmt.Sprintf("Total: %d items", len(todo.Items)))

		completed := 0
		for _, item := range todo.Items {
			if item.Completed {
				completed++
			}
		}

		if len(todo.Items) > 0 {
			completionRate := float32(completed) / float32(len(todo.Items)) * 100
			imgui.SameLine()
			imgui.Text(fmt.Sprintf("Completed: %.1f%%", completionRate))
		}
	}
	imgui.PopStyleVarV(3)
	imgui.PopStyleColorV(13)
	imgui.End()
}
