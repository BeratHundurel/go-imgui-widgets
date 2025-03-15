package types

import (
	"github.com/AllenDang/cimgui-go/imgui"
)

type AppState struct {
	IsModalOpen bool
	NewTodoText string
	NewListTitle string
	CurrentListIds []int
	Todos []TodoLists
}


var (
	WindowWidth  = 800
	WindowHeight = 600
	DockspaceID  imgui.ID
	State    AppState
)
