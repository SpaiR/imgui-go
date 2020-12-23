package imgui

// #include "wrapper/Viewport.h"
import "C"

func SetNextWindowViewport(viewportId int) {
	C.iggSetNextWindowViewport(C.int(viewportId))
}

type Viewport uintptr

func GetWindowViewport() Viewport {
	return Viewport(C.iggGetWindowViewport())
}

func GetMainViewport() Viewport {
	return Viewport(C.iggGetMainViewport())
}

func (viewport Viewport) handle() C.IggViewport {
	return C.IggViewport(viewport)
}

func (viewport Viewport) GetID() int {
	return int(C.iggViewportGetID(viewport.handle()))
}

func (viewport Viewport) GetPos() Vec2 {
	var value Vec2
	valueArg, valueFin := value.wrapped()
	C.iggViewportGetPos(viewport.handle(), valueArg)
	valueFin()
	return value
}

func (viewport Viewport) GetSize() Vec2 {
	var value Vec2
	valueArg, valueFin := value.wrapped()
	C.iggViewportGetSize(viewport.handle(), valueArg)
	valueFin()
	return value
}
