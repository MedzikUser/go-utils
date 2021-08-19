package stats

import (
	"runtime"

	"github.com/MedzikUser/go-utils/utils"
)

type StatMemory struct {
	// Alloc is bytes of allocated heap objects.
	Alloc string
	// TotalAlloc is cumulative bytes allocated for heap objects.
	//
	// TotalAlloc increases as heap objects are allocated, but unlike Alloc and HeapAlloc, it does not decrease when objects are freed.
	TotalAlloc string
	// Sys is the total bytes of memory obtained from the OS.
	//
	// Sys is the sum of the XSys fields below. Sys measures the virtual address space reserved by the Go runtime for the heap, stacks, and other internal data structures. It's likely that not all of the virtual address space is backed by physical memory at any given moment, though in general it all was at some point.
	Sys string
	// NumGC is the number of completed GC cycles.
	NumGC uint32
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
