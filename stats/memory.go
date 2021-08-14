package stats

import (
	"runtime"

	"github.com/MedzikUser/go-utils/utils"
)

type StatMemory struct {
	Alloc      string
	TotalAlloc string
	Sys        string
	NumGC      uint32
}

// Get Memory stats
func Memory() StatMemory {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	stat := StatMemory{
		Alloc:      utils.Bytes(m.Alloc),
		TotalAlloc: utils.Bytes(m.TotalAlloc),
		Sys:        utils.Bytes(m.Sys),
		NumGC:      m.NumGC,
	}

	return stat
}
