package segment

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"

	"crdx.org/panel/internal/colour"
	"crdx.org/panel/internal/util"
)

type CpuLoad struct{}

func (CpuLoad) Optional() bool {
	return false
}

func (CpuLoad) Thresholds() []float64 {
	n := float64(runtime.NumCPU())

	if n == 1 {
		return []float64{n}
	} else {
		return []float64{n / 2}
	}
}

func (CpuLoad) Values() ([]float64, error) {
	b, err := os.ReadFile("/proc/loadavg")
	if err != nil {
		return nil, fmt.Errorf("unable to read load average")
	}

	s, _, _ := strings.Cut(string(b), " ")

	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return nil, fmt.Errorf("unable to parse load average as float")
	}

	return []float64{f}, nil
}

func (CpuLoad) Format(value float64) string {
	return util.FormatSf(value, 3)
}

func (CpuLoad) Colours() []string {
	return []string{colour.Purple}
}

func (CpuLoad) Icon(_ []float64) string {
	return ""
}
