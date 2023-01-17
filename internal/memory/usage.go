package memory

import (
	"os"
	"regexp"
	"strconv"
)

type Info struct {
	Total   float64
	Free    float64
	Caches  float64
	Buffers float64
	Used    float64
}

func GetUsage() (*Info, error) {
	b, err := os.ReadFile("/proc/meminfo")
	if err != nil {
		return nil, err
	}

	meminfo := string(b)

	total := get("MemTotal", meminfo)
	free := get("MemFree", meminfo)
	caches := get("Cached", meminfo) + get("SReclaimable", meminfo)
	buffers := get("Buffers", meminfo)

	// Algorithm taken from `free`
	// https://gitlab.com/procps-ng/procps/-/blob/b652c35f7f4d0fe644df4c756ef96cb4fd08a9f8/proc/meminfo.c#L694-704
	used := total - free - caches - buffers

	if used < 0 {
		used = total - free
	}

	return &Info{
		Total:   total,
		Free:    free,
		Caches:  caches,
		Buffers: buffers,
		Used:    used,
	}, nil
}

func get(key string, meminfo string) float64 {
	re := regexp.MustCompile(key + `:\s*(\d+) kB`)
	matches := re.FindStringSubmatch(meminfo)

	if len(matches) > 1 {
		f, err := strconv.ParseFloat(matches[1], 64)
		if err != nil {
			return 0
		}
		return f * 1024
	}

	return 0
}
