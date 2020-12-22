#pragma once

#include "Types.h"

#ifdef __cplusplus
extern "C" {
#endif

extern void iggDockSpace(int id, IggVec2 const *size, int flags);
extern void iggSetNextWindowDockID(int id, int cond);
extern int iggGetWindowDockID();
extern IggBool iggIsWindowDocked();

#ifdef __cplusplus
}
#endif
