#pragma once

#include "Types.h"

#ifdef __cplusplus
extern "C" {
#endif

extern void iggDockSpace(int id, IggVec2 const *size, int flags);
extern int iggDockSpaceOverViewport(IggViewport vp, int flags);
extern void iggSetNextWindowDockID(int id, int cond);
extern int iggGetWindowDockID();
extern IggBool iggIsWindowDocked();
extern void iggDockBuilderDockWindow(const char *windowName, int nodeId);
extern int iggDockBuilderAddNode(int nodeId, int flags);
extern void iggDockBuilderRemoveNode(int nodeId);
extern void iggDockBuilderSetNodePos(int nodeId, IggVec2 const *pos);
extern void iggDockBuilderSetNodeSize(int nodeId, IggVec2 const *size);
extern int iggDockBuilderSplitNode(int nodeId, int splitDir, float sizeRatioForNodeAtDir, int *outIdAtDir, int *outIdAtOppositeDir);
extern void iggDockBuilderFinish(int nodeId);

#ifdef __cplusplus
}
#endif
