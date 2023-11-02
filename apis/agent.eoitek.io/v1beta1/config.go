package v1beta1

import (
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
)

const (
	JAVAASYNC = 0
	GOPPROF   = 1
	NATIVE    = 2
)

type ProfileKind string

const (
	GO          ProfileKind = "go"
	JAVA        ProfileKind = "java"
	C           ProfileKind = "native"
	UNKNOWN     ProfileKind = "unknown"
	GoPrefix    string      = "go_"
	GoCPU       string      = GoPrefix + "cpu" //profile
	GoGoRoutine string      = GoPrefix + "goroutine"
	GoHeap      string      = GoPrefix + "heap"
	GoAllocs    string      = GoPrefix + "allocs"
	GoMutex     string      = GoPrefix + "mutex"
	GoBlock     string      = GoPrefix + "block"
	JavaPrefix  string      = ""
	JavaCPU     string      = JavaPrefix + "cpu"
	JavaAlloc   string      = JavaPrefix + "alloc"
	JavaLock    string      = JavaPrefix + "lock"
)

type PID uint32

type Java struct {
	StartTime string   `json:"startTime,omitempty"`
	Duration  uint32   `json:"duration,omitempty"`
	Events    []string `json:"events,omitempty"`
	Cpu       uint64   `json:"cpu,omitempty"`
	Alloc     uint64   `json:"alloc,omitempty"`
	Lock      uint64   `json:"lock,omitempty"`
	Jfrsync   string   `json:"jfrsync,omitempty"`
	PID       `json:"pid,omitempty"`
	Service   string            `json:"service,omitempty"`
	Tags      map[string]string `json:"tags,omitempty"`
}

func (j *Java) GetType() int {
	return JAVAASYNC
}

func (j *Java) SetPid(pid PID) {
	j.PID = pid
}

func (j *Java) GetPid() PID {
	return j.PID
}

type Golang struct {
	StartTime         string           `json:"startTime,omitempty"`
	Duration          uint32           `json:"duration,omitempty"`
	Url               string           `json:"url,omitempty"`    // http://127.0.0.1:5000
	Router            string           `json:"router,omitempty"` // http://127.0.0.1:5000
	Port              uint16           `json:"port,omitempty"`   // 5000
	EnabledTypes      []string         `json:"enabled_types"`    // cpu,goroutine,heap,mutex,block
	Interval          map[string]int32 `json:"interval"`         //{"cpu":10,"goroutine":10,"heap":10,"mutex":10,"block":10}
	BasicAuthUser     string           `json:"basicAuthUser,omitempty"`
	BasicAuthPassword string           `json:"basicAuthPassword,omitempty"`
	PID               `json:"pid,omitempty"`
	Service           string            `json:"service,omitempty"`
	Tags              map[string]string `json:"tags,omitempty"`
}

func (g *Golang) GetType() int {
	return GOPPROF
}

func (g *Golang) SetPid(pid PID) {
	g.PID = pid
}

func (g *Golang) GetPid() PID {
	return g.PID
}

type OnCPUConfig struct {
	// Enabled enables on-cpu profiling
	// Default: false
	// +optional
	Enabled *bool `json:"enabled"`

	// Frequency specifies the sampling frequency per second
	// Default: 100
	// +optional
	Frequency *int `json:"frequency"`
}

type OffCPUConfig struct {
	// MaxBlock specifies the amount of time in nanoseconds under which we store traces
	// Default: U64_MAX
	// +optional
	MaxBlock *uint64 `json:"maxblock"`
	// MinBlock specifies the amount of time in nanoseconds over which we store traces
	// Default: 1
	// +optional
	MinBlock *uint64 `json:"minblock"`
	// Kernel threads only (no user threads)
	// Default: false
	// +optional
	KernelOnly *bool `json:"kernelonly"`
	// User threads only (no kernel threads)
	// Default: false
	// +optional
	UserOnly *bool `json:"useronly"`
	// Enabled enables off-cpu profiling
	// Default: false
	// +optional
	Enabled *bool `json:"enabled"`
}

func (c *OnCPUConfig) Normalize() error {
	if *c.Frequency <= 0 {
		return errors.New("the ON_CPU dump frequency should be a positive number")
	}
	return nil
}

func (c *OffCPUConfig) Normalize() error {
	if *c.MinBlock == 0 {
		*c.MinBlock = defaultMinBlock
	}
	if *c.MaxBlock == 0 {
		*c.MaxBlock = defaultMaxBlock
	}

	return nil
}

type Native struct {
	// Native profiling start timestamp
	// Default: time.Now()
	// +optional
	StartTime *string `json:"startTime,omitempty"`
	// Profiler work duration in seconds
	// Default: 60
	// +optional
	Duration *uint32 `json:"duration,omitempty"`
	// On-CPU profiler config
	// +optional
	OnCpu *OnCPUConfig `json:"oncpu,omitempty"`
	// Off-CPU profiler config
	// +optional
	OffCpu *OffCPUConfig `json:"offcpu,omitempty"`
	// Process ID of profiling target
	PID `json:"pid,omitempty"`
}

func (n *Native) GetType() int {
	return NATIVE
}

func (n *Native) SetPid(pid PID) {
	n.PID = pid
}

func (n *Native) GetPid() PID {
	return n.PID
}

func Hash(c ProfileConfig) (string, error) {
	idHash := sha256.New()
	cb, err := json.Marshal(c)
	if err != nil {
		return "", err
	}
	idHash.Write(cb)
	configID := fmt.Sprintf("%x", idHash.Sum(nil))
	return configID, nil
}

func GetKindString(i int) string {
	switch i {
	case JAVAASYNC:
		return string(JAVA)
	case GOPPROF:
		return string(GO)
	case NATIVE:
		return string(C)
	default:
		return string(UNKNOWN)
	}
}
