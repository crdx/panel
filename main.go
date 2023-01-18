package main

import (
	"fmt"

	"crdx.org/panel/internal/segment"
)

func main() {
	segments := []Segment{
		segment.CpuLoad{},
		segment.Temperature{},
		segment.NetworkDown{},
		segment.NetworkUp{},
		segment.MemoryUsage{},
	}

	command := "kitty --start-as=fullscreen sudo htop --sort-key=PERCENT_CPU"
	rendered := renderSegments(segments)

	fmt.Printf("<txt> %s </txt>\n", rendered)
	fmt.Printf("<txtclick>%s</txtclick>\n", command)
	fmt.Printf("<tool></tool>\n")
}
