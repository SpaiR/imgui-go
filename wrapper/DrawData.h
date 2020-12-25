#pragma once

#include "Types.h"

#ifdef __cplusplus
extern "C" {
#endif

extern IggDrawData iggGetDrawData(void);
extern IggBool iggDrawDataValid(IggDrawData handle);
extern void iggDrawDataGetCommandLists(IggDrawData handle, void **handles, int *count);
extern void iggDrawDataScaleClipRects(IggDrawData handle, IggVec2 const *scale);
extern void iggDrawDataGetDisplayPos(IggDrawData handle, IggVec2 *value);
extern void iggDrawDataSetDisplayPos(IggDrawData handle, IggVec2 const *value);
extern void iggDrawDataGetDisplaySize(IggDrawData handle, IggVec2 *value);
extern void iggDrawDataSetDisplaySize(IggDrawData handle, IggVec2 const *value);
extern void iggDrawDataGetFramebufferScale(IggDrawData handle, IggVec2 *value);
extern void iggDrawDataSetFramebufferScale(IggDrawData handle, IggVec2 const *value);

#ifdef __cplusplus
}
#endif
