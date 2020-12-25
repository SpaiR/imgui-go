package imgui

// #include "wrapper/DrawData.h"
import "C"
import "unsafe"

// RenderedDrawData returns the created draw commands, which are valid after Render() and
// until the next call to NewFrame(). This is what you have to render.
func RenderedDrawData() DrawData {
	return DrawData(C.iggGetDrawData())
}

// DrawData contains all draw data to render an ImGui frame.
type DrawData uintptr

func (data DrawData) handle() C.IggDrawData {
	return C.IggDrawData(data)
}

// Valid indicates whether the structure is usable.
// It is valid only after Render() is called and before the next NewFrame() is called.
func (data DrawData) Valid() bool {
	return (data.handle() != nil) && (C.iggDrawDataValid(data.handle()) != 0)
}

// CommandLists is an array of DrawList to render.
// The DrawList are owned by the context and only pointed to from here.
func (data DrawData) CommandLists() []DrawList {
	var handles unsafe.Pointer
	var count C.int

	C.iggDrawDataGetCommandLists(data.handle(), &handles, &count)
	list := make([]DrawList, int(count))
	for i := 0; i < int(count); i++ {
		list[i] = DrawList(*((*uintptr)(handles)))
		handles = unsafe.Pointer(uintptr(handles) + unsafe.Sizeof(handles)) // nolint: gas
	}

	return list
}

// ScaleClipRects is a helper to scale the ClipRect field of each DrawCmd.
// Use if your final output buffer is at a different scale than ImGui expects,
// or if there is a difference between your window resolution and framebuffer resolution.
func (data DrawData) ScaleClipRects(scale Vec2) {
	scaleArg, _ := scale.wrapped()
	C.iggDrawDataScaleClipRects(data.handle(), scaleArg)
}

func (data DrawData) GetDisplayPos() Vec2 {
	var value Vec2
	valueArg, valueFin := value.wrapped()
	C.iggDrawDataGetDisplayPos(data.handle(), valueArg)
	valueFin()
	return value
}

func (data DrawData) SetDisplayPos(value Vec2) {
	valueArg, _ := value.wrapped()
	C.iggDrawDataSetDisplayPos(data.handle(), valueArg)
}

func (data DrawData) GetDisplaySize() Vec2 {
	var value Vec2
	valueArg, valueFin := value.wrapped()
	C.iggDrawDataGetDisplaySize(data.handle(), valueArg)
	valueFin()
	return value
}

func (data DrawData) SetDisplaySize(value Vec2) {
	valueArg, _ := value.wrapped()
	C.iggDrawDataSetDisplaySize(data.handle(), valueArg)
}

func (data DrawData) GetFramebufferScale() Vec2 {
	var value Vec2
	valueArg, valueFin := value.wrapped()
	C.iggDrawDataGetFramebufferScale(data.handle(), valueArg)
	valueFin()
	return value
}

func (data DrawData) SetFramebufferScale(value Vec2) {
	valueArg, _ := value.wrapped()
	C.iggDrawDataSetFramebufferScale(data.handle(), valueArg)
}
