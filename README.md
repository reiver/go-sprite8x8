# go-sprite8x8

Package **sprite8x8** provides a type that represents a **8×8 sprite**,
that seamlessly works with Go's built-in `"image"`, `"image/color"`, and `"image/draw"` packages.


## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-sprite8x8

[![GoDoc](https://godoc.org/github.com/reiver/go-sprite8x8?status.svg)](https://godoc.org/github.com/reiver/go-sprite8x8)


## Example
```go
import (
	"github.com/reiver/go-sprite8x8"
)

// ...

sprite := sprite8x8.Paletted{
	Pix:      pix,
	Palette:  palette,
	Category: "tiles",
	ID:       5,
}

// ...

draw.Draw(dst, rect, sprite, image.ZP, draw.Src)
```

## See Also

Also see:
* https://github.com/reiver/go-sprite32x32
