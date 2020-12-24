#include "ConfiguredImGui.h"

#include "Docking.h"
#include "WrapperConverter.h"

void iggDockSpace(int id, IggVec2 const *size, int flags)
{
    Vec2Wrapper sizeArg(size);
    ImGui::DockSpace(id, *sizeArg, flags);
}

int iggDockSpaceOverViewport(IggViewport vp, int flags)
{
    ImGuiViewport *viewport = reinterpret_cast<ImGuiViewport *>(vp);  
    return ImGui::DockSpaceOverViewport(viewport, flags);
}

void iggSetNextWindowDockID(int id, int cond)
{
    ImGui::SetNextWindowDockID(id, cond);
}

int iggGetWindowDockID()
{
    return ImGui::GetWindowDockID();
}

IggBool iggIsWindowDocked()
{
    return ImGui::IsWindowDocked() ? 1 : 0;
}

void iggDockBuilderDockWindow(const char *windowName, int nodeId)
{
    ImGui::DockBuilderDockWindow(windowName, nodeId);
}

int iggDockBuilderAddNode(int nodeId, int flags)
{
    return ImGui::DockBuilderAddNode(nodeId, flags);
}

void iggDockBuilderRemoveNode(int nodeId)
{
    ImGui::DockBuilderRemoveNode(nodeId);
}

void iggDockBuilderSetNodePos(int nodeId, IggVec2 const *pos)
{
    Vec2Wrapper posArg(pos);
    ImGui::DockBuilderSetNodePos(nodeId, *posArg);
}

void iggDockBuilderSetNodeSize(int nodeId, IggVec2 const *size)
{
    Vec2Wrapper sizeArg(size);
    ImGui::DockBuilderSetNodeSize(nodeId, *sizeArg);
}

int iggDockBuilderSplitNode(int nodeId, int splitDir, float sizeRatioForNodeAtDir, int *outIdAtDir, int *outIdAtOppositeDir)
{
    return ImGui::DockBuilderSplitNode(nodeId, splitDir, sizeRatioForNodeAtDir, (ImGuiID*)outIdAtDir, (ImGuiID*)outIdAtOppositeDir);
}

void iggDockBuilderFinish(int nodeId)
{
    ImGui::DockBuilderFinish(nodeId);
}
