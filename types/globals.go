package types

import (
	"github.com/AllenDang/cimgui-go/imgui"
)

type AppState struct {
	IsModalOpen bool
	NewListTitle string
	CurrentListIds []int
	Todos []TodoLists
	NewTodoTexts map[int]string
}


var (
	WindowWidth  = 800
	WindowHeight = 600
	DockspaceID  imgui.ID
	State    AppState
)

