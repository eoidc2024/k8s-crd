package v1beta1

import (
	"math"
)

var (
	defaultSampleFrequency = 100
	defaultMinBlock        = uint64(1e6) // 1 millisecond
	defaultMaxBlock        = uint64(math.MaxUint64)
	defaultKernelOnly      = false
	defaultUserOnly        = false
	defaultDuration        = uint32(60)
)
