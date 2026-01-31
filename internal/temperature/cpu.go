package temperature

import (
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
)

func Cpu() (float64, error) {
	sensors := []string{
		"coretemp-isa-0000",
		"k10temp-pci-00c3",
	}

	var stdout []byte
	for _, sensor := range sensors {
		var err error
		if stdout, err = exec.Command("sensors", sensor).Output(); err == nil { //nolint:gosec // sensor is from a hardcoded slice
			break
		}
	}

	if len(stdout) == 0 {
		return 0, fmt.Errorf("unable to get sensor data")
	}

	regexes := []*regexp.Regexp{
		regexp.MustCompile(`Tdie:\s*\+(.*?)°C`),
		regexp.MustCompile(`Package id 0:\s*\+(.*?)°C`),
	}

	for _, re := range regexes {
		if matches := re.FindStringSubmatch(string(stdout)); len(matches) > 1 {
			if f, err := strconv.ParseFloat(matches[1], 64); err == nil {
				return f, nil
			}
		}
	}

	return 0, fmt.Errorf("unable to find temperature")
}
