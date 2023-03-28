//go:build cgo

package goroutine

import (
	"runtime"
	"strconv"
	"strings"
	"testing"
)

func BenchmarkGetGoroutineID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Id()
	}
}

func legacyGoID() int {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(err)
	}
	return id
}

func BenchmarkGetGoroutineIDLegacy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = legacyGoID()
	}
}
