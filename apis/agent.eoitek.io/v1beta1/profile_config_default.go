package v1beta1

import (
	"math"
	"time"

	"code.eoitek.net/monitor/dc/k8s-crd/utils"
)

var (
	defaultSampleFrequency = 100
	defaultMinBlock        = uint64(1e6) // 1 millisecond
	defaultMaxBlock        = uint64(math.MaxUint64)
	defaultKernelOnly      = false
	defaultUserOnly        = false
	defaultDuration        = uint32(60)
)

func NewDefaultNativeConfig() *Native {
	c := &Native{}
	DefaultNativeConfig(c)
	return c
}

func NewDefaultNativeConfigFrom(n *Native) *Native {
	c := &Native{
		StartTime: n.StartTime,
		Duration:  n.Duration,
		PID:       n.PID,
	}
	if n.OnCpu != nil {
		c.OnCpu = &OnCPUConfig{
			Frequency: n.OnCpu.Frequency,
			Enabled:   n.OnCpu.Enabled,
		}
	}
	if n.OffCpu != nil {
		c.OffCpu = &OffCPUConfig{
			MaxBlock:   n.OffCpu.MaxBlock,
			MinBlock:   n.OffCpu.MinBlock,
			KernelOnly: n.OffCpu.KernelOnly,
			UserOnly:   n.OffCpu.UserOnly,
			Enabled:    n.OffCpu.Enabled,
		}
	}
	DefaultNativeConfig(c)
	return c
}

func DefaultNativeConfig(c *Native) {
	if c.OnCpu != nil {
		defualtOnCPUConfig(c.OnCpu)
	}
	if c.OffCpu != nil {
		defaultOffCPUConfig(c.OffCpu)
	}
	utils.DefaultIfUnset[string](&c.StartTime, time.Now().Format(time.RFC3339))
	utils.DefaultIfUnset[uint32](&c.Duration, defaultDuration)
}

func defualtOnCPUConfig(c *OnCPUConfig) {
	if c.Enabled == nil {
		c.Enabled = utils.NewBoolPointer(false)
		return
	}
	utils.DefaultIfUnset[int](&c.Frequency, defaultSampleFrequency)
}

func defaultOffCPUConfig(c *OffCPUConfig) {
	if c.Enabled == nil {
		c.Enabled = utils.NewBoolPointer(false)
		return
	}
	utils.DefaultIfUnset[uint64](&c.MaxBlock, defaultMaxBlock)
	utils.DefaultIfUnset[uint64](&c.MinBlock, defaultMinBlock)
	utils.DefaultBoolIfUnset(&c.KernelOnly, defaultKernelOnly)
	utils.DefaultBoolIfUnset(&c.UserOnly, defaultUserOnly)
}
