package temperature

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

func Gpu() (float64, error) {
	stdout, err := exec.Command("nvidia-smi", "--query-gpu=temperature.gpu", "--format=csv,noheader,nounits").Output()
	if err != nil {
		return 0, fmt.Errorf("unable to call nvidia-smi")
	}

	f, err := strconv.ParseFloat(strings.TrimSpace(string(stdout)), 64)
	if err != nil {
		return 0, fmt.Errorf("unable to parse temperature as float")
	}

	return f, nil
}
