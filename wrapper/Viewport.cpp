#include "ConfiguredImGui.h"

#include "Viewport.h"
#include "WrapperConverter.h"

void iggSetNextWindowViewport(int viewportId)
{
    ImGui::SetNextWindowViewport(viewportId);
}

IggViewport iggGetWindowViewport()
{
    return reinterpret_cast<IggViewport>(ImGui::GetWindowViewport());
}

IggViewport iggGetMainViewport()
{
    return reinterpret_cast<IggViewport>(ImGui::GetMainViewport());
}

int iggViewportGetID(IggViewport handle)
{
    ImGuiViewport *viewport = reinterpret_cast<ImGuiViewport *>(handle);
    return viewport->ID;
}

void iggViewportGetPos(IggViewport handle, IggVec2 *value)
{
    ImGuiViewport *viewport = reinterpret_cast<ImGuiViewport *>(handle);
    exportValue(*value, viewport->Pos);
}

void iggViewportGetSize(IggViewport handle, IggVec2 *value)
{
    ImGuiViewport *viewport = reinterpret_cast<ImGuiViewport *>(handle);
    exportValue(*value, viewport->Size);
}
