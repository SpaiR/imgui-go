package imgui

// #include "wrapper/Docking.h"
import "C"

type DockNodeFlags int

const (
	DockNodeFlagsNone                   = 0
	DockNodeFlagsKeepAliveOnly          = 1 << 0 // Shared       // Don't display the dockspace node but keep it alive. Windows docked into this dockspace node won't be undocked.
	DockNodeFlagsNoDockingInCentralNode = 1 << 2 // Shared       // Disable docking inside the Central Node, which will be always kept empty.
	DockNodeFlagsPassthruCentralNode    = 1 << 3 // Shared       // Enable passthru dockspace: 1) DockSpace() will render a ImGuiCol_WindowBg background covering everything excepted the Central Node when empty. Meaning the host window should probably use SetNextWindowBgAlpha(0.0f) prior to Begin() when using this. 2) When Central Node is empty: let inputs pass-through + won't display a DockingEmptyBg background. See demo for details.
	DockNodeFlagsNoSplit                = 1 << 4 // Shared/Local // Disable splitting the node into smaller nodes. Useful e.g. when embedding dockspaces into a main root one (the root one may have splitting disabled to reduce confusion). Note: when turned off, existing splits will be preserved.
	DockNodeFlagsNoResize               = 1 << 5 // Shared/Local // Disable resizing node using the splitter/separators. Useful with programatically setup dockspaces.
	DockNodeFlagsAutoHideTabBar         = 1 << 6 // Shared/Local // Tab bar will automatically hide when there is a single window in the dock node.
)

func DockSpaceV(id int, size Vec2, flags DockNodeFlags) {
	sizeArg, _ := size.wrapped()
	C.iggDockSpace(C.int(id), sizeArg, C.int(flags))
}

func DockSpace(id int) {
	DockSpaceV(id, Vec2{0, 0}, DockNodeFlagsNone)
}

func SetNextWindowDockIDV(id int, cond Condition) {
	C.iggSetNextWindowDockID(C.int(id), C.int(cond))
}

func SetNextWindowDockID(id int) {
	SetNextWindowDockIDV(id, 0)
}

func GetWindowDockID() int {
	return int(C.iggGetWindowDockID())
}

func IsWindowDocked() bool {
	return C.iggIsWindowDocked() != 0
}
