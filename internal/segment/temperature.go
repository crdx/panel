package segment

import (
	"fmt"

	"crdx.org/panel/internal/colour"
	"crdx.org/panel/internal/temperature"
)

type Temperature struct{}

func (Temperature) Optional() bool {
	return true
}

func (Temperature) Thresholds() []float64 {
	return []float64{
		70, // CPU
		70, // GPU
	}
}

func (self Temperature) Values() ([]float64, error) {
	cpu, err := temperature.Cpu()
	if err != nil {
		return nil, err
	}

	temperatures := []float64{cpu}

	// No problem if we can't find a GPU temperature. It may not be relevant to this system.
	if gpu, err := temperature.Gpu(); err == nil {
		temperatures = append(temperatures, gpu)
	}

	return temperatures, nil
}

func (Temperature) Format(value float64) string {
	return fmt.Sprintf("%d", int(value))
}

func (Temperature) Colours() []string {
	return []string{
		colour.Blue,  // CPU
		colour.Green, // GPU
	}
}

func (Temperature) Icon(values []float64) string {
	// https://fontawesome.com/v5.15/icons?d=gallery&m=free

	icons := []string{"", "", "", ""}
	thresholds := []float64{60, 50, 40, 0}

	for i, icon := range icons {
		for _, value := range values {
			if value >= thresholds[i] {
				return icon + "°"
			}
		}
	}

	return ""
}
