#include "ConfiguredImGui.h"

#include "Storage.h"
#include "WrapperConverter.h"

IggStorage iggGetStateStorage()
{
    return reinterpret_cast<IggStorage>(ImGui::GetStateStorage());
}

void iggSetAllInt(IggStorage handle, int val)
{
    ImGuiStorage *storage = reinterpret_cast<ImGuiStorage *>(handle);
    storage->SetAllInt(val);
}
