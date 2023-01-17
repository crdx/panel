package temperature

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

func Gpu() (float64, error) {
	nvidiaSmi, err := exec.LookPath("nvidia-smi")
	if err != nil {
		return 0, fmt.Errorf("unable to find nvidia-smi")
	}

	stdout, err := exec.Command(nvidiaSmi, "--query-gpu=temperature.gpu", "--format=csv,noheader,nounits").Output()
	if err != nil {
		return 0, fmt.Errorf("unable to call nvidia-smi")
	}

	f, err := strconv.ParseFloat(strings.TrimSpace(string(stdout)), 64)
	if err != nil {
		return 0, fmt.Errorf("unable to parse temperature as float")
	}

	return f, nil
}
