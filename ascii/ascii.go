package ascii

import (
	. "github.com/klauspost/cpuid/v2"
)

var asm struct {
	valid16      func(*byte, uintptr) int
	validPrint16 func(*byte, uintptr) int
}

func init() {
	if CPU.Supports(AVX, AVX2) {
		asm.valid16 = validAVX2
		asm.validPrint16 = validPrintAVX2
	} else {
		asm.valid16 = valid16
		asm.validPrint16 = validPrint16
	}
}

// https://graphics.stanford.edu/~seander/bithacks.html#HasLessInWord
const (
	hasLessConstL64 = (^uint64(0)) / 255
	hasLessConstR64 = hasLessConstL64 * 128

	hasLessConstL32 = (^uint32(0)) / 255
	hasLessConstR32 = hasLessConstL32 * 128

	hasMoreConstL64 = (^uint64(0)) / 255
	hasMoreConstR64 = hasMoreConstL64 * 128

	hasMoreConstL32 = (^uint32(0)) / 255
	hasMoreConstR32 = hasMoreConstL32 * 128
)

//go:nosplit
func hasLess64(x, n uint64) bool {
	return ((x - (hasLessConstL64 * n)) & ^x & hasLessConstR64) != 0
}

//go:nosplit
func hasLess32(x, n uint32) bool {
	return ((x - (hasLessConstL32 * n)) & ^x & hasLessConstR32) != 0
}

//go:nosplit
func hasMore64(x, n uint64) bool {
	return (((x + (hasMoreConstL64 * (127 - n))) | x) & hasMoreConstR64) != 0
}

//go:nosplit
func hasMore32(x, n uint32) bool {
	return (((x + (hasMoreConstL32 * (127 - n))) | x) & hasMoreConstR32) != 0
}
