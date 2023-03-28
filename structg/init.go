package structg

import (
	"debug/elf"
)

func init() {
	exe, err := elf.Open("/proc/self/exe")
	if err != nil {
		panic(err)
	}
	defer exe.Close()

	offset, err := getGOffsetFromELF(exe)
	if err != nil {
		panic(err)
	}
	gOffset = offset

	err = initGFieldOffsets(exe)
	if err != nil {
		panic(err)
	}
}
