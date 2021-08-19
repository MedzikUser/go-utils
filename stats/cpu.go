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
}

// Get CPU stats
func CPU() (StatCPU, error) {
	pid := os.Getpid()

	stat := StatCPU{
		Num:  runtime.NumCPU(),
		Arch: runtime.GOARCH,
	}

	sysInfo, err := pidusage.GetStat(pid)
	if err != nil {
		stat.Usage = "error"
		return stat, err
	}

	stat.Usage = fmt.Sprint(math.Round(sysInfo.CPU*100)/100, "%")

	return stat, nil
}
