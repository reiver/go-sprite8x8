package sprite8x8_test

import (
	"github.com/reiver/go-sprite8x8"

	"github.com/reiver/go-palette2048"
	"github.com/reiver/go-palette2048_tia"

	"image/color"

	"testing"
)

func TestPaletted_Set(t *testing.T) {

		var palette palette2048.Slice = palette2048.Slice(palette2048_tia.Palette[:])

		const category = "does not matter category"
		const id = 7

		var pix [8*8*1]uint8
		var sprite sprite8x8.Paletted = sprite8x8.Paletted{
			Pix: pix[:],
			Palette: palette,
			Category: category,
			ID: id,
		}

		{
			sprite.Set(0,0, color.RGBA{  0,  0,  0, 255})
			sprite.Set(1,0, color.RGBA{255,255,255, 255})
			sprite.Set(2,0, color.RGBA{  0,  0,  0, 255})
			sprite.Set(3,0, color.RGBA{255,255,255, 255})
			sprite.Set(4,0, color.RGBA{  0,  0,  0, 255})
			sprite.Set(5,0, color.RGBA{255,255,255, 255})
			sprite.Set(6,0, color.RGBA{  0,  0,  0, 255})
			sprite.Set(7,0, color.RGBA{255,255,255, 255})

			sprite.Set(0,1, color.RGBA{255,  0,  0, 255})
			sprite.Set(1,1, color.RGBA{255,  0,  0, 255})
			sprite.Set(2,1, color.RGBA{255,  0,  0, 255})
			sprite.Set(3,1, color.RGBA{255,  0,  0, 255})
			sprite.Set(4,1, color.RGBA{255,  0,  0, 255})
			sprite.Set(5,1, color.RGBA{255,  0,  0, 255})
			sprite.Set(6,1, color.RGBA{255,  0,  0, 255})
			sprite.Set(7,1, color.RGBA{255,  0,  0, 255})

			sprite.Set(0,2, color.RGBA{255,165,  0, 255})
			sprite.Set(1,2, color.RGBA{255,165,  0, 255})
			sprite.Set(2,2, color.RGBA{255,165,  0, 255})
			sprite.Set(3,2, color.RGBA{255,165,  0, 255})
			sprite.Set(4,2, color.RGBA{255,165,  0, 255})
			sprite.Set(5,2, color.RGBA{255,165,  0, 255})
			sprite.Set(6,2, color.RGBA{255,165,  0, 255})
			sprite.Set(7,2, color.RGBA{255,165,  0, 255})

			sprite.Set(0,3, color.RGBA{255,255,  0, 255})
			sprite.Set(1,3, color.RGBA{255,255,  0, 255})
			sprite.Set(2,3, color.RGBA{255,255,  0, 255})
			sprite.Set(3,3, color.RGBA{255,255,  0, 255})
			sprite.Set(4,3, color.RGBA{255,255,  0, 255})
			sprite.Set(5,3, color.RGBA{255,255,  0, 255})
			sprite.Set(6,3, color.RGBA{255,255,  0, 255})
			sprite.Set(7,3, color.RGBA{255,255,  0, 255})

			sprite.Set(0,4, color.RGBA{  0,255,  0, 255})
			sprite.Set(1,4, color.RGBA{  0,255,  0, 255})
			sprite.Set(2,4, color.RGBA{  0,255,  0, 255})
			sprite.Set(3,4, color.RGBA{  0,255,  0, 255})
			sprite.Set(4,4, color.RGBA{  0,255,  0, 255})
			sprite.Set(5,4, color.RGBA{  0,255,  0, 255})
			sprite.Set(6,4, color.RGBA{  0,255,  0, 255})
			sprite.Set(7,4, color.RGBA{  0,255,  0, 255})

			sprite.Set(0,5, color.RGBA{  0,  0,255, 255})
			sprite.Set(1,5, color.RGBA{  0,  0,255, 255})
			sprite.Set(2,5, color.RGBA{  0,  0,255, 255})
			sprite.Set(3,5, color.RGBA{  0,  0,255, 255})
			sprite.Set(4,5, color.RGBA{  0,  0,255, 255})
			sprite.Set(5,5, color.RGBA{  0,  0,255, 255})
			sprite.Set(6,5, color.RGBA{  0,  0,255, 255})
			sprite.Set(7,5, color.RGBA{  0,  0,255, 255})

			sprite.Set(0,6, color.RGBA{ 75,  0,130, 255})
			sprite.Set(1,6, color.RGBA{ 75,  0,130, 255})
			sprite.Set(2,6, color.RGBA{ 75,  0,130, 255})
			sprite.Set(3,6, color.RGBA{ 75,  0,130, 255})
			sprite.Set(4,6, color.RGBA{ 75,  0,130, 255})
			sprite.Set(5,6, color.RGBA{ 75,  0,130, 255})
			sprite.Set(6,6, color.RGBA{ 75,  0,130, 255})
			sprite.Set(7,6, color.RGBA{ 75,  0,130, 255})

			sprite.Set(0,7, color.RGBA{238,130,238, 255})
			sprite.Set(1,7, color.RGBA{238,130,238, 255})
			sprite.Set(2,7, color.RGBA{238,130,238, 255})
			sprite.Set(3,7, color.RGBA{238,130,238, 255})
			sprite.Set(4,7, color.RGBA{238,130,238, 255})
			sprite.Set(5,7, color.RGBA{238,130,238, 255})
			sprite.Set(6,7, color.RGBA{238,130,238, 255})
			sprite.Set(7,7, color.RGBA{238,130,238, 255})
		}

		var indexes [8][8]uint8
		{
			indexes[0][0] = sprite.ColorIndexAt(0,0)
			indexes[1][0] = sprite.ColorIndexAt(1,0)
			indexes[2][0] = sprite.ColorIndexAt(2,0)
			indexes[3][0] = sprite.ColorIndexAt(3,0)
			indexes[4][0] = sprite.ColorIndexAt(4,0)
			indexes[5][0] = sprite.ColorIndexAt(5,0)
			indexes[6][0] = sprite.ColorIndexAt(6,0)
			indexes[7][0] = sprite.ColorIndexAt(7,0)

			indexes[0][1] = sprite.ColorIndexAt(0,1)
			indexes[1][1] = sprite.ColorIndexAt(1,1)
			indexes[2][1] = sprite.ColorIndexAt(2,1)
			indexes[3][1] = sprite.ColorIndexAt(3,1)
			indexes[4][1] = sprite.ColorIndexAt(4,1)
			indexes[5][1] = sprite.ColorIndexAt(5,1)
			indexes[6][1] = sprite.ColorIndexAt(6,1)
			indexes[7][1] = sprite.ColorIndexAt(7,1)

			indexes[0][2] = sprite.ColorIndexAt(0,2)
			indexes[1][2] = sprite.ColorIndexAt(1,2)
			indexes[2][2] = sprite.ColorIndexAt(2,2)
			indexes[3][2] = sprite.ColorIndexAt(3,2)
			indexes[4][2] = sprite.ColorIndexAt(4,2)
			indexes[5][2] = sprite.ColorIndexAt(5,2)
			indexes[6][2] = sprite.ColorIndexAt(6,2)
			indexes[7][2] = sprite.ColorIndexAt(7,2)

			indexes[0][3] = sprite.ColorIndexAt(0,3)
			indexes[1][3] = sprite.ColorIndexAt(1,3)
			indexes[2][3] = sprite.ColorIndexAt(2,3)
			indexes[3][3] = sprite.ColorIndexAt(3,3)
			indexes[4][3] = sprite.ColorIndexAt(4,3)
			indexes[5][3] = sprite.ColorIndexAt(5,3)
			indexes[6][3] = sprite.ColorIndexAt(6,3)
			indexes[7][3] = sprite.ColorIndexAt(7,3)

			indexes[0][4] = sprite.ColorIndexAt(0,4)
			indexes[1][4] = sprite.ColorIndexAt(1,4)
			indexes[2][4] = sprite.ColorIndexAt(2,4)
			indexes[3][4] = sprite.ColorIndexAt(3,4)
			indexes[4][4] = sprite.ColorIndexAt(4,4)
			indexes[5][4] = sprite.ColorIndexAt(5,4)
			indexes[6][4] = sprite.ColorIndexAt(6,4)
			indexes[7][4] = sprite.ColorIndexAt(7,4)

			indexes[0][5] = sprite.ColorIndexAt(0,5)
			indexes[1][5] = sprite.ColorIndexAt(1,5)
			indexes[2][5] = sprite.ColorIndexAt(2,5)
			indexes[3][5] = sprite.ColorIndexAt(3,5)
			indexes[4][5] = sprite.ColorIndexAt(4,5)
			indexes[5][5] = sprite.ColorIndexAt(5,5)
			indexes[6][5] = sprite.ColorIndexAt(6,5)
			indexes[7][5] = sprite.ColorIndexAt(7,5)

			indexes[0][6] = sprite.ColorIndexAt(0,6)
			indexes[1][6] = sprite.ColorIndexAt(1,6)
			indexes[2][6] = sprite.ColorIndexAt(2,6)
			indexes[3][6] = sprite.ColorIndexAt(3,6)
			indexes[4][6] = sprite.ColorIndexAt(4,6)
			indexes[5][6] = sprite.ColorIndexAt(5,6)
			indexes[6][6] = sprite.ColorIndexAt(6,6)
			indexes[7][6] = sprite.ColorIndexAt(7,6)

			indexes[0][7] = sprite.ColorIndexAt(0,7)
			indexes[1][7] = sprite.ColorIndexAt(1,7)
			indexes[2][7] = sprite.ColorIndexAt(2,7)
			indexes[3][7] = sprite.ColorIndexAt(3,7)
			indexes[4][7] = sprite.ColorIndexAt(4,7)
			indexes[5][7] = sprite.ColorIndexAt(5,7)
			indexes[6][7] = sprite.ColorIndexAt(6,7)
			indexes[7][7] = sprite.ColorIndexAt(7,7)
		}
		{
			for x:=0; x<8; x++ {
				const y = 0

				actual := indexes[x][y]
				var expected uint8
				if 0==x%2 {
					expected = 0
				} else {
					expected = 7
				}

				if expected != actual {
					t.Errorf("The actual color index at (x,y)=(%d,%d) is not what was expected.", x,y)
					t.Logf("EXPECTED: %d -> %s", expected, palette.Color(expected))
					t.Logf("ACTUAL:   %d -> %s", actual,   palette.Color(actual))
				}
			}

			for x:=0; x<8; x++ {
				const y = 1

				actual := indexes[x][y]
				const expected = 33

				if expected != actual {
					t.Errorf("The actual color index at (x,y)=(%d,%d) is not what was expected.", x,y)
					t.Logf("EXPECTED: %d -> %s", expected, palette.Color(expected))
					t.Logf("ACTUAL:   %d -> %s", actual,   palette.Color(actual))
				}
			}

			for x:=0; x<8; x++ {
				const y = 2

				actual := indexes[x][y]
				const expected = 12

				if expected != actual {
					t.Errorf("The actual color index at (x,y)=(%d,%d) is not what was expected.", x,y)
					t.Logf("EXPECTED: %d -> %s", expected, palette.Color(expected))
					t.Logf("ACTUAL:   %d -> %s", actual,   palette.Color(actual))
				}
			}

			for x:=0; x<8; x++ {
				const y = 3

				actual := indexes[x][y]
				const expected = 14

				if expected != actual {
					t.Errorf("The actual color index at (x,y)=(%d,%d) is not what was expected.", x,y)
					t.Logf("EXPECTED: %d -> %s", expected, palette.Color(expected))
					t.Logf("ACTUAL:   %d -> %s", actual,   palette.Color(actual))
				}
			}

			for x:=0; x<8; x++ {
				const y = 4

				actual := indexes[x][y]
				const expected = 98

				if expected != actual {
					t.Errorf("The actual color index at (x,y)=(%d,%d) is not what was expected.", x,y)
					t.Logf("EXPECTED: %d -> %s", expected, palette.Color(expected))
					t.Logf("ACTUAL:   %d -> %s", actual,   palette.Color(actual))
				}
			}

			for x:=0; x<8; x++ {
				const y = 5

				actual := indexes[x][y]
				const expected = 65

				if expected != actual {
					t.Errorf("The actual color index at (x,y)=(%d,%d) is not what was expected.", x,y)
					t.Logf("EXPECTED: %d -> %s", expected, palette.Color(expected))
					t.Logf("ACTUAL:   %d -> %s", actual,   palette.Color(actual))
				}
			}

			for x:=0; x<8; x++ {
				const y = 6

				actual := indexes[x][y]
				const expected = 48

				if expected != actual {
					t.Errorf("The actual color index at (x,y)=(%d,%d) is not what was expected.", x,y)
					t.Logf("EXPECTED: %d -> %s", expected, palette.Color(expected))
					t.Logf("ACTUAL:   %d -> %s", actual,   palette.Color(actual))
				}
			}

			for x:=0; x<8; x++ {
				const y = 7

				actual := indexes[x][y]
				const expected = 46

				if expected != actual {
					t.Errorf("The actual color index at (x,y)=(%d,%d) is not what was expected.", x,y)
					t.Logf("EXPECTED: %d -> %s", expected, palette.Color(expected))
					t.Logf("ACTUAL:   %d -> %s", actual,   palette.Color(actual))
				}
			}
		}
}
