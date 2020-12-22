#pragma once

#include "Types.h"

#ifdef __cplusplus
extern "C" {
#endif

extern IggViewport iggGetWindowViewport(void);
extern IggViewport iggGetMainViewport(void);

extern int iggViewportGetID(IggViewport handle);
extern void iggViewportGetPos(IggViewport handle, IggVec2 *value);
extern void iggViewportGetSize(IggViewport handle, IggVec2 *value);

#ifdef __cplusplus
}
#endif
