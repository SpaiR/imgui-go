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

	DockNodeFlagsDockSpace           = 1 << 10 // Local, Saved  // A dockspace is a node that occupy space within an existing user window. Otherwise the node is floating and create its own window.
	DockNodeFlagsCentralNode         = 1 << 11 // Local, Saved  // The central node has 2 main properties: stay visible when empty, only use "remaining" spaces from its neighbor.
	DockNodeFlagsNoTabBar            = 1 << 12 // Local, Saved  // Tab bar is completely unavailable. No triangle in the corner to enable it back.
	DockNodeFlagsHiddenTabBar        = 1 << 13 // Local, Saved  // Tab bar is hidden, with a triangle in the corner to show it again (NB: actual tab-bar instance may be destroyed as this is only used for single-window tab bar)
	DockNodeFlagsNoWindowMenuButton  = 1 << 14 // Local, Saved  // Disable window/docking menu (that one that appears instead of the collapse button)
	DockNodeFlagsNoCloseButton       = 1 << 15 // Local, Saved  //
	DockNodeFlagsNoDocking           = 1 << 16 // Local, Saved  // Disable any form of docking in this dockspace or individual node. (On a whole dockspace, this pretty much defeat the purpose of using a dockspace at all). Note: when turned on, existing docked nodes will be preserved.
	DockNodeFlagsNoDockingSplitMe    = 1 << 17 // [EXPERIMENTAL] Prevent another window/node from splitting this node.
	DockNodeFlagsNoDockingSplitOther = 1 << 18 // [EXPERIMENTAL] Prevent this node from splitting another window/node.
	DockNodeFlagsNoDockingOverMe     = 1 << 19 // [EXPERIMENTAL] Prevent another window/node to be docked over this node.
	DockNodeFlagsNoDockingOverOther  = 1 << 20 // [EXPERIMENTAL] Prevent this node to be docked over another window/node.
	DockNodeFlagsNoResizeX           = 1 << 21 // [EXPERIMENTAL]
	DockNodeFlagsNoResizeY           = 1 << 22 // [EXPERIMENTAL]
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

func DockBuilderDockWindow(windowName string, nodeId int) {
	windowNameArg, windowNameFin := wrapString(windowName)
	defer windowNameFin()
	C.iggDockBuilderDockWindow(windowNameArg, C.int(nodeId))
}

func DockBuilderAddNodeV(nodeId int, flags DockNodeFlags) int {
	return int(C.iggDockBuilderAddNode(C.int(nodeId), C.int(flags)))
}

func DockBuilderAddNode() int {
	return DockBuilderAddNodeV(0, 0)
}

func DockBuilderRemoveNode(nodeId int) {
	C.iggDockBuilderRemoveNode(C.int(nodeId))
}

func DockBuilderSetNodePos(nodeId int, pos Vec2) {
	posArg, _ := pos.wrapped()
	C.iggDockBuilderSetNodePos(C.int(nodeId), posArg)
}

func DockBuilderSetNodeSize(nodeId int, size Vec2) {
	sizeArg, _ := size.wrapped()
	C.iggDockBuilderSetNodeSize(C.int(nodeId), sizeArg)
}

type Dir int

const (
	DirNone  = -1
	DirLeft  = 0
	DirRight = 1
	DirUp    = 2
	DirDown  = 3
	DirCount = 5
)

func DockBuilderSplitNode(nodeId int, splitDir Dir, sizeRatioForNodeAtDir float32, outIdAtDir *int32, outIdAtOppositeDir *int32) int {
	outIdAtDirArg, outIdAtDirFin := wrapInt32(outIdAtDir)
	defer outIdAtDirFin()
	outIdAtOppositeDirArg, outIdAtOppositeDirFin := wrapInt32(outIdAtOppositeDir)
	defer outIdAtOppositeDirFin()
	return int(C.iggDockBuilderSplitNode(C.int(nodeId), C.int(splitDir), C.float(sizeRatioForNodeAtDir), outIdAtDirArg, outIdAtOppositeDirArg))
}

func DockBuilderFinish(nodeId int) {
	C.iggDockBuilderFinish(C.int(nodeId))
}
