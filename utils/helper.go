package utils

import (
	"strconv"
	"github.com/AllenDang/cimgui-go/imgui"
	"imgui_try/types"
)

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

func LoadedFonts() []types.Font {
	fonts := []types.Font{
		{FontName: "Normal", FontSize: 18, Font: getImguiFont("../fonts/FiraMonoNerdFont-Regular.otf", 18)},
		{FontName: "Bold", FontSize: 16, Font: getImguiFont("../fonts/FiraMonoNerdFont-Bold.otf", 16)},
		{FontName: "Medium", FontSize: 16, Font: getImguiFont("../fonts/FiraMonoNerdFont-Medium.otf", 16)},
	}

	return fonts
}

func getImguiFont(path string, size float32) *imgui.Font {
	io := imgui.CurrentIO()
	imgui_fonts := io.Fonts()

	return imgui_fonts.AddFontFromFileTTF(path, size)
}