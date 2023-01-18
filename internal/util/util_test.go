package util_test

import (
	"testing"

	"crdx.org/assert"
	"crdx.org/panel/internal/util"
)

func TestFormatSf(t *testing.T) {
	assert.Equal(t, util.FormatSf(1.234, 0), "1")
	assert.Equal(t, util.FormatSf(1.23, 0), "1")
	assert.Equal(t, util.FormatSf(1.2, 0), "1")
	assert.Equal(t, util.FormatSf(1, 0), "1")
	assert.Equal(t, util.FormatSf(12, 0), "12")
	assert.Equal(t, util.FormatSf(1.234, 1), "1")
	assert.Equal(t, util.FormatSf(1.23, 1), "1")
	assert.Equal(t, util.FormatSf(1.2, 1), "1")
	assert.Equal(t, util.FormatSf(1, 1), "1")
	assert.Equal(t, util.FormatSf(12, 1), "12")
	assert.Equal(t, util.FormatSf(1.234, 2), "1.2")
	assert.Equal(t, util.FormatSf(1.23, 2), "1.2")
	assert.Equal(t, util.FormatSf(1.2, 2), "1.2")
	assert.Equal(t, util.FormatSf(1, 2), "1.0")
	assert.Equal(t, util.FormatSf(12, 2), "12")
	assert.Equal(t, util.FormatSf(1.234, 3), "1.23")
	assert.Equal(t, util.FormatSf(1.23, 3), "1.23")
	assert.Equal(t, util.FormatSf(1.2, 3), "1.20")
	assert.Equal(t, util.FormatSf(1, 3), "1.00")
	assert.Equal(t, util.FormatSf(12, 3), "12.0")
	assert.Equal(t, util.FormatSf(1.234, 4), "1.234")
	assert.Equal(t, util.FormatSf(1.23, 4), "1.230")
	assert.Equal(t, util.FormatSf(1.2, 4), "1.200")
	assert.Equal(t, util.FormatSf(1, 4), "1.000")
	assert.Equal(t, util.FormatSf(12, 4), "12.00")
	assert.Equal(t, util.FormatSf(1.234, 5), "1.2340")
	assert.Equal(t, util.FormatSf(1.23, 5), "1.2300")
	assert.Equal(t, util.FormatSf(1.2, 5), "1.2000")
	assert.Equal(t, util.FormatSf(1, 5), "1.0000")
	assert.Equal(t, util.FormatSf(12, 5), "12.000")
}

func TestFormatBytes(t *testing.T) {
	assert.Equal(t, util.FormatBytes(1024*1024*10, 3), "0.01")
	assert.Equal(t, util.FormatBytes(1024*1024*100, 3), "0.10")
	assert.Equal(t, util.FormatBytes(1024*1024*1024, 3), "1.00")
	assert.Equal(t, util.FormatBytes(1024*1024*1024*10, 3), "10.0")
}

func TestAdjustColour(t *testing.T) {
	assert.Equal(t, util.AdjustColour("#000000", -10), "#000000")
	assert.Equal(t, util.AdjustColour("#a6cdee", -10), "#9cc3e4")
	assert.Equal(t, util.AdjustColour("#de935f", -10), "#d48955")
	assert.Equal(t, util.AdjustColour("#99ea86", -10), "#8fe07c")
	assert.Equal(t, util.AdjustColour("#f9b384", -10), "#efa97a")
	assert.Equal(t, util.AdjustColour("#e4a8f9", -10), "#da9eef")
	assert.Equal(t, util.AdjustColour("#cc6666", -10), "#c25c5c")
	assert.Equal(t, util.AdjustColour("#e1d276", -10), "#d7c86c")
}
