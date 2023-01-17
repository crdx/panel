package segment

import (
	"github.com/crdx/panel/internal/colour"
	"github.com/crdx/panel/internal/network"
	"github.com/crdx/panel/internal/util"
)

type NetworkDown struct{}

func (NetworkDown) Optional() bool {
	return false
}

func (NetworkDown) Thresholds() []float64 {
	return nil
}

func (NetworkDown) Values() ([]float64, error) {
	usage, err := network.GetUsage()
	return []float64{usage.Rx}, err
}

func (NetworkDown) Format(value float64) string {
	return util.FormatBytes(value, 4)
}

func (NetworkDown) Colours() []string {
	return []string{colour.Orange}
}

func (NetworkDown) Icon(_ []float64) string {
	// https://fontawesome.com/v5.15/icons?d=gallery&m=free
	return "↓"
}
