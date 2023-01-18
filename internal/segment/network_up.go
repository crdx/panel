package segment

import (
	"crdx.org/panel/internal/colour"
	"crdx.org/panel/internal/network"
	"crdx.org/panel/internal/util"
)

type NetworkUp struct{}

func (NetworkUp) Optional() bool {
	return false
}

func (NetworkUp) Thresholds() []float64 {
	return nil
}

func (NetworkUp) Values() ([]float64, error) {
	usage, err := network.GetUsage()
	return []float64{usage.Tx}, err
}

func (NetworkUp) Format(value float64) string {
	return util.FormatBytes(value, 4)
}

func (NetworkUp) Colours() []string {
	return []string{colour.Orange}
}

func (NetworkUp) Icon(_ []float64) string {
	// https://fontawesome.com/v5.15/icons?d=gallery&m=free
	return "↑"
}
