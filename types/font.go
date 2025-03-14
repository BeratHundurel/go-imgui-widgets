package types

import "github.com/AllenDang/cimgui-go/imgui"

type Font struct {
	FontName string
	FontSize float32
	*imgui.Font
}