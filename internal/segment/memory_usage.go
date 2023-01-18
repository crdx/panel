package segment

import (
	"crdx.org/panel/internal/colour"
	"crdx.org/panel/internal/memory"
	"crdx.org/panel/internal/util"
)

type MemoryUsage struct{}

func (MemoryUsage) Optional() bool {
	return false
}

func (MemoryUsage) Thresholds() []float64 {
	info, err := memory.GetUsage()
	if err != nil {
		return []float64{0}
	}

	return []float64{info.Total * 0.8} // 80%
}

func (MemoryUsage) Values() ([]float64, error) {
	info, err := memory.GetUsage()
	if err != nil {
		return nil, err
	}

	return []float64{info.Used}, nil
}

func (MemoryUsage) Format(value float64) string {
	return util.FormatBytes(value, 3)
}

func (MemoryUsage) Colours() []string {
	return []string{colour.Yellow}
}

func (MemoryUsage) Icon(_ []float64) string {
	// https://fontawesome.com/v5.15/icons?d=gallery&m=free
	// ""
	// ""
	// ""
	// ""
	return ""
}
