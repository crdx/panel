package network

import (
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Info struct {
	Tx float64
	Rx float64
}

func GetUsage() (*Info, error) {
	re := regexp.MustCompile(`(?:enp|eno|wlp)\w+:\s*(\d+)\s*\d+\s*\d+\s*\d+\s*\d+\s*\d+\s*\d+\s*\d+\s*(\d+)`)

	b, err := os.ReadFile("/proc/net/dev")
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(b), "\n")

	var rx, tx float64

	for _, line := range lines {
		if matches := re.FindStringSubmatch(line); len(matches) > 0 {
			if f, err := strconv.ParseFloat(matches[1], 64); err == nil {
				rx += f
			}

			if f, err := strconv.ParseFloat(matches[2], 64); err == nil {
				tx += f
			}
		}
	}

	return &Info{
		Tx: tx,
		Rx: rx,
	}, nil
}
