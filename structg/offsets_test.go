package structg

import (
	"debug/elf"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetGOffsetFromELF(t *testing.T) {
	exe, err := elf.Open("/proc/self/exe")
	require.NoError(t, err)
	defer exe.Close()

	offset, err := getGOffsetFromELF(exe)
	require.NoError(t, err)

	require.Greater(t, offset, uint64(math.MaxInt32), "offset should represent a negative number")
}

func TestGetGAddr(t *testing.T) {
	exe, err := elf.Open("/proc/self/exe")
	require.NoError(t, err)
	defer exe.Close()

	_, err = getGOffsetFromELF(exe)
	require.NoError(t, err)
}

func TestGetGIDOffset(t *testing.T) {
	exe, err := elf.Open("/proc/self/exe")
	require.NoError(t, err)
	defer exe.Close()

	err = initGFieldOffsets(exe)
	require.NoError(t, err)
}
