//go:build cgo

package goroutine

import (
	"debug/elf"
	"runtime"
	"unsafe"

	"github.com/YangKeao/stdext/structg"
)

// Id returns the id of current goroutine
func Id() uint64 {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	addr := structg.GetGAddr()

	addr += uintptr(structg.FieldOffsets.GoID)
	return *(*uint64)(unsafe.Pointer(addr))
}

func init() {
	exe, err := elf.Open("/proc/self/exe")
	if err != nil {
		panic(err)
	}
	defer exe.Close()
}
