package util_test

import (
	"testing"

	"crdx.org/panel/internal/util"
	"github.com/stretchr/testify/assert"
)

func TestFormatSf(t *testing.T) {
	assert.Equal(t, "1", util.FormatSf(1.234, 0))
	assert.Equal(t, "1", util.FormatSf(1.23, 0))
	assert.Equal(t, "1", util.FormatSf(1.2, 0))
	assert.Equal(t, "1", util.FormatSf(1, 0))
	assert.Equal(t, "12", util.FormatSf(12, 0))
	assert.Equal(t, "1", util.FormatSf(1.234, 1))
	assert.Equal(t, "1", util.FormatSf(1.23, 1))
	assert.Equal(t, "1", util.FormatSf(1.2, 1))
	assert.Equal(t, "1", util.FormatSf(1, 1))
	assert.Equal(t, "12", util.FormatSf(12, 1))
	assert.Equal(t, "1.2", util.FormatSf(1.234, 2))
	assert.Equal(t, "1.2", util.FormatSf(1.23, 2))
	assert.Equal(t, "1.2", util.FormatSf(1.2, 2))
	assert.Equal(t, "1.0", util.FormatSf(1, 2))
	assert.Equal(t, "12", util.FormatSf(12, 2))
	assert.Equal(t, "1.23", util.FormatSf(1.234, 3))
	assert.Equal(t, "1.23", util.FormatSf(1.23, 3))
	assert.Equal(t, "1.20", util.FormatSf(1.2, 3))
	assert.Equal(t, "1.00", util.FormatSf(1, 3))
	assert.Equal(t, "12.0", util.FormatSf(12, 3))
	assert.Equal(t, "1.234", util.FormatSf(1.234, 4))
	assert.Equal(t, "1.230", util.FormatSf(1.23, 4))
	assert.Equal(t, "1.200", util.FormatSf(1.2, 4))
	assert.Equal(t, "1.000", util.FormatSf(1, 4))
	assert.Equal(t, "12.00", util.FormatSf(12, 4))
	assert.Equal(t, "1.2340", util.FormatSf(1.234, 5))
	assert.Equal(t, "1.2300", util.FormatSf(1.23, 5))
	assert.Equal(t, "1.2000", util.FormatSf(1.2, 5))
	assert.Equal(t, "1.0000", util.FormatSf(1, 5))
	assert.Equal(t, "12.000", util.FormatSf(12, 5))
}

func TestFormatBytes(t *testing.T) {
	assert.Equal(t, "0.01", util.FormatBytes(1024*1024*10, 3))
	assert.Equal(t, "0.10", util.FormatBytes(1024*1024*100, 3))
	assert.Equal(t, "1.00", util.FormatBytes(1024*1024*1024, 3))
	assert.Equal(t, "10.0", util.FormatBytes(1024*1024*1024*10, 3))
}

func TestAdjustColour(t *testing.T) {
	assert.Equal(t, "#000000", util.AdjustColour("#000000", -10))
	assert.Equal(t, "#9cc3e4", util.AdjustColour("#a6cdee", -10))
	assert.Equal(t, "#d48955", util.AdjustColour("#de935f", -10))
	assert.Equal(t, "#8fe07c", util.AdjustColour("#99ea86", -10))
	assert.Equal(t, "#efa97a", util.AdjustColour("#f9b384", -10))
	assert.Equal(t, "#da9eef", util.AdjustColour("#e4a8f9", -10))
	assert.Equal(t, "#c25c5c", util.AdjustColour("#cc6666", -10))
	assert.Equal(t, "#d7c86c", util.AdjustColour("#e1d276", -10))
}
