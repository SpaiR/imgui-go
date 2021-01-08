#pragma once

#include "Types.h"

#ifdef __cplusplus
extern "C" {
#endif

extern void iggSetNextWindowViewport(int viewportId);

extern IggViewport iggGetWindowViewport(void);
extern IggViewport iggGetMainViewport(void);

extern int iggViewportGetID(IggViewport handle);
extern void iggViewportGetPos(IggViewport handle, IggVec2 *value);
extern void iggViewportGetSize(IggViewport handle, IggVec2 *value);

extern void iggViewportGetCenter(IggViewport handle, IggVec2 *value);
extern void iggViewportGetWorkPos(IggViewport handle, IggVec2 *value);
extern void iggViewportGetWorkSize(IggViewport handle, IggVec2 *value);

#ifdef __cplusplus
}
#endif
