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
	// NumCPU returns the number of logical CPUs usable by the current process.
	//
	// The set of available CPUs is checked by querying the operating system at process startup. Changes to operating system CPU allocation after process startup are not reflected.
	Num int
	// GOARCH is the running program's architecture target: one of 386, amd64, arm, s390x, and so on.
	Arch string
	// Process ID
	PID int
}

// Get CPU stats
func CPU() (StatCPU, error) {
	pid := os.Getpid()

	stat := StatCPU{
		Num:  runtime.NumCPU(),
		Arch: runtime.GOARCH,
		PID:  pid,
	}

	sysInfo, err := pidusage.GetStat(pid)
	if err != nil {
		stat.Usage = "error"
		return stat, err
	}

	stat.Usage = fmt.Sprint(math.Round(sysInfo.CPU*100)/100, "%")

	return stat, nil
}
