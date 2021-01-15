#pragma once

#include "Types.h"

#ifdef __cplusplus
extern "C" {
#endif

extern IggStorage iggGetStateStorage();

extern void iggSetAllInt(IggStorage handle, int val);

#ifdef __cplusplus
}
#endif
