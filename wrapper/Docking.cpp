#include "ConfiguredImGui.h"

#include "Docking.h"
#include "WrapperConverter.h"

void iggDockSpace(int id, IggVec2 const *size, int flags)
{
    Vec2Wrapper sizeArg(size);
    ImGui::DockSpace(id, *sizeArg, flags);
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
