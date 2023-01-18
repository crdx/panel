package main

import (
	"fmt"
	"strings"

	"crdx.org/panel/internal/util"
)

type Segment interface {
	Optional() bool
	Thresholds() []float64
	Values() ([]float64, error)
	Format(value float64) string
	Colours() []string
	Icon(values []float64) string
}

const (
	separator   = "<span fgcolor='#aaaaaa'>/</span>"
	alertColour = "red"
)

func renderString(s string, colour string) string {
	return fmt.Sprintf("<span fgcolor='%s'>%s</span>", colour, s)
}

func renderValue(threshold float64, value float64, formattedValue string, icon string, colour string) string {
	var renderedValue string
	var iconColour string

	if threshold > 0 && value >= threshold {
		colour = alertColour
		iconColour = alertColour
	} else {
		iconColour = util.AdjustColour(colour, -10)
	}

	renderedValue = renderString(formattedValue, colour)

	if icon != "" {
		renderedValue = fmt.Sprintf("%s %s", renderString(icon, iconColour), renderedValue)
	}

	return renderedValue
}

func renderSegment(segment Segment) string {
	var renderedValues []string

	values, err := segment.Values()
	if err != nil {
		return ""
	}

	for i, value := range values {
		if value == 0 && segment.Optional() {
			continue
		}

		formattedValue := segment.Format(value)
		colour := util.GetAtIndex(segment.Colours(), i)
		threshold := util.GetAtIndex(segment.Thresholds(), i)

		// Only show an icon for the first value.
		var icon string
		if i == 0 {
			icon = segment.Icon(values)
		}

		renderedValue := renderValue(threshold, value, formattedValue, icon, colour)
		renderedValues = append(renderedValues, renderedValue)
	}

	return strings.Join(renderedValues, separator)
}

func renderSegments(segments []Segment) string {
	var renderedSegments []string

	for _, segment := range segments {
		if renderedSegment := renderSegment(segment); renderedSegment != "" {
			renderedSegments = append(renderedSegments, renderedSegment)
		}
	}

	return strings.Join(renderedSegments, "  ")
}
