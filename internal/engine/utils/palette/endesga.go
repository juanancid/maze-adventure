package palette

import "image/color"

// ENDESGA16 is the canonical palette.
// Source: https://lospec.com/palette-list/endesga-16
var ENDESGA16 = struct {
	// Neutrals (dark -> light)
	Ink0   color.RGBA // #0B1320 (near-black)
	Ink1   color.RGBA // #1F2933
	Ink2   color.RGBA // #374151
	Ink3   color.RGBA // #4B5563
	Ink4   color.RGBA // #6B7280
	Ink5   color.RGBA // #9CA3AF
	Paper1 color.RGBA // #E5E7EB (near-white)
	Paper0 color.RGBA // #FFFFFF (white)

	// Hues / accents
	Navy   color.RGBA // #1E3A8A
	Azure  color.RGBA // #2563EB
	Cyan   color.RGBA // #06B6D4
	Green  color.RGBA // #22C55E
	Lime   color.RGBA // #A3E635
	Orange color.RGBA // #F97316
	Red    color.RGBA // #EF4444
	Yellow color.RGBA // #FACC15
}{
	Ink0:   color.RGBA{R: 0x0B, G: 0x13, B: 0x20, A: 0xFF},
	Ink1:   color.RGBA{R: 0x1F, G: 0x29, B: 0x33, A: 0xFF},
	Ink2:   color.RGBA{R: 0x37, G: 0x41, B: 0x51, A: 0xFF},
	Ink3:   color.RGBA{R: 0x4B, G: 0x55, B: 0x63, A: 0xFF},
	Ink4:   color.RGBA{R: 0x6B, G: 0x72, B: 0x80, A: 0xFF},
	Ink5:   color.RGBA{R: 0x9C, G: 0xA3, B: 0xAF, A: 0xFF},
	Paper1: color.RGBA{R: 0xE5, G: 0xE7, B: 0xEB, A: 0xFF},
	Paper0: color.RGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xFF},

	Navy:   color.RGBA{R: 0x1E, G: 0x3A, B: 0x8A, A: 0xFF},
	Azure:  color.RGBA{R: 0x25, G: 0x63, B: 0xEB, A: 0xFF},
	Cyan:   color.RGBA{R: 0x06, G: 0xB6, B: 0xD4, A: 0xFF},
	Green:  color.RGBA{R: 0x22, G: 0xC5, B: 0x5E, A: 0xFF},
	Lime:   color.RGBA{R: 0xA3, G: 0xE6, B: 0x35, A: 0xFF},
	Orange: color.RGBA{R: 0xF9, G: 0x73, B: 0x16, A: 0xFF},
	Red:    color.RGBA{R: 0xEF, G: 0x44, B: 0x44, A: 0xFF},
	Yellow: color.RGBA{R: 0xFA, G: 0xCC, B: 0x15, A: 0xFF},
}
