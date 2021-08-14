package stats

import (
	"fmt"
	"math"
	"os"
	"runtime"

	"github.com/struCoder/pidusage"
)

type StatCPU struct {
	// Percent of CPU Usage
	Usage string
	// Number of CPUs
	Num int
	// CPU Arch
	Arch string
	// Process ID
	pid int
}

// Get CPU stats
func CPU() (*StatCPU, error) {
	pid := os.Getpid()

	sysInfo, err := pidusage.GetStat(pid)
	if err != nil {
		return nil, err
	}

	stat := StatCPU{
		Usage: fmt.Sprint(math.Round(sysInfo.CPU*100)/100, "%"),
		Num:   runtime.NumCPU(),
		Arch:  runtime.GOARCH,
		pid: pid,
	}

	return &stat, nil
}
