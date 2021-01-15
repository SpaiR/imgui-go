package imgui

// #include "wrapper/Storage.h"
import "C"

type Storage uintptr

func (storage Storage) handle() C.IggStorage {
	return C.IggStorage(storage)
}

func StateStorage() Storage {
	return Storage(C.iggGetStateStorage())
}

func (storage Storage) SetAllInt(val int) {
	C.iggSetAllInt(storage.handle(), C.int(val))
}
